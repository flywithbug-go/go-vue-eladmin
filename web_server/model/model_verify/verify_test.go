package model_verify

import (
	"fmt"
	"testing"
	"vue-admin/web_server/core/mongo"
)

func TestGenerVerifyData(t *testing.T) {
	mongo.DialMgo("127.0.0.1:27017")

	source := "addadad232323a"
	vCode, err := GeneralVerifyData(source)
	fmt.Println(vCode)

	if err != nil {
		panic(err)
	}

	fmt.Println("result:", CheckVerify(source, vCode))
	//time.Sleep(time.Second * 2)
	fmt.Println("result:", CheckVerify(source, vCode))

}
