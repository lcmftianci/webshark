package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    service := ":5000"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkErr(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkErr(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(conn)
    }
}

func checkErr(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()
    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:])
        if err != nil {
            return
        }
        rAddr := conn.RemoteAddr()
        fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))
        _, err2 := conn.Write([]byte("Welcome client!"))
        if err2 != nil {
            return
        }
    }
}
