package main

const (
	UserAgent     = "Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0"
	ContentType   = "application/x-www-form-urlencoded"
	RefererLogin  = "https://uis.nwpu.edu.cn/cas/login"
	RefererReport = "https://yqtb.nwpu.edu.cn/wx/ry/jrsb_xs.jsp"
	IndexUrl      = "https://uis.nwpu.edu.cn/cas/login"
	LoginUrl      = "https://uis.nwpu.edu.cn/cas/login?service=http://yqtb.nwpu.edu.cn//sso/login.jsp?targetUrl=base64aHR0cDovL3lxdGIubndwdS5lZHUuY24vL3d4L3hnL3l6LW1vYmlsZS9pbmRleC5qc3A="
	// just need to get jsessionid from the url.
	JUrl = "http://yqtb.nwpu.edu.cn//sso/login.jsp"
	// get req req suffix from the url.
	SuffixUrl   = "http://yqtb.nwpu.edu.cn/wx/ry/jrsb_xs.jsp"
	ReportUrl   = "https://yqtb.nwpu.edu.cn/wx/ry/ry_util.jsp?sign=%s&timeStamp=%s"
	UserInfoUrl = "http://yqtb.nwpu.edu.cn/wx/xg/yz-mobile/userInfo.jsp"
)
