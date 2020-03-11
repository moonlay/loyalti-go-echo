package getToken

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func GetToken(username string, password string) *model.Response {

	respon := model.Response{}

	fmt.Println("masuk ke getTOken")
	apiUrl := "https://11.11.5.146:9443"
	resource := "/oauth2/token"
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Add("username", username)
	data.Add("password", password)
	data.Set("scope", "openid")
	fmt.Println("berhasil dikirim")

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
	u.Path = resource
	urlStr := u.String()
	fmt.Println("lewati u")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	//r, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	r, err := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", "Basic cm1SSF95MmxqVUZfR0ZFWlFlX0xha09fWWl3YTpEYng1VG82YjFDQ2dMd2R1MWdaRHkzdnJ6YjRh")

	fmt.Println("melewati header")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal([]byte(bodyBytes), &respon)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	fmt.Println("Respon Body : ", respon)
	return &respon
}
