package basecode

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	Id           int    `json:"id"`
	CodeType     string `json:"code_type"`
	CodeNo       string `json:"code_no"`
	ParentCodeNo string `json:"parent_code_no"`
	CodeVal      string `json:"code_val"`
	Flag         int    `json:"flag"`
	InsertTime   string `json:"insert_time"`
	UpdateTime   string `json:"update_time"`
	Version      int    `json:"version"`
	PageNo       int    `json:"page_no"`
	PageSize     int    `json:"page_size"`
	SortFld      string `json:"sort_fld"`
}

type BasecodeList struct {
	DB        *sql.DB
	Level     int
	Total     int        `json:"total"`
	Basecodes []Basecode `json:"Basecode"`
}

type Basecode struct {
	Id           int    `json:"id"`
	CodeType     string `json:"code_type"`
	CodeNo       string `json:"code_no"`
	ParentCodeNo string `json:"parent_code_no"`
	CodeVal      string `json:"code_val"`
	Flag         int    `json:"flag"`
	InsertTime   string `json:"insert_time"`
	UpdateTime   string `json:"update_time"`
	Version      int    `json:"version"`
}

/*
	说明：创建实例对象
	入参：db:数据库sql.DB, 数据库已经连接, level:日志级别
	出参：实例对象
*/

func New(db *sql.DB, level int) *BasecodeList {
	if db == nil {
		log.Println(SQL_SELECT, "Database is nil")
		return nil
	}
	return &BasecodeList{DB: db, Total: 0, Basecodes: make([]Basecode, 0), Level: level}
}

/*
	说明：创建实例对象
	入参：url:连接数据的url, 数据库还没有CONNECTED, level:日志级别
	出参：实例对象
*/

func NewUrl(url string, level int) *BasecodeList {
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
	return &BasecodeList{DB: db, Total: 0, Basecodes: make([]Basecode, 0), Level: level}
}

/*
	说明：得到符合条件的总条数
	入参：s: 查询条件
	出参：参数1：返回符合条件的总条件, 参数2：如果错误返回错误对象
*/

func (r *BasecodeList) GetTotal(s Search) (int, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.CodeType != "" {
		where += " and code_type='" + s.CodeType + "'"
	}

	if s.CodeNo != "" {
		where += " and code_no='" + s.CodeNo + "'"
	}

	if s.ParentCodeNo != "" {
		where += " and parent_code_no='" + s.ParentCodeNo + "'"
	}

	if s.CodeVal != "" {
		where += " and code_val='" + s.CodeVal + "'"
	}

	if s.Flag != 0 {
		where += " and flag=" + fmt.Sprintf("%d", s.Flag)
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

	qrySql := fmt.Sprintf("Select count(1) as total from tba_base_codes   where 1=1 %s", where)
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

func (r BasecodeList) Get(s Search) (*Basecode, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.CodeType != "" {
		where += " and code_type='" + s.CodeType + "'"
	}

	if s.CodeNo != "" {
		where += " and code_no='" + s.CodeNo + "'"
	}

	if s.ParentCodeNo != "" {
		where += " and parent_code_no='" + s.ParentCodeNo + "'"
	}

	if s.CodeVal != "" {
		where += " and code_val='" + s.CodeVal + "'"
	}

	if s.Flag != 0 {
		where += " and flag=" + fmt.Sprintf("%d", s.Flag)
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

	qrySql := fmt.Sprintf("Select id,code_type,code_no,parent_code_no,code_val,flag,insert_time,update_time,version from tba_base_codes where 1=1 %s Limit %d offset %d", where, s.PageSize, (s.PageNo-1)*s.PageSize)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, nil
	}
	defer rows.Close()

	var p Basecode
	for rows.Next() {
		rows.Scan(&p.Id, &p.CodeType, &p.CodeNo, &p.ParentCodeNo, &p.CodeVal, &p.Flag, &p.InsertTime, &p.UpdateTime, &p.Version)
		break
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

func (r *BasecodeList) GetList(s Search) ([]Basecode, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.CodeType != "" {
		where += " and code_type='" + s.CodeType + "'"
	}

	if s.CodeNo != "" {
		where += " and code_no='" + s.CodeNo + "'"
	}

	if s.ParentCodeNo != "" {
		where += " and parent_code_no='" + s.ParentCodeNo + "'"
	}

	if s.CodeVal != "" {
		where += " and code_val='" + s.CodeVal + "'"
	}

	if s.Flag != 0 {
		where += " and flag=" + fmt.Sprintf("%d", s.Flag)
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

	var qrySql string
	if s.PageSize == 0 && s.PageNo == 0 {
		qrySql = fmt.Sprintf("Select id,code_type,code_no,parent_code_no,code_val,flag,insert_time,update_time,version from tba_base_codes where 1=1 %s", where)
	} else {
		qrySql = fmt.Sprintf("Select id,code_type,code_no,parent_code_no,code_val,flag,insert_time,update_time,version from tba_base_codes where 1=1 %s Limit %d offset %d", where, s.PageSize, (s.PageNo-1)*s.PageSize)
	}
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, nil
	}
	defer rows.Close()

	var p Basecode
	for rows.Next() {
		rows.Scan(&p.Id, &p.CodeType, &p.CodeNo, &p.ParentCodeNo, &p.CodeVal, &p.Flag, &p.InsertTime, &p.UpdateTime, &p.Version)
		r.Basecodes = append(r.Basecodes, p)
	}
	log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return r.Basecodes, nil
}

/*
	说明：插入对象到数据表中，这个方法要求对象的各个属性必须赋值
	入参：p:插入的对象
	出参：参数1：如果出错，返回错误对象；成功返回nil
*/

func (r BasecodeList) Insert(p Basecode) error {
	l := time.Now()
	exeSql := fmt.Sprintf("Insert into  tba_base_codes(code_type,code_no,parent_code_no,code_val,flag,version)  values(?,?,?,?,?,?,?,?,?)")
	if r.Level == DEBUG {
		log.Println(SQL_INSERT, exeSql)
	}
	_, err := r.DB.Exec(exeSql, p.CodeType, p.CodeNo, p.ParentCodeNo, p.CodeVal, p.Flag, p.Version)
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

func (r BasecodeList) InsertEntity(p Basecode) error {
	l := time.Now()
	var colNames, colTags string
	valSlice := make([]interface{}, 0)

	if p.Id != 0 {
		colNames += "id,"
		colTags += "?,"
		valSlice = append(valSlice, p.Id)
	}

	if p.CodeType != "" {
		colNames += "code_type,"
		colTags += "?,"
		valSlice = append(valSlice, p.CodeType)
	}

	if p.CodeNo != "" {
		colNames += "code_no,"
		colTags += "?,"
		valSlice = append(valSlice, p.CodeNo)
	}

	if p.ParentCodeNo != "" {
		colNames += "parent_code_no,"
		colTags += "?,"
		valSlice = append(valSlice, p.ParentCodeNo)
	}

	if p.CodeVal != "" {
		colNames += "code_val,"
		colTags += "?,"
		valSlice = append(valSlice, p.CodeVal)
	}

	if p.Flag != 0 {
		colNames += "flag,"
		colTags += "?,"
		valSlice = append(valSlice, p.Flag)
	}

	if p.InsertTime != "" {
		colNames += "insert_time,"
		colTags += "?,"
		valSlice = append(valSlice, p.InsertTime)
	}

	if p.UpdateTime != "" {
		colNames += "update_time,"
		colTags += "?,"
		valSlice = append(valSlice, p.UpdateTime)
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
	stmt, err := r.DB.Prepare(exeSql)
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

func (r BasecodeList) InsertMap(m map[string]interface{}) error {
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

	exeSql := fmt.Sprintf("Insert into  tba_base_codes(%s)  values(%s)", colNames, colTags)
	if r.Level == DEBUG {
		log.Println(SQL_INSERT, exeSql)
	}
	stmt, err := r.DB.Prepare(exeSql)
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
	说明：根据更新主键及更新Map值更新数据表；
	入参：keyNo:更新数据的关键条件，m:更新数据列的Map
	出参：参数1：如果出错，返回错误对象；成功返回nil
*/

func (r BasecodeList) UpdateMap(keyNo string, m map[string]interface{}) error {
	l := time.Now()

	var colNames string
	valSlice := make([]interface{}, 0)
	for k, v := range m {
		colNames += k + "=?,"
		valSlice = append(valSlice, v)
	}
	valSlice = append(valSlice, keyNo)
	colNames = strings.TrimRight(colNames, ",")
	updateSql := fmt.Sprintf("Update tba_base_codes set %s where code_type=?", colNames)

	if r.Level == DEBUG {
		log.Println(SQL_UPDATE, updateSql)
	}
	stmt, err := r.DB.Prepare(updateSql)
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

func (r BasecodeList) Delete(keyNo string) error {
	l := time.Now()
	delSql := fmt.Sprintf("Delete from  tba_base_codes  where code_type=?")
	if r.Level == DEBUG {
		log.Println(SQL_UPDATE, delSql)
	}
	stmt, err := r.DB.Prepare(delSql)
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
