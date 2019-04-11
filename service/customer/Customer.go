package customer

import (
	"database/sql"
	"fmt"
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
	Id           int64  `json:"id"`
	CustomerId   int64  `json:"customer_id"`
	CustomerType int64  `json:"customer_type"`
	CustomerNo   string `json:"customer_no"`
	CustomerName string `json:"customer_name"`
	Gender       int64  `json:"gender"`
	Mail         string `json:"mail"`
	Url          string `json:"url"`
	CustomerDesc string `json:"customer_desc"`
	InsertTime   string `json:"insert_time"`
	UpdateTime   string `json:"update_time"`
	Version      int64  `json:"version"`
	PageNo       int    `json:"page_no"`
	PageSize     int    `json:"page_size"`
	ExtraWhere   string `json:"extra_where"`
	SortFld      string `json:"sort_fld"`
}

type CustomersList struct {
	DB         *sql.DB
	Level      int
	Total      int         `json:"total"`
	Customerss []Customers `json:"Customers"`
}

type Customers struct {
	Id           int64  `json:"id"`
	CustomerId   int64  `json:"customer_id"`
	CustomerType int64  `json:"customer_type"`
	CustomerNo   string `json:"customer_no"`
	CustomerName string `json:"customer_name"`
	Gender       int64  `json:"gender"`
	Mail         string `json:"mail"`
	Url          string `json:"url"`
	CustomerDesc string `json:"customer_desc"`
	InsertTime   string `json:"insert_time"`
	UpdateTime   string `json:"update_time"`
	Version      int64  `json:"version"`
}

type Form struct {
	Form Customers `json:"Customers"`
}

/*
	说明：创建实例对象
	入参：db:数据库sql.DB, 数据库已经连接, level:日志级别
	出参：实例对象
*/

func New(db *sql.DB, level int) *CustomersList {
	if db == nil {
		log.Println(SQL_SELECT, "Database is nil")
		return nil
	}
	return &CustomersList{DB: db, Total: 0, Customerss: make([]Customers, 0), Level: level}
}

/*
	说明：创建实例对象
	入参：url:连接数据的url, 数据库还没有CONNECTED, level:日志级别
	出参：实例对象
*/

func NewUrl(url string, level int) *CustomersList {
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
	return &CustomersList{DB: db, Total: 0, Customerss: make([]Customers, 0), Level: level}
}

/*
	说明：得到符合条件的总条数
	入参：s: 查询条件
	出参：参数1：返回符合条件的总条件, 参数2：如果错误返回错误对象
*/

func (r *CustomersList) GetTotal(s Search) (int, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.CustomerId != 0 {
		where += " and customer_id=" + fmt.Sprintf("%d", s.CustomerId)
	}

	if s.CustomerType != 0 {
		where += " and customer_type=" + fmt.Sprintf("%d", s.CustomerType)
	}

	if s.CustomerNo != "" {
		where += " and customer_no='" + s.CustomerNo + "'"
	}

	if s.CustomerName != "" {
		where += " and customer_name='" + s.CustomerName + "'"
	}

	if s.Gender != 0 {
		where += " and gender=" + fmt.Sprintf("%d", s.Gender)
	}

	if s.Mail != "" {
		where += " and mail='" + s.Mail + "'"
	}

	if s.Url != "" {
		where += " and url='" + s.Url + "'"
	}

	if s.CustomerDesc != "" {
		where += " and customer_desc='" + s.CustomerDesc + "'"
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

	qrySql := fmt.Sprintf("Select count(1) as total from tba_customers   where 1=1 %s", where)
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

func (r CustomersList) Get(s Search) (*Customers, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.CustomerId != 0 {
		where += " and customer_id=" + fmt.Sprintf("%d", s.CustomerId)
	}

	if s.CustomerType != 0 {
		where += " and customer_type=" + fmt.Sprintf("%d", s.CustomerType)
	}

	if s.CustomerNo != "" {
		where += " and customer_no='" + s.CustomerNo + "'"
	}

	if s.CustomerName != "" {
		where += " and customer_name='" + s.CustomerName + "'"
	}

	if s.Gender != 0 {
		where += " and gender=" + fmt.Sprintf("%d", s.Gender)
	}

	if s.Mail != "" {
		where += " and mail='" + s.Mail + "'"
	}

	if s.Url != "" {
		where += " and url='" + s.Url + "'"
	}

	if s.CustomerDesc != "" {
		where += " and customer_desc='" + s.CustomerDesc + "'"
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

	qrySql := fmt.Sprintf("Select id,customer_id,customer_type,customer_no,customer_name,gender,mail,url,customer_desc,insert_time,update_time,version from tba_customers where 1=1 %s ", where)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, err
	}
	defer rows.Close()

	var p Customers
	if !rows.Next() {
		return nil, fmt.Errorf("Not Finded Record")
	} else {
		err := rows.Scan(&p.Id, &p.CustomerId, &p.CustomerType, &p.CustomerNo, &p.CustomerName, &p.Gender, &p.Mail, &p.Url, &p.CustomerDesc, &p.InsertTime, &p.UpdateTime, &p.Version)
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

func (r *CustomersList) GetList(s Search) ([]Customers, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.CustomerId != 0 {
		where += " and customer_id=" + fmt.Sprintf("%d", s.CustomerId)
	}

	if s.CustomerType != 0 {
		where += " and customer_type=" + fmt.Sprintf("%d", s.CustomerType)
	}

	if s.CustomerNo != "" {
		where += " and customer_no='" + s.CustomerNo + "'"
	}

	if s.CustomerName != "" {
		where += " and customer_name='" + s.CustomerName + "'"
	}

	if s.Gender != 0 {
		where += " and gender=" + fmt.Sprintf("%d", s.Gender)
	}

	if s.Mail != "" {
		where += " and mail='" + s.Mail + "'"
	}

	if s.Url != "" {
		where += " and url='" + s.Url + "'"
	}

	if s.CustomerDesc != "" {
		where += " and customer_desc='" + s.CustomerDesc + "'"
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
		qrySql = fmt.Sprintf("Select id,customer_id,customer_type,customer_no,customer_name,gender,mail,url,customer_desc,insert_time,update_time,version from tba_customers where 1=1 %s", where)
	} else {
		qrySql = fmt.Sprintf("Select id,customer_id,customer_type,customer_no,customer_name,gender,mail,url,customer_desc,insert_time,update_time,version from tba_customers where 1=1 %s Limit %d offset %d", where, s.PageSize, (s.PageNo-1)*s.PageSize)
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

	var p Customers
	for rows.Next() {
		rows.Scan(&p.Id, &p.CustomerId, &p.CustomerType, &p.CustomerNo, &p.CustomerName, &p.Gender, &p.Mail, &p.Url, &p.CustomerDesc, &p.InsertTime, &p.UpdateTime, &p.Version)
		r.Customerss = append(r.Customerss, p)
	}
	log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return r.Customerss, nil
}

/*
	说明：根据主键查询符合条件的记录，并保持成MAP
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象, 参数2：如果错误返回错误对象
*/

func (r *CustomersList) GetExt(s Search) (map[string]string, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.CustomerId != 0 {
		where += " and customer_id=" + fmt.Sprintf("%d", s.CustomerId)
	}

	if s.CustomerType != 0 {
		where += " and customer_type=" + fmt.Sprintf("%d", s.CustomerType)
	}

	if s.CustomerNo != "" {
		where += " and customer_no='" + s.CustomerNo + "'"
	}

	if s.CustomerName != "" {
		where += " and customer_name='" + s.CustomerName + "'"
	}

	if s.Gender != 0 {
		where += " and gender=" + fmt.Sprintf("%d", s.Gender)
	}

	if s.Mail != "" {
		where += " and mail='" + s.Mail + "'"
	}

	if s.Url != "" {
		where += " and url='" + s.Url + "'"
	}

	if s.CustomerDesc != "" {
		where += " and customer_desc='" + s.CustomerDesc + "'"
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

	qrySql := fmt.Sprintf("Select id,customer_id,customer_type,customer_no,customer_name,gender,mail,url,customer_desc,insert_time,update_time,version from tba_customers where 1=1 %s ", where)
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

func (r CustomersList) Insert(p Customers) error {
	l := time.Now()
	exeSql := fmt.Sprintf("Insert into  tba_customers(customer_id,customer_type,customer_no,customer_name,gender,mail,url,customer_desc,version)  values(?,?,?,?,?,?,?,?,?,?,?,?)")
	if r.Level == DEBUG {
		log.Println(SQL_INSERT, exeSql)
	}
	_, err := r.DB.Exec(exeSql, p.CustomerId, p.CustomerType, p.CustomerNo, p.CustomerName, p.Gender, p.Mail, p.Url, p.CustomerDesc, p.Version)
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

func (r CustomersList) InsertEntity(p Customers, tr *sql.Tx) error {
	l := time.Now()
	var colNames, colTags string
	valSlice := make([]interface{}, 0)

	if p.CustomerId != 0 {
		colNames += "customer_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.CustomerId)
	}

	if p.CustomerType != 0 {
		colNames += "customer_type,"
		colTags += "?,"
		valSlice = append(valSlice, p.CustomerType)
	}

	if p.CustomerNo != "" {
		colNames += "customer_no,"
		colTags += "?,"
		valSlice = append(valSlice, p.CustomerNo)
	}

	if p.CustomerName != "" {
		colNames += "customer_name,"
		colTags += "?,"
		valSlice = append(valSlice, p.CustomerName)
	}

	if p.Gender != 0 {
		colNames += "gender,"
		colTags += "?,"
		valSlice = append(valSlice, p.Gender)
	}

	if p.Mail != "" {
		colNames += "mail,"
		colTags += "?,"
		valSlice = append(valSlice, p.Mail)
	}

	if p.Url != "" {
		colNames += "url,"
		colTags += "?,"
		valSlice = append(valSlice, p.Url)
	}

	if p.CustomerDesc != "" {
		colNames += "customer_desc,"
		colTags += "?,"
		valSlice = append(valSlice, p.CustomerDesc)
	}

	if p.Version != 0 {
		colNames += "version,"
		colTags += "?,"
		valSlice = append(valSlice, p.Version)
	}

	colNames = strings.TrimRight(colNames, ",")
	colTags = strings.TrimRight(colTags, ",")
	exeSql := fmt.Sprintf("Insert into  tba_customers(%s)  values(%s)", colNames, colTags)
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

func (r CustomersList) InsertMap(m map[string]interface{}, tr *sql.Tx) error {
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

	exeSql := fmt.Sprintf("Insert into  tba_customers(%s)  values(%s)", colNames, colTags)
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

func (r CustomersList) UpdataEntity(keyNo string, p Customers, tr *sql.Tx) error {
	l := time.Now()
	var colNames string
	valSlice := make([]interface{}, 0)

	if p.Id != 0 {
		colNames += "id=?,"
		valSlice = append(valSlice, p.Id)
	}

	if p.CustomerId != 0 {
		colNames += "customer_id=?,"
		valSlice = append(valSlice, p.CustomerId)
	}

	if p.CustomerType != 0 {
		colNames += "customer_type=?,"
		valSlice = append(valSlice, p.CustomerType)
	}

	if p.CustomerNo != "" {
		colNames += "customer_no=?,"

		valSlice = append(valSlice, p.CustomerNo)
	}

	if p.CustomerName != "" {
		colNames += "customer_name=?,"

		valSlice = append(valSlice, p.CustomerName)
	}

	if p.Gender != 0 {
		colNames += "gender=?,"
		valSlice = append(valSlice, p.Gender)
	}

	if p.Mail != "" {
		colNames += "mail=?,"

		valSlice = append(valSlice, p.Mail)
	}

	if p.Url != "" {
		colNames += "url=?,"

		valSlice = append(valSlice, p.Url)
	}

	if p.CustomerDesc != "" {
		colNames += "customer_desc=?,"

		valSlice = append(valSlice, p.CustomerDesc)
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

	exeSql := fmt.Sprintf("update  tba_customers  set %s  where id=? ", colNames)
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

func (r CustomersList) UpdateMap(keyNo string, m map[string]interface{}, tr *sql.Tx) error {
	l := time.Now()

	var colNames string
	valSlice := make([]interface{}, 0)
	for k, v := range m {
		colNames += k + "=?,"
		valSlice = append(valSlice, v)
	}
	valSlice = append(valSlice, keyNo)
	colNames = strings.TrimRight(colNames, ",")
	updateSql := fmt.Sprintf("Update tba_customers set %s where id=?", colNames)
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

func (r CustomersList) Delete(keyNo string, tr *sql.Tx) error {
	l := time.Now()
	delSql := fmt.Sprintf("Delete from  tba_customers  where id=?")
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
