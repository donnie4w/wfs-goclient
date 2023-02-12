/**
 * Copyright 2017 wfs-goclient Author. All Rights Reserved.
 * email: donnie4w@gmail.com
 */
package client

import (
	"context"
	"wfs-goclient/protocol"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/donnie4w/simplelog/logging"
)

type WfsClient struct {
	ServerUrl string
}

func (this *WfsClient) PostFile(bs []byte, name, fileType string) (err error) {
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err := thrift.NewTHttpPostClient(this.ServerUrl)
	if err != nil {
		logging.Error("err:", err.Error())
	}
	client := protocol.NewIWfsClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		logging.Error("err:", err.Error())
	}
	defer transport.Close()
	wf := protocol.NewWfsFile()
	wf.FileBody = bs
	wf.Name = &name
	wf.FileType = &fileType
	_, er := client.WfsPost(context.Background(), wf)
	if er != nil {
		err = er
		logging.Debug("err:", err.Error())
	}
	return
}

func (this *WfsClient) GetFile(name string) (bs []byte, err error) {
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err := thrift.NewTHttpPostClient(this.ServerUrl)
	if err != nil {
		logging.Error("err:", err.Error())
	}
	client := protocol.NewIWfsClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		logging.Error("err:", err.Error())
	}
	defer transport.Close()
	wf, er := client.WfsRead(context.Background(), name)
	if er == nil {
		bs = wf.GetFileBody()
		logging.Debug("len(bs):", len(bs))
	}
	err = er
	return
}

func (this *WfsClient) DelFile(name string) (err error) {
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err := thrift.NewTHttpPostClient(this.ServerUrl)
	if err != nil {
		logging.Error("err:", err.Error())
	}
	client := protocol.NewIWfsClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		logging.Error("err:", err.Error())
	}
	defer transport.Close()
	ack, er := client.WfsDel(context.Background(), name)
	if er == nil {
		logging.Debug("ack:", ack)
	}
	err = er
	return
}
