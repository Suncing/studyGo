package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(ExecPath())
	fmt.Println(GetCurrentDirectory())
}

/**
 * ExecPath
 * @author wangcy
 * @description: 获取项目执行文件
 * @date 2022-08-30 08:57:34
 * @return string
 * @return error
 */func ExecPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatalln(err)
	}
	re, err := filepath.Abs(file)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("The path is", re)
	return re
}

/**
 * GetCurrentDirectory
 * @author wangcy
 * @description: 获取项目执行路径
 * @date 2022-08-30 08:57:48
 * @return string
 */
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
