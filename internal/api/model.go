package api

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func (in *OnlyFans) GetPinnedPosts(id int) (Posts, error) {
	url := fmt.Sprintf(PINNED_POSTS_URL, id)
	in.Sign(url)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	in.Mutex.Lock()
	for k, v := range in.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range in.Cookies {
		req.Header.SetCookie(k, v)
	}
	in.Mutex.Unlock()
	resp := fasthttp.AcquireResponse()
	err := in.Client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	if err != nil {
		return Posts{}, err
	}
	defer fasthttp.ReleaseResponse(resp)

	var posts Posts
	err = json.Unmarshal(resp.Body(), &posts)
	return posts, nil
}
func (in *OnlyFans) GetTimeLinePosts(id int, timestamp string) (Posts, error) {
	var url string
	if timestamp == "" {
		url = fmt.Sprintf(TIMELINE_POSTS_URL, id)
	} else {
		url = fmt.Sprintf(TIMELINE_NEXT_POSTS_URL, id, timestamp)
	}

	in.Sign(url)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)

	in.Mutex.Lock()
	for k, v := range in.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range in.Cookies {
		req.Header.SetCookie(k, v)
	}
	in.Mutex.Unlock()
	resp := fasthttp.AcquireResponse()
	err := in.Client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	if err != nil {
		return Posts{}, err
	}
	defer fasthttp.ReleaseResponse(resp)
	var posts Posts
	err = json.Unmarshal(resp.Body(), &posts)
	if len(posts.List) == 0 {
		return posts, nil
	} else {
		p, err := in.GetTimeLinePosts(id, posts.List[len(posts.List)-1].PostedAtPrecise)
		if err != nil {
			return Posts{}, err
		}
		posts.List = append(posts.List, p.List...)
	}

	return posts, nil
}

func (in *OnlyFans) GetArchivedPosts(id int, timestamp string) (Posts, error) {
	var url string
	if timestamp == "" {
		url = fmt.Sprintf(ARCHIVED_POSTS_URL, id)
	} else {
		url = fmt.Sprintf(ARCHIVED_NEXT_POSTS_URL, id, timestamp)
	}

	in.Sign(url)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)

	in.Mutex.Lock()
	for k, v := range in.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range in.Cookies {
		req.Header.SetCookie(k, v)
	}
	in.Mutex.Unlock()
	resp := fasthttp.AcquireResponse()
	err := in.Client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	if err != nil {
		return Posts{}, err
	}
	defer fasthttp.ReleaseResponse(resp)
	var posts Posts
	err = json.Unmarshal(resp.Body(), &posts)
	if len(posts.List) == 0 {
		return posts, nil
	} else {
		p, err := in.GetTimeLinePosts(id, posts.List[len(posts.List)-1].PostedAtPrecise)
		if err != nil {
			return Posts{}, err
		}
		posts.List = append(posts.List, p.List...)
	}

	return posts, nil
}
func (in *OnlyFans) GetMessages(id int, messageId int) (Posts, error) {
	var url string
	if messageId == 0 {
		url = fmt.Sprintf(MESSAGES_URL, id)
	} else {
		url = fmt.Sprintf(MESSAGES_URL_NEXT, id, messageId)
	}

	in.Sign(url)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)

	in.Mutex.Lock()
	for k, v := range in.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range in.Cookies {
		req.Header.SetCookie(k, v)
	}
	in.Mutex.Unlock()
	resp := fasthttp.AcquireResponse()
	err := in.Client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	if err != nil {
		return Posts{}, err
	}
	defer fasthttp.ReleaseResponse(resp)
	var messages Posts
	err = json.Unmarshal(resp.Body(), &messages)
	if len(messages.List) == 0 {
		return messages, nil
	} else {
		m, err := in.GetMessages(id, messages.List[len(messages.List)-1].ID)
		if err != nil {
			return Posts{}, err
		}
		messages.List = append(messages.List, m.List...)
	}

	return messages, nil
}
func (in *OnlyFans) GetHighlights(id int) (Highlights, error) {
	url := fmt.Sprintf(HIGHLIGHTS_URL, id)
	in.Sign(url)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)

	in.Mutex.Lock()
	for k, v := range in.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range in.Cookies {
		req.Header.SetCookie(k, v)
	}
	in.Mutex.Unlock()
	resp := fasthttp.AcquireResponse()
	err := in.Client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	if err != nil {
		return Highlights{}, err
	}
	defer fasthttp.ReleaseResponse(resp)

	var highlights Highlights
	err = json.Unmarshal(resp.Body(), &highlights)
	return highlights, nil
}

func (in *OnlyFans) GetHighlight(id int) (Highlight, error) {
	url := fmt.Sprintf(HIGHLIGHT_URL, id)
	in.Sign(url)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)

	in.Mutex.Lock()
	for k, v := range in.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range in.Cookies {
		req.Header.SetCookie(k, v)
	}
	in.Mutex.Unlock()
	resp := fasthttp.AcquireResponse()
	err := in.Client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	if err != nil {
		return Highlight{}, err
	}
	defer fasthttp.ReleaseResponse(resp)

	var highlight Highlight
	err = json.Unmarshal(resp.Body(), &highlight)
	return highlight, nil
}

func (in *OnlyFans) GetStories(id int) ([]Story, error) {
	url := fmt.Sprintf(STORIES_URL, id)
	in.Sign(url)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)

	in.Mutex.Lock()
	for k, v := range in.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range in.Cookies {
		req.Header.SetCookie(k, v)
	}
	in.Mutex.Unlock()
	resp := fasthttp.AcquireResponse()
	err := in.Client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp)

	var stories []Story
	err = json.Unmarshal(resp.Body(), &stories)
	return stories, nil
}

func (s *Story) ExtractMedia() []Media {
	var media []Media
	if len(s.Media) == 0 {
		return []Media{}
	}
	for _, m := range s.Media {
		if !m.CanView {
			continue
		}
		media = append(media, Media{
			URL:     m.Files.Source.URL,
			MediaID: int(m.ID),
			PostID:  s.ID,
		})
	}
	return media
}
func (p *Posts) ExtractMedia() []Media {
	var media []Media
	for _, p := range p.List {
		if len(p.Media) == 0 {
			continue
		}
		for _, m := range p.Media {
			if !m.CanView {
				continue
			}
			media = append(media, Media{URL: m.Source.Source, MediaID: int(m.ID), PostID: p.ID})
		}
	}
	return media
}
