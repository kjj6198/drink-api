package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type H map[string]interface{}

type Options struct {
	Method  string
	Headers H
	Params  H
	Body    H
	Cookie  *http.Cookie
}

// GetHttpRequest return a http request by given URL and method.
func GetHttpRequest(method string, URL string, reader io.Reader) (req *http.Request, err error) {
	if (method == http.MethodPost || method == http.MethodPut) && reader != nil {
		return http.NewRequest(strings.ToUpper(method), URL, reader)
	}

	return http.NewRequest(strings.ToUpper(method), URL, nil)
}

// DoRequest return response in string.
func DoRequest(URL string, options *Options) (result []byte, err error) {
	client := http.DefaultClient

	if options != nil && options.Params != nil {
		qs := url.Values{}
		for k, v := range options.Params {
			qs.Set(k, v.(string))
		}
		URL = fmt.Sprintf("%s?%s", URL, qs.Encode())
	}

	d, _ := json.Marshal(options.Body)

	request, err := GetHttpRequest(options.Method, URL, bytes.NewBuffer(d))

	if err != nil {
		return nil, err
	}

	if options != nil && options.Cookie != nil {
		request.AddCookie(options.Cookie)
	}

	if options != nil && options.Headers != nil {
		for k, v := range options.Headers {
			request.Header.Add(k, v.(string))
		}
	}

	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	return data, nil
}
