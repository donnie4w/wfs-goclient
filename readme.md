wfs提供了thrift访问接口
  WfsPost()    //上传文件
  wfsRead()    //拉取文件
  wfsDel       //删除文件


   	bs, _ := ioutil.ReadFile(`1.jpg`)
	client := &WfsClient{"http://127.0.0.1:3434/thrift"}
	client.PostFile(bs, "1.jpg", "")
	
	