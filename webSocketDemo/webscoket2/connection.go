package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type connection struct {
	//ws连接器
	ws   *websocket.Conn
	send chan []byte
	data *Data
}

func (c *connection) writer() {
	for message := range c.send {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

var user_list = []string{}

func (c *connection) reader() {
	//不断读取websocket数据
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			//读不进去数据可能出问题了，所以将用户连接移除
			h.unregister <- c //将该对象放到unregister管道中
			break
		}
		//读取数据，反序列化
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			//登陆后，将用户加入用户列表
			user_list = append(user_list, c.data.User)
			//每个用户都加载所有登录了的列表
			c.data.UserList = user_list
			//数据序列化
			data_b, _ := json.Marshal(c.data)
			h.broadcast <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.broadcast <- data_b
		case "logout":
			c.data.Type = "logout"
			//用户列表删除
			user_list = remove(user_list, c.data.User)
			c.data.UserList = user_list
			c.data.Content = c.data.User
			//数据序列化，让所有人看到xxx下线了
			data_b, _ := json.Marshal(c.data)
			h.broadcast <- data_b
			h.unregister <- c
		default:
			fmt.Print("========default================")
		}
	}
}

func remove(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	//新的返回切片
	var n_slice = []string{}
	//删除传入切片中的指定用户，其他用户放到新的切片
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(n_slice)
	return n_slice
}

//定义ws升级，将http请求升级为ws请求
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//1.获取ws对象
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	//2.创捷连接对象去做事情，也就是初始化连接对象
	c := &connection{send: make(chan []byte, 256), ws: ws, data: &Data{}}
	//早ws中注册一些
	h.register <- c
	//ws将数据读写跑出来
	go c.writer()
	c.reader()
	defer func() {
		c.data.Type = "logout"
		//用户列表删除
		user_list = remove(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		//数据序列化，让所有人看到xxx下线了
		data_b, _ := json.Marshal(c.data)
		h.broadcast <- data_b
		h.unregister <- c
	}()
}
