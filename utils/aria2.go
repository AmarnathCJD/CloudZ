package utils

import (
	"fmt"

	"github.com/siku2/arigo"
)

const AriaCmd = "aria2c --enable-rpc --rpc-listen-all"

func init() {
	if _, err := RunCommand(AriaCmd); err != nil {
		Logger.Error(fmt.Sprintf("Error starting aria2: %s", err.Error()))
	}
	cli, err := Aria2()
	if err != nil {
		Logger.Error(fmt.Sprintf("Error connecting to aria2: %s", err.Error()))
		client = &arigo.Client{}
	} else {
		client = cli
	}
}

func Aria2() (*arigo.Client, error) {
	aria2, err := arigo.Dial("ws://localhost:6800/jsonrpc", "")
	if err != nil {
		return nil, err
	}
	return &aria2, nil
}

var client *arigo.Client

func AddUri(uri string) (*arigo.Status, error) {
	gid, err := client.Download(arigo.URIs(uri), nil)
	if err != nil {
		return nil, err
	}
	return &gid, nil
}
