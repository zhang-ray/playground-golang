// http://blog.csdn.net/wowzai/article/details/9936659
package main

import (
 "net"
 "os"
 "fmt"
 "io"
)

const BUFF_SIZE = 1024
var buff = make([]byte,BUFF_SIZE)

// 接受一个TCPConn处理内容
func handleConn(tcpConn *net.TCPConn){
  println("handleConn __00")
  if tcpConn == nil {
    return
  }
  for {
    szRead, errRead := tcpConn.Read(buff)
    handleError(errRead)
    if errRead == io.EOF {
      fmt.Printf("The RemoteAddr:%s is closed!\n",tcpConn.RemoteAddr().String())
      return
    }

    // tcpConn.Write 的输出长度要注意    
    szWrite, errWrite := tcpConn.Write(buff[:szRead])
    // szWrite, errWrite := tcpConn.Write(buff, szRead)
    if szRead != szWrite {
      fmt.Printf("%d, %d\n", szRead, szWrite)
    }
    handleError(errWrite)
    // if string(buff[:n]) == "exit" {
    //   fmt.Printf("The client:%s has exited\n",tcpConn.RemoteAddr().String())
    // }
    // if n > 0 {
    //   fmt.Printf("Read:%s",string(buff[:n]))
    // }
      // println("handleConn __02")
  }
  println("handleConn __00")
}
// 错误处理
func handleError(err error) {
  if err == nil {
    return
  }
  fmt.Printf("error:%s\n",err.Error())
}

func main() {
  var port string
  if len(os.Args) > 2 {
    fmt.Println("Usage:<command> <port>")
    return
  } else if len(os.Args) == 2 {
      port = os.Args[1]
  } else{
   port = "7777"
 }
  tcpAddr,err := net.ResolveTCPAddr("tcp4",":"+port)
  handleError(err)
  tcpListener,err := net.ListenTCP("tcp4",tcpAddr)  //监听
  handleError(err)
  defer tcpListener.Close()
  for {
    println("tcpListener.AcceptTCP() before")
    tcpConn, err := tcpListener.AcceptTCP()
    println("tcpListener.AcceptTCP() after")
    fmt.Printf("The client:%s has connected!\n", tcpConn.RemoteAddr().String())
    handleError(err)
    defer tcpConn.Close()
    println("go handleConn(tcpConn) before")
    go handleConn(tcpConn)    //起一个goroutine处理
    println("go handleConn(tcpConn) after")
  }
}