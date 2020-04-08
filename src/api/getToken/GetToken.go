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
	apiUrl := "http://identityserver-loyalti.azurewebsites.net/connect/token"
	//resource := "/connect/token"
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
	//u.Path = resource
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
	r.SetBasicAuth("roclient","secret")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//r.Header.Add("Authorization", "Basic cm9jbGllbnQ6c2VjcmV0")

	fmt.Println("melewati header")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error ini : ", resp.StatusCode)

		//bytesResp, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Errornya" ,resp)
		//fmt.Println(string(bytesResp))
		//fmt.Println(bytesResp)
		os.Exit(1)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error ini : ", err.Error())
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal([]byte(bodyBytes), &respon)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	fmt.Println("Respon Body : ", respon)
	os.Exit(1)
	return &respon
}
