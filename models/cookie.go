package models

import (
	"net/http"
	"net/url"
	"sync"
)

type jar struct {
	lk      sync.Mutex
	cookies map[string][]*http.Cookie
}

func newJar() *jar {
	jar := new(jar)
	jar.cookies = make(map[string][]*http.Cookie)
	return jar
}

func (jar *jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.lk.Lock()
	if _, ok := jar.cookies[u.Host]; ok {
		for _, c := range cookies {
			jar.cookies[u.Host] = append(jar.cookies[u.Host], c)
		}
	} else {
		jar.cookies[u.Host] = cookies
	}
	jar.lk.Unlock()
}

func (jar *jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies[u.Host]
}
