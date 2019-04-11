// rcs_contract_mgr project main.go
package main

import (
	"flag"

	"io"

	"log"
	"net/http"
	"os"
	"tbactl/control/account"
	"tbactl/control/basecode"
	"tbactl/control/flow"
	"tbactl/control/login"
	"tbactl/control/position"
	"tbactl/control/resume"

	"tbactl/service/dbcomm"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	http_srv *http.Server
)

func go_WebServer() {
	log.Println("Listen Service start...")

	http.HandleFunc("/wxLogin", login.WxLogin)

	http.HandleFunc("/getcodeList", basecode.GetBasecodeList)

	http.HandleFunc("/getResumeList", resume.GetResumeList)
	http.HandleFunc("/addResume", resume.AddResume)
	http.HandleFunc("/edtResume", resume.EdtResume)
	http.HandleFunc("/getResume", resume.GetResume)

	http.HandleFunc("/getPositionList", position.GetPositionList)
	http.HandleFunc("/addPosition", position.AddPosition)
	http.HandleFunc("/edtPosition", position.EdtPosition)
	http.HandleFunc("/getPosition", position.GetPosition)

	http.HandleFunc("/getFlowList", flows.GetFlowsList)
	http.HandleFunc("/addFlow", flows.AddFlow)

	http.HandleFunc("/getAccount", account.GetAccount)
	http.HandleFunc("/updateAccount", account.UpdateAccount)

	http_srv = &http.Server{
		Addr: ":8000",
	}
	log.Printf("listen:")
	if err := http_srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}
}

func init() {
	log.Println("System Paras Init......")
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
	log.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "rcs_sync.log",
		MaxSize:    500, // megabytes
		MaxBackups: 50,
		MaxAge:     90, //days
	}))
	envConf := flag.String("env", "config-ci.json", "select a environment config file")
	flag.Parse()
	log.Println("config file ==", *envConf)

}

func main() {
	dbcomm.InitDB()
	go_WebServer()
}
