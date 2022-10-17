package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	// 转换器，允许跨域
	upGrader = websocket.Upgrader{
		WriteBufferSize: 10240,
		ReadBufferSize:  10240,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// 当客户端访问/ws接口的时候，会自动调用这个wsHandler接口
func wsHandler(w http.ResponseWriter, r *http.Request) {
	//完成协议转换，http握手应答
	//w.Write([]byte("Hello,我是自动触发的函数"))
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)
	if conn, err = upGrader.Upgrade(w, r, nil); err != nil {
		return
	}
	//没有错误进行数据收发
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		fmt.Println(string(data))
		text := []byte("这就是爱")
		if err = conn.WriteMessage(websocket.TextMessage, text); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("127.0.0.1:7777", nil)
}
