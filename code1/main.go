package main

import (
	"github.com/go-kit/kit/transport/http"
	"gokit-demo/code1/endpoint"
	"gokit-demo/code1/service"
	"gokit-demo/code1/transport"
	"log"
	syshttp "net/http"
)

func main() {
	s := service.UserService{}                       //业务接口服务
	getName := endpoint.MakeServerEndPointGetName(s) //使用myEndPoints创建业务服务
	//使用 kit 创建 handler
	// 固定格式
	// 传入 业务服务 以及 定义的 加密解密方法
	server := http.NewServer(getName, transport.GetNameDecodeRequest, transport.GetNameEncodeResponse)
	log.Println(syshttp.ListenAndServe(":8080", server)) //监听服务
}
