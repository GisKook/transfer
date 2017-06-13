package http_server

const (
	HTTP_PLC_ID     string = "plc_id"
	HTTP_PLC_SERIAL string = "serial"

	// query all routers
	HTTP_QUERY_ALL_ROUTERS string = "/plc/query_all_routers"

	//////////////RESPONSE////////////////
	//HTTP_RESPONSE_RESULT               string = "result"
	HTTP_RESPONSE_RESULT_SUCCESS uint8 = 0
	HTTP_RESPONSE_RESULT_FAILED  uint8 = 1

	HTTP_RESPONSE_RESULT_PARAMTER_ERR  uint8 = 255
	HTTP_RESPONSE_RESULT_TIMEOUT       uint8 = 254
	HTTP_RESPONSE_RESULT_SERVER_FAILED uint8 = 253
)

//var HTTP_RESULT []string = []string{"成功", "失败,路由器反馈失败 或 dps服务器内部错误", "参数错误", "超时,路由器掉线 或 路由器反馈慢"}
var HTTP_RESULT map[uint8]string = map[uint8]string{
	HTTP_RESPONSE_RESULT_SUCCESS:       "成功",
	HTTP_RESPONSE_RESULT_FAILED:        "失败",
	HTTP_RESPONSE_RESULT_PARAMTER_ERR:  "参数错误",
	HTTP_RESPONSE_RESULT_TIMEOUT:       "超时",
	HTTP_RESPONSE_RESULT_SERVER_FAILED: "失败,dps服务器内部错误"}
