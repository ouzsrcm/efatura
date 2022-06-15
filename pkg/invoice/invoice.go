package invoice

import (
	"fmt"
	"invoicer/pkg/common"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func Init() {
	fmt.Println("started...")
}

func GetToken(username string, password string) string {
	client := &http.Client{}

	// r, w := io.Pipe()
	// data := map[string]string{"username": username, "password": password}
	// go func() {
	// 	w.CloseWithError(json.NewEncoder(w).Encode(data))
	// }()
	// defer r.Close()

	loginType := "anologin" //anologin:prod or login:test
	if !common.Configuration().Env.PRODUCTION {
		loginType = "login"
	}

	r := ioutil.NopCloser(strings.NewReader("assoscmd=" + loginType + "&rtype=json&userid=" + username + "&sifre=" + password + "&sifre2=" + password + "&parola=1&"))

	req := getRequest("/earsiv-services/assos-login/", common.Configuration().Env.TESTURL+"/earsiv-services/intragiris.html", r, "POST")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func getRequest(url string, referrer string, body io.ReadCloser, method string) *http.Request {

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
