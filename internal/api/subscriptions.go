package api

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
)

func (in *OnlyFans) GetSubscriptions() (Subscriptions, error) {
	s := Subscriptions{}
	for offset := 0; offset < in.Me.SubscribesCount; offset += 10 {
		url := fmt.Sprintf(SUBSCRIPTIONS_URL, strconv.FormatInt(int64(offset), 10))
		in.Sign(url)

		req := fasthttp.AcquireRequest()
		req.SetRequestURI(url)
		req.Header.SetMethod(fasthttp.MethodGet)

		for k, v := range in.Headers {
			req.Header.Set(k, v)
		}
		for k, v := range in.Cookies {
			req.Header.SetCookie(k, v)
		}
		resp := fasthttp.AcquireResponse()
		err := in.Client.Do(req, resp)

		fasthttp.ReleaseRequest(req)
		if err != nil {
			return Subscriptions{}, err
		}

		if resp.StatusCode() != http.StatusOK {
			if resp.StatusCode() == http.StatusUnauthorized {
				return Subscriptions{}, fmt.Errorf("unauthorized, check your auth.json")
			}
			return Subscriptions{}, fmt.Errorf("unexpected status code: %v", resp.StatusCode())
		}

		var subs Subscriptions
		err = json.Unmarshal(resp.Body(), &subs)
		if err != nil {
			return Subscriptions{}, err
		}

		s.List = append(s.List, subs.List...)

		fasthttp.ReleaseResponse(resp)
	}
	return s, nil
}
