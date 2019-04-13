package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tbactl/service/basecode"
	"tbactl/service/dbcomm"
)

const (
	USER_CHARGE = "用户充值"
	FLOW_CHARGE = "charge"
	FLOW_INIT   = "i"
	FLOW_SUCC   = "s"
	FLOW_FAIL   = "f"

	NOW_TIME_FORMAT    = "2006-01-02 15:04:05"
	FIELD_ACCOUNT_BAL  = "Account_bal"
	FIELD_UPDATED_TIME = "Updated_time"

	CODE_SUCC    = "0000"
	CODE_NOEXIST = "1000"
	CODE_FAIL    = "2000"

	RESP_SUCC = "0000"
	RESP_FAIL = "1000"

	EMPTY_STRING = ""

	CODE_TYPE_EDU       = "EDU"
	CODE_TYPE_POSITION  = "POSITION"
	CODE_TYPE_SALARY    = "SALARY"
	CODE_TYPE_WORKYEARS = "WORKYEARS"
	CODE_TYPE_POSICLASS = "POSICLASS"
	CODE_TYPE_REWARDS   = "REWARDS"
)

func PrintHead(a ...interface{}) {
	log.Println("========》", a)
}

func PrintTail(a ...interface{}) {
	log.Println("《========", a)
}

func Write_Response(response interface{}, w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Max-Age", "1728000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "content-type,Action, Module,Authorization")
	fmt.Fprintf(w, string(json))
}

func GetCodeValue(codeType string, codeNo string) string {

	var search basecode.Search
	search.CodeType = codeType
	search.CodeNo = codeNo
	r := basecode.New(dbcomm.GetDB(), basecode.DEBUG)
	e, err := r.Get(search)
	if err != nil {
		log.Println(err)
		return EMPTY_STRING
	}
	return e.CodeVal
}
