package thirdcall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
	发送短信
 */
func Send(phone string) (code int) {
	client := &http.Client{}

	data := make(map[string]interface{})
	data["apikey"] = "N56568a4ae"
	data["secret"] = "56568410a24499c"
	data["sign_id"] = "130248"
	data["mobile"] = phone
	data["content"] = "您好，您的验证码是：6666"



	bytesData ,_ := json.Marshal(&data)

	fmt.Println(data)


	req ,_ := http.NewRequest("POST","https://api.4321.sh/sms/send",bytes.NewReader(bytesData))
	resp ,_ := client.Do(req)

	body ,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))




	return 1
}