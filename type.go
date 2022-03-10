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
		CurrentMenu string `naive:"currentMenu"`
		Execution   string `naive:"execution"`
		EventId     string `naive:"_eventId"`
		Geolocation string `naive:"geolocation"`
	}

	ReportForm struct {
		Hsjc        string `naive:"hsjc"`
		Xasymt      string `naive:"xasymt"`
		ActionType  string `naive:"actionType"`
		UserLoginId string `naive:"userLoginId"`
		Szcsbm      string `naive:"szcsbm"`
		Bdzt        string `naive:"bdzt"`
		Szcsmc      string `naive:"szcsmc"`
		Sfyzz       string `naive:"sfyzz"`
		Sfqz        string `naive:"sfqz"`
		Tbly        string `naive:"tbly"`
		Qtqksm      string `naive:"qtqksm"`
		Ycqksm      string `naive:"ycqksm"`
		UserType    string `naive:"userType"`
		Username    string `naive:"userName"`
	}

	ReportResponse struct {
		Status string `json:"state"`
	}
)
