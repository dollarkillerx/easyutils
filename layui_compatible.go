package easyutils

// layui兼容文件

// 注意使用layui 请将模板文件设置为 <<< >>> 以免冲突

// 对Layui数据结构封装


// 分页返回数据结构体
type LayUiPage struct {
	Code int `json:"code"`
	Count int `json:"count"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
}


// 包装分页数据
func LayPage(count int,data interface{}) *LayUiPage {
	result := &LayUiPage{}
	result.Code = 0
	result.Data = data
	result.Count = count
	result.Msg = ""
	return result
}

// 分页客户端请求结构体
type LayGet struct {
	Page int `form:"page"`
	Limit int `form:"limit"`
}

// 定义空数据  当无数据 或者数据异常时返回
var (
	LayEmptyData = &LayUiPage{Code:0,Count:0,Msg:"Empty Data"} // 空数据
	LayErrorData = &LayUiPage{Code:-1,Count:0,Msg:"Error Data"} // 异常数据
)