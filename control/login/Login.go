package login

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tbactl/service/account"
	"tbactl/service/dbcomm"
	"tbactl/service/login"
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
	log.Printf("%#v", resp)
	return nil, resp.Openid, resp.Session_key
}

/*
	微信登录
*/
func WxLogin(w http.ResponseWriter, req *http.Request) {
	log.Println("========》WxLogin")
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
	regUser(openId, sessionKey)
	w.Write([]byte(sessionKey))
	log.Println("《========WxLogin")
}

/*
	登记用户注册信息
*/
func regUser(openId string, sessionKey string) error {
	var search account.Search
	search.PartnerUserId = openId
	r := account.New(dbcomm.GetDB(), account.DEBUG)
	if e, err := r.Get(search); err != nil {
		var a account.Account
		a.PartnerUserId = openId
		a.UserId = time.Now().Unix()
		r.InsertEntity(a, nil)
	} else {
		r := login.New(dbcomm.GetDB(), login.DEBUG)
		var l login.Login
		l.UserId = e.UserId
		l.LoginTime = time.Now().Format("2006-01-02 15:04:05")
		l.LoginDesc = "wechat login successful!"
		l.LoginNo = time.Now().Unix()
		r.InsertEntity(l, nil)
	}
	return nil
}
