package position

import (
	"encoding/json"
	"log"
	"net/http"
	"tbactl/control/common"
	"tbactl/service/dbcomm"
	"tbactl/service/position"
	//	"time"
)

/*
	说明：得到我的协议列表
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表
*/

func GetPositionList(w http.ResponseWriter, req *http.Request) {
	log.Println("GetPositionList===============>")
	var search position.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Search====>", search)
	r := position.New(dbcomm.GetDB(), position.DEBUG)
	list, err := r.GetList(search)
	tl, err := r.GetTotal(search)
	r.Total = tl
	common.Write_Response(list, w, req)
}

/*
	说明：得到我的协议列表
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表
*/

func AddPosition(w http.ResponseWriter, req *http.Request) {
	log.Println("AddPosition===============>")
	var form position.Form
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("form====>", form)
	r := position.New(dbcomm.GetDB(), position.DEBUG)
	form.Form.PublishNo = "2222"
	r.InsertEntity(form.Form)
	common.Write_Response("OK", w, req)
}

/*
	说明：得到我的协议列表
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表
*/

func EdtPosition(w http.ResponseWriter, req *http.Request) {
	log.Println("EdtPosition===============>")
	var search position.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Search====>", search)
	r := position.New(dbcomm.GetDB(), position.DEBUG)
	list, err := r.GetList(search)
	tl, err := r.GetTotal(search)
	r.Total = tl
	common.Write_Response(list, w, req)
}

/*
	说明：得到我的协议列表
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表
*/

func GetPosition(w http.ResponseWriter, req *http.Request) {
	log.Println("GetPosition===============>")
	var search position.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Search====>", search)
	r := position.New(dbcomm.GetDB(), position.DEBUG)
	l, err := r.Get(search)
	common.Write_Response(l, w, req)
}
