package api

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

func (in *OnlyFans) GetProfile(username string) (Profile, error) {
	url := fmt.Sprintf(PROFILE_URL, username)
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
		return Profile{}, err
	}
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() != http.StatusOK {
		if resp.StatusCode() == 400 && username == "me" {
			return Profile{}, fmt.Errorf("unauthorized, check your auth.json")
		}
		return Profile{}, fmt.Errorf("unexpected status code: %v", resp.StatusCode())
	}

	var profile Profile
	err = json.Unmarshal(resp.Body(), &profile)
	return profile, nil
}
