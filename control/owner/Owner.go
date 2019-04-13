package owner

import (
	"encoding/json"
	"tbactl/control/common"
	"tbactl/service/dbcomm"
	"tbactl/service/owner"

	"log"
	"net/http"
)

type OwnerResp struct {
	ErrCode string `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
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
	}
	common.PrintTail("GetOwner")
	common.Write_Response(e, w, req)
}
