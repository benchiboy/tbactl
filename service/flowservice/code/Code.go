package code

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	SQL_INSERT  = "Insert ===>"
	SQL_UPDATE  = "Update ===>"
	SQL_SELECT  = "Select ===>"
	SQL_DELETE  = "Delete ===>"
	SQL_ELAPSED = "Elapsed===>"
	SQL_ERROR   = "Error  ===>"
	SQL_TITLE   = "===================================="
	DEBUG       = 1
	INFO        = 2
)

type Search struct {
	Id         string `json:"id"`
	Table_name string `json:"table_name"`
}

type CodeList struct {
	DB    *sql.DB
	Level int
	Total int    `json:"total"`
	Codes []Code `json:"Code"`
}

type Code struct {
	Id             string `json:id`
	Code_type      string `json:code_type`
	Code_no        string `json:code_no`
	Parent_code_no string `json:parent_code_no`
	Code_val       string `json:code_val`
	Flag           string `json:flag`
	Insert_time    string `json:insert_time`
	Update_time    string `json:update_time`
	Version        string `json:version`
}

/*
	创建产品对象
*/
func New(url string, level int) *CodeList {
	var err error
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Println("Open database error:", err)
		return nil
	}
	if err = db.Ping(); err != nil {
		log.Println("Ping database error:", err)
		return nil
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return &CodeList{DB: db, Total: 0, Codes: make([]Code, 0), Level: level}
}

/*
	得到字段定义的MAP
*/
func (r *CodeList) GetCodeMap(s Search) (*sync.Map, error) {
	l := time.Now()

	qrySql := fmt.Sprintf("select b.code_type from tba_applicant_codes b group by b.code_type")
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, nil
	}

	defer rows.Close()
	var codeType string
	CodeMap := new(sync.Map)

	for rows.Next() {
		rows.Scan(&codeType)
		subSql := fmt.Sprintf("select id,code_type,code_no,parent_code_no,code_val,flag,insert_time,update_time,version from tba_applicant_codes where code_type='" + codeType + "'")
		subRows, err := r.DB.Query(subSql)
		if err != nil {
			log.Println(SQL_ERROR, err.Error())
			break
		}
		defer subRows.Close()

		codes := make([]Code, 0)
		for subRows.Next() {
			var p Code
			subRows.Scan(&p.Id, &p.Code_type, &p.Code_no, &p.Parent_code_no, &p.Code_val, &p.Flag, &p.Insert_time, &p.Update_time, &p.Version)
			codes = append(codes, p)
		}
		CodeMap.Store(codeType, codes)
	}

	log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return CodeMap, nil
}
