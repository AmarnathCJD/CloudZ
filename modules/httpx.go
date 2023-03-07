package modules

import (
	"bytes"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"
)

type Coookies []*http.Cookie

func (c Coookies) Get(name string) (string, bool) {
	for _, item := range c {
		if item.Name == name {
			return item.Value, true
		}
	}

	return "", false
}

func (c Coookies) Set(name string, val string) bool {
	for _, item := range c {
		if item.Name == name {
			item.Value = val
			return true
		}
	}

	return false
}

type HttpSession struct {
	client  *http.Client
	headers http.Header
}

func NewHttpSession() (*HttpSession, error) {
	opts := &cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}

	cookieJar, err := cookiejar.New(opts)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
		Jar:     cookieJar,
	}
	return &HttpSession{
		client: httpClient,
	}, nil
}

func (httpx *HttpSession) Get(url string, headers http.Header, allowRedirects bool) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if httpx.headers != nil {
		for key, value := range httpx.headers {
			req.Header.Set(key, value[0])
		}
	}

	for key, value := range headers {
		req.Header.Set(key, value[0])
	}

	if !allowRedirects {
		httpx.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}

		defer func() {
			httpx.client.CheckRedirect = nil
		}()
	}

	resp, err := httpx.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (httpx *HttpSession) Post(url string, headers http.Header, data []byte, allowRedirects bool) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	if httpx.headers != nil {
		for key, value := range httpx.headers {
			req.Header.Set(key, value[0])
		}
	}

	for key, value := range headers {
		req.Header.Set(key, value[0])
	}

	if !allowRedirects {
		httpx.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}

		defer func() {
			httpx.client.CheckRedirect = nil
		}()
	}

	resp, err := httpx.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (httpx *HttpSession) Cookies(host string) Coookies {
	domain := &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   "/",
	}

	rawCookies := httpx.client.Jar.Cookies(domain)
	if rawCookies != nil {
		return Coookies(rawCookies)
	}

	return nil
}

func (httpx *HttpSession) SetHeaders(headers http.Header) {
	httpx.headers = headers
}
