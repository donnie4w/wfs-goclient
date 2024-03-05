package main

import (
	"os"
	"testing"

	"github.com/donnie4w/simplelog/logging"
	wfsclient "github.com/donnie4w/wfs-goclient/client"
	"github.com/donnie4w/wfs-goclient/stub"
)

func Test_append(t *testing.T) {
	if client, err := newclient(); err == nil {
		bs, _ := os.ReadFile("1.jpg")
		wf := &stub.WfsFile{Data: bs, Name: "test/go/1.jpg"}
		if ack, err := client.Append(wf); err == nil {
			if !ack.Ok {
				logging.Error(ack.Error.GetCode(), " >>", ack.Error.GetInfo())
			} else {
				logging.Debug("ok")
			}
		}
	}
}

func Test_get(t *testing.T) {
	if client, err := newclient(); err == nil {
		if wd, err := client.Get("test/go/1.jpg"); err == nil {
			logging.Debug(len(wd.Data))
		}
	}
}

func Test_del(t *testing.T) {
	if client, err := newclient(); err == nil {
		if ack, err := client.Delete("test/go/1.jpg"); err == nil {
			if !ack.Ok {
				logging.Error(ack.Error.GetCode(), " >>", ack.Error.GetInfo())
			} else {
				logging.Debug("ok")
			}
		}
	}
}

func newclient() (client *wfsclient.Client, err error) {
	client, err = wfsclient.NewConnect(false, "127.0.0.1", 6802, "admin", "123")
	return
}
