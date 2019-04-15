package owner

import (
	"encoding/json"
	"fmt"
	"tbactl/control/common"
	"tbactl/service/dbcomm"
	"tbactl/service/owner"

	"log"
	"net/http"
)

type OwnerResp struct {
	ErrCode string      `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Owner   owner.Owner `json:"Owner"`
}

/*
	说明：得到账号的主体信息
	入参：
	出参：
*/

func GetOwner(w http.ResponseWriter, req *http.Request) {
	common.PrintHead("GetOwner")
	var search owner.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	var resp OwnerResp
	log.Println("Search====>", search)
	r := owner.New(dbcomm.GetDB(), owner.DEBUG)
	e, err := r.Get(search)
	if err != nil {
		resp.ErrCode = common.CODE_NOEXIST
		resp.ErrMsg = "账户实体不存在"

	} else {
		resp.ErrCode = common.CODE_SUCC
		resp.ErrMsg = "账户实体存在"
		resp.Owner = *e
		log.Println(resp.Owner.OwnerName)
	}
	common.PrintTail("GetOwner")
	common.Write_Response(resp, w, req)
}

/*
	说明：存储账号实体信息
	入参：
	出参：
*/

func SetOwner(w http.ResponseWriter, req *http.Request) {
	common.PrintHead("SetOwner")
	var form owner.Form
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	r := owner.New(dbcomm.GetDB(), owner.DEBUG)
	var search owner.Search
	search.UserId = form.Form.UserId
	e, err := r.Get(search)
	if err != nil {
		err = r.InsertEntity(form.Form, nil)

	} else {
		err = r.UpdataEntity(fmt.Sprintf("%d", e.Id), form.Form, nil)
	}
	common.PrintHead("SetOwner")
	var resultTip string
	if err != nil {
		resultTip = "更新企业信息失败！"
	} else {
		resultTip = "更新企业信息成功！"
	}

	common.Write_Response(resultTip, w, req)
}
