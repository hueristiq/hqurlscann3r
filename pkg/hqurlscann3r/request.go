package hqurlscann3r

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
	"unicode/utf8"
)

func (hqurlscann3r *Sigurlx) initClient() error {
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(hqurlscann3r.Options.Timeout) * time.Second,
			KeepAlive: time.Second,
		}).DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if hqurlscann3r.Options.HTTPProxy != "" {
		if proxyURL, err := url.Parse(hqurlscann3r.Options.HTTPProxy); err == nil {
			tr.Proxy = http.ProxyURL(proxyURL)
		}
	}

	re := func(_ *http.Request, _ []*http.Request) error {
		return http.ErrUseLastResponse
	}

	if hqurlscann3r.Options.FollowRedirects {
		re = nil
	}

	if hqurlscann3r.Options.FollowHostRedirects {
		re = func(redirectedRequest *http.Request, previousRequest []*http.Request) error {
			newHost := redirectedRequest.URL.Host
			oldHost := previousRequest[0].URL.Host

			if newHost != oldHost {
				return http.ErrUseLastResponse
			}

			return nil
		}
	}

	hqurlscann3r.Client = &http.Client{
		Timeout:       time.Duration(hqurlscann3r.Options.Timeout) * time.Second,
		Transport:     tr,
		CheckRedirect: re,
	}

	return nil
}

func (hqurlscann3r *Sigurlx) DoHTTP(URL string) (Response, error) {
	var response Response

	headers := map[string]string{
		"User-Agent": hqurlscann3r.Options.UserAgent,
	}

	res, err := hqurlscann3r.httpRequest(http.MethodGet, URL, headers)
	if err != nil {
		return response, err
	}

	response.Headers = res.Header.Clone()

	// websockets don't have a readable body
	if res.StatusCode != http.StatusSwitchingProtocols {
		// always read the full body so we can re-use the tcp connection
		if response.Body, err = ioutil.ReadAll(res.Body); err != nil {
			return response, err
		}
	}

	if err := res.Body.Close(); err != nil {
		return response, err
	}

	response.StatusCode = res.StatusCode
	response.ContentType = response.GetHeaderPart("Content-Type", ";")
	response.ContentLength = utf8.RuneCountInString(string(response.Body))
	response.RedirectLocation = response.GetHeaderPart("Location", ";")

	return response, nil
}

func (hqurlscann3r *Sigurlx) httpRequest(method, URL string, headers map[string]string) (res *http.Response, err error) {
	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		return
	}

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	res, err = hqurlscann3r.Client.Do(req)
	if err != nil {
		return
	}

	return
}
