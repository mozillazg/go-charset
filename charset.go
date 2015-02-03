package charset

import (
	"net/http"
	"regexp"
	"strings"
)

var charsetRe = regexp.MustCompile(`(?i:(?:charset|encoding)\s*=\s*['\"]? *([-\w]+))`)

var PeekSize int = 255

type Parser struct {
	PeekSize int
}

func NewParser() *Parser {
	return &Parser{PeekSize}
}

func (p *Parser) Parse(obj interface{}, data []byte) string {
	var matchs []string
	end := 0
	contentType := ""
	ct := ""

	if data != nil {
		if len(data) > p.PeekSize {
			end = p.PeekSize
		} else {
			end = len(data)
		}
	}

	// Parse("text/html;charset=gbk", nil)
	switch obj.(type) {
	case string:
		contentType, _ = obj.(string)
	case http.Header:
		header, _ := obj.(http.Header)
		contentType = header.Get("Content-Type")
	}

	if contentType != "" {
		// guest from obj first
		matchs = charsetRe.FindStringSubmatch(contentType)
	}
	if len(matchs) == 0 && end > 0 {
		// guest from content body (html/xml) header
		contentType = string(data[0:end])
		matchs = charsetRe.FindStringSubmatch(contentType)
	}

	if len(matchs) != 0 {
		ct = strings.ToLower(matchs[1])
		if ct == "utf-8" {
			ct = "utf8"
		}
	}
	return ct
}

func Parse(obj interface{}, data []byte) string {
	return NewParser().Parse(obj, data)
}
