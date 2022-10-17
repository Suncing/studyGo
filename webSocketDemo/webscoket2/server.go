package main

/**
*   @author wangchenyang
*   @date 2022/8/17 15:57
*   @description:
 */
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	//ws控制器不断去处理管道数据，进行同步数据
	go h.run()
	//指定回调函数
	router.HandleFunc("/ws", wsHandler)
	//开启服务端监听
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err", err)
	}
}

//【go聊天室】2019年最新go实战视频教程聊天室实战channel，协程，websocket实现一个小型聊天室
