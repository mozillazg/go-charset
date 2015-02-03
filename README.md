# go-charset
Get the content charset from header and html content-type.

[![Build Status](https://travis-ci.org/mozillazg/go-charset.svg?branch=master)](https://travis-ci.org/mozillazg/go-charset)
[![Coverage Status](https://coveralls.io/repos/mozillazg/go-charset/badge.svg?branch=master)](https://coveralls.io/r/mozillazg/go-charset?branch=master)
[![GoDoc](https://godoc.org/github.com/mozillazg/go-charset?status.svg)](https://godoc.org/github.com/mozillazg/go-charset)


## Installation

```
go get -u github.com/mozillazg/go-charset
```


## Usage

```go
s := `<meta http-equiv="Content-Type" content="text/html; charset=gbk"/>'
fmt.Println(charset.Parse(s, nil))
//gbk
```

```go
b := []byte(`<meta http-equiv="Content-Type" content="text/html; charset=gbk"/>')`)
fmt.Println(charset.Parse(nil, b))
//gbk
```

```go
resp, err := http.Get("http://www.mozilla.gr.jp/standards/contact.html")
if err == nil {
	b, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(charset.Parse(resp.Header, b))
	//euc-jp
}
defer resp.Body.Close()
```
