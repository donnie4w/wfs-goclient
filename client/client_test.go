/**
 * Copyright 2017 wfs-goclient Author. All Rights Reserved.
 * email: donnie4w@gmail.com
 */
package client

import (
	"io/ioutil"
	"testing"

	"github.com/donnie4w/simplelog/logging"
)

func Test_post(t *testing.T) {
	client, err := NewWfsClient("http://127.0.0.1:3434/thrift")
	defer client.Close()
	bs, _ := ioutil.ReadFile(`1.jpg`)
	err = client.PostFile(bs, "222", "")
	logging.Debug(err)

}

func Test_del(t *testing.T) {
	client, err := NewWfsClient("http://127.0.0.1:3434/thrift")
	defer client.Close()
	err = client.DelFile("22")
	logging.Debug(err)
}

func Test_read(t *testing.T) {
	client, err := NewWfsClient("http://127.0.0.1:3434/thrift")
	defer client.Close()
	var bs []byte
	bs, err = client.GetFile("222?imageView2/0/w/100")
	if err == nil {
		logging.Debug(len(bs))
	} else {
		logging.Error("err:", err.Error())
	}
}
