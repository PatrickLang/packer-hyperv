package main

import (
	"github.com/MSOpenTech/packer-hyperv/packer/builder/hyperv/iso"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(iso.Builder))
	server.Serve()
}
