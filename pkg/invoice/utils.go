package invoice

import (
	"invoicer/pkg/common"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RequestOptsHeader struct {
	Key   string
	Value string
}

type RequestOpts struct {
	Credentials string              `json:"credentials"`
	Headers     []RequestOptsHeader `json:"headers"`
	Method      string              `json:"method"`
	Mode        string              `json:"mode"`
}

func DefaultRequestOpts() RequestOpts {
	config := common.Configuration()
	return RequestOpts{
		Credentials: "include",
		Headers: []RequestOptsHeader{
			{Key: "Content-Type", Value: "application/x-www-form-urlencoded;charset=UTF-8"},
			{Key: "Referrer", Value: config.Env.TESTURL + "/intragiris.html"},
			{Key: "ReferrerPolicy", Value: "no-referrer-when-downgrade"},
			{Key: "Accept", Value: "*/*"},
			{Key: "Accept-Language", Value: "tr,en-US;q=0.9,en;q=0.8"},
			{Key: "cache-control", Value: "no-cache"},
			{Key: "Connection", Value: "keep-alive"},
			{Key: "pragma", Value: "no-cache"},
			{Key: "x-requested-with", Value: "XMLHttpRequest"},
			{Key: "sec-fetch-mode", Value: "cors"},
			{Key: "sec-fetch-site", Value: "same-origin"},
			{Key: "Mode", Value: "cors"},
		},
		Method: "POST",
		Mode:   "cors",
	}
}

func GetToken(username string, password string) string {
	loginType := "anologin" //anologin:prod or login:test
	if !common.Configuration().Env.PRODUCTION {
		loginType = "login"
	}

	r := ioutil.NopCloser(strings.NewReader("assoscmd=" + loginType + "&rtype=json&userid=" + username + "&sifre=" + password + "&sifre2=" + password + "&parola=1&"))

	req := GetRequest("/earsiv-services/assos-login/", common.Configuration().Env.TESTURL+"/earsiv-services/intragiris.html", r)
	res := RunRequest(*req)
	return res
}

func Logout(token string) string {
	env := "logout"
	if common.Configuration().Env.PRODUCTION {
		env = "anologin"
	}
	body := ioutil.NopCloser((strings.NewReader("assoscmd=" + env + "&rtype=json&token=" + token + "&")))
	req := GetRequest("/earsiv-services/assos-login", common.Configuration().Env.TESTURL+"/earsiv-services/intragiris.html", body)
	res := RunRequest(*req)
	return res
}

func RunCommand(token string, command string, pageName string, data string) string {
	body_text := "cmd=" + command + "&callid=" + common.GetUUID() + "&pageName=" + pageName + "&token=" + token + "&jp=" + url.QueryEscape(data)
	req := GetRequest("/earsiv-services/dispatch", common.Configuration().Env.TESTURL+"/earsiv-services/login.jsp", ioutil.NopCloser(strings.NewReader(body_text)))
	res := RunRequest(*req)
	return res
}

func RunRequest(req http.Request) string {
	client := &http.Client{}
	res, err := client.Do(&req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func GetRequest(url string, referrer string, body io.ReadCloser) *http.Request {
	defaults := DefaultRequestOpts()
	req, err := http.NewRequest(defaults.Method, common.Configuration().Env.TESTURL+url, body)
	if err != nil {
		panic(err)
	}
	for _, header := range defaults.Headers {
		req.Header.Set(header.Key, header.Value)
	}
	req.Method = defaults.Method
	return req
}
