package main

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

var (
	ErrorNotStruct = errors.New("not struct")
)

// extract `sign` and `timestamp`
// then need use the two params to sign the request.
func reqSuffix(source string) (string, string, bool) {
	pattern := `sign=([A-Za-z0-9]{40})&timeStamp=([\d]{13})`
	reg := regexp.MustCompile(pattern)
	params := reg.FindStringSubmatch(source)

	// log.Println("extract sign and timestamp params: ", params)

	if len(params) != 3 {
		return "", "", false
	}

	return params[1], params[2], true
}

// extract `name`
func stuName(source string) (string, bool) {
	pattern := "<div class=\"weui-cell__ft\">([\u4e00-\u9fa5_a-zA-Z0-9]+)</div>"
	reg := regexp.MustCompile(pattern)
	params := reg.FindStringSubmatch(source)

	// log.Printf("extract student name params: %v", params)

	if len(params) != 2 {
		return "", false
	}

	return params[1], true
}

// convert struct to raw req stirng.
func convertStruct2RawReqStr(s interface{}) (string, error) {
	ts, vs := reflect.TypeOf(s), reflect.ValueOf(s)
	if ts.Kind() != reflect.Struct {
		return "", ErrorNotStruct
	}

	strs := []string{}
	for i := 0; i < vs.NumField(); i++ {
		fv := vs.Field(i).String()
		if fv == "" {
			fv = ts.Field(i).Tag.Get("default")
		}
		s := ts.Field(i).Tag.Get("naive") + "=" + fv
		strs = append(strs, s)
	}

	return strings.Join(strs, "&"), nil
}
