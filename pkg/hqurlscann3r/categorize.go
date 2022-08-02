package hqurlscann3r

func (hqurlscann3r *Sigurlx) initCategories() {
	hqurlscann3r.JSRegex, _ = newRegex(`(?m).*?\.(js)(\?.*?|)$`)
	hqurlscann3r.DOCRegex, _ = newRegex(`(?m).*?\.(pdf|xlsx|doc|docx|txt)(\?.*?|)$`)
	hqurlscann3r.DATARegex, _ = newRegex(`(?m).*?\.(json|xml|csv)(\?.*?|)$`)
	hqurlscann3r.STYLERegex, _ = newRegex(`(?m).*?\.(css)(\?.*?|)$`)
	hqurlscann3r.MEDIARegex, _ = newRegex(`(?m).*?\.(jpg|jpeg|png|ico|svg|gif|webp|mp3|mp4|woff|woff2|ttf|eot|tif|tiff)(\?.*?|)$`)
	hqurlscann3r.ARCHIVERegex, _ = newRegex(`(?m).*?\.(zip|tar|tar\.gz)(\?.*?|)$`)
}

func (hqurlscann3r *Sigurlx) categorize(URL string) (category string, err error) {
	if match := hqurlscann3r.JSRegex.MatchString(URL); match {
		category = "js"
	}

	if category == "" {
		if match := hqurlscann3r.DOCRegex.MatchString(URL); match {
			category = "doc"
		}
	}

	if category == "" {
		if match := hqurlscann3r.DATARegex.MatchString(URL); match {
			category = "data"
		}
	}

	if category == "" {
		if match := hqurlscann3r.STYLERegex.MatchString(URL); match {
			category = "style"
		}
	}

	if category == "" {
		if match := hqurlscann3r.MEDIARegex.MatchString(URL); match {
			category = "media"
		}
	}

	if category == "" {
		if match := hqurlscann3r.ARCHIVERegex.MatchString(URL); match {
			category = "archive"
		}
	}

	if category == "" {
		category = "endpoint"
	}

	return category, nil
}
