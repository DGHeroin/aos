package aos

import (
	"fmt"
	"io"
)

var(
	NewVendorFunc func(name, endpoint, args string) (Client, error)
	errVendorNotSupport  = fmt.Errorf("vendor not import")
)

type Client interface {
	Put(name string, reader io.Reader) error
	Get(name string) (io.ReadCloser, error)
	Delete(name string) error
}

func New(name, uri, args string) (Client, error) {
	if NewVendorFunc == nil {
		return nil, errVendorNotSupport
	}
	return NewVendorFunc(name, uri, args)
}