package sigurlscann3r

import (
	"net/http"
	"net/url"
	"regexp"
)

type Sigurlx struct {
	Client       *http.Client
	Params       []CommonVulnParam
	Options      *Options
	JSRegex      *regexp.Regexp
	DOCRegex     *regexp.Regexp
	DATARegex    *regexp.Regexp
	STYLERegex   *regexp.Regexp
	MEDIARegex   *regexp.Regexp
	ARCHIVERegex *regexp.Regexp
	DOMXSSRegex  *regexp.Regexp
}

func New(options *Options) (Sigurlx, error) {
	sigurlscann3r := Sigurlx{}
	sigurlscann3r.Options = options
	sigurlscann3r.initCategories()
	sigurlscann3r.initParams()
	sigurlscann3r.initClient()

	return sigurlscann3r, nil
}

func (sigurlscann3r *Sigurlx) Process(URL string) (result Result, err error) {
	var res Response

	parsedURL, err := url.Parse(URL)
	if err != nil {
		return result, err
	}

	result.URL = parsedURL.String()

	if result.Category, err = sigurlscann3r.categorize(URL); err != nil {
		return result, err
	}

	if res, err = sigurlscann3r.DoHTTP(parsedURL.String()); err != nil {
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
			if result.CommonVulnParams, err = sigurlscann3r.CommonVulnParamsProbe(query); err != nil {
				return result, err
			}

			if res.IsEmpty() {
				res, _ = sigurlscann3r.DoHTTP(parsedURL.String())
			}

			if result.ReflectedParams, err = sigurlscann3r.ReflectedParamsProbe(parsedURL, query, res); err != nil {
				return result, err
			}
		}
	}

	return result, nil
}
