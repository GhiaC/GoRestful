package Controler

import (
	"../Models/Struct"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func getToken() string {
	url := "http://RestfulSms.com/api/Token"
	var jsonData []byte
	jsonData, err := json.Marshal(Struct.GetSMSTokenRequest{
		SecretKey:  Configuration.SecretKey,
		UserApiKey: Configuration.UserApiKey,
	})
	if err != nil {
		log.Println(err)
	}
	var jsonStr = []byte(string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	var pro Struct.GetTokenResponse
	json.Unmarshal([]byte(buf.Bytes()), &pro)
	token := pro.TokenKey
	return string(token)
}

func SendSms(getter string, id int, text string, answer bool) {
	token := getToken()
	url := "http://RestfulSms.com/api/MessageSend"

	message := make([]string, 1)
	if answer {
		message[0] = "#پاسخ_پیام" + "\n" + text
	} else {
		message[0] = "#سوال" + "\n" + "#کاربر" + string(id) + "\n" + text
	}
	phonenumbers := make([]string, 1)
	phonenumbers[0] = getter

	var jsonData []byte
	jsonData, err := json.Marshal(Struct.SendSMSMessageRequest{
		Messages:                 message,
		MobileNumbers:            phonenumbers,
		CanContinueInCaseOfError: "false",
		LineNumber:               GetLineNumber(),
		SendDateTime:             "",
	})
	if err != nil {
		log.Println(err)
	}
	var jsonStr = []byte(string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("x-sms-ir-secure-token", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}
