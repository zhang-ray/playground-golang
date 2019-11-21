// 备注：
// 1. 预设的 SPS, PPS NALU 已经预存在 Android 客户端上

package main

import (
	"fmt"
	"net"
	"time"
)

const (
	UDP_PORT_RPI     = 8888
	UDP_PORT_ANDROID = 6666
)

type Messages struct {
	srcAddr *net.UDPAddr
	data    []byte
}

var (
	streamDst  net.UDPAddr
	streamSrc  net.UDPAddr
	gotAndroid bool
)

func sendAVC(socket *net.UDPConn, messages chan Messages) {
	for {
		// 下面一句是阻塞的
		msg := <-messages
		// length, err := socket.WriteToUDP(msg.data, connsToSend)
		if !gotAndroid {
			time.Sleep(1 * time.Second)
			continue
		}
		_, err := socket.WriteToUDP(msg.data, &streamDst)
		if err != nil {
			fmt.Println("发送失败", err)
			return
		}
	}
}

func printState() {
	for {
		time.Sleep(1 * time.Minute)

		fmt.Println("last Android addr:      " + streamDst.String())
		fmt.Println("last Raspberry Pi addr: " + streamSrc.String())
	}
}

func getNaluType(data []byte) uint8 {
	return data[4] & 0x1F
}

func forAndroid(socketANDROID *net.UDPConn) {
	// 接收的数据
	data := make([]byte, 1<<10)
	for {
		_, remote, err := socketANDROID.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败", err)
			continue
		}
		fmt.Println("收到 Android 的数据了")

		if data[0] == 0x41 && data[1] == 0x44 { // AD, Android 手机
			fmt.Println("AD, Android 手机")
			streamDst = *remote
			gotAndroid = true
		}
	}
}

func main() {
	var naluType uint8

	// 树莓派 的 socket
	socketRPI, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: UDP_PORT_RPI,
	})
	if err != nil {
		fmt.Println("监听失败", err)
		return
	}
	defer socketRPI.Close()

	// Android 手机 的 socket
	socketANDROID, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: UDP_PORT_ANDROID,
	})
	if err != nil {
		fmt.Println("监听失败", err)
		return
	}
	defer socketANDROID.Close()

	messages := make(chan Messages, 100) // 100是消息队列长度

	//启动服务器广播协程
	go sendAVC(socketANDROID, messages)
	//启动打印状态的协程
	go printState()

	// 接收的数据
	data := make([]byte, 1<<10)

	fmt.Println("开始循环接收数据")

	go forAndroid(socketANDROID)

	// 树莓派循环接收数据
	for {
		length, remote, err := socketRPI.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败", err)
			continue
		}
		// fmt.Println(data[:length])
		// fmt.Print("收到数据：")
		// fmt.Println(remote)

		// 判断包头, RP, 树莓派
		if data[0] == 0x52 && data[1] == 0x50 {
			messages <- Messages{srcAddr: remote, data: data[:length]}
			streamSrc = *remote
			if data[7] == 1 { // first sub sequence
				naluType = getNaluType(data[8:length])
				if naluType != 1 && naluType != 5 {
				}
			}
			//else {
			//fmt.Println("RP, 树莓派")
			//fmt.Println(data[:15])
			//}
		} else {
			fmt.Print("未知数据类型：")
			fmt.Println(data[0:1])
		}
	}
}
