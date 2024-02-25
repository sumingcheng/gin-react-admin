package util

import (
	"backend/util/sms"
	"fmt"
)

func Sms() {
	params := make(map[string]interface{})
	params["code"] = 1517

	client := sms.NewClient()
	client.SetAppId("xxxxx")
	client.SetSecretKey("xxxxxx")

	request := sms.NewRequest()
	request.SetMethod("sms.message.send")
	request.SetBizContent(sms.TemplateMessage{
		Mobile:     []string{"13800138000"},
		Type:       0,
		Sign:       "闪速码",
		TemplateId: "ST_2019043000000001",
		SendTime:   "",
		Params:     params,
	})

	buf, err := client.Execute(request)
	fmt.Println(buf, err)

}
