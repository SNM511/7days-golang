package codec

import (
	"io"
)

// Header 将请求和响应中的参数和返回值抽象为 body，剩余的信息放在 header 中
type Header struct {
	ServiceMethod string // 服务名和方法名
	Seq           uint64 // sequence number chosen by client，用于区分请求
	Error         string
}

// Codec 对消息体进行编解码的接口的抽象
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
