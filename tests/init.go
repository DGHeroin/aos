package tests

import (
	"flag"
	"fmt"
)

var(
	localTest     = flag.Bool("localTest", false, "enable local test")
	localEndpoint = flag.String("localEP", "", "local endpoint")
	localArgs     = flag.String("localArgs", "", "local args")

	qcloudTest     = flag.Bool("qcloudTest", false, "enable qcloud test")
	qcloudEndpoint = flag.String("qcloudEP", "", "qcloud endpoint")
	qcloudArgs     = flag.String("qcloudArgs", "", "qcloud args")

	aliyunTest     = flag.Bool("aliyunTest", false, "enable aliyun test")
	aliyunEndpoint = flag.String("aliyunEP", "", "aliyun endpoint")
	aliyunArgs     = flag.String("aliyunArgs", "", "aliyun args")
)

func init()  {
	flag.Parse()
}

func checkErr(err error) bool {
	if err == nil {
		return false
	}
	fmt.Println(err)
	return true
}
