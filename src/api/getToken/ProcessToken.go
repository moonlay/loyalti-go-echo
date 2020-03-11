package getToken

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"io/ioutil"
	"net/http"
	"os"
)

func ProcessToken (token string) *model.ResponseProcess {
	//accessToken := model.Token{}
	respons := model.ResponseProcess{}

	apiUrl := "https://11.11.5.146:9443/oauth2/userinfo?schema=openid"

	bearer :=  token

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
		fmt.Println("Bearer : ", bearer)
	req.Header.Add("Authorization", bearer)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport:tr}
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println("Error : ", err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal([]byte(bodyBytes), &respons)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	fmt.Println("Ini respon : ", bodyBytes)
	return &respons
}