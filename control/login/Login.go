package login

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
 */
type Login struct {
	Code string `json:"code"`
}

/*
 */
type LoginResp struct {
	Openid      string `json:"openid"`
	Session_key string `json:"session_key"`
	Unionid     string `json:"unionid"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

/*
	根据CODE 得到OPENID
*/
func wxGetOpenid(code string) (error, string, string) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	getUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=wxcc7ef55685a5221c&secret=4d53e212c52cd1955703cf45600f7472&js_code=" + code + "&grant_type=authorization_code"
	res, err := httpClient.Get(getUrl)
	if err != nil {
		return fmt.Errorf("访问微信认证服务出错！"), "", ""
	}
	defer res.Body.Close()

	var resp LoginResp
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return fmt.Errorf("解析JSON出错"), "", ""
	}
	return nil, resp.Openid, resp.Session_key
}

/*
	微信登录
*/
func WxLogin(w http.ResponseWriter, req *http.Request) {
	log.Println("WxLogin===============>")
	keys, ok := req.URL.Query()["code"]
	if !ok || len(keys) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	code := keys[0]
	err, openId, sessionKey := wxGetOpenid(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("=====>", openId)
	w.Write([]byte(sessionKey))
}
