# wfs-goclient
**WFS的golang实现访问接口**


------------
上传文件

    获取wfs客户端实例：
	client, err := NewWfsClient("http://127.0.0.1:3434/thrift")
	defer client.Close()
	
    bs, _ := ioutil.ReadFile(`1.jpg`)
	//上传 文件bytes, 文件名
    err = client.PostFile(bs, "22", "")
	//相当于：curl -F "file=@1.jpg" "http://127.0.0.1:3434/u/22.jpg"
	//1.jpg 是本地文件，22.jpg是上传到服务自定义的文件名，
	//也可以：
	err = client.PostFile(bs, "fff/ggg/1.jpg", "")
	//访问则为：http://127.0.0.1:3434/r/fff/ggg/1.jpg
	
拉取 文件

	var bs []byte
	bs, err = client.GetFile("22.jpg")
	//相当于：http://127.0.0.1:3434/r/22.jpg
	bs, err = client.GetFile("fff/ggg/22.jpg")
	//相当于：http://127.0.0.1:3434/r/fff/ggg/22.jpg

删除文件

    err = client.DelFile("22.jpg")

