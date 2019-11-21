package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

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
//客户端发送线程
//参数
//		发送连接 conn
//
////////////////////////////////////////////////////////
func SendData(conn net.Conn) {

	username := conn.LocalAddr().String()
	fmt.Println(username)

	var i, szBody uint16
	szBody = 160

	buffer := make([]byte, 2048)
	buffer[0] = 65 //'A'
	buffer[1] = 85 //'U'
	buffer[2] = 0
	buffer[3] = 0
	binary.BigEndian.PutUint16(buffer[4:6], szBody)
	buffer[6] = 0
	buffer[7] = 0

	for i = 0; i < szBody; i++ {
		buffer[8+i] = 0
	}

	buffer[8+szBody] = 85 //'U'
	buffer[9+szBody] = 65 //'A'

	for {
		_, err := conn.Write(buffer[0 : szBody+8+2])
		// if length != int(szBody)+8+2 {
		// 	fmt.Println(length, "!=", szBody+8+2)
		// }
		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			break
		}
		time.Sleep(1e9)
	}
}

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "121.40.209.116:7777")
	checkError(err, "ResolveTCPAddr")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "DialTCP")

	//启动客户端发送线程
	go SendData(conn)

	//开始客户端循环接收数据
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if checkError(err, "Connection") == false {
			conn.Close()
			fmt.Println("Server is dead.")
			os.Exit(0)
		}
	}
}
