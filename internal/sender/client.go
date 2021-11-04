package sender

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

/**
 *
 * @author hushengdong
 */
func sendTcp(mqReq *mqProtocol, server string) {

	// 连接服务器
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Printf("connect to server error: [%s]", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	client := &TcpClient{}
	client.conn = conn
	client.stopChan = make(chan struct{})
	client.dst = server
	client.src = conn.LocalAddr().String()
	client.mqReq = mqReq

	body := toByteArray(mqReq)
	desc := client.src + " ~ " + client.dst
	reqMsg, _ := getReqCodeMsg(int(client.mqReq.Header.Code))
	fmt.Println(now(), "发送包------------------------[", desc, "]--------------------")
	fmt.Println(now(), "[", reqMsg, "][请求][", client.mqReq.Header.Code, "] 消息体 :", toJson(client.mqReq))
	client.sendReq(body)

	go client.receiveResp()
	// 等待关闭
	<-client.stopChan
}

//客户端对象
type TcpClient struct {
	conn     net.Conn
	stopChan chan struct{}
	src      string
	dst      string
	mqReq    *mqProtocol
}

// 接收数据包 直接阻塞读取
func (client *TcpClient) receiveResp() {

	now := now()
	for {
		length := client.readInt()
		serializeTypeHeaderLength := client.readInt()
		headerLengh := (serializeTypeHeaderLength & 0xffffffff)
		header := client.readString(headerLengh)
		mqHeader := &mqProtocolHeader{}
		err := json.Unmarshal(header, mqHeader)
		if err != nil {
			fmt.Println(now, "消息头数据 :", header)
		}
		bodyLen := length - 4 - 4 - headerLengh
		body := client.readString(bodyLen)
		msg, _ := getRespCodeMsg(int(mqHeader.Code))

		desc := client.src + " ~ " + client.dst
		fmt.Println(now, "------------------------[", desc, "]--------------------")
		fmt.Println(now, "[", msg, "][响应][header][", mqHeader.Code, "] 消息头数据 :", header)
		fmt.Println(now, "[", msg, "][响应][body][", body)
		close(client.stopChan)
	}

}

func (client *TcpClient) readString(len int) []byte {

	buf := make([]byte, 0, len)
	temp := make([]byte, len)
	total := 0
	for total < len { // 可能一次读不到所有的数据
		n, err := client.conn.Read(temp)
		if err != nil {
			panic(err)
		}
		buf = append(buf, temp[0:n]...)
		total = total + n
	}
	return buf
}

func (client *TcpClient) readInt() int {

	total := 0
	buf := make([]byte, 0, 4)
	temp := make([]byte, 4)
	for total < 4 {
		n, err := client.conn.Read(temp)
		if err != nil {
			panic(err)
		}
		buf = append(buf, temp...)
		total = total + n
	}
	return bytesToInt(buf)
}

// 发送数据包
func (client *TcpClient) sendReq(body []byte) {

	n, err := client.conn.Write(body)
	if err != nil {
		panic(err)
		close(client.stopChan)
	}
	fmt.Println("发送成功：", n)
}

func bytesToInt(buf []byte) int {
	return int(binary.BigEndian.Uint32(buf))
}

func now() string {
	now := time.Now()      //获取当前时间
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d ", year, month, day, hour, minute, second)
}

func toJson(mq *mqProtocol) string {

	res, _ := json.Marshal(mq)
	return string(res)
}
