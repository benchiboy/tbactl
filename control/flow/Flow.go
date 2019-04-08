package flows

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tbactl/control/common"
	"tbactl/service/account"
	"tbactl/service/dbcomm"
	"tbactl/service/flow"
	"time"
)

/*
	说明：增加交易流水
	入参：
	出参：
*/

func AddFlow(w http.ResponseWriter, req *http.Request) {
	common.PrintHead("AddFlow")
	var form flows.Form
	err := json.NewDecoder(req.Body).Decode(&form)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	var balance float64
	var search account.Search
	search.UserId = form.Form.UserId
	r := account.New(dbcomm.GetDB(), account.DEBUG)
	tr, err := r.DB.Begin()
	if err != nil {
		log.Println(err)
		return
	}
	if e, err := r.Get(search); err == nil {
		balance = e.AccountBal + form.Form.TrxnAmt
		keyVal := map[string]interface{}{common.FIELD_ACCOUNT_BAL: balance, common.FIELD_UPDATED_TIME: time.Now().Format(common.NOW_TIME_FORMAT)}
		if err := r.UpdateMap(fmt.Sprintf("%d", e.Id), keyVal, tr); err != nil {
			log.Println(err)
			tr.Rollback()
		}
	}
	rr := flows.New(r.DB, flows.DEBUG)
	form.Form.TrxnNo = time.Now().UnixNano()
	form.Form.TrxnMemo = common.USER_CHARGE
	form.Form.ProcStatus = common.FLOW_INIT
	form.Form.TrxnType = common.FLOW_CHARGE
	form.Form.InsertTime = time.Now().Format(common.NOW_TIME_FORMAT)
	form.Form.TrxnDate = time.Now().Format(common.NOW_TIME_FORMAT)
	form.Form.AccountBal = balance
	if err := rr.InsertEntity(form.Form, tr); err != nil {
		log.Println(err)
		tr.Rollback()
	}
	tr.Commit()
	common.PrintTail("AddFlow")
	common.Write_Response("OK", w, req)
}

/*
	说明：得到交易流水列表
	入参：
	出参：
*/

func GetFlowsList(w http.ResponseWriter, req *http.Request) {
	common.PrintHead("GetFlowsList")
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
	common.PrintTail("GetFlowsList")
	common.Write_Response(list, w, req)
}
