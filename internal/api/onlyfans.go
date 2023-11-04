package api

import (
	"OnlyScraper/internal/config"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type OnlyFans struct {
	DRO     DRO
	Client  *fasthttp.Client
	Auth    *config.Auth
	Headers map[string]string
	Cookies map[string]string
	Mutex   sync.Mutex
	Me      Profile
}
type DRO struct {
	StaticParam       string   `json:"static_param"`
	Format            string   `json:"format"`
	ChecksumIndexes   []int    `json:"checksum_indexes"`
	ChecksumConstants []int    `json:"checksum_constants"`
	ChecksumConstant  int      `json:"checksum_constant"`
	AppToken          string   `json:"app_token"`
	RemoveHeaders     []string `json:"remove_headers"`
	ErrorCode         int      `json:"error_code"`
	Message           string   `json:"message"`
}

func Init(auth *config.Auth) (*OnlyFans, error) {
	readTimeout := 10 * time.Second
	writeTimeout := 10 * time.Second
	maxIdleConnDuration, _ := time.ParseDuration("1h")
	client := &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}
	in := OnlyFans{
		Client: client,
		Auth:   auth,
	}
	err := in.GetDRO()
	if err != nil {
		return nil, err
	}
	in.SetHeaders()
	in.SetCookies()

	in.Me, err = in.GetProfile("me")

	return &in, err
}

func (in *OnlyFans) SetHeaders() {
	in.Headers = map[string]string{
		"accept":     "application/json, text/plain, */*",
		"app-token":  in.DRO.AppToken,
		"user-id":    in.Auth.AuthID,
		"x-bc":       in.Auth.XBC,
		"user-agent": in.Auth.UserAgent,
		"refer":      "https://onlyfans.com",
	}
}

func (in *OnlyFans) SetCookies() {
	in.Cookies = map[string]string{
		"sess":    in.Auth.Session,
		"auth_id": in.Auth.AuthID,
	}
}

func (in *OnlyFans) Sign(link string) {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()
	t := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	parsedURL, _ := url.Parse(link)
	path := parsedURL.Path
	query := parsedURL.RawQuery
	if query != "" {
		path += "?" + query
	}
	a := []string{in.DRO.StaticParam, t, path, in.Auth.AuthID}
	msg := strings.Join(a, "\n")
	message := []byte(msg)
	hashObject := sha1.Sum(message)
	sha1Sign := hex.EncodeToString(hashObject[:])
	sha1B := []byte(sha1Sign)
	checksum := 0
	for _, i := range in.DRO.ChecksumIndexes {
		checksum += int(sha1B[i])
	}
	checksum += in.DRO.ChecksumConstant
	formatString := in.DRO.Format
	formatString = strings.Replace(formatString, "{}", sha1Sign, 1)
	formatString = strings.Replace(formatString, "{:x}", fmt.Sprintf("%x", checksum), 1)
	in.Headers["sign"] = formatString
	in.Headers["time"] = t
}

func (in *OnlyFans) GetDRO() error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(DRO_URL)
	req.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	err := in.Client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	if err != nil {
		return err
	}
	defer fasthttp.ReleaseResponse(resp)

	var dro DRO
	err = json.Unmarshal(resp.Body(), &dro)
	if err != nil {
		return err
	}
	in.DRO = dro

	return nil
}
