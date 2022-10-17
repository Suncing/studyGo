package json

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
*   @author wangchenyang
*   @date 2022/8/29 19:41
*   @description:
 */
type Person struct {
	Name  string `json:"name"`
	Age   string `json:"age"`
	Hobby string `json:"hobby"`
	Pet   Pet    `json:"pet"`
}

type Pet struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Call  string `json:"call"`
}

func main() {
	file, _ := os.Open("redis.json")
	decoder := json.NewDecoder(file)
	fmt.Println(decoder)
	conf := Person{}
	_ = decoder.Decode(&conf)
	fmt.Println(conf)
	fmt.Println(conf.Age)
	fmt.Println(conf.Name)
	fmt.Println(conf.Hobby)
	fmt.Println(conf.Pet.Color)
	fmt.Println(conf.Pet.Name)
	fmt.Println(conf.Pet.Call)
}
