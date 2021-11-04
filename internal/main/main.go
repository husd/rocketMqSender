package main

import (
	"flag"
	"husd.com/sender/sender"
)

/**
 * 模拟给rocketMQ发送各种TCP报文的工具，视角是底层的网络通讯。
 * 也可以用来学习rocketMQ的底层实现原理。
 * 谨慎操作
 * @author hushengdong
 */
func main() {

	configFile := flag.String("f", "./config.properties", "配置文件地址")
	code := flag.Int("code", 34, "请求代码 默认给nameserver发送一个心跳消息")
	targetType := flag.Int("target", 0, "向谁请求消息 0nameserver 1broker")
	sender.SendMqReqByCode(*configFile, *code, *targetType)
}