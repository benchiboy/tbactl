package position

import (
	"database/sql"
	"fmt"
	"hcd-gate/service/pubtype"
	"log"
	"strings"
	"time"
)

const (
	SQL_NEWDB   = "NewDB  ===>"
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
	Id            int64  `json:"id"`
	PublishNo     string `json:"publish_no"`
	UserId        int64  `json:"user_id"`
	PositionName  string `json:"position_name"`
	PositionEmps  int64  `json:"position_emps"`
	PositionDesc  string `json:"position_desc"`
	PositionClass string `json:"position_class"`
	PublishTime   string `json:"publish_time"`
	ExpireTime    string `json:"expire_time"`
	Rewards       string `json:"rewards"`

	SalaryMin    float64 `json:"salary_min"`
	SalaryMax    float64 `json:"salary_max"`
	City         string  `json:"city"`
	Area         string  `json:"area"`
	SubArea      string  `json:"sub_area"`
	SchoolLevel  string  `json:"school_level"`
	WorkYears    string  `json:"work_years"`
	WorkAddr     string  `json:"work_addr"`
	ContactName  string  `json:"contact_name"`
	ContactPhone string  `json:"contact_phone"`
	InsertTime   string  `json:"insert_time"`
	UpdateTime   string  `json:"update_time"`
	Version      int64   `json:"version"`
	SearchText   string  `json:"search_text"`

	PageNo     int    `json:"page_no"`
	PageSize   int    `json:"page_size"`
	ExtraWhere string `json:"extra_where"`
	SortFld    string `json:"sort_fld"`
}

type PostPositionList struct {
	DB            *sql.DB
	Level         int
	Total         int            `json:"total"`
	PostPositions []PostPosition `json:"PostPosition"`
}

type PostPosition struct {
	Id            int64   `json:"id"`
	PublishNo     string  `json:"publish_no"`
	UserId        int64   `json:"user_id"`
	PositionName  string  `json:"position_name"`
	PositionEmps  int64   `json:"position_emps"`
	PositionDesc  string  `json:"position_desc"`
	PositionClass string  `json:"position_class"`
	PublishTime   string  `json:"publish_time"`
	ExpireTime    string  `json:"expire_time"`
	Rewards       string  `json:"rewards"`
	SalaryMin     float64 `json:"salary_min"`
	SalaryMax     float64 `json:"salary_max"`
	City          string  `json:"city"`
	Area          string  `json:"area"`
	SubArea       string  `json:"sub_area"`
	SchoolLevel   string  `json:"school_level"`
	WorkYears     string  `json:"work_years"`
	WorkAddr      string  `json:"work_addr"`
	ContactName   string  `json:"contact_name"`
	ContactPhone  string  `json:"contact_phone"`
	InsertTime    string  `json:"insert_time"`
	UpdateTime    string  `json:"update_time"`
	Version       int64   `json:"version"`
}

type Form struct {
	Form PostPosition `json:"PostPosition"`
}

/*
	说明：创建实例对象
	入参：db:数据库sql.DB, 数据库已经连接, level:日志级别
	出参：实例对象
*/

func New(db *sql.DB, level int) *PostPositionList {
	if db == nil {
		log.Println(SQL_SELECT, "Database is nil")
		return nil
	}
	return &PostPositionList{DB: db, Total: 0, PostPositions: make([]PostPosition, 0), Level: level}
}

/*
	说明：创建实例对象
	入参：url:连接数据的url, 数据库还没有CONNECTED, level:日志级别
	出参：实例对象
*/

func NewUrl(url string, level int) *PostPositionList {
	var err error
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Println(SQL_SELECT, "Open database error:", err)
		return nil
	}
	if err = db.Ping(); err != nil {
		log.Println(SQL_SELECT, "Ping database error:", err)
		return nil
	}
	return &PostPositionList{DB: db, Total: 0, PostPositions: make([]PostPosition, 0), Level: level}
}

/*
	说明：得到符合条件的总条数
	入参：s: 查询条件
	出参：参数1：返回符合条件的总条件, 参数2：如果错误返回错误对象
*/

func (r *PostPositionList) GetTotal(s Search) (int, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.PublishNo != "" {
		where += " and publish_no='" + s.PublishNo + "'"
	}

	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}

	if s.PositionName != "" {
		where += " and position_name='" + s.PositionName + "'"
	}

	if s.SearchText != "" {
		where += " and position_name like '%" + s.SearchText + "%'"
	}

	if s.PositionEmps != 0 {
		where += " and position_emps=" + fmt.Sprintf("%d", s.PositionEmps)
	}

	if s.PositionDesc != "" {
		where += " and position_desc='" + s.PositionDesc + "'"
	}

	if s.PositionClass != "" {
		where += " and position_class='" + s.PositionClass + "'"
	}

	if s.PublishTime != "" {
		where += " and publish_time='" + s.PublishTime + "'"
	}

	if s.ExpireTime != "" {
		where += " and expire_time='" + s.ExpireTime + "'"
	}

	if s.Rewards != "" {
		where += " and rewards='" + s.Rewards + "'"
	}

	if s.SalaryMin != 0 {
		where += " and salary_min=" + fmt.Sprintf("%f", s.SalaryMin)
	}

	if s.SalaryMax != 0 {
		where += " and salary_max=" + fmt.Sprintf("%f", s.SalaryMax)
	}

	if s.City != "" {
		where += " and city='" + s.City + "'"
	}

	if s.Area != "" {
		where += " and area='" + s.Area + "'"
	}

	if s.SubArea != "" {
		where += " and sub_area='" + s.SubArea + "'"
	}

	if s.SchoolLevel != "" {
		where += " and school_level='" + s.SchoolLevel + "'"
	}

	if s.WorkYears != "" {
		where += " and work_years='" + s.WorkYears + "'"
	}

	if s.WorkAddr != "" {
		where += " and work_addr='" + s.WorkAddr + "'"
	}

	if s.ContactName != "" {
		where += " and contact_name='" + s.ContactName + "'"
	}

	if s.ContactPhone != "" {
		where += " and contact_phone='" + s.ContactPhone + "'"
	}

	if s.InsertTime != "" {
		where += " and insert_time='" + s.InsertTime + "'"
	}

	if s.UpdateTime != "" {
		where += " and update_time='" + s.UpdateTime + "'"
	}

	if s.Version != 0 {
		where += " and version=" + fmt.Sprintf("%d", s.Version)
	}

	if s.ExtraWhere != "" {
		where += s.ExtraWhere
	}

	qrySql := fmt.Sprintf("Select count(1) as total from tba_account_post_positions   where 1=1 %s", where)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return 0, err
	}
	defer rows.Close()
	var total int
	for rows.Next() {
		rows.Scan(&total)
	}
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return total, nil
}

/*
	说明：根据主键查询符合条件的条数
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象, 参数2：如果错误返回错误对象
*/

func (r PostPositionList) Get(s Search) (*PostPosition, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.PublishNo != "" {
		where += " and publish_no='" + s.PublishNo + "'"
	}

	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}

	if s.PositionName != "" {
		where += " and position_name='" + s.PositionName + "'"
	}

	if s.PositionEmps != 0 {
		where += " and position_emps=" + fmt.Sprintf("%d", s.PositionEmps)
	}

	if s.PositionDesc != "" {
		where += " and position_desc='" + s.PositionDesc + "'"
	}

	if s.PositionClass != "" {
		where += " and position_class='" + s.PositionClass + "'"
	}

	if s.PublishTime != "" {
		where += " and publish_time='" + s.PublishTime + "'"
	}

	if s.ExpireTime != "" {
		where += " and expire_time='" + s.ExpireTime + "'"
	}

	if s.Rewards != "" {
		where += " and rewards='" + s.Rewards + "'"
	}

	if s.SalaryMin != 0 {
		where += " and salary_min=" + fmt.Sprintf("%f", s.SalaryMin)
	}

	if s.SalaryMax != 0 {
		where += " and salary_max=" + fmt.Sprintf("%f", s.SalaryMax)
	}

	if s.City != "" {
		where += " and city='" + s.City + "'"
	}

	if s.Area != "" {
		where += " and area='" + s.Area + "'"
	}

	if s.SubArea != "" {
		where += " and sub_area='" + s.SubArea + "'"
	}

	if s.SchoolLevel != "" {
		where += " and school_level='" + s.SchoolLevel + "'"
	}

	if s.WorkYears != "" {
		where += " and work_years='" + s.WorkYears + "'"
	}

	if s.WorkAddr != "" {
		where += " and work_addr='" + s.WorkAddr + "'"
	}

	if s.ContactName != "" {
		where += " and contact_name='" + s.ContactName + "'"
	}

	if s.ContactPhone != "" {
		where += " and contact_phone='" + s.ContactPhone + "'"
	}

	if s.InsertTime != "" {
		where += " and insert_time='" + s.InsertTime + "'"
	}

	if s.UpdateTime != "" {
		where += " and update_time='" + s.UpdateTime + "'"
	}

	if s.Version != 0 {
		where += " and version=" + fmt.Sprintf("%d", s.Version)
	}

	if s.ExtraWhere != "" {
		where += s.ExtraWhere
	}

	qrySql := fmt.Sprintf("Select id,publish_no,user_id,position_name,position_emps,position_desc,position_class,publish_time,expire_time,rewards,salary_min,salary_max,city,area,sub_area,school_level,work_years,work_addr,contact_name,contact_phone,version from tba_account_post_positions where 1=1 %s ", where)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, err
	}
	defer rows.Close()

	var p PostPosition
	if !rows.Next() {
		return nil, fmt.Errorf("Not Finded Record")
	} else {
		err := rows.Scan(&p.Id, &p.PublishNo, &p.UserId, &p.PositionName, &p.PositionEmps, &p.PositionDesc, &p.PositionClass, &p.PublishTime, &p.ExpireTime, &p.Rewards, &p.SalaryMin, &p.SalaryMax, &p.City, &p.Area, &p.SubArea, &p.SchoolLevel, &p.WorkYears, &p.WorkAddr, &p.ContactName, &p.ContactPhone, &p.Version)
		if err != nil {
			log.Println(SQL_ERROR, err.Error())
			return nil, err
		}
	}
	log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return &p, nil
}

/*
	说明：根据条件查询复核条件对象列表，支持分页查询
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表, 参数2：如果错误返回错误对象
*/

func (r *PostPositionList) GetList(s Search) ([]PostPosition, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.PublishNo != "" {
		where += " and publish_no='" + s.PublishNo + "'"
	}

	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}

	if s.PositionName != "" {
		where += " and position_name='" + s.PositionName + "'"
	}

	if s.SearchText != "" {
		where += " and position_name like '%" + s.SearchText + "%'"
	}

	if s.PositionEmps != 0 {
		where += " and position_emps=" + fmt.Sprintf("%d", s.PositionEmps)
	}

	if s.PositionDesc != "" {
		where += " and position_desc='" + s.PositionDesc + "'"
	}

	if s.PositionClass != "" {
		where += " and position_class='" + s.PositionClass + "'"
	}

	if s.PublishTime != "" {
		where += " and publish_time='" + s.PublishTime + "'"
	}

	if s.ExpireTime != "" {
		where += " and expire_time='" + s.ExpireTime + "'"
	}

	if s.Rewards != "" {
		where += " and rewards='" + s.Rewards + "'"
	}

	if s.SalaryMin != 0 {
		where += " and salary_min=" + fmt.Sprintf("%f", s.SalaryMin)
	}

	if s.SalaryMax != 0 {
		where += " and salary_max=" + fmt.Sprintf("%f", s.SalaryMax)
	}

	if s.City != "" {
		where += " and city='" + s.City + "'"
	}

	if s.Area != "" {
		where += " and area='" + s.Area + "'"
	}

	if s.SubArea != "" {
		where += " and sub_area='" + s.SubArea + "'"
	}

	if s.SchoolLevel != "" {
		where += " and school_level='" + s.SchoolLevel + "'"
	}

	if s.WorkYears != "" {
		where += " and work_years='" + s.WorkYears + "'"
	}

	if s.WorkAddr != "" {
		where += " and work_addr='" + s.WorkAddr + "'"
	}

	if s.ContactName != "" {
		where += " and contact_name='" + s.ContactName + "'"
	}

	if s.ContactPhone != "" {
		where += " and contact_phone='" + s.ContactPhone + "'"
	}

	if s.InsertTime != "" {
		where += " and insert_time='" + s.InsertTime + "'"
	}

	if s.UpdateTime != "" {
		where += " and update_time='" + s.UpdateTime + "'"
	}

	if s.Version != 0 {
		where += " and version=" + fmt.Sprintf("%d", s.Version)
	}

	if s.ExtraWhere != "" {
		where += s.ExtraWhere
	}

	var qrySql string
	if s.PageSize == 0 && s.PageNo == 0 {
		qrySql = fmt.Sprintf("Select id,publish_no,user_id,position_name,position_emps,position_desc,position_class,publish_time,expire_time,rewards,salary_min,salary_max,city,area,sub_area,school_level,work_years,work_addr,contact_name,contact_phone,insert_time,update_time,version from tba_account_post_positions where 1=1 %s", where)
	} else {
		qrySql = fmt.Sprintf("Select id,publish_no,user_id,position_name,position_emps,position_desc,position_class,publish_time,expire_time,rewards,salary_min,salary_max,city,area,sub_area,school_level,work_years,work_addr,contact_name,contact_phone,insert_time,update_time,version from tba_account_post_positions where 1=1 %s Limit %d offset %d", where, s.PageSize, (s.PageNo-1)*s.PageSize)
	}
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, err
	}
	defer rows.Close()

	var p PostPosition
	for rows.Next() {
		rows.Scan(&p.Id, &p.PublishNo, &p.UserId, &p.PositionName, &p.PositionEmps, &p.PositionDesc, &p.PositionClass, &p.PublishTime, &p.ExpireTime, &p.Rewards, &p.SalaryMin, &p.SalaryMax, &p.City, &p.Area, &p.SubArea, &p.SchoolLevel, &p.WorkYears, &p.WorkAddr, &p.ContactName, &p.ContactPhone, &p.InsertTime, &p.UpdateTime, &p.Version)
		r.PostPositions = append(r.PostPositions, p)
	}
	log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return r.PostPositions, nil
}

/*
	说明：根据条件查询复核条件对象列表，支持分页查询
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表, 参数2：如果错误返回错误对象
*/

func (r *PostPositionList) GetListExt(s Search, fList []string) ([][]pubtype.Data, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.PublishNo != "" {
		where += " and publish_no='" + s.PublishNo + "'"
	}

	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}

	if s.PositionName != "" {
		where += " and position_name='" + s.PositionName + "'"
	}

	if s.PositionEmps != 0 {
		where += " and position_emps=" + fmt.Sprintf("%d", s.PositionEmps)
	}

	if s.PositionDesc != "" {
		where += " and position_desc='" + s.PositionDesc + "'"
	}

	if s.PositionClass != "" {
		where += " and position_class='" + s.PositionClass + "'"
	}

	if s.PublishTime != "" {
		where += " and publish_time='" + s.PublishTime + "'"
	}

	if s.ExpireTime != "" {
		where += " and expire_time='" + s.ExpireTime + "'"
	}

	if s.Rewards != "" {
		where += " and rewards='" + s.Rewards + "'"
	}

	if s.SalaryMin != 0 {
		where += " and salary_min=" + fmt.Sprintf("%f", s.SalaryMin)
	}

	if s.SalaryMax != 0 {
		where += " and salary_max=" + fmt.Sprintf("%f", s.SalaryMax)
	}

	if s.City != "" {
		where += " and city='" + s.City + "'"
	}

	if s.Area != "" {
		where += " and area='" + s.Area + "'"
	}

	if s.SubArea != "" {
		where += " and sub_area='" + s.SubArea + "'"
	}

	if s.SchoolLevel != "" {
		where += " and school_level='" + s.SchoolLevel + "'"
	}

	if s.WorkYears != "" {
		where += " and work_years='" + s.WorkYears + "'"
	}

	if s.WorkAddr != "" {
		where += " and work_addr='" + s.WorkAddr + "'"
	}

	if s.ContactName != "" {
		where += " and contact_name='" + s.ContactName + "'"
	}

	if s.ContactPhone != "" {
		where += " and contact_phone='" + s.ContactPhone + "'"
	}

	if s.InsertTime != "" {
		where += " and insert_time='" + s.InsertTime + "'"
	}

	if s.UpdateTime != "" {
		where += " and update_time='" + s.UpdateTime + "'"
	}

	if s.Version != 0 {
		where += " and version=" + fmt.Sprintf("%d", s.Version)
	}

	if s.ExtraWhere != "" {
		where += s.ExtraWhere
	}

	colNames := ""
	for _, v := range fList {
		colNames += v + ","

	}
	colNames = strings.TrimRight(colNames, ",")

	var qrySql string
	if s.PageSize == 0 && s.PageNo == 0 {
		qrySql = fmt.Sprintf("Select %s from tba_account_post_positions where 1=1 %s", colNames, where)
	} else {
		qrySql = fmt.Sprintf("Select %s from tba_account_post_positions where 1=1 %s Limit %d offset %d", colNames, where, s.PageSize, (s.PageNo-1)*s.PageSize)
	}
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, err
	}
	defer rows.Close()

	Columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(Columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	rowData := make([][]pubtype.Data, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		colData := make([]pubtype.Data, 0)
		for k, _ := range values {
			d := new(pubtype.Data)
			d.FieldName = Columns[k]
			d.FieldValue = string(values[k])
			colData = append(colData, *d)
		}
		//extra flow_batch_id
		d2 := new(pubtype.Data)
		d2.FieldName = "flow_batch_id"
		d2.FieldValue = string(values[0])
		colData = append(colData, *d2)

		rowData = append(rowData, colData)
	}

	log.Println(SQL_ELAPSED, "==========>>>>>>>>>>>", rowData)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return rowData, nil
}

/*
	说明：根据主键查询符合条件的记录，并保持成MAP
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象, 参数2：如果错误返回错误对象
*/

func (r *PostPositionList) GetExt(s Search) (map[string]string, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.PublishNo != "" {
		where += " and publish_no='" + s.PublishNo + "'"
	}

	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}

	if s.PositionName != "" {
		where += " and position_name='" + s.PositionName + "'"
	}

	if s.PositionEmps != 0 {
		where += " and position_emps=" + fmt.Sprintf("%d", s.PositionEmps)
	}

	if s.PositionDesc != "" {
		where += " and position_desc='" + s.PositionDesc + "'"
	}

	if s.PositionClass != "" {
		where += " and position_class='" + s.PositionClass + "'"
	}

	if s.PublishTime != "" {
		where += " and publish_time='" + s.PublishTime + "'"
	}

	if s.ExpireTime != "" {
		where += " and expire_time='" + s.ExpireTime + "'"
	}

	if s.Rewards != "" {
		where += " and rewards='" + s.Rewards + "'"
	}

	if s.SalaryMin != 0 {
		where += " and salary_min=" + fmt.Sprintf("%f", s.SalaryMin)
	}

	if s.SalaryMax != 0 {
		where += " and salary_max=" + fmt.Sprintf("%f", s.SalaryMax)
	}

	if s.City != "" {
		where += " and city='" + s.City + "'"
	}

	if s.Area != "" {
		where += " and area='" + s.Area + "'"
	}

	if s.SubArea != "" {
		where += " and sub_area='" + s.SubArea + "'"
	}

	if s.SchoolLevel != "" {
		where += " and school_level='" + s.SchoolLevel + "'"
	}

	if s.WorkYears != "" {
		where += " and work_years='" + s.WorkYears + "'"
	}

	if s.WorkAddr != "" {
		where += " and work_addr='" + s.WorkAddr + "'"
	}

	if s.ContactName != "" {
		where += " and contact_name='" + s.ContactName + "'"
	}

	if s.ContactPhone != "" {
		where += " and contact_phone='" + s.ContactPhone + "'"
	}

	if s.InsertTime != "" {
		where += " and insert_time='" + s.InsertTime + "'"
	}

	if s.UpdateTime != "" {
		where += " and update_time='" + s.UpdateTime + "'"
	}

	if s.Version != 0 {
		where += " and version=" + fmt.Sprintf("%d", s.Version)
	}

	qrySql := fmt.Sprintf("Select id,publish_no,user_id,position_name,position_emps,position_desc,position_class,publish_time,expire_time,rewards,salary_min,salary_max,city,area,sub_area,school_level,work_years,work_addr,contact_name,contact_phone,insert_time,update_time,version from tba_account_post_positions where 1=1 %s ", where)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, err
	}
	defer rows.Close()

	Columns, _ := rows.Columns()

	values := make([]sql.RawBytes, len(Columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	if !rows.Next() {
		return nil, fmt.Errorf("Not Finded Record")
	} else {
		err = rows.Scan(scanArgs...)
	}

	fldValMap := make(map[string]string)
	for k, v := range Columns {
		fldValMap[v] = string(values[k])
	}

	log.Println(SQL_ELAPSED, "==========>>>>>>>>>>>", fldValMap)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return fldValMap, nil

}

/*
	说明：插入对象到数据表中，这个方法要求对象的各个属性必须赋值
	入参：p:插入的对象
	出参：参数1：如果出错，返回错误对象；成功返回nil
*/

func (r PostPositionList) Insert(p PostPosition) error {
	l := time.Now()
	exeSql := fmt.Sprintf("Insert into  tba_account_post_positions(publish_no,user_id,position_name,position_emps,position_desc,position_class,publish_time,expire_time,rewards,salary_min,salary_max,city,area,sub_area,school_level,work_years,work_addr,contact_name,contact_phone,version)  values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if r.Level == DEBUG {
		log.Println(SQL_INSERT, exeSql)
	}
	_, err := r.DB.Exec(exeSql, p.PublishNo, p.UserId, p.PositionName, p.PositionEmps, p.PositionDesc, p.PositionClass, p.PublishTime, p.ExpireTime, p.Rewards, p.SalaryMin, p.SalaryMax, p.City, p.Area, p.SubArea, p.SchoolLevel, p.WorkYears, p.WorkAddr, p.ContactName, p.ContactPhone, p.Version)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return err
	}
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return nil
}

/*
	说明：插入对象到数据表中，这个方法会判读对象的各个属性，如果属性不为空，才加入插入列中；
	入参：p:插入的对象
	出参：参数1：如果出错，返回错误对象；成功返回nil
*/

func (r PostPositionList) InsertEntity(p PostPosition, tr *sql.Tx) error {
	l := time.Now()
	var colNames, colTags string
	valSlice := make([]interface{}, 0)

	if p.PublishNo != "" {
		colNames += "publish_no,"
		colTags += "?,"
		valSlice = append(valSlice, p.PublishNo)
	}

	if p.UserId != 0 {
		colNames += "user_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.UserId)
	}

	if p.PositionName != "" {
		colNames += "position_name,"
		colTags += "?,"
		valSlice = append(valSlice, p.PositionName)
	}

	if p.PositionEmps != 0 {
		colNames += "position_emps,"
		colTags += "?,"
		valSlice = append(valSlice, p.PositionEmps)
	}

	if p.PositionDesc != "" {
		colNames += "position_desc,"
		colTags += "?,"
		valSlice = append(valSlice, p.PositionDesc)
	}

	if p.PositionClass != "" {
		colNames += "position_class,"
		colTags += "?,"
		valSlice = append(valSlice, p.PositionClass)
	}

	if p.PublishTime != "" {
		colNames += "publish_time,"
		colTags += "?,"
		valSlice = append(valSlice, p.PublishTime)
	}

	if p.ExpireTime != "" {
		colNames += "expire_time,"
		colTags += "?,"
		valSlice = append(valSlice, p.ExpireTime)
	}

	if p.Rewards != "" {
		colNames += "rewards,"
		colTags += "?,"
		valSlice = append(valSlice, p.Rewards)
	}

	if p.SalaryMin != 0.00 {
		colNames += "salary_min,"
		colTags += "?,"
		valSlice = append(valSlice, p.SalaryMin)
	}

	if p.SalaryMax != 0.00 {
		colNames += "salary_max,"
		colTags += "?,"
		valSlice = append(valSlice, p.SalaryMax)
	}

	if p.City != "" {
		colNames += "city,"
		colTags += "?,"
		valSlice = append(valSlice, p.City)
	}

	if p.Area != "" {
		colNames += "area,"
		colTags += "?,"
		valSlice = append(valSlice, p.Area)
	}

	if p.SubArea != "" {
		colNames += "sub_area,"
		colTags += "?,"
		valSlice = append(valSlice, p.SubArea)
	}

	if p.SchoolLevel != "" {
		colNames += "school_level,"
		colTags += "?,"
		valSlice = append(valSlice, p.SchoolLevel)
	}

	if p.WorkYears != "" {
		colNames += "work_years,"
		colTags += "?,"
		valSlice = append(valSlice, p.WorkYears)
	}

	if p.WorkAddr != "" {
		colNames += "work_addr,"
		colTags += "?,"
		valSlice = append(valSlice, p.WorkAddr)
	}

	if p.ContactName != "" {
		colNames += "contact_name,"
		colTags += "?,"
		valSlice = append(valSlice, p.ContactName)
	}

	if p.ContactPhone != "" {
		colNames += "contact_phone,"
		colTags += "?,"
		valSlice = append(valSlice, p.ContactPhone)
	}

	if p.Version != 0 {
		colNames += "version,"
		colTags += "?,"
		valSlice = append(valSlice, p.Version)
	}

	colNames = strings.TrimRight(colNames, ",")
	colTags = strings.TrimRight(colTags, ",")
	exeSql := fmt.Sprintf("Insert into  tba_account_post_positions(%s)  values(%s)", colNames, colTags)
	if r.Level == DEBUG {
		log.Println(SQL_INSERT, exeSql)
	}

	var stmt *sql.Stmt
	var err error
	if tr == nil {
		stmt, err = r.DB.Prepare(exeSql)
	} else {
		stmt, err = tr.Prepare(exeSql)
	}
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(valSlice...)
	if err != nil {
		log.Println(SQL_INSERT, "Insert data error: %v\n", err)
		return err
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		log.Println(SQL_INSERT, "LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		log.Println(SQL_INSERT, "RowsAffected:", RowsAffected)
	}

	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return nil
}

/*
	说明：插入一个MAP到数据表中；
	入参：m:插入的Map
	出参：参数1：如果出错，返回错误对象；成功返回nil
*/

func (r PostPositionList) InsertMap(m map[string]interface{}, tr *sql.Tx) error {
	l := time.Now()
	var colNames, colTags string
	valSlice := make([]interface{}, 0)
	for k, v := range m {
		colNames += k + ","
		colTags += "?,"
		valSlice = append(valSlice, v)
	}
	colNames = strings.TrimRight(colNames, ",")
	colTags = strings.TrimRight(colTags, ",")

	exeSql := fmt.Sprintf("Insert into  tba_account_post_positions(%s)  values(%s)", colNames, colTags)
	if r.Level == DEBUG {
		log.Println(SQL_INSERT, exeSql)
	}

	var stmt *sql.Stmt
	var err error
	if tr == nil {
		stmt, err = r.DB.Prepare(exeSql)
	} else {
		stmt, err = tr.Prepare(exeSql)
	}

	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(valSlice...)
	if err != nil {
		log.Println(SQL_INSERT, "insert data error: %v\n", err)
		return err
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		log.Println(SQL_INSERT, "LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		log.Println(SQL_INSERT, "RowsAffected:", RowsAffected)
	}

	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return nil
}

/*
	说明：插入对象到数据表中，这个方法会判读对象的各个属性，如果属性不为空，才加入插入列中；
	入参：p:插入的对象
	出参：参数1：如果出错，返回错误对象；成功返回nil
*/

func (r PostPositionList) UpdataEntity(keyNo string, p PostPosition, tr *sql.Tx) error {
	l := time.Now()
	var colNames string
	valSlice := make([]interface{}, 0)

	if p.Id != 0 {
		colNames += "id=?,"
		valSlice = append(valSlice, p.Id)
	}

	if p.PublishNo != "" {
		colNames += "publish_no=?,"

		valSlice = append(valSlice, p.PublishNo)
	}

	if p.UserId != 0 {
		colNames += "user_id=?,"
		valSlice = append(valSlice, p.UserId)
	}

	if p.PositionName != "" {
		colNames += "position_name=?,"

		valSlice = append(valSlice, p.PositionName)
	}

	if p.PositionEmps != 0 {
		colNames += "position_emps=?,"
		valSlice = append(valSlice, p.PositionEmps)
	}

	if p.PositionDesc != "" {
		colNames += "position_desc=?,"

		valSlice = append(valSlice, p.PositionDesc)
	}

	if p.PositionClass != "" {
		colNames += "position_class=?,"

		valSlice = append(valSlice, p.PositionClass)
	}

	if p.PublishTime != "" {
		colNames += "publish_time=?,"

		valSlice = append(valSlice, p.PublishTime)
	}

	if p.ExpireTime != "" {
		colNames += "expire_time=?,"

		valSlice = append(valSlice, p.ExpireTime)
	}

	if p.Rewards != "" {
		colNames += "rewards=?,"

		valSlice = append(valSlice, p.Rewards)
	}

	if p.SalaryMin != 0.00 {
		colNames += "salary_min=?,"
		valSlice = append(valSlice, p.SalaryMin)
	}

	if p.SalaryMax != 0.00 {
		colNames += "salary_max=?,"
		valSlice = append(valSlice, p.SalaryMax)
	}

	if p.City != "" {
		colNames += "city=?,"

		valSlice = append(valSlice, p.City)
	}

	if p.Area != "" {
		colNames += "area=?,"

		valSlice = append(valSlice, p.Area)
	}

	if p.SubArea != "" {
		colNames += "sub_area=?,"

		valSlice = append(valSlice, p.SubArea)
	}

	if p.SchoolLevel != "" {
		colNames += "school_level=?,"

		valSlice = append(valSlice, p.SchoolLevel)
	}

	if p.WorkYears != "" {
		colNames += "work_years=?,"

		valSlice = append(valSlice, p.WorkYears)
	}

	if p.WorkAddr != "" {
		colNames += "work_addr=?,"

		valSlice = append(valSlice, p.WorkAddr)
	}

	if p.ContactName != "" {
		colNames += "contact_name=?,"

		valSlice = append(valSlice, p.ContactName)
	}

	if p.ContactPhone != "" {
		colNames += "contact_phone=?,"

		valSlice = append(valSlice, p.ContactPhone)
	}

	if p.InsertTime != "" {
		colNames += "insert_time=?,"

		valSlice = append(valSlice, p.InsertTime)
	}

	if p.UpdateTime != "" {
		colNames += "update_time=?,"

		valSlice = append(valSlice, p.UpdateTime)
	}

	if p.Version != 0 {
		colNames += "version=?,"
		valSlice = append(valSlice, p.Version)
	}

	colNames = strings.TrimRight(colNames, ",")
	valSlice = append(valSlice, keyNo)

	exeSql := fmt.Sprintf("update  tba_account_post_positions  set %s  where id=? ", colNames)
	if r.Level == DEBUG {
		log.Println(SQL_INSERT, exeSql)
	}

	var stmt *sql.Stmt
	var err error
	if tr == nil {
		stmt, err = r.DB.Prepare(exeSql)
	} else {
		stmt, err = tr.Prepare(exeSql)
	}

	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(valSlice...)
	if err != nil {
		log.Println(SQL_INSERT, "Update data error: %v\n", err)
		return err
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		log.Println(SQL_INSERT, "LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		log.Println(SQL_INSERT, "RowsAffected:", RowsAffected)
	}

	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return nil
}

/*
	说明：根据更新主键及更新Map值更新数据表；
	入参：keyNo:更新数据的关键条件，m:更新数据列的Map
	出参：参数1：如果出错，返回错误对象；成功返回nil
*/

func (r PostPositionList) UpdateMap(keyNo string, m map[string]interface{}, tr *sql.Tx) error {
	l := time.Now()

	var colNames string
	valSlice := make([]interface{}, 0)
	for k, v := range m {
		colNames += k + "=?,"
		valSlice = append(valSlice, v)
	}
	valSlice = append(valSlice, keyNo)
	colNames = strings.TrimRight(colNames, ",")
	updateSql := fmt.Sprintf("Update tba_account_post_positions set %s where id=?", colNames)
	if r.Level == DEBUG {
		log.Println(SQL_UPDATE, updateSql)
	}
	var stmt *sql.Stmt
	var err error
	if tr == nil {
		stmt, err = r.DB.Prepare(updateSql)
	} else {
		stmt, err = tr.Prepare(updateSql)
	}

	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return err
	}
	ret, err := stmt.Exec(valSlice...)
	if err != nil {
		log.Println(SQL_UPDATE, "Update data error: %v\n", err)
		return err
	}
	defer stmt.Close()

	if LastInsertId, err := ret.LastInsertId(); nil == err {
		log.Println(SQL_UPDATE, "LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		log.Println(SQL_UPDATE, "RowsAffected:", RowsAffected)
	}
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return nil
}

/*
	说明：根据主键删除一条数据；
	入参：keyNo:要删除的主键值
	出参：参数1：如果出错，返回错误对象；成功返回nil
*/

func (r PostPositionList) Delete(keyNo string, tr *sql.Tx) error {
	l := time.Now()
	delSql := fmt.Sprintf("Delete from  tba_account_post_positions  where id=?")
	if r.Level == DEBUG {
		log.Println(SQL_UPDATE, delSql)
	}

	var stmt *sql.Stmt
	var err error
	if tr == nil {
		stmt, err = r.DB.Prepare(delSql)
	} else {
		stmt, err = tr.Prepare(delSql)
	}

	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return err
	}
	ret, err := stmt.Exec(keyNo)
	if err != nil {
		log.Println(SQL_DELETE, "Delete error: %v\n", err)
		return err
	}
	defer stmt.Close()

	if LastInsertId, err := ret.LastInsertId(); nil == err {
		log.Println(SQL_DELETE, "LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		log.Println(SQL_DELETE, "RowsAffected:", RowsAffected)
	}
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return nil
}
