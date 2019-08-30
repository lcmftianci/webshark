package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "io/ioutil"
)

func queryHtml(strUrl string, strPath string, flag int)(strCtx string){
    resp, err := http.Get(strUrl)
    if err != nil {
        fmt.Println(err)
        log.Fatal(err)
    }
    if resp.StatusCode == http.StatusOK {
        fmt.Println(resp.StatusCode)
    }
    defer resp.Body.Close()

    if flag == 1{
    buf := make([]byte, 1024)
    f, err1 := os.OpenFile(strPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)//可读写，追加的方式打开（或创建文件）
    if err1 != nil {
        panic(err1)
        return ""
    }
    defer f.Close()

    for {
        n, _ := resp.Body.Read(buf)
        if 0 == n {
            break
        }
        f.WriteString(string(buf[:n]))
	return ""
    }	
  }else{
   s, _ := ioutil.ReadAll(resp.Body) //把  body 内容读入字符串 s
   return string(s)
  }
  return ""
}

func main(){
	
	fmt.Println("Hello go!")
	strCtx := ""
	strCtx = queryHtml("https://www.baidu.com", "", 0)
	strCtx = queryHtml("https://blog.csdn.net/quicmous/article/details/80068161", "", 0)
	fmt.Print(strCtx)
}
