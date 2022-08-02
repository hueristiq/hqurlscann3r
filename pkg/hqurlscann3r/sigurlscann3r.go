package hqurlscann3r

import (
	"net/http"
	"regexp"

	"github.com/enenumxela/urlx/pkg/urlx"
	"github.com/hueristiq/hqurlscann3r/internal/configuration"
)

type Sigurlx struct {
	Client       *http.Client
	Params       []CommonVulnerableParameters
	Options      *configuration.Options
	JSRegex      *regexp.Regexp
	DOCRegex     *regexp.Regexp
	DATARegex    *regexp.Regexp
	STYLERegex   *regexp.Regexp
	MEDIARegex   *regexp.Regexp
	ARCHIVERegex *regexp.Regexp
	DOMXSSRegex  *regexp.Regexp
}

func New(options *configuration.Options) (Sigurlx, error) {
	hqurlscann3r := Sigurlx{}
	hqurlscann3r.Options = options
	hqurlscann3r.initCategories()
	hqurlscann3r.initParams()
	hqurlscann3r.initClient()

	return hqurlscann3r, nil
}

func (hqurlscann3r *Sigurlx) Process(URL string) (result Result, err error) {
	var res Response

	parsedURL, err := urlx.Parse(URL)
	if err != nil {
		return result, err
	}

	result.URL = parsedURL.String()

	if result.Category, err = hqurlscann3r.categorize(URL); err != nil {
		return result, err
	}

	if res, err = hqurlscann3r.DoHTTP(parsedURL.String()); err != nil {
		return result, err
	}

	result.StatusCode = res.StatusCode
	result.ContentType = res.ContentType
	result.ContentLength = res.ContentLength
	result.RedirectLocation = res.RedirectLocation

	query, err := getQuery(parsedURL.String())
	if err != nil {
		return result, err
	}

	if len(query) > 0 {
		if result.Category == "endpoint" {
			if res.IsEmpty() {
				res, _ = hqurlscann3r.DoHTTP(parsedURL.String())
			}

			if result.StatusCode == http.StatusForbidden {
				if result.ClietErrorBypass, err = hqurlscann3r.bypass4xx(parsedURL); err != nil {
					return result, err
				}
			}

			if result.ReflectedParameters, err = hqurlscann3r.ReflectedParamsProbe(parsedURL, query, res); err != nil {
				return result, err
			}

			if result.CommonVulnerableParameters, err = hqurlscann3r.CommonVulnParamsProbe(query); err != nil {
				return result, err
			}
		}
	}

	return result, nil
}
