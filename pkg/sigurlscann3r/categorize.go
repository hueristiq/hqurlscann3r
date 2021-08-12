package sigurlscann3r

func (sigurlscann3r *Sigurlx) initCategories() {
	sigurlscann3r.JSRegex, _ = newRegex(`(?m).*?\.(js)(\?.*?|)$`)
	sigurlscann3r.DOCRegex, _ = newRegex(`(?m).*?\.(pdf|xlsx|doc|docx|txt)(\?.*?|)$`)
	sigurlscann3r.DATARegex, _ = newRegex(`(?m).*?\.(json|xml|csv)(\?.*?|)$`)
	sigurlscann3r.STYLERegex, _ = newRegex(`(?m).*?\.(css)(\?.*?|)$`)
	sigurlscann3r.MEDIARegex, _ = newRegex(`(?m).*?\.(jpg|jpeg|png|ico|svg|gif|webp|mp3|mp4|woff|woff2|ttf|eot|tif|tiff)(\?.*?|)$`)
	sigurlscann3r.ARCHIVERegex, _ = newRegex(`(?m).*?\.(zip|tar|tar\.gz)(\?.*?|)$`)
}

func (sigurlscann3r *Sigurlx) categorize(URL string) (category string, err error) {
	if match := sigurlscann3r.JSRegex.MatchString(URL); match {
		category = "js"
	}

	if category == "" {
		if match := sigurlscann3r.DOCRegex.MatchString(URL); match {
			category = "doc"
		}
	}

	if category == "" {
		if match := sigurlscann3r.DATARegex.MatchString(URL); match {
			category = "data"
		}
	}

	if category == "" {
		if match := sigurlscann3r.STYLERegex.MatchString(URL); match {
			category = "style"
		}
	}

	if category == "" {
		if match := sigurlscann3r.MEDIARegex.MatchString(URL); match {
			category = "media"
		}
	}

	if category == "" {
		if match := sigurlscann3r.ARCHIVERegex.MatchString(URL); match {
			category = "archive"
		}
	}

	if category == "" {
		category = "endpoint"
	}

	return category, nil
}
