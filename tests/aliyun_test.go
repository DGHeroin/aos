package tests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/DGHeroin/aos"
	_ "github.com/DGHeroin/aos/aliyun"
)

func TestAliyun(t *testing.T) {
	if !*aliyunTest { return }
	cli, err := aos.New("aliyun", *aliyunEndpoint, *aliyunArgs)
	if err != nil {
		fmt.Println(err)
		return
	}
	var (
		saveName = "saved/helloworld.txt"
	)
	reader := bytes.NewReader([]byte("helloworld!"))
	// 写入
	err = cli.Put(saveName, reader)
	if checkErr(err) {
		return
	}

	// 读取
	resp, err := cli.Get(saveName)
	if checkErr(err) {
		return
	}
	data, err := ioutil.ReadAll(resp)
	if checkErr(err) {
		return
	}
	fmt.Printf("%s\n", data)
	err = resp.Close()
	if checkErr(err) {
		return
	}

	// 删除
	err = cli.Delete(saveName)
	if checkErr(err) { return }
}
