package account

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"tbactl/control/common"
	"tbactl/service/account"
	"tbactl/service/dbcomm"
)

type EncryptedDataUserInfo struct {
	OpenID    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	UnionID   string `json:"unionId"`
	Watermark struct {
		Timestamp int    `json:"timestamp"`
		Appid     string `json:"appid"`
	} `json:"watermark"`
}

/*
	说明：得到账户信息
	出参：参数1：返回符合条件的对象列表
*/

func GetAccount(w http.ResponseWriter, req *http.Request) {
	common.PrintHead("GetAccount")
	var form account.Form
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	var search account.Search
	search.UserId = form.Form.UserId
	r := account.New(dbcomm.GetDB(), account.DEBUG)
	e, _ := r.Get(search)
	common.PrintHead("GetAccount")
	common.Write_Response(e, w, req)
}

/*
	说明：更新账号信息
	出参：参数1：返回符合条件的对象列表
*/

func UpdateAccount(w http.ResponseWriter, req *http.Request) {
	common.PrintHead("UpdateAccount")
	var form account.Form
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Printf("%#v====>", form)
	var search account.Search
	search.UserId = form.Form.UserId
	r := account.New(dbcomm.GetDB(), account.DEBUG)
	if e, err := r.Get(search); err == nil {
		u := getWechatUserInfo(form.Form.EncryptedData, form.Form.Iv, e.SessionKey)
		e.AvatarUrl = u.AvatarURL
		e.Province = u.Province
		e.City = u.City
		e.Country = u.Country
		e.Language = u.Language
		e.Gender = u.Gender
		r.UpdataEntity(fmt.Sprintf("%d", e.Id), *e, nil)
	} else {
		r.InsertEntity(form.Form, nil)
	}
	common.Write_Response("OK", w, req)
	common.PrintTail("UpdateAccount")
}

/*
	说明：得到微信的基本信息
	出参：参数1：返回符合条件的对象列表
*/

func getWechatUserInfo(inEncryptedData string, inIv string, inSessionKey string) *EncryptedDataUserInfo {
	common.PrintHead("getWechatUserInfo")
	encryptedData, _ := base64.StdEncoding.DecodeString(inEncryptedData)
	iv, _ := base64.StdEncoding.DecodeString(inIv)
	sessionKey, _ := base64.StdEncoding.DecodeString(inSessionKey)

	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err := aes.NewCipher([]byte(sessionKey))
	if err != nil {
		return nil
	}
	decrypted := make([]byte, len(encryptedData))
	aesDecrypter := cipher.NewCBCDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.CryptBlocks(decrypted, encryptedData)
	var userInfo EncryptedDataUserInfo
	t := string(decrypted)
	fmt.Println(t)
	total := strings.Index(t, "}}") + 2
	err = json.Unmarshal(decrypted[:total], &userInfo)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Println(userInfo.OpenID)
	common.PrintTail("getWechatUserInfo")
	return &userInfo

}
