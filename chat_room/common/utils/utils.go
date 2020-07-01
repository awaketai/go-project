package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gostudy/chatroom/common/message"
	"io"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte // 传输时使用的缓冲

}
// 发送数据
func (tr *Transfer) WritePkg(data []byte) (err error) {
	// 1.先发送数据长度
	var bufLen uint32
	bufLen = uint32(len(data))
	binary.BigEndian.PutUint32(tr.Buf[0:4],bufLen) // 将长度转换为二进制字节
	n,err := tr.Conn.Write(tr.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) failed",err)
		return
	}
	// 发送data
	n,err = tr.Conn.Write(data)
	if uint32(n) != bufLen || err != nil {
		fmt.Println("send data error",err)
		return
	}
	return
}

// 读取发送的数据
func (tr *Transfer) ReadPkg()(msg message.Message,err error)  {
	buf := tr.Buf
	fmt.Println("read data...")
	n,err := tr.Conn.Read(buf[:4])
	if err != nil{
		if err == io.EOF {

		}
		fmt.Println("read data length error",err)
		return
	}
	var bufLen uint32
	bufLen = binary.BigEndian.Uint32(buf[0:4])
	n,err = tr.Conn.Read(buf[:bufLen])
	if n != int(bufLen) || err != nil {
		fmt.Println("read data error",err)
		return
	}
	// 反序列化接收到的数据
	err = json.Unmarshal(buf[:bufLen],&msg)
	if err != nil {
		fmt.Println("json unmarshal error",err)
		return
	}
	return
}