package main

import "github.com/PinkPinkPigg/dora/services/task/impl"

func main() {
	imp := impl.GetServiceImp()
	//处理连接关闭的问题
	defer imp.Close()
	//	todo:注册imp到rpc服务器

}
