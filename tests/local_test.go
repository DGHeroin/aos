package tests

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "testing"

    "github.com/DGHeroin/aos"
    _ "github.com/DGHeroin/aos/local"
)

func TestLocal(t *testing.T) {
    if !*localTest { return }
    cli, err := aos.New("local", *localEndpoint, *localArgs)
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
