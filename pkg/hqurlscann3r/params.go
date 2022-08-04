package hqurlscann3r

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	urlz "net/url"
	"strings"

	"github.com/hueristiq/hqurlscann3r/pkg/params"
	"github.com/hueristiq/url"
)

func (hqurlscann3r *Sigurlx) initParams() error {
	raw, err := ioutil.ReadFile(params.File())
	if err != nil {
		return err
	}

	if err = json.Unmarshal(raw, &hqurlscann3r.Params); err != nil {
		return err
	}

	return nil
}

func (hqurlscann3r *Sigurlx) CommonVulnParamsProbe(query urlz.Values) ([]CommonVulnerableParameters, error) {
	var commonVulnParams []CommonVulnerableParameters

	for parameter := range query {
		for i := range hqurlscann3r.Params {
			if strings.ToLower(hqurlscann3r.Params[i].Param) == strings.ToLower(parameter) {
				commonVulnParams = append(commonVulnParams, hqurlscann3r.Params[i])

				break
			}
		}
	}

	return commonVulnParams, nil
}

func (hqurlscann3r *Sigurlx) ReflectedParamsProbe(parsedURL *url.URL, query urlz.Values, res Response) ([]ReflectedParameters, error) {
	var reflectedParams []ReflectedParameters

	reflected, err := hqurlscann3r.checkReflection(parsedURL.String(), query, res)
	if err != nil {
		return reflectedParams, err
	}

	if len(reflected) > 0 {
		for _, parameter := range reflected {
			characters := []string{"\"", "'", "<", ">", "/"}

			var reflectedCharacters []string

			for _, char := range characters {
				wasReflected, err := hqurlscann3r.checkAppend(parsedURL, query, parameter, "aprefix"+char+"asuffix")
				if err != nil {
					continue
				}

				if wasReflected {
					reflectedCharacters = append(reflectedCharacters, char)
				}
			}

			if len(reflectedCharacters) > 2 {
				reflectedParams = append(reflectedParams, ReflectedParameters{Param: parameter, Characters: reflectedCharacters})
			}
		}
	}

	return reflectedParams, nil
}

func getQuery(URL string) (urlz.Values, error) {
	var query urlz.Values

	queryUnescaped, err := urlz.QueryUnescape(URL)
	if err != nil {
		return query, err
	}

	parsedURL, err := urlz.Parse(queryUnescaped)
	if err != nil {
		return query, err
	}

	query, err = urlz.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return query, err
	}

	return query, nil
}

func (hqurlscann3r *Sigurlx) checkReflection(URL string, query urlz.Values, res Response) ([]string, error) {
	var reflected []string

	if res.IsEmpty() {
		res, _ = hqurlscann3r.DoHTTP(URL)
	}

	if res.StatusCode >= http.StatusMultipleChoices && res.StatusCode < http.StatusBadRequest {
		return reflected, nil
	}

	if res.ContentType != "" && !strings.Contains(res.ContentType, "html") {
		return reflected, nil
	}

	for param, value := range query {
		for _, v := range value {
			if !strings.Contains(string(res.Body), v) {
				continue
			}

			reflected = append(reflected, param)
		}
	}

	return reflected, nil
}

func (hqurlscann3r *Sigurlx) checkAppend(parsedURL *url.URL, query urlz.Values, param, suffix string) (bool, error) {
	val := query.Get(param)

	query.Set(param, val+suffix)
	parsedURL.RawQuery = query.Encode()

	reflected, err := hqurlscann3r.checkReflection(parsedURL.String(), query, Response{})
	if err != nil {
		return false, err
	}

	for _, r := range reflected {
		if r == param {
			return true, nil
		}
	}

	query.Set(param, val)

	return false, nil
}
