package main

import (
	"testing"
)

type TestCase struct {
	Name     string `naive:"name"`
	Password string `naive:"password"`
}

func TestConvert(t *testing.T) {
	res, err := convertStruct2RawReqStr(TestCase{
		Name:     "qianxi",
		Password: "notsobad",
	})
	if err != nil {
		t.Fail()
	}
	if res != "name=qianxi&password=notsobad" {
		t.Fail()
	}

	res, err = convertStruct2RawReqStr(TestCase{
		Name:     "qianxi",
		Password: "",
	})
	if err != nil {
		t.Fail()
	}
	if res != "name=qianxi&password=" {
		t.Fail()
	}
}
