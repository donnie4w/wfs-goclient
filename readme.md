# wfs-goclient

###### WFS的golang实现访问接口

------------

创建wfsclient实例对象

	func newclient() (client *wfsclient.Client, err error) {
		client, err = wfsclient.NewConnect(false, "127.0.0.1",6802, "admin", "123")
		return
	}

参数说明：wfsclient.NewConnect(false, "127.0.0.1", 6802, "admin", "123")
1. 第一个参数：是否TLS
2. 第二个参数：wfs thrift 服务ip或域名
3. 第三个参数：端口
4. 第四个参数：后台用户名
5. 第五个参数：后台密码

上传文件

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

拉取文件

	func Test_get(t *testing.T) {
		if client, err := newclient(); err == nil {
			if wd, err := client.Get("test/go/1.jpg"); err == nil {
				logging.Debug(len(wd.Data))
			}
		}
	}

删除文件

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

重命名

	func Test_rename(t *testing.T) {
		if client, err := newclient(); err == nil {
			if ack, err := client.Rename("test/111.jpeg","c9e8efcfcd.jpeg"); err == nil {
				if !ack.Ok {
					logging.Error(ack.Error.GetCode(), " >>", ack.Error.GetInfo())
				} else {
					logging.Debug("ok")
				}
			}
		}
	}



