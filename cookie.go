package main

import (
	"errors"
	"net/http"
)

var ErrNoCookie = errors.New("no such cookie")

// keys.
const (
	JSESSIONID = "JSESSIONID"
	TGC        = "TGC"
)

// Cookies stores cookies for a client.
type Cookies map[string]string

func NewCookies() Cookies {
	return make(Cookies)
}

// Get returns the value of the named cookie.
func (c Cookies) Get(name string) (string, error) {
	if v, ok := c[name]; ok {
		return v, nil
	}
	return "", ErrNoCookie
}

// Set sets the named cookie value.
func (c Cookies) Set(name string, cookies []*http.Cookie) {
	if c == nil {
		c = make(map[string]string)
	}
	for _, cookie := range cookies {
		if cookie.Name == name {
			c[name] = cookie.Value
			break
		}
	}
}
