// http://blog.csdn.net/wowzai/article/details/9936659

package main

import (
	"fmt"
	"net"
	"os"
	// "bufio"
)

const BUFF_SIZE = 1024

var input = make([]byte, BUFF_SIZE)

func handleError(err error) {
	if err == nil {
		return
	}
	fmt.Printf("error:%s\n", err.Error())
}

func main() {
	var port string
	if len(os.Args) > 2 {
		fmt.Println("Usage:<command> <port>")
		return
	} else if len(os.Args) == 2 {
		port = os.Args[1]
	} else {
		port = "7777"
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:"+port)
	handleError(err)
	tcpConn, err := net.DialTCP("tcp4", nil, tcpAddr)
	handleError(err)
	defer tcpConn.Close()
	// reader :=  bufio.NewReader(os.Stdin)
	var continued = true
	var inputStr string
	buffer := make([]byte, 1024)
	inputStr = "What do you whant me to send?!?!"
	for continued {
		// n, err := reader.Read(input)
		// handleError(err)
		//if n > 0 {
		szRead, _ := tcpConn.Write([]byte(inputStr))
		if szRead < 0 {
			// inputStr = string(input[:k])
			// fmt.Printf("Write:%s",inputStr)
			// if inputStr == "exit\n" {  //在比对时由于有个回车符，所以加上\n
			continued = false //也可以将inputStr = TrimRight(inputStr,"\n")
			// }
			//}
		}
		szWrite, _ := tcpConn.Read(buffer)
		if szWrite < 0 {
			continued = false
		}
		// else {
		//   println("Not Equal!")
		// }
		// print(string(buffer))
		// println()
	}
}
