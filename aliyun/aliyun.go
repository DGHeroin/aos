package aliyun

import (
	"github.com/DGHeroin/aos"
	"github.com/DGHeroin/aos/registry"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"net/url"
)

func init()  {
	registry.RegisterVendor("aliyun", newAliYunClient)
}

type aliyunClient struct {
	aos.Client
	client *oss.Client
	bucketName string
}

func newAliYunClient(endpoint string, uriArgs string) (aos.Client, error) {
	client := &aliyunClient{}
	// parse key
	values, err := url.ParseQuery(uriArgs)
	if err != nil { return nil, err}
	id := values.Get("id")
	key := values.Get("key")
	client.bucketName = values.Get("bucket")

	client.client, err = oss.New(endpoint, id, key)
	return client, err
}

func (this* aliyunClient) Put(name string, reader io.Reader) error {
	b, err := this.client.Bucket(this.bucketName)
	if err != nil {
		return err
	}
	return b.PutObject(name, reader)
}

func (this* aliyunClient) Get(name string) (io.ReadCloser, error) {
	b, err := this.client.Bucket(this.bucketName)
	if err != nil {
		return nil, err
	}
	return  b.GetObject(name)
}

func (this *aliyunClient) Delete(name string) error {
	b, err := this.client.Bucket(this.bucketName)
	if err != nil {
		return err
	}
	return b.DeleteObject(name)
}