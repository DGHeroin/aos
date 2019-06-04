package local

import (
    "github.com/DGHeroin/aos"
    "github.com/DGHeroin/aos/registry"
    "io"
    "os"
    "path/filepath"
)

func init()  {
    registry.RegisterVendor("local", newLocalClient)
}

type localClient struct {
    aos.Client
    endpoint string
}

func newLocalClient(endpoint string, _ string) (aos.Client, error) {
    client := &localClient{}
    client.endpoint = endpoint
    return client, nil
}

func (this* localClient) Put(name string, reader io.Reader) error {
    name = filepath.Join(this.endpoint, name)
    base := filepath.Dir(name)
    if err := os.MkdirAll(base, os.ModePerm); err != nil {
        return err
    }
    f, err := os.Create(name)
    if err != nil { return err}

    _, err = io.Copy(f, reader)
    if err != nil { return err}
    return f.Close()
}

func (this* localClient) Get(name string) (io.ReadCloser, error) {
    name = filepath.Join(this.endpoint, name)
    return os.OpenFile(name, os.O_RDONLY, os.ModePerm)
}

func (this *localClient) Delete(name string) error {
    name = filepath.Join(this.endpoint, name)
    return os.Remove(name)
}