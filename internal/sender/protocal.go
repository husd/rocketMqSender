package sender

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

/**
 *
 * @author hushengdong
 */
type mqProtocol struct {
	Length                    int32 // 总长度
	SerializeTypeHeaderLength int32 // 序列化类型&headerLength
	Header                    *mqProtocolHeader
	Body                      string
}

type mqProtocolHeader struct {
	Code      int32
	Language  string
	Version   int32
	Opaque    int32
	Flag      int32
	Remark    string
	ExtFields string
}

// 这里并没有计算长度，后续转字节数组的时候再计算
func NewMqReq(code reqCode) *mqProtocol {

	mq := &mqProtocol{}

	header := &mqProtocolHeader{}
	header.Code = int32(code)
	header.Language = "JAVA"
	header.Version = getVersion()
	header.Opaque = 123456789
	header.Flag = 0
	header.Remark = ""
	header.ExtFields = getExtFieldsByCode(code)

	mq.Length = 0
	mq.SerializeTypeHeaderLength = 0
	mq.Header = header
	mq.Body = getBodyByCode(code)

	return mq
}

func calcSeriAndHeaderLen(serializeType int, length int) int {

	return (serializeType << 24) | (length & 0xffffffff)
}

func toByteArray(mq *mqProtocol) []byte {

	header := mq.Header
	headerByte, err := json.Marshal(header)
	if err != nil {
		panic("转换header的json错误")
	}
	headerLen := len(headerByte)
	const seriType int = 0

	body := mq.Body
	bodyByte := []byte(body)
	bodyLen := len(bodyByte)

	length := 4 + 4 + headerLen + bodyLen
	serializeTypeHeaderLength := calcSeriAndHeaderLen(seriType, headerLen)

	mq.Length = int32(length)
	mq.SerializeTypeHeaderLength = int32(serializeTypeHeaderLength)

	res := make([]byte, 0, length)
	res = append(res, int2Byte(int32(length))...)
	res = append(res, int2Byte(int32(serializeTypeHeaderLength))...)
	res = append(res, headerByte...)
	res = append(res, bodyByte...)

	return res
}

// 先保证正确 大端模式
func int2Byte(n int32) []byte {

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}
