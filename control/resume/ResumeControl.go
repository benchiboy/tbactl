package resume

import (
	"encoding/json"
	"tbactl/control/common"
	"tbactl/service/dbcomm"
	"tbactl/service/resume"

	"log"
	"net/http"
)

/*
	说明：得到我的协议列表
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表
*/

func GetResumeList(w http.ResponseWriter, req *http.Request) {
	log.Println("GetResumeList===============>")
	var search resume.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Search====>", search)
	r := resume.New(dbcomm.GetDB(), resume.DEBUG)
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

func AddResume(w http.ResponseWriter, req *http.Request) {
	log.Println("AddResume===============>")
	var search resume.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Search====>", search)
	r := resume.New(dbcomm.GetDB(), resume.DEBUG)
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

func EdtResume(w http.ResponseWriter, req *http.Request) {
	log.Println("EdtResume===============>")
	var search resume.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Search====>", search)
	r := resume.New(dbcomm.GetDB(), resume.DEBUG)
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

func GetResume(w http.ResponseWriter, req *http.Request) {
	log.Println("GetResume===============>")
	var search resume.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Search====>", search)
	r := resume.New(dbcomm.GetDB(), resume.DEBUG)
	list, err := r.GetList(search)
	tl, err := r.GetTotal(search)
	r.Total = tl
	common.Write_Response(list, w, req)
}
