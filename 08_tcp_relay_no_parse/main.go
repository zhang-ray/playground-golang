/*
 * TODO: 解决粘包，半包问题
 */

package main

import (
	"fmt"
	"net"
	"os"
)

type Head struct {
	flag     [2]byte // 标识 AU
	sequence uint16  // 内容长度,指的是具体的包体内容
	body_len uint16  // 内容长度,指的是具体的包体内容
	identify uint16  // 身份识别 号
}

type Messages struct {
	con  net.Conn
	data []byte
}

////////////////////////////////////////////////////////
//
//错误检查
//
////////////////////////////////////////////////////////
func checkError(err error, info string) (res bool) {
	if err != nil {
		fmt.Println(info + "  " + err.Error())
		return false
	}
	return true
}

////////////////////////////////////////////////////////
//
//服务器端接收数据线程
//参数：
//		数据连接 conn
//		通讯通道 messages
//
////////////////////////////////////////////////////////
func Handler(conn net.Conn, messages chan Messages) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())

	buffer := make([]byte, 1024) // header+body+footer
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			println(err.Error())
		}

		msg := Messages{con: conn, data: buffer[0 : length-1]} // same as msg := Messages{conn, buf[0:length]}
		messages <- msg                                        // to be modified
	}
}

////////////////////////////////////////////////////////
//
//服务器发送数据的线程
//
//参数
//		连接字典 conns
//		数据通道 messages
//
////////////////////////////////////////////////////////
func echoHandler(conns *map[string]net.Conn, messages chan Messages) {
	for {
		// 下面一句是阻塞的
		msg := <-messages
		for key, value := range *conns {
			if value == msg.con {
				continue
			}
			_, err := value.Write(msg.data)
			if err != nil {
				fmt.Println(err.Error())
				delete(*conns, key)
			}
		}
	}
}

////////////////////////////////////////////////////////
//
//启动服务器
//参数
//	端口 port
//
////////////////////////////////////////////////////////
func StartServer(port string) {
	service := ":" + port //strconv.Itoa(port);
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, "ResolveTCPAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")

	conns := make(map[string]net.Conn)
	// messages := make(chan []byte, 10)
	messages := make(chan Messages, 10)

	//启动服务器广播线程
	go echoHandler(&conns, messages)

	for {
		//fmt.Println("Listening ...")
		conn, err := l.Accept()
		checkError(err, "Accept")
		//fmt.Println("Accepting ...")
		conns[conn.RemoteAddr().String()] = conn
		//启动一个新线程
		go Handler(conn, messages)
	}
}

////////////////////////////////////////////////////////
//
//主程序
//
//参数说明：
//	启动服务器端：  Chat server [port]				eg: Chat server 9090
//	启动客户端：    Chat client [Server Ip Addr]:[Server Port]  	eg: Chat client 192.168.0.74:9090
//
////////////////////////////////////////////////////////
func main() {
	if len(os.Args) == 2 {
		StartServer(os.Args[1])
	} else {
		StartServer("7777")
	}
}
