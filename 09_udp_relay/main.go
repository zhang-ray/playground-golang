package main

import (
	"fmt"
	"net"
	"time"
)

const (
	UDP_PORT = 8888
)

type Messages struct {
	srcAddr *net.UDPAddr
	data    []byte
}

func sendProxy(socket *net.UDPConn, conns *map[string]net.UDPAddr, messages chan Messages) {
	for {
		// 下面一句是阻塞的
		msg := <-messages
		for key, value := range *conns {
			// fmt.Println("key=", key)
			if key == msg.srcAddr.String() {
				continue
			}
			length, err := socket.WriteToUDP(msg.data, &value)
			if length != len(msg.data) || err != nil {
				fmt.Println(err.Error())
				delete(*conns, key)
			}
		}
	}
}

func printState(conns *map[string]net.UDPAddr) {
	for {
		time.Sleep(1 * time.Minute)

		fmt.Println("last conn list:")
		for key, _ := range *conns {
			fmt.Println("key=", key)
		}
		fmt.Println()
	}
}
func main() {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: UDP_PORT,
	})

	if err != nil {
		fmt.Println("监听失败", err)
		return
	}

	defer socket.Close()

	conns := make(map[string]net.UDPAddr)

	messages := make(chan Messages, 100) // 100是消息队列长度

	//启动服务器广播线程
	go sendProxy(socket, &conns, messages)
	go printState(&conns)

	data := make([]byte, 1<<10)

	for {
		length, remote, err := socket.ReadFromUDP(data)

		if err != nil {
			fmt.Println("读取数据失败", err)
			continue
		}
		fmt.Println(data[:length])
		fmt.Println(remote)

		conns[remote.String()] = *remote
		messages <- Messages{srcAddr: remote, data: data[:length]}

	}
}
