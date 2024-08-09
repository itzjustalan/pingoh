package services

import "regexp"

func HeaderForEncoding(enc string) string {
	h := ""
	switch enc {
	case "text":
		h = "text/plain"
	case "html":
		h = "text/html"
	case "xml":
		h = "text/xml"
	case "json":
		h = "application/json"
	case "form":
		h = "application/x-www-form-urlencoded"
	}
	return h
}

func IsAlphanumeric(str string) bool {
	alphanumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return alphanumericRegex.MatchString(str)
}
