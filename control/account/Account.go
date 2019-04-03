package account

import (
	"encoding/json"
	"log"
	"net/http"
	"tbactl/control/common"
	"tbactl/service/account"
	"tbactl/service/dbcomm"
	"time"
)

/*
	说明：添加账号信息
	出参：参数1：返回符合条件的对象列表
*/

func AddAccount(w http.ResponseWriter, req *http.Request) {
	log.Println("<========AddAccount========>")
	var form account.Account
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("form====>", form)
	r := account.New(dbcomm.GetDB(), account.DEBUG)
	form.UserId = time.Now().UnixNano()
	r.InsertEntity(form, nil)
	common.Write_Response("OK", w, req)
}

/*
	说明：更新账号信息
	出参：参数1：返回符合条件的对象列表
*/

func UpdateAccount(w http.ResponseWriter, req *http.Request) {
	log.Println("========》UpdateAccount")
	var form account.Account
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("form====>", form)
	r := account.New(dbcomm.GetDB(), account.DEBUG)
	form.UserId = time.Now().UnixNano()
	r.InsertEntity(form, nil)
	common.Write_Response("OK", w, req)
	log.Println("《========UpdateAccount")
}
