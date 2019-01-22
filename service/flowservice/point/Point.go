package point

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"tbactl/service/flowservice/applicant"
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
	Id               string `json:"id"`
	Service_point_no string `json:"service_point_no"`
	Flow_no          string `json:"flow_no"`
}

type ServicePointList struct {
	DB            *sql.DB
	Level         int
	Total         int            `json:"total"`
	ServicePoints []ServicePoint `json:"servicePoints"`
}

type ServicePoint struct {
	Id                 string `json:id`
	Flow_no            string `json:flow_no`
	Service_point_no   string `json:service_point_no`
	Service_point_tips string `json:service_point_tips`
	Is_flow_finish     string `json:is_flow_finish`
	Is_reject          string `json:is_reject`
	Enter_condition    string `json:enter_condition`
	Flag               string `json:flag`
	Insert_time        string `json:insert_time`
	Update_time        string `json:update_time`
	Version            string `json:version`
}

/*
	创建产品对象
*/
func New(url string, level int) *ServicePointList {
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

	return &ServicePointList{DB: db, Total: 0, ServicePoints: make([]ServicePoint, 0), Level: level}
}

/*
	得到服务节点
*/
func (r *ServicePointList) GetServicePointList() ([]ServicePoint, error) {
	//var where string
	l := time.Now()

	qrySql := fmt.Sprintf("Select id,flow_no,service_point_no,service_point_tips,is_flow_finish,is_reject,enter_condition,flag,insert_time,update_time,version  from tba_flow_service_points where 1=1 ")
	if r.Level == DEBUG {
		//log.Println(SQL_SELECT, qrySql)
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

	var p ServicePoint
	for rows.Next() {
		rows.Scan(&p.Id, &p.Flow_no, &p.Service_point_no, &p.Service_point_tips, &p.Is_flow_finish, &p.Is_reject, &p.Enter_condition, &p.Flag, &p.Insert_time, &p.Update_time, &p.Version)
		r.ServicePoints = append(r.ServicePoints, p)
	}
	//log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}

	return r.ServicePoints, nil
}

/*
	得到当前输出的服务节点
*/
func (r *ServicePointList) GetServicePointFields(s Search) ([]applicant.ApplicantFields, error) {
	var where string
	l := time.Now()
	if s.Service_point_no != "" {
		where = " and service_point_no='" + s.Service_point_no + "'"
	}
	if s.Flow_no != "" {
		where = " and flow_no='" + s.Flow_no + "'"
	}

	qrySql := fmt.Sprintf("Select a.id,a.table_name,a.field_seq,a.field_name,a.field_desc,a.field_type,a.special_type,a.field_lenth,a.field_dec,a.defult_value,a.code_type,a.is_basic,a.is_interact,a.is_filter,a.is_show,a.is_secret,a.flag,a.insert_time,a.update_time,a.version   from tba_service_point_fields  b ,tba_applicant_fields a where a.field_name=b.field_name  " + where)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, "=======>", qrySql)
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

	ls := make([]applicant.ApplicantFields, 0)
	for rows.Next() {
		var p applicant.ApplicantFields
		rows.Scan(&p.Id, &p.Table_name, &p.Field_seq, &p.Field_name, &p.Field_desc, &p.Field_type, &p.Special_type, &p.Field_lenth, &p.Field_dec, &p.Defult_value, &p.Code_type, &p.Is_basic, &p.Is_interact, &p.Is_filter, &p.Is_show, &p.Is_secret, &p.Flag, &p.Insert_time, &p.Update_time, &p.Version)
		ls = append(ls, p)
	}

	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}

	return ls, nil
}

/*
	得到当前输出的服务节点
*/
func (r *ServicePointList) GetServicePointsMap() (*sync.Map, error) {
	servicePoints, _ := r.GetServicePointList()
	pointMap := new(sync.Map)
	for _, v := range servicePoints {
		s := Search{Service_point_no: v.Service_point_no}
		entity, err := r.GetServicePointFields(s)
		if err != nil {
			log.Panicf(err.Error())
		}
		pointMap.Store(v.Service_point_no, entity)
	}
	return pointMap, nil
}
