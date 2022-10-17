package utils

/**
*   @author  王晨阳
*   @date 2022/8/24 8:45
*   @description:toString
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ToString(value interface{}) {
	marshal, _ := json.Marshal(value)
	var out bytes.Buffer
	json.Indent(&out, marshal, "", "\t")
	fmt.Println(out.String())
}
