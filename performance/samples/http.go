package samples

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"bytes"
	"time"
	. "tyrannosaurs/util"
)

type Http struct {
	Name       string
	Annotation string
	Protocol   string
	Host       string
	Port       string
	Method     string
	Path       string
	Coding     string
	Consuming  int64
	error      string
	Query      map[string]string
	Header     map[string]string
	Form       map[string]string
	PostForm   map[string]string
	Body       string
	Params     *[]Param
	*CookieData
	*Auth
}

type Param struct {
	Key   string
	Value string
}

type CookieData struct {
	Clean   bool
	Cookies []Cookie
}

type Cookie struct {
	Key     string
	Value   string
	Domain  string
	path    string
	Secure  bool
	Expires time.Time
}

type Auth struct {
	Username string
	Password string
}

func (h *Http) Fetch() []byte {
	httpClient := &http.Client{}
	request := &http.Request{
		Method: h.Method,
		URL: &url.URL{
			Host: h.getHost(),
			Path: h.Path,
		},
	}
	h.setRequest(request)
	start := NowMillisecond()
	res, err := httpClient.Do(request)
	h.Consuming = NowMillisecond() - start
	if err != nil {
		h.error = err.Error()
		return nil
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		h.error = err.Error()
		return nil
	}
	return result
}

func (h *Http) getHost() string {
	if h.Port != "" {
		return fmt.Sprintf("%s:%s", h.Host, h.Port)
	}
	return h.Host
}

func (h *Http) setRequest(req *http.Request) {
	if h.Auth != nil {
		req.SetBasicAuth(h.Username, h.Password)
	}
	if h.Form != nil {
		for k, v := range h.Form {
			req.Form.Add(k, v)
		}
	}
	if h.Body != "" {
		req.Body = ioutil.NopCloser(bytes.NewReader([]byte(h.Body)))
	}
	if h.PostForm != nil {
		for k, v := range h.PostForm {
			req.PostForm.Add(k, v)
		}
	}
	if h.Query != nil {
		values := req.URL.Query()
		for k, v := range h.Query {
			values.Add(k, v)
		}
		req.URL.RawQuery = values.Encode()
	}
	if h.Header != nil {
		for k, v := range h.Header {
			req.Header.Add(k, v)
		}
	}
	if len(h.Cookies) > 0 {
		for _, cookieData := range h.Cookies {
			c := http.Cookie{
				Name:    cookieData.Key,
				Value:   cookieData.Value,
				Path:    cookieData.path,
				Domain:  cookieData.Domain,
				Secure:  cookieData.Secure,
				Expires: cookieData.Expires,
			}
			req.AddCookie(&c)
		}
	}
}
