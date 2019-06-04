package qcloud

import (
	"context"
	"github.com/DGHeroin/aos"
	"github.com/DGHeroin/aos/registry"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
)

func init()  {
	// register qcloud
	registry.RegisterVendor("qcloud", newQCloudClient)
}

type qcloudClient struct {
	aos.Client
	client *cos.Client
}

func newQCloudClient(endpoint string, uriArgs string) (aos.Client, error) {
	client := &qcloudClient{}
	u, err := url.Parse(endpoint)
	if err != nil { return nil, err}
	b := &cos.BaseURL{BucketURL:u}
	// parse key
	values, err := url.ParseQuery(uriArgs)
	if err != nil { return nil, err}
	id := values.Get("id")
	key := values.Get("key")

	client.client = cos.NewClient(b, &http.Client{
		Transport:&cos.AuthorizationTransport{
			SecretID:id,
			SecretKey:key,
		},
	})
	return client, nil
}

func (this* qcloudClient) Put(name string, reader io.Reader) error {
	_, err := this.client.Object.Put(context.Background(), name, reader, nil)
	if err != nil {
		return err
	}
	return nil
}

func (this* qcloudClient) Get(name string) (io.ReadCloser, error) {
	resp, err := this.client.Object.Get(context.Background(), name, nil)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return resp.Body, err
	}
	return nil, nil
}

func (this *qcloudClient) Delete(name string) error {
	_, err := this.client.Object.Delete(context.Background(), name)
	return err
}