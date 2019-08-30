// UDPClient project main.go
package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "strings"
)

func main() {
    var buf [512]byte
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkErr(err)
    conn, err := net.DialUDP("udp", nil, udpAddr)
    defer conn.Close()
    checkErr(err)
    rAddr := conn.RemoteAddr()
    n, err := conn.Write([]byte("Hello server!"))
    checkErr(err)
    for{
    	n, err = conn.Read(buf[0:])
    	checkErr(err)
    	fmt.Println("Reply from server ", rAddr.String(), string(buf[0:n]))
        reader := bufio.NewReader(os.Stdin)
        text, _ := reader.ReadString('\n')
        // convert CRLF to LF
        text = strings.Replace(text, "\n", "", -1)
        if strings.Compare("hi", text) == 0 {
            fmt.Println("hello, Yourself")
	    os.Exit(0)
        }
	conn.Write([]byte(text))
    }
    os.Exit(0)
}

func checkErr(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
