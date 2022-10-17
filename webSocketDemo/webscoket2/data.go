package main

/**
*   @author wangchenyang
*   @date 2022/8/17 8:55
*   @description:
 */

type Data struct {
	Ip   string `json:"ip"`
	User string `json:"user"`
	Type string `json:"type"`
	//代表哪个用户说的
	From string `json:"from"`
	//传输内容
	Content string `json:"content"`
	//用户列表
	UserList []string `json:"user_list"`
}
