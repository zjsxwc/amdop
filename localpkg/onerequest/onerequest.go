package onerequest

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"strings"
)

type MainController struct {
	beego.Controller
}

type SubRequest struct {
	Api    string `json:"api"`    //必须是包含schema、port以及query参数的完整URL地址，比如 https://api.xxx.com:3000/someapi?param1=value1&param2=value2
	Method string `json:"method"` // post 或者 get
	Params string `json:"params"` //post 参数 json字符串
}

type SubResponse struct {
	Api      string                 `json:"api"`
	Method   string                 `json:"method"`
	Params   string                 `json:"params"`
	Response map[string]interface{} `json:"response"`
	Error    interface{}            `json:"error"`
}

func (c *MainController) processSubRequest(subRequest SubRequest) SubResponse {
	sid := c.Ctx.Input.Header(beego.BConfig.WebConfig.Session.SessionNameInHTTPHeader)
	var req *httplib.BeegoHTTPRequest
	if strings.ToLower(subRequest.Method) == "get" {
		req = httplib.Get(subRequest.Api).Header(beego.BConfig.WebConfig.Session.SessionNameInHTTPHeader, sid).Header("X-Requested-With", "XMLHttpRequest")
	} else {
		req = httplib.Post(subRequest.Api).Header(beego.BConfig.WebConfig.Session.SessionNameInHTTPHeader, sid).Header("X-Requested-With", "XMLHttpRequest")

		if len(subRequest.Params) > 0 {
			var params map[string]string
			err := json.Unmarshal([]byte(subRequest.Params), &params)
			if err != nil {
				return SubResponse{
					Api:    subRequest.Api,
					Method: strings.ToLower(subRequest.Method),
					Params: subRequest.Params,
					Error:  "不能解析json参数",
				}
			}
			for k, v := range params {
				req.Param(k, v)
			}
		}
	}
	respStr, err := req.String()
	if err != nil {
		return SubResponse{
			Api:    subRequest.Api,
			Method: strings.ToLower(subRequest.Method),
			Params: subRequest.Params,
			Error:  "不能获取接口返回",
		}
	}

	var response map[string]interface{}
	err = json.Unmarshal([]byte(respStr), &response)
	if err != nil {
		return SubResponse{
			Api:    subRequest.Api,
			Method: strings.ToLower(subRequest.Method),
			Params: subRequest.Params,
			Error:  "解析接口返回不是json",
		}
	}
	return SubResponse{
		Api:      subRequest.Api,
		Method:   strings.ToLower(subRequest.Method),
		Params:   subRequest.Params,
		Response: response,
		Error:    nil,
	}

}

func (c *MainController) Onerequest() {
	status := "SUCCESS"
	var subRequestList []SubRequest
	var subResponseList []SubResponse
	sessionId := c.Ctx.Input.Header(beego.BConfig.WebConfig.Session.SessionNameInHTTPHeader)

	subRequestListJson := c.GetString("subRequestList", "[]")
	err := json.Unmarshal([]byte(subRequestListJson), &subRequestList)
	if err != nil {
		status = "ERROR_PARSE_PARAMS"
	} else {
		for _, subRequest := range subRequestList {
			subResponseList = append(subResponseList, c.processSubRequest(subRequest))
		}
	}

	c.Data["json"] = &map[string]interface{}{"subResponseList": subResponseList, "status": status, "sessionId": sessionId}
	c.ServeJSON()
}
func init() {
	beego.Router("/onerequest", &MainController{}, "post:Onerequest")
}
