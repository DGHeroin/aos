# aos is (是什么)
aos is abstract object storage <br>
aos是一个抽象的对象存储库

# support list (支持列表)
* qcloud
* aliyun
* local

# usage (使用方式)

```go
// import
import (
  "github.com/DGHeroin/aos"
	_ "github.com/DGHeroin/aos/qcloud"   // import the vender what you want to use
//  _ "github.com/DGHeroin/aos/aliyun"
//  _ "github.com/DGHeroin/aos/local"
)

func new() {
    cli, err := aos.New("qcloud", "xxx", "id=111&key=222")
}

func put() {
    reader := bytes.NewReader([]byte("helloworld!"))
    err = cli.Put("some-folder/filename", reader)
}

func get() {
    resp, err := cli.Get("some-folder/filename")
    data, err := ioutil.ReadAll(resp)
    resp.Close()
}

func delete() {
    err = cli.Delete("some-folder/filename")
}

```

# Args (参数)

## qcloud

| name   |      Desc      |  Example |
|----------|:-------------:|------:|
| id |  SecretID | xxxx |
| key |    SecretKey   |   xxx |

## aliyun

| name   |      Desc      |  Example |
|----------|:-------------:|------:|
| id |  SecretID | xxxx |
| key |    SecretKey   |   xxx |
| bucket |    BucketName   |   my-bucket |

## local

no args