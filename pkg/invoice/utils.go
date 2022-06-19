package invoice

import (
	"invoicer/config"
	"invoicer/pkg/common"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func DefaultRequestOpts() common.RequestOpts {
	config := common.Configuration()
	return common.RequestOpts{
		Credentials: "include",
		Headers: []common.RequestOptsHeader{
			{Key: "Content-Type", Value: "application/x-www-form-urlencoded;charset=UTF-8"},
			{Key: "Referrer", Value: config.Env.BASEURL + "/intragiris.html"},
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

func GetToken() string {
	assoscmd := "login"
	cfg := config.GetConfig()
	var username, password string
	username, password = cfg.Credentials.USERNAME, cfg.Credentials.PASSWORD
	if config.CURRENT_ENV == "production" {
		assoscmd = "anologin"
	}
	body_text := "assoscmd=" + assoscmd + "&rtype=json&userid=" + username + "&sifre2=" + password + "&parola=1&"
	req := GetRequest("/earsiv-services/assos-login", AsReadCloser(body_text))
	res := RunRequest(*req)
	return res
}

func Logout(token string) string {
	env := "logout"
	if common.Configuration().Env.PRODUCTION {
		env = "anologin"
	}
	body := "assoscmd=" + env + "&rtype=json&token=" + token + "&"
	req := GetRequest("/earsiv-services/assos-login", AsReadCloser(body))
	res := RunRequest(*req)
	return res
}

func RunCommand(token string, command string, pageName string, data string) string {
	body_text := "cmd=" + command + "&callid=" + common.GetUUID() + "&pageName=" + pageName + "&token=" + token + "&jp=" + url.QueryEscape(data)
	req := GetRequest("/earsiv-services/dispatch", AsReadCloser(body_text))
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

func AsReadCloser(data string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(data))
}

func GetRequest(url string, body io.ReadCloser) *http.Request {
	defaults := DefaultRequestOpts()
	req, err := http.NewRequest(defaults.Method, common.Configuration().Env.BASEURL+url, body)
	if err != nil {
		panic(err)
	}
	for _, header := range defaults.Headers {
		req.Header.Set(header.Key, header.Value)
	}
	req.Method = defaults.Method
	return req
}
