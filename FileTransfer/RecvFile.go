package main

import (
    "fmt"
    "io"
    "net"
    "os"
)

//接受文件
func ReciveFile(fileName string, conn net.Conn) {
    f, err := os.Create(fileName)
    if err != nil {
        fmt.Println("os.Create err = ", err)
        return
    }
    buf := make([]byte, 1024*4)
    for {
        n, err := conn.Read(buf) //接受对方发送的文件内容
        if err != nil {
            if err == io.EOF {
                fmt.Println("文件接受 完毕")
            } else {
                fmt.Println("f.Read err = ", err)
            }

            return
        }
        f.Write(buf[:n])
    }
}

func main() {
    //监听
    listenner, err := net.Listen("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println("net.Listen err = ", err)
        return
    }
    defer listenner.Close()

    // 阻塞等待用户连接
    conn, err1 := listenner.Accept()
    if err1 != nil {
        fmt.Println("listenner.Accept err1 = ", err1)
        return
    }
    buf := make([]byte, 1024)
    var n int
    n, err = conn.Read(buf) //读取对方发送的文件名
    if err != nil {
        fmt.Println("conn.Read err = ", err)
        return
    }
    fileName := string(buf[:n])
    defer conn.Close()

    // 回复ok
    conn.Write([]byte("ok"))

    //接受文件
    ReciveFile(fileName, conn)
}

