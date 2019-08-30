// UDPServer project main.go
package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    service := ":5000"
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkErr(err)
    conn, err := net.ListenUDP("udp", udpAddr)
    checkErr(err)
    for {
        handleClient(conn)
    }
}

func checkErr(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func handleClient(conn *net.UDPConn) {
    defer conn.Close()
    var buf [512]byte
    for {
        n, rAddr, err := conn.ReadFromUDP(buf[0:])
        if err != nil {
            return
        }
        fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))
        _, err2 := conn.WriteToUDP([]byte("ServerSay:" + string(buf[0:n])), rAddr)
        if err2 != nil {
            return
        }
    }
}
