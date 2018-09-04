/*
    @Time : 2018/8/31 下午2:22 
    @Author : frozenChen
    @File : microServer
    @Software: dailyserver2
*/
package rpc

import (
	"dailyserver2/commons/glog"
	"dailyserver2/proto"
	"github.com/micro/go-micro"
)

func InitRpc() {
	go func() {
		microS := micro.NewService(micro.Name("module_user"))
		microS.Init()

		proto.RegisterUserServiceHandler(microS.Server(), new(UserHandler))
		err := microS.Run()
		if err != nil {
			glog.Sugar().Panicf("microServer run is error :%s", err)
		}
	}()
}