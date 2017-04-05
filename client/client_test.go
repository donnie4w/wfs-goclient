/**
 * Copyright 2017 wfs-goclient Author. All Rights Reserved.
 * email: donnie4w@gmail.com
 */
package client

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func _Test_post(t *testing.T) {
	bs, _ := ioutil.ReadFile(`C:\Users\dong\Desktop\wfs\1.jpg`)
	client := &WfsClient{"http://127.0.0.1:3434/thrift"}
	client.PostFile(bs, "11", "")
}

func _Test_del(t *testing.T) {
	client := &WfsClient{"http://127.0.0.1:3434/thrift"}
	client.DelFile("1122")
}

func Test_read(t *testing.T) {
	client := &WfsClient{"http://127.0.0.1:3434/thrift"}
	bs, err := client.GetFile("11?imageView2/0/w/100")
	if err == nil {
		fmt.Println(len(bs))
	} else {
		fmt.Println("err:", err.Error())
	}
}
