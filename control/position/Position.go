package position

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tbactl/control/common"
	"tbactl/service/dbcomm"
	"tbactl/service/position"
	"time"
)

/*
	说明：得到职位列表
	入参：s: 查询条件
	出参：
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
	说明：增加职位
	入参：
	出参：
*/

func AddPosition(w http.ResponseWriter, req *http.Request) {
	log.Println("========》AddPosition")
	var form position.Form
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("form====>", form)
	form.Form.PublishTime = time.Now().Format("2006-01-02 15:04:05")
	form.Form.ExpireTime = time.Now().Add(time.Duration(30 * 24 * time.Hour)).Format("2006-01-02")
	r := position.New(dbcomm.GetDB(), position.DEBUG)
	form.Form.PublishNo = fmt.Sprintf("%d", time.Now().UnixNano())
	form.Form.InsertTime = time.Now().Format("2006-01-02 15:04:05")
	r.InsertEntity(form.Form)

	log.Println("《========AddPosition")
	common.Write_Response("OK", w, req)
}

/*
	说明：修改职位信息
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
	说明：得到某个职位信息
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
