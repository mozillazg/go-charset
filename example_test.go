package charset_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mozillazg/go-charset"
)

func Example_header() {
	p := charset.NewParser()
	h := http.Header{"Content-Type": {"text/html;charset=utf8"}}
	fmt.Println(p.Parse(h, nil))
	fmt.Println(charset.Parse(h, nil))
	// Output:
	//utf8
	//utf8
}

func Example_string() {
	s := "text/html;charset=utf8"
	fmt.Println(charset.Parse(s, nil))
	// Output:
	//utf8
}

func Example_body() {
	b := []byte(`<meta http-equiv="Content-Type" content="text/html; charset=gbk"/>')`)
	fmt.Println(charset.Parse(nil, b))
	resp, err := http.Get("http://www.mozilla.gr.jp/standards/contact.html")
	if err == nil {
		b, _ = ioutil.ReadAll(resp.Body)
		fmt.Println(charset.Parse(resp.Header, b))
	}
	defer resp.Body.Close()
	// Output:
	//gbk
	//euc-jp
}
