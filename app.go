package main

import (
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"strings"
	"time"
)	

//go:generate cqcfg .
// cqp: 名称: Cool
// cqp: 版本: 1.0.0:1
// cqp: 作者: fucora
// cqp: 简介: 一个超棒的Go语言插件Demo，它会回复你的私聊消息~
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "me.cqp.fucora.cool" // TODO: 修改为这个插件的ID
	cqp.PrivateMsg = onPrivateMsg
	cqp.GroupMsg = onGroupMsg
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	cqp.SendPrivateMsg(fromQQ, msg) //复读机
	return 0
}

func onGroupMsg (subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32{
	if strings.Contains(msg, "jpg") {
		cqp.AddLog(cqp.Debug, "调试", "调试99：" + msg)
		time.Sleep(5*time.Second)
		cqp.SendGroupMsg(fromGroup, "哈哈")
		
	}
	cqp.AddLog(cqp.Debug, "调试", "调试88："+msg)
	return 0
}