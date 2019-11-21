/*
 * TODO: 变量待加锁
 */

package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

type Head struct {
	flag     [2]byte // 标识 AU
	sequence uint16  //
	body_len uint16  // payload 内容长度(字节数)
	identify uint16  // 身份识别 号
}

type Messages struct {
	con  net.Conn
	data []byte
}

////////////////////////////////////////////////////////
// 非业务相关的 Golang 错误检查函数
////////////////////////////////////////////////////////
func checkError(err error, info string) (res bool) {
	if err != nil {
		fmt.Println(info + "  " + err.Error())
		return false
	}
	return true
}

////////////////////////////////////////////////////////
// 读取指定区间的字节流 ([]byte)
////////////////////////////////////////////////////////
func ReadBuffer(conn net.Conn, buffer []byte, index int, szRead int) (res bool) {
	if len(buffer) < index+szRead {
		fmt.Print("ReadBuffer Error: index+szRead is too large")
		return false
	}

	var indexBefore, lengthBefore int

	indexBefore = index
	lengthBefore = szRead

	for {
		if lengthBefore <= 0 {
			break
		}

		lengthAfter, err := conn.Read(buffer[indexBefore : indexBefore+lengthBefore])
		if checkError(err, "Connection") == false {
			conn.Close()
			return false
		}
		if lengthAfter == lengthBefore {
			break
		}
		indexBefore += lengthAfter
		lengthBefore -= lengthAfter
	}

	return true
}

////////////////////////////////////////////////////////
// 接收数据的 goroutine 执行体
// 参数：
//		数据连接 conn
//		通讯通道 messages
////////////////////////////////////////////////////////
func Receiver(conn net.Conn, messages chan Messages) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())

	buffer := make([]byte, 1024) // header+body+footer
	var nbBadPack uint8
	for {
		// close bad TCP.socket
		if nbBadPack > 100 {
			conn.Close()
			break
		}

		if ReadBuffer(conn, buffer, 0, 1024) == false {
			conn.Close()
			break
		}

		// 接下来是一系列的数据包合法性检查
		// 'A'=65, 'U'=85
		if buffer[0] != 65 || buffer[1] != 85 {
			fmt.Println("Bad Header: ", buffer[0], ", ", buffer[1])
			nbBadPack++
			continue
		}

		szBody := binary.BigEndian.Uint16(buffer[4:6])

		if ReadBuffer(conn, buffer, 8, int(szBody)) == false {
			conn.Close()
			break
		}

		if ReadBuffer(conn, buffer, 8+int(szBody), 2) == false {
			conn.Close()
			break
		}

		// 'U'=85, 'A'=65
		if buffer[8+szBody] != 85 || buffer[8+szBody+1] != 65 {
			fmt.Println("Bad Footer: ", buffer[0], ", ", buffer[1])
			nbBadPack++
			continue
		}

		msg := Messages{con: conn, data: buffer[0 : szBody+8+2]} // same as msg := Messages{conn, buf[0:length]}
		messages <- msg                                          // to be modified
	}
}

////////////////////////////////////////////////////////
// 发送数据的 goroutine 执行体
// 参数：
//		连接字典 conns
//		数据通道 messages
////////////////////////////////////////////////////////
func Sender(conns *map[string]net.Conn, messages chan Messages) {
	for {
		msg := <-messages // 在本 goroutine 中，阻塞等待消息
		for key, value := range *conns {
			if value == msg.con {
				// 不转发给自己
				continue
			}
			// TCP 发送操作
			_, err := value.Write(msg.data)
			// 发送失败则移除该 TCP 连接
			if err != nil {
				fmt.Println(err.Error())
				delete(*conns, key)
			}
		}
	}
}

////////////////////////////////////////////////////////
// 启动服务器
// 参数:
//		TCP 端口
////////////////////////////////////////////////////////
func StartServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+port)
	checkError(err, "ResolveTCPAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")

	// conns 是所有的 C/S 连接的 map
	conns := make(map[string]net.Conn)

	// messages 是 Server 收到的 所有消息的 缓冲队列，在此预留长度为 10
	messages := make(chan Messages, 10)

	// 发送的 goroutine，对于所有连接
	go Sender(&conns, messages)

	// 循环接收新的 TCP 连接
	for {
		conn, err := l.Accept()
		checkError(err, "Accept")
		conns[conn.RemoteAddr().String()] = conn // 新的 conn 连接 加入 conns 连接队列

		// 发送的 goroutine，对于一个连接
		go Receiver(conn, messages)
	}
}

////////////////////////////////////////////////////////
// main 函数
// 默认使用 7777 端口
////////////////////////////////////////////////////////
func main() {
	if len(os.Args) == 2 {
		StartServer(os.Args[1])
	} else {
		StartServer("7777")
	}
}
