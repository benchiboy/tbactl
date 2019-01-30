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
	说明：得到我的协议列表
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表
*/

func AddAccount(w http.ResponseWriter, req *http.Request) {
	log.Println("AddAccount===============>")
	var form account.Form
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("form====>", form)
	r := account.New(dbcomm.GetDB(), account.DEBUG)
	form.Form.UserId = time.Now().UnixNano()
	r.InsertEntity(form.Form)
	common.Write_Response("OK", w, req)
}
