package applicant

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

type ApplicantFieldsList struct {
	DB              *sql.DB
	Level           int
	Total           int               `json:"total"`
	ApplicantFields []ApplicantFields `json:"ApplicantFields"`
}

type ApplicantFields struct {
	Id           string `json:id`
	Table_name   string `json:table_name`
	Field_seq    string `json:field_seq`
	Field_name   string `json:field_name`
	Field_desc   string `json:field_desc`
	Field_type   string `json:field_type`
	Special_type string `json:special_type`
	Field_lenth  string `json:field_lenth`
	Field_dec    string `json:field_dec`
	Defult_value string `json:defult_value`
	Code_type    string `json:code_type`
	Is_basic     string `json:is_basic`
	Is_interact  string `json:is_interact`
	Is_filter    string `json:is_filter`
	Is_show      string `json:is_show`
	Is_secret    string `json:is_secret`
	Flag         string `json:flag`
	Insert_time  string `json:insert_time`
	Update_time  string `json:update_time`
	Version      string `json:version`
}

/*
	创建产品对象
*/
func New(url string, level int) *ApplicantFieldsList {
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

	return &ApplicantFieldsList{DB: db, Total: 0, ApplicantFields: make([]ApplicantFields, 0), Level: level}
}

/*
	得到字段定义列表
*/
func (r *ApplicantFieldsList) GetFieldList(s Search) ([]ApplicantFields, error) {
	var where string
	l := time.Now()

	if s.Table_name != "" {
		where = " and table_name='" + s.Table_name + "'"
	}

	qrySql := fmt.Sprintf("Select id,table_name,field_seq,field_name,field_desc,field_type,special_type,field_lenth,field_dec,defult_value,code_type,is_basic,is_interact,is_filter,is_show,is_secret,flag,insert_time,update_time,version  from tba_applicant_fields where 1=1 %s ", where)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	if r.DB == nil {
		log.Println("=====================")
	}
	rows, err := r.DB.Query(qrySql)

	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, nil
	}
	defer rows.Close()

	var p ApplicantFields
	for rows.Next() {
		rows.Scan(&p.Id, &p.Table_name, &p.Field_seq, &p.Field_name, &p.Field_desc, &p.Field_type, &p.Special_type, &p.Field_lenth, &p.Field_dec, &p.Defult_value, &p.Code_type, &p.Is_basic, &p.Is_interact, &p.Is_filter, &p.Is_show, &p.Is_secret, &p.Flag, &p.Insert_time, &p.Update_time, &p.Version)
		r.ApplicantFields = append(r.ApplicantFields, p)
	}
	log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}

	return r.ApplicantFields, nil
}

/*
	得到字段定义的MAP
*/
func (r *ApplicantFieldsList) GetFieldMap(s Search) (*sync.Map, error) {
	var where string
	l := time.Now()

	if s.Table_name != "" {
		where = " and table_name='" + s.Table_name + "'"
	}

	qrySql := fmt.Sprintf("Select id,table_name,field_seq,field_name,field_desc,field_type,special_type,field_lenth,field_dec,defult_value,code_type,is_basic,is_interact,is_filter,is_show,is_secret,flag,insert_time,update_time,version  from tba_applicant_fields where 1=1 %s ", where)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	if r.DB == nil {
		log.Println("=====================")
	}
	rows, err := r.DB.Query(qrySql)

	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, nil
	}
	defer rows.Close()

	ApplicantMap := new(sync.Map)

	for rows.Next() {
		p := new(ApplicantFields)
		rows.Scan(&p.Id, &p.Table_name, &p.Field_seq, &p.Field_name, &p.Field_desc, &p.Field_type, &p.Special_type, &p.Field_lenth, &p.Field_dec, &p.Defult_value, &p.Code_type, &p.Is_basic, &p.Is_interact, &p.Is_filter, &p.Is_show, &p.Is_secret, &p.Flag, &p.Insert_time, &p.Update_time, &p.Version)
		ApplicantMap.Store(p.Field_name, *p)
	}

	log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}

	return ApplicantMap, nil
}
