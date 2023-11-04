package actions

import (
	"OnlyScraper/internal/api"
	"OnlyScraper/internal/config"
	"OnlyScraper/internal/database"
	"fmt"
	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

func ScrapeAction(c *cli.Context) error {
	in := c.Context.Value("instance").(*api.OnlyFans)
	logger := c.Context.Value("logger").(pterm.Logger)
	cfg := c.Context.Value("config").(*config.Config)
	f := c.String("model")
	var models []string
	if f == "" {
		subscriptions, err := in.GetSubscriptions()
		if err != nil {
			return err
		}
		for _, sub := range subscriptions.List {
			models = append(models, sub.Username)
		}
	} else {
		split := strings.Split(f, ",")
		for _, s := range split {
			models = append(models, strings.Replace(s, " ", "", -1))
		}
	}

	logger.Info("Scraping", logger.ArgsFromMap(map[string]any{
		"models": models,
	}))

	for _, m := range models {
		err := os.MkdirAll(path.Join("output", m, "Metadata"), 0777)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		db, err := database.CreateAndConnect(path.Join("output", m, "Metadata", "model.db"))
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		existing, err := db.GetExisting()
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		model, err := in.GetProfile(m)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		logger.Info("Start Scraping", logger.ArgsFromMap(map[string]any{
			"model":  model.Username,
			"photos": model.PhotosCount,
			"videos": model.VideosCount,
			"audios": model.AudiosCount,
		}))
		sm := ysmrr.NewSpinnerManager(
			ysmrr.WithSpinnerColor(colors.FgHiBlue),
			ysmrr.WithCompleteColor(colors.FgHiBlue),
		)
		sm.Start()
		pinnedSpinner := sm.AddSpinner("Fetching Pinned Posts")
		pinnedPosts, err := in.GetPinnedPosts(model.ID)
		pinnedSpinner.Complete()
		timelineSpinner := sm.AddSpinner("Fetching Timeline Posts")
		timelinePosts, err := in.GetTimeLinePosts(model.ID, "")
		timelineSpinner.Complete()
		archivedSpinner := sm.AddSpinner("Fetching Archived Posts")
		archivedPosts, err := in.GetArchivedPosts(model.ID, "")
		archivedSpinner.Complete()
		highlightsSpinner := sm.AddSpinner("Fetching Highlights")
		highlights, err := in.GetHighlights(model.ID)
		highlightsSpinner.Complete()
		messagesSpinner := sm.AddSpinner("Fetching Messages")
		messages, err := in.GetMessages(model.ID, 0)
		messagesSpinner.Complete()
		storiesSpinner := sm.AddSpinner("Fetching Stories")
		stories, err := in.GetStories(model.ID)
		storiesSpinner.Complete()
		var media []api.Media

		media = append(media, pinnedPosts.ExtractMedia(cfg.Types.Pinned, cfg.Types.Paid, cfg.Types.PaidPreview)...)
		media = append(media, timelinePosts.ExtractMedia(cfg.Types.Timeline, cfg.Types.Paid, cfg.Types.PaidPreview)...)
		media = append(media, archivedPosts.ExtractMedia(cfg.Types.Archived, cfg.Types.Paid, cfg.Types.PaidPreview)...)
		media = append(media, messages.ExtractMedia(cfg.Types.Message, cfg.Types.Paid, cfg.Types.PaidPreview)...)

		highlightDetailSpinner := sm.AddSpinner("Getting Highlight Details")
		for _, h := range highlights.List {
			highlight, err := in.GetHighlight(h.ID)
			if err != nil {
				logger.Error(err.Error())
				continue
			}
			for _, story := range highlight.Stories {
				media = append(media, story.ExtractMedia(cfg.Types.Highlights)...)
			}
		}
		for _, story := range stories {
			media = append(media, story.ExtractMedia(cfg.Types.Highlights)...)
		}
		highlightDetailSpinner.Complete()
		sm.Stop()
		logger.Info("Fetched Posts", logger.ArgsFromMap(map[string]any{
			"pinned":     len(pinnedPosts.List),
			"timeline":   len(timelinePosts.List),
			"archived":   len(archivedPosts.List),
			"highlights": len(highlights.List),
			"messages":   len(messages.List),
			"stories":    len(stories),
			"urls":       len(media),
		}))

		threads := 10

		existingIDs := make(map[int]struct{})

		for _, exist := range existing {
			if exist.Link != "" {
				existingIDs[exist.MediaId] = struct{}{}
			}
		}

		var filtered []api.Media

		for _, m := range media {
			if _, found := existingIDs[m.MediaID]; !found {
				filtered = append(filtered, m)
			}
		}

		if len(filtered) == 0 {
			logger.Info("All content has already been scraped")

		} else {
			mediaChannel := make(chan api.Media, len(filtered))
			done := make(chan bool, threads)
			var wg sync.WaitGroup

			p := mpb.New(
				mpb.WithOutput(color.Output),
			)
			bars := make([]*mpb.Bar, 10)
			for i := range bars {
				bars[i] = p.AddBar(0,
					mpb.PrependDecorators(
						decor.CountersKiloByte("% .2f / % .2f"),
					),
					mpb.AppendDecorators(
						decor.AverageSpeed(decor.UnitKB, "% .2f")),
				)
			}

			for i := 0; i < threads; i++ {
				wg.Add(1)
				go downloadWorker(model, mediaChannel, done, &wg, logger, bars[i], db, cfg)
			}

			for _, m := range filtered {
				mediaChannel <- m
			}
			close(mediaChannel)
			wg.Wait()
			p.Wait()
			close(done)
			for i := 0; i < threads; i++ {
				<-done
			}

		}

		logger.Info(fmt.Sprintf("Finished scraping %v", model.Username))
	}

	return nil
}

func downloadWorker(model api.Profile, media <-chan api.Media, done chan<- bool, wg *sync.WaitGroup, logger pterm.Logger, bar *mpb.Bar, db *database.Database, cfg *config.Config) {
	for m := range media {
		client := http.Client{Timeout: 10 * time.Second}
		resp, err := client.Get(m.URL)
		if err != nil {
			continue
		}
		if resp.StatusCode != http.StatusOK {
			continue
		}
		contentType := resp.Header.Get("Content-Type")
		var dir, fileExt string
		switch {
		case strings.Contains(contentType, "image/jpeg"):
			fileExt = ".jpg"
			dir = "output/" + model.Username + "/" + m.Type + "/" + cfg.MediaType.Images + "/"
		case strings.Contains(contentType, "image/png"):
			fileExt = ".png"
			dir = "output/" + model.Username + "/" + m.Type + "/" + cfg.MediaType.Images + "/"
		case strings.Contains(contentType, "image/gif"):
			fileExt = ".gif"
			dir = "output/" + model.Username + "/" + m.Type + "/" + cfg.MediaType.Videos + "/"
		case strings.Contains(contentType, "video/mp4"):
			fileExt = ".mp4"
			dir = "output/" + model.Username + "/" + m.Type + "/" + cfg.MediaType.Videos + "/"
		case strings.Contains(contentType, "audio/mpeg"):
			fileExt = ".mp3"
			dir = "output/" + model.Username + "/" + m.Type + "/" + cfg.MediaType.Audios + "/"
		default:
			fileExt = "" // Unknown or non-media file type
			dir = ""
		}
		filename := uuid.NewString() + fileExt
		dir = strings.Replace(dir, "{username}", model.Username, -1)
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			continue
		}
		filepath := path.Join(dir, filename)
		out, err := os.Create(filepath)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		bar.SetCurrent(0)
		bar.SetTotal(resp.ContentLength, false)
		proxyReader := bar.ProxyReader(resp.Body)
		_, err = out.ReadFrom(proxyReader)
		if err != nil {
			continue
		}
		err = db.InsertMedia(m.MediaID, m.PostID, m.URL, dir, filename)
		err = resp.Body.Close()
		if err != nil {
			continue
		}
	}
	bar.SetTotal(bar.Current(), true)
	wg.Done()
	done <- true
}
