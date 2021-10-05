package emailVerification

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// SendEmail takes email that receives verification code,
//connects to sendgrid api and use a template to send code
func SendEmail(to string) (string,error){
	from := os.Getenv("EMAIL_SENDER")
	code:=CodeGenerator()

	template_id := os.Getenv("TEMPLATE_ID")
	jsonText:=fmt.Sprintf("{\n  \"from\":{\"email\":\"%s\"},\n  \"personalizations\":[\n    {\n      \"to\":[\n        {\n          \"email\":\"%s\"\n        }\n      ],\n      \"dynamic_template_data\":{\n        \"code\": \"%s\"\n      }\n    }\n  ],\n  \"template_id\":\"%s\"\n}\n",from,to,code,template_id)
	var jsonReq2 = []byte(jsonText)
	url :="https://api.sendgrid.com/v3/mail/send"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq2))
	req.Header.Set("Content-Type", "application/json")
	bearer:="Bearer " + os.Getenv("SENDGRID_API_KEY")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		fmt.Println("error while making request to sendgrid api")
		return "",err
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error while making request to sendgrid api")
		return "",err
	}
	defer resp.Body.Close()
	//_ code be replaced with response variable
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in response of sendgrid api")
		return "",err
	}
	return code,nil
}