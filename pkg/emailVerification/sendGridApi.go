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
	//TODO remove email to be env variable
	//TODO rename variables
	//TODO add comments
	from := "ryasser@egirna.net"
	code:=CodeGenerator()
	//TODO template id could be deleted
	jsonText:=fmt.Sprintf("{\n  \"from\":{\"email\":\"%s\"},\n  \"personalizations\":[\n    {\n      \"to\":[\n        {\n          \"email\":\"%s\"\n        }\n      ],\n      \"dynamic_template_data\":{\n        \"code\": \"%s\"\n      }\n    }\n  ],\n  \"template_id\":\"d-de1aefebe42f43939bac714a95b8779e\"\n}\n",from,to,code)
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