package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

var stu UserInfo

func init() {
	stu = UserInfo{
		StudentID: os.Getenv("student_id"),
		Password:  os.Getenv("password"),
	}
	if stu.Password == "" || stu.StudentID == "" {
		log.Fatalf("please set student_id and password in env")
	}
}

func main() {
	clinet := resty.New()
	cookies := NewCookies()

	clinet.SetRedirectPolicy(resty.RedirectPolicyFunc(func(r1 *http.Request, _ []*http.Request) error {
		log.Println("redirect to: ", r1.URL.String())
		if len(r1.Response.Cookies()) != 0 {
			cookies.Set(JSESSIONID, r1.Response.Cookies())
		}
		return nil
	}))

	resp, err := clinet.R().Get(IndexUrl)
	if err != nil {
		log.Fatalf("get session request failed. %v", err)
	}
	// set login SESSION cookie.
	cookies.Set(SESSION, resp.Cookies())

	session, err := cookies.Get(SESSION)
	if err != nil {
		log.Fatalf("get session from cookie failed. %v", err)
	}
	// log.Printf("[cookie] session: %v", session)

	loginForm := LoginForm{
		Username:    stu.StudentID,
		Password:    stu.Password,
		CurrentMenu: "1",
		Execution:   "e1s1",
		EventId:     "submit",
		Geolocation: "",
	}

	loginFormStr, err := convertStruct2RawReqStr(loginForm)
	if err != nil {
		log.Fatalf("convert loginform struct to raw string failed. %v", err)
	}

	resp, err = clinet.R().SetCookie(&http.Cookie{
		Name:  SESSION,
		Value: session,
	}).
		SetHeader("User-Agent", UserAgent).
		SetHeader("Content-Type", ContentType).
		SetHeader("Referer", RefererLogin).
		SetBody(loginFormStr).
		Post(LoginUrl)

	if err != nil {
		log.Fatalf("login failed. %v", err)
	}

	// set TGC cookie.
	cookies.Set(TGC, resp.Cookies())

	_, err = clinet.R().Get(JUrl)
	if err != nil {
		log.Fatalf("get jsessionid request failed. %v", err)
	}

	jsessionid, err := cookies.Get(JSESSIONID)
	if err != nil {
		log.Fatalf("get jsessionid from cookie failed. %v", err)
	}

	// log.Println("[cookie] jsessionid: ", jsessionid)

	resp, err = clinet.R().
		SetCookie(&http.Cookie{
			Name:  JSESSIONID,
			Value: jsessionid,
		}).
		SetHeader("User-Agent", UserAgent).
		SetHeader("Content-Type", ContentType).
		SetHeader("Referer", RefererLogin).
		Get(SuffixUrl)
	if err != nil {
		log.Fatalf("get suffix request failed. %v", err)
	}

	// log.Printf("[suffix resp] %v", resp.String())
	sign, timestamp, ok := reqSuffix(resp.String())
	if !ok {
		log.Fatalf("extract sign and timestamp failed.")
	}

	resp, err = clinet.R().SetCookie(&http.Cookie{
		Name:  JSESSIONID,
		Value: jsessionid,
	}).
		SetHeader("User-Agent", UserAgent).
		SetHeader("Content-Type", ContentType).
		Post(UserInfoUrl)

	if err != nil {
		log.Fatalf("get userinfo req failed. %v", err)
	}

	// log.Printf("[userinfo resp] %v", resp.String())
	name, ok := stuName(resp.String())
	if !ok {
		log.Fatalf("extract stu name failed. %v", err)
	}
	stu.Name = name

	reportForm := ReportForm{
		Hsjc:        "1",
		Xasymt:      "1",
		ActionType:  "addRbxx",
		UserLoginId: stu.StudentID,
		Szcsbm:      "1",
		Bdzt:        "1",
		Szcsmc:      "在学校",
		Sfyzz:       "0",
		Sfqz:        "0",
		Tbly:        "sso",
		Qtqksm:      "",
		Ycqksm:      "",
		UserType:    "2",
		Username:    stu.Name,
	}

	reportFormStr, err := convertStruct2RawReqStr(reportForm)
	if err != nil {
		log.Fatalf("convert reportform struct to raw string failed. %v", err)
	}

	resp, err = clinet.R().SetCookie(&http.Cookie{
		Name:  JSESSIONID,
		Value: jsessionid,
	}).
		SetHeader("User-Agent", UserAgent).
		SetHeader("Content-Type", ContentType).
		SetHeader("Referer", RefererReport).
		SetBody(reportFormStr).
		Post(fmt.Sprintf(ReportUrl, sign, timestamp))

	if err != nil {
		log.Fatalf("report request failed. %v", err)
	}

	res := &ReportResponse{}
	err = json.Unmarshal(resp.Body(), res)
	if err != nil {
		log.Fatalf("unmarshal response failed. %v", err)
	}

	_, month, day := time.Now().Add(time.Hour * 8).Date()
	if res.Status != "1" {
		log.Printf("today: [%v:%v] report failed. please manually report.", month, day)
		return
	}

	log.Printf("today: [%v.%v] report success.", month, day)
}
