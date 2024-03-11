// Copyright (c) , donnie <donnie4w@gmail.com>
// All rights reserved.

package client

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	thrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/donnie4w/simplelog/logging"
	. "github.com/donnie4w/wfs-goclient/stub"
)

var ConnectTimeout = 60 * time.Second
var SocketTimeout = 60 * time.Second
var _transportFactory = thrift.NewTBufferedTransportFactory(1 << 13)

type Client struct {
	Conn      *WfsIfaceClient
	transport thrift.TTransport
	hostPort  string
	tls       bool
	_auth     *WfsAuth
	_connid   int32
	mux       *sync.Mutex
	_pingNum  int32
}

func (t *Client) Link(hostPort, name, pwd string, TLS bool) (err error) {
	defer _recover()
	t.tls, t.hostPort, t._auth = TLS, hostPort, &WfsAuth{Name: &name, Pwd: &pwd}
	tcf := &thrift.TConfiguration{ConnectTimeout: ConnectTimeout, SocketTimeout: SocketTimeout}
	var transport thrift.TTransport
	if t.tls {
		tcf.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		transport = thrift.NewTSSLSocketConf(hostPort, tcf)
	} else {
		transport = thrift.NewTSocketConf(hostPort, tcf)
	}
	var useTransport thrift.TTransport
	if err == nil && transport != nil {
		if useTransport, err = _transportFactory.GetTransport(transport); err == nil {
			if err = useTransport.Open(); err == nil {
				t.hostPort = hostPort
				t.transport = useTransport
				t.Conn = NewWfsIfaceClientFactory(useTransport, thrift.NewTCompactProtocolFactoryConf(&thrift.TConfiguration{}))
				<-time.After(time.Second)
				t._pingNum = 0
				err = t.auth()
			}
		}
	}
	if err != nil {
		logging.Error("client to [", hostPort, "] Error:", err)
	}
	return
}

func (t *Client) auth() (_err error) {
	if ack, err := t.auth0(t._auth); err == nil {
		if !ack.Ok {
			_err = errors.New(ack.Error.GetInfo())
		}
	} else {
		_err = err
	}
	return
}

func (t *Client) timer(i int32) {
	ticker := time.NewTicker(3 * time.Second)
	for i == t._connid {
		<-ticker.C
		func() {
			defer _recover()
			if t._pingNum > 5 && i == t._connid {
				t.reLink()
				return
			}
			atomic.AddInt32(&t._pingNum, 1)
			if i, err := t.ping(); err == nil && i == 1 {
				atomic.AddInt32(&t._pingNum, -1)
			}
		}()
	}
}

func (t *Client) Close() (err error) {
	atomic.AddInt32(&t._connid, 1)
	if t.transport != nil {
		err = t.transport.Close()
	}
	return
}

func (t *Client) reLink() error {
	logging.Warn("reconnect")
	if t.transport != nil {
		t.transport.Close()
	}
	return t.Link(t.hostPort, t._auth.GetName(), t._auth.GetPwd(), t.tls)
}

func NewConnect(tls bool, host string, port int, name, pwd string) (client *Client, err error) {
	client = &Client{mux: &sync.Mutex{}}
	i := atomic.AddInt32(&client._connid, 1)
	if err = client.Link(fmt.Sprint(host, ":", port), name, pwd, tls); err == nil {
		go client.timer(i)
	}
	return
}

func (t *Client) ping() (_r int8, _err error) {
	defer _recover()
	defer t.mux.Unlock()
	t.mux.Lock()
	return t.Conn.Ping(context.TODO())
}

func (t *Client) auth0(wa *WfsAuth) (_r *WfsAck, _err error) {
	defer _recover()
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.Conn.Auth(context.TODO(), wa)
}

func (t *Client) Append(wf *WfsFile) (_r *WfsAck, _err error) {
	defer _recover()
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.Conn.Append(context.TODO(), wf)
}

func (t *Client) Delete(path string) (_r *WfsAck, _err error) {
	defer _recover()
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.Conn.Delete(context.TODO(), path)
}

func (t *Client) Rename(path, newPath string) (_r *WfsAck, _err error) {
	defer _recover()
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.Conn.Rename(context.TODO(), path, newPath)
}

// Parameters:
//   - Name
func (t *Client) Get(path string) (_r *WfsData, _err error) {
	defer _recover()
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.Conn.Get(context.TODO(), path)
}

func _recover() {
	if err := recover(); err != nil {
		logging.Error(err)
	}
}
