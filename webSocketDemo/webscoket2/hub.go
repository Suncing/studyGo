package main

import "encoding/json"

/**
*   @author wangchenyang
*   @date 2022/8/17 9:50
*   @description:
 */
type hub struct {
	//connections注册了连接器
	connections map[*connection]bool
	//从连接器发送的消息
	broadcast chan []byte
	//从连接器注册请求
	register chan *connection
	//销毁请求
	unregister chan *connection
}

var h = hub{
	connections: make(map[*connection]bool),
	//从连接器发送的信息
	broadcast: make(chan []byte),
	//从连接器注册请求
	register: make(chan *connection),
	//销毁请求
	unregister: make(chan *connection),
}

func (h *hub) run() {
	//监听数据管道，在后端能不断处理管道数据
	for {
		//根据不同的数据管道，处理不同的逻辑
		select {
		case c := <-h.register:
			//标志注册了
			h.connections[c] = true
			//组装data数据，因为要不断给其他人传播数据
			c.data.Ip = c.ws.RemoteAddr().String()
			//更新类型
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)

			c.send <- data_b
		case c := <-h.unregister:
			//判断map里是否存在要删的数据
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case data := <-h.broadcast:
			//处理数据流转，将数据同步到所有用户
			for c := range h.connections {
				select {
				case c.send <- data:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}
