package invoice

import (
	"invoicer/pkg/common"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetToken(username string, password string) string {
	client := &http.Client{}

	loginType := "anologin" //anologin:prod or login:test
	if !common.Configuration().Env.PRODUCTION {
		loginType = "login"
	}

	r := ioutil.NopCloser(strings.NewReader("assoscmd=" + loginType + "&rtype=json&userid=" + username + "&sifre=" + password + "&sifre2=" + password + "&parola=1&"))

	req := GetRequest("/earsiv-services/assos-login/", common.Configuration().Env.TESTURL+"/earsiv-services/intragiris.html", r, "POST")
	res, err := client.Do(req)
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

func Logout(token string) string {
	if token == "" {
		panic("token is empty")
	}
	env := "logout"
	if common.Configuration().Env.PRODUCTION {
		env = "anologin"
	}
	body := ioutil.NopCloser((strings.NewReader("assoscmd=" + env + "&rtype=json&token=" + token + "&")))
	req := GetRequest("/earsiv-services/assos-login", common.Configuration().Env.TESTURL+"/earsiv-services/intragiris.html", body, "GET")
	client := &http.Client{}
	res, err := client.Do(req)
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

///TODO: ...
func RunCommand(token string, command string) {

}

func GetRequest(url string, referrer string, body io.ReadCloser, method string) *http.Request {

	req, err := http.NewRequest(method, common.Configuration().Env.TESTURL+url, body)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Referrer", referrer)
	req.Header.Add("ReferrerPolicy", "no-referrer-when-downgrade")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "tr,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("Mode", "cors")

	return req
}
