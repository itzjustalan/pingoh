package services

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
