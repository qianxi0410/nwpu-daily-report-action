package main

type (
	UserInfo struct {
		Name      string
		StudentID string
		Password  string
	}

	LoginForm struct {
		Username    string `naive:"username"`
		Password    string `naive:"password"`
		CurrentMenu string `naive:"currentMenu" default:"1"`
		Execution   string `naive:"execution"`
		EventId     string `naive:"_eventId" default:"submit"`
		Geolocation string `naive:"geolocation"`
	}

	ReportForm struct {
		Hsjc        string `naive:"hsjc" default:"1"`
		Xasymt      string `naive:"xasymt" default:"1"`
		ActionType  string `naive:"actionType" default:"addRbxx"`
		UserLoginId string `naive:"userLoginId"`
		Szcsbm      string `naive:"szcsbm" default:"1"`
		Bdzt        string `naive:"bdzt" default:"1"`
		Szcsmc      string `naive:"szcsmc" default:"在学校"`
		Szcsmc1     string `naive:"szcsmc1" default:"在学校"`
		Sfyzz       string `naive:"sfyzz" default:"0"`
		Sfqz        string `naive:"sfqz" default:"0"`
		Tbly        string `naive:"tbly" default:"pc"`
		Qtqksm      string `naive:"qtqksm"`
		Ycqksm      string `naive:"ycqksm"`
		Sfxn        string `naive:"sfxn" default:"0"`
		Sfdw        string `naive:"sfdw" default:"0"`
		Longlat     string `naive:"longlat"`
		UserType    string `naive:"userType" default:"2"`
		Username    string `naive:"userName"`
	}

	ReportResponse struct {
		Status string `json:"state"`
	}
)
