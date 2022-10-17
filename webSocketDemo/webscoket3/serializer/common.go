package serializer

/**
*   @author wangchenyang
*   @date 2022/8/17 10:15
*   @description:
 */
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
