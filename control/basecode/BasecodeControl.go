package basecode

import (
	"encoding/json"
	"tbactl/control/common"
	"tbactl/service/basecode"
	"tbactl/service/dbcomm"

	"log"
	"net/http"
)

/*
	说明：得到我的协议列表
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表
*/

func GetBasecodeList(w http.ResponseWriter, req *http.Request) {
	log.Println("GetBasecodeList===============>")
	var search basecode.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Search====>", search)
	r := basecode.New(dbcomm.GetDB(), basecode.DEBUG)
	list, err := r.GetList(search)
	tl, err := r.GetTotal(search)
	r.Total = tl
	common.Write_Response(list, w, req)
}
