package registry

import (
	"fmt"
	"github.com/DGHeroin/aos"
)

var (
	vendorFuncs  map[string] NewVendorFunc
)

type NewVendorFunc func(endpoint string, uriArgs string) (aos.Client, error)

func init()  {
	vendorFuncs = make(map[string] NewVendorFunc)
	aos.NewVendorFunc = newVendor
}

func RegisterVendor(name string, cb NewVendorFunc) {
	vendorFuncs[name] = cb
}

func newVendor(name string, endpoint string, args string) (aos.Client, error) {
	if cb, ok := vendorFuncs[name]; ok {
		return cb(endpoint, args)
	} else {
		return nil, fmt.Errorf("vendor(%s) not support", name)
	}
}