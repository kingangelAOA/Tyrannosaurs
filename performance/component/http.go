package component

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"bytes"
	"time"
	. "tyrannosaurs/util"
	. "tyrannosaurs/constant"
)

type Http struct {
	Name       string
	Comment string
	Protocol   string
	Host       string
	Port       string
	Method     string
	Path       string
	Coding     string
	Error      string
	Consuming  int64
	Body       string
	Query      map[string]string
	Header     map[string]string
	Form       map[string]string
	PostForm   map[string]string
	Cache      map[string]string
	CookieData
	Auth
	Base
}

type CookieData struct {
	clean   bool
	cookies []Cookie
}

type Cookie struct {
	key     string
	value   string
	domain  string
	Path    string
	secure  bool
	expires time.Time
}

type Auth struct {
	Username string
	Password string
}

func (h *Http) fetch() []byte {
	httpClient := &http.Client{}
	request := &http.Request{
		Method: h.Method,
		URL: &url.URL{
			Host: h.getHost(),
			Path: h.Path,
		},
	}
	h.preData()
	h.setRequest(request)
	start := NowMillisecond()
	res, err := httpClient.Do(request)
	h.Consuming = NowMillisecond() - start
	if err != nil {
		h.Error = err.Error()
		return nil
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		h.Error = err.Error()
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
	if &(h.Auth) != nil {
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
	if len(h.cookies) > 0 {
		for _, cookieData := range h.cookies {
			c := http.Cookie{
				Name:    cookieData.key,
				Value:   cookieData.value,
				Path:    cookieData.Path,
				Domain:  cookieData.domain,
				Secure:  cookieData.secure,
				Expires: cookieData.expires,
			}
			req.AddCookie(&c)
		}
	}
}

func (h *Http) preData() error {
	if err := ReplaceData(&(h.Name), h.Cache); err != nil {
		return err
	}
	if err := ReplaceData(&(h.Comment), h.Cache); err != nil {
		return err
	}
	if err := ReplaceData(&(h.Protocol), h.Cache); err != nil {
		return err
	}
	if err := ReplaceData(&(h.Host), h.Cache); err != nil {
		return err
	}
	if err := ReplaceData(&(h.Port), h.Cache); err != nil {
		return err
	}
	if err := ReplaceData(&(h.Path), h.Cache); err != nil {
		return err
	}
	if err := ReplaceData(&(h.Body), h.Cache); err != nil {
		return err
	}
	if err := h.replaceMap(h.Query); err != nil {
		return err
	}
	if err := h.replaceMap(h.Header); err != nil {
		return err
	}
	if err := h.replaceMap(h.Form); err != nil {
		return err
	}
	if err := h.replaceMap(h.PostForm); err != nil {
		return err
	}
	return nil
}

func (h *Http) replaceMap(target map[string]string) error {
	if target == nil {
		return nil
	}
	for k, v := range target {
		result, err := ReplaceDataHasResult(v, h.Cache)
		if err != nil {
			return err
		}
		target[k] = result
	}
	return nil
}

func (h *Http) Receive(message interface{}) error {
	if m, ok := message.(map[string]string); ok {
		for k, v := range m {
			if _, ok := h.Cache[k]; !ok {
				return HttpMergeCacheError
			}
			h.Cache[k] = v
		}
	}
	return nil
}

func (h *Http) Notify() error {
	for _, o := range h.ObserverList {
		res := h.fetch()
		data := ""
		if nil != res {
			data = string(res)
		}
		if err := o.Receive(data); err != nil {
			return err
		}
	}
	return nil
}
