package sigurlscann3r

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/enenumxela/urlx/pkg/urlx"
	"github.com/signedsecurity/sigurlscann3r/pkg/params"
)

func (sigurlscann3r *Sigurlx) initParams() error {
	raw, err := ioutil.ReadFile(params.File())
	if err != nil {
		return err
	}

	if err = json.Unmarshal(raw, &sigurlscann3r.Params); err != nil {
		return err
	}

	return nil
}

func (sigurlscann3r *Sigurlx) CommonVulnParamsProbe(query url.Values) ([]CommonVulnerableParameters, error) {
	var commonVulnParams []CommonVulnerableParameters

	for parameter := range query {
		for i := range sigurlscann3r.Params {
			if strings.ToLower(sigurlscann3r.Params[i].Param) == strings.ToLower(parameter) {
				commonVulnParams = append(commonVulnParams, sigurlscann3r.Params[i])

				break
			}
		}
	}

	return commonVulnParams, nil
}

func (sigurlscann3r *Sigurlx) ReflectedParamsProbe(parsedURL *urlx.URL, query url.Values, res Response) ([]ReflectedParameters, error) {
	var reflectedParams []ReflectedParameters

	reflected, err := sigurlscann3r.checkReflection(parsedURL.String(), query, res)
	if err != nil {
		return reflectedParams, err
	}

	if len(reflected) > 0 {
		for _, parameter := range reflected {
			characters := []string{"\"", "'", "<", ">", "/"}

			var reflectedCharacters []string

			for _, char := range characters {
				wasReflected, err := sigurlscann3r.checkAppend(parsedURL, query, parameter, "aprefix"+char+"asuffix")
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

func getQuery(URL string) (url.Values, error) {
	var query url.Values

	queryUnescaped, err := url.QueryUnescape(URL)
	if err != nil {
		return query, err
	}

	parsedURL, err := url.Parse(queryUnescaped)
	if err != nil {
		return query, err
	}

	query, err = url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return query, err
	}

	return query, nil
}

func (sigurlscann3r *Sigurlx) checkReflection(URL string, query url.Values, res Response) ([]string, error) {
	var reflected []string

	if res.IsEmpty() {
		res, _ = sigurlscann3r.DoHTTP(URL)
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

func (sigurlscann3r *Sigurlx) checkAppend(parsedURL *urlx.URL, query url.Values, param, suffix string) (bool, error) {
	val := query.Get(param)

	query.Set(param, val+suffix)
	parsedURL.RawQuery = query.Encode()

	reflected, err := sigurlscann3r.checkReflection(parsedURL.String(), query, Response{})
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
