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
	transport thrift.TTransport
	client    *protocol.IWfsClient
}

func NewWfsClient(_serverUrl string) (wfsclient *WfsClient, err error) {
	wfsclient = new(WfsClient)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	wfsclient.transport, err = thrift.NewTHttpPostClient(_serverUrl)
	if err != nil {
		logging.Error("err:", err.Error())
	}
	wfsclient.client = protocol.NewIWfsClientFactory(wfsclient.transport, protocolFactory)
	if err = wfsclient.transport.Open(); err != nil {
		logging.Error("err:", err.Error())
	}
	return
}

func (this *WfsClient) Close() error {
	return this.transport.Close()
}

func (this *WfsClient) PostFile(bs []byte, name, fileType string) (err error) {
	wf := protocol.NewWfsFile()
	wf.FileBody, wf.Name, wf.FileType = bs, &name, &fileType
	_, err = this.client.WfsPost(context.Background(), wf)
	if err != nil {
		logging.Debug("err:", err.Error())
	}
	return
}

func (this *WfsClient) GetFile(name string) (bs []byte, err error) {
	wf, er := this.client.WfsRead(context.Background(), name)
	if er == nil {
		bs = wf.GetFileBody()
	}
	return bs, er
}

func (this *WfsClient) DelFile(name string) (err error) {
	ack, er := this.client.WfsDel(context.Background(), name)
	if er == nil {
		logging.Info("DelFile[", name, "] and Ack:", ack.GetStatus())
	}
	return er
}
