package charset

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
)

var testCasesHeaders = map[string][]http.Header{
	"gbk": []http.Header{
		{"Content-Type": {"text/html;charset=gbk"}},
		{"Content-Type": {"text/html;charset=gBk"}},
	},
	"utf8": []http.Header{
		{"Content-Type": {"text/html;charset=utf8"}},
		{"Content-Type": {"text/html;charset=UTF8"}},
		{"Content-Type": {"text/html;charset=UTF-8"}},
	},
	"gb2312": []http.Header{
		{"Content-Type": {"text/html;charset=gb2312"}},
	},
	"": []http.Header{
		{"content-type": {"text/html;charset=utf8"}},
		{"Content-Type": {"text/html;charset=|||"}},
		{"Content-Type": {"text/html;charset=编码"}},
		{"Content-Type": {"text/html"}},
	},
}
var testCasesStrings = map[string][]string{
	"gbk":  {"text/html;charset=GBK", "text/html;charset=gbk"},
	"utf8": {"text/html;charset=utf8", "text/html;charset=utf-8"},
	"":     {"text/html", "foobar"},
}

func TestHeader(t *testing.T) {
	for charset, v := range testCasesHeaders {
		for _, x := range v {
			assert.Equal(t, Parse(x, nil), charset)
		}
	}
}

func TestString(t *testing.T) {
	for charset, v := range testCasesStrings {
		for _, x := range v {
			assert.Equal(t, Parse(x, nil), charset)
		}
	}
}

func TestBody(t *testing.T) {
	b1 := []byte(`<meta http-equiv="Content-Type" content="text/html; charset=gBk"/>')`)
	assert.Equal(t, Parse(nil, b1), "gbk")
	b2 := []byte(`<meta charset=UTF8>`)
	assert.Equal(t, Parse(nil, b2), "utf8")
}

func TestXMLBody(t *testing.T) {
	b1 := []byte(`<?xml version="1.0" encoding="utf-8"?>`)
	b2 := []byte(`<?xml version="1.0" encoding =  "utf-8"?>`)
	b3 := []byte(`<?xml version="1.0" encoding =  "utf-8"?>`)
	b4 := []byte(`<?xml version="1.0" encoding=" utf-8 "?>`)
	for _, b := range [][]byte{b1, b2, b3, b4} {
		assert.Equal(t, Parse(nil, b), "utf8")
	}
}

func TestEmpty(t *testing.T) {
	assert.Equal(t, Parse(nil, []byte{}), "")
	assert.Equal(t, Parse(nil, []byte("foobar")), "")
	resp, err := http.Get("http://httpbin.org/html")
	if err == nil {
		p := NewParser()
		b, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, p.Parse(nil, b), "")
		b = append(b, "<meta charset=utf8>"...)
		p.PeekSize = len(b) + 10
		assert.Equal(t, p.Parse(nil, b), "utf8")
	}
	defer resp.Body.Close()
}
