package flow

import (
	"encoding/json"
	"log"
	"net/http"
	"tbactl/control/common"
	"tbactl/service/dbcomm"
	"tbactl/service/flow"
	"time"
)

/*
	说明：得到我的协议列表
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表
*/

func AddFlow(w http.ResponseWriter, req *http.Request) {
	log.Println("AddFlow===============>")
	var form flows.Form
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("form====>", form)
	r := flows.New(dbcomm.GetDB(), flows.DEBUG)
	form.Form.UserId = time.Now().UnixNano()
	r.InsertEntity(form.Form)
	common.Write_Response("OK", w, req)
}

/*
	说明：得到我的协议列表
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表
*/

func GetFlowsList(w http.ResponseWriter, req *http.Request) {
	log.Println("GetFlowsList===============>")
	var search flows.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Search====>", search)
	r := flows.New(dbcomm.GetDB(), flows.DEBUG)
	list, err := r.GetList(search)
	tl, err := r.GetTotal(search)
	r.Total = tl
	common.Write_Response(list, w, req)
}
