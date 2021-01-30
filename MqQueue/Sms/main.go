package main

import (
	"elearn100/MqQueue/Mq"
	"elearn100/MqQueue/Sms/SmsServices"
	"encoding/json"
	"fmt"
)

func main() {
	Mq.Consumer("", "smsinfo", sendSmsToUser)
}

// @Title 发送短信
// @Param s string 内容
func sendSmsToUser(s string) {
	type InfoSms struct {
		Msg string
		Tel string
	}
	var data InfoSms
	err := json.Unmarshal([]byte(s), &data)
	if err == nil {
		SmsServices.SendSms(data.Tel, data.Msg)
	} else {
		fmt.Println(err, s)
	}
}
