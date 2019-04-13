package position

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tbactl/control/common"
	"tbactl/service/dbcomm"
	"tbactl/service/followed_position"
	"tbactl/service/liked_position"
	"tbactl/service/position"
	"time"
)

type PositionResp struct {
	ErrCode   string                  `json:"err_code"`
	ErrMsg    string                  `json:"err_msg"`
	Total     int                     `json:"total"`
	Positions []position.PostPosition `json:"positions"`
}

type ActionResp struct {
	ErrCode string `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

/*
	说明：得到职位列表
	入参：s: 查询条件
	出参：
*/

func GetPositionList(w http.ResponseWriter, req *http.Request) {
	common.PrintHead("GetPositionList")
	var search position.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	r := position.New(dbcomm.GetDB(), position.DEBUG)
	var posiResp PositionResp
	l, err := r.GetList(search)
	tl, err := r.GetTotal(search)
	posiResp.Positions = l
	posiResp.Total = tl
	posiResp.ErrCode = common.RESP_SUCC
	common.PrintTail("GetPositionList")
	common.Write_Response(posiResp, w, req)
}

/*
	说明：增加职位
	入参：
	出参：
*/

func AddPosition(w http.ResponseWriter, req *http.Request) {
	common.PrintHead("AddPosition")
	var form position.Form
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	form.Form.PublishTime = time.Now().Format("2006-01-02 15:04:05")
	form.Form.ExpireTime = time.Now().Add(time.Duration(30 * 24 * time.Hour)).Format("2006-01-02")
	r := position.New(dbcomm.GetDB(), position.DEBUG)
	form.Form.PublishNo = fmt.Sprintf("%d", time.Now().UnixNano())
	form.Form.InsertTime = time.Now().Format("2006-01-02 15:04:05")
	r.InsertEntity(form.Form, nil)
	common.PrintTail("AddPosition")
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
	common.PrintHead("GetPosition")
	var search position.Search
	err := json.NewDecoder(req.Body).Decode(&search)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Search====>", search)
	r := position.New(dbcomm.GetDB(), position.DEBUG)
	e, err := r.Get(search)

	e.SchoolLevel = common.GetCodeValue(common.CODE_TYPE_EDU, e.SchoolLevel)
	e.WorkYears = common.GetCodeValue(common.CODE_TYPE_WORKYEARS, e.WorkYears)

	common.PrintTail("GetPosition")
	common.Write_Response(e, w, req)
}

/*
	说明：申请该职位
	出参：参数1：返回符合条件的对象列表
*/

func LikePosition(w http.ResponseWriter, req *http.Request) {
	common.PrintHead("LikedPosition")
	var form likedposition.Search
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	var resp ActionResp
	log.Println("Form ====>", form)
	r := likedposition.New(dbcomm.GetDB(), likedposition.DEBUG)
	_, err = r.Get(form)
	if err == nil {
		resp.ErrCode = common.RESP_FAIL
		resp.ErrMsg = "已经申请该职位"
		log.Println("已经申请该职位", err)
		common.Write_Response(resp, w, req)
		return
	}
	var e likedposition.LikedPosition
	e.UserId = form.UserId
	e.PublishNo = form.PublishNo
	e.InsertTime = time.Now().Format(common.NOW_TIME_FORMAT)
	err = r.InsertEntity(e, nil)
	if err != nil {
		log.Println("申请职位", err)
		return
	}
	common.PrintTail("LikedPosition")
	resp.ErrCode = common.RESP_FAIL
	resp.ErrMsg = "申请成功"
	common.Write_Response(resp, w, req)
}

/*
	说明：申请该职位
	出参：参数1：返回符合条件的对象列表
*/

func FollowPosition(w http.ResponseWriter, req *http.Request) {
	common.PrintHead("FollowPosition")
	var form followedposition.Search
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	log.Println("Form ====>", form)
	r := followedposition.New(dbcomm.GetDB(), followedposition.DEBUG)
	var resp ActionResp
	_, err = r.Get(form)
	if err == nil {
		resp.ErrCode = common.RESP_FAIL
		resp.ErrMsg = "已经收藏该职位"
		log.Println("已经收藏该职位", err)
		return
	}
	var e followedposition.FollowedPosition
	e.UserId = form.UserId
	e.PublishNo = form.PublishNo
	err = r.InsertEntity(e, nil)
	if err != nil {
		log.Println("申请职位", err)
		return
	}
	common.PrintTail("FollowPosition")
	common.Write_Response("申请成功", w, req)
}
