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
	"tbactl/control/owner"
	"tbactl/control/position"

	"tbactl/control/resume"

	"tbactl/service/dbcomm"

	goconf "github.com/pantsing/goconf"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	http_srv   *http.Server
	dbUrl      string
	listenPort int
	idleConns  int
	openConns  int
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
	http.HandleFunc("/likePosition", position.LikePosition)
	http.HandleFunc("/followPosition", position.FollowPosition)

	http.HandleFunc("/getFlowList", flows.GetFlowsList)
	http.HandleFunc("/addFlow", flows.AddFlow)

	http.HandleFunc("/getOwner", owner.GetOwner)
	http.HandleFunc("/setOwner", owner.SetOwner)

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
		Filename:   "tbactl.log",
		MaxSize:    500, // megabytes
		MaxBackups: 50,
		MaxAge:     90, //days
	}))
	envConf := flag.String("env", "config-ci.json", "select a environment config file")
	flag.Parse()
	log.Println("config file ==", *envConf)
	c, err := goconf.New(*envConf)
	if err != nil {
		log.Fatalln("读配置文件出错", err)
	}

	//填充配置文件
	c.Get("/config/LISTEN_PORT", &listenPort)
	c.Get("/config/DB_URL", &dbUrl)
	c.Get("/config/OPEN_CONNS", &openConns)
	c.Get("/config/IDLE_CONNS", &idleConns)

}

func main() {
	dbcomm.InitDB(dbUrl, idleConns, openConns)
	go_WebServer()
}
