package owner

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
	"hcd-gate/service/pubtype"
	
)

const (
	SQL_NEWDB	= "NewDB  ===>"
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
	
	Id	int64	`json:"id"`
	UserId	int64	`json:"user_id"`
	OwnerType	int64	`json:"owner_type"`
	OwnerNo	string	`json:"owner_no"`
	OwnerName	string	`json:"owner_name"`
	Gender	int64	`json:"gender"`
	Mail	string	`json:"mail"`
	Logo	string	`json:"logo"`
	Url	string	`json:"url"`
	OwnerAddr	string	`json:"owner_addr"`
	OwnerDesc	string	`json:"owner_desc"`
	InsertTime	string	`json:"insert_time"`
	UpdateTime	string	`json:"update_time"`
	Version	int64	`json:"version"`
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
	ExtraWhere   string `json:"extra_where"`
	SortFld  string `json:"sort_fld"`
}

type OwnerList struct {
	DB      *sql.DB
	Level   int
	Total   int      `json:"total"`
	Owners []Owner `json:"Owner"`
}

type Owner struct {
	
	Id	int64	`json:"id"`
	UserId	int64	`json:"user_id"`
	OwnerType	int64	`json:"owner_type"`
	OwnerNo	string	`json:"owner_no"`
	OwnerName	string	`json:"owner_name"`
	Gender	int64	`json:"gender"`
	Mail	string	`json:"mail"`
	Logo	string	`json:"logo"`
	Url	string	`json:"url"`
	OwnerAddr	string	`json:"owner_addr"`
	OwnerDesc	string	`json:"owner_desc"`
	InsertTime	string	`json:"insert_time"`
	UpdateTime	string	`json:"update_time"`
	Version	int64	`json:"version"`
}


type Form struct {
	Form   Owner `json:"Owner"`
}

/*
	说明：创建实例对象
	入参：db:数据库sql.DB, 数据库已经连接, level:日志级别
	出参：实例对象
*/

func New(db *sql.DB, level int) *OwnerList {
	if db==nil{
		log.Println(SQL_SELECT,"Database is nil")
		return nil
	}
	return &OwnerList{DB: db, Total: 0, Owners: make([]Owner, 0), Level: level}
}

/*
	说明：创建实例对象
	入参：url:连接数据的url, 数据库还没有CONNECTED, level:日志级别
	出参：实例对象
*/

func NewUrl(url string, level int) *OwnerList {
	var err error
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Println(SQL_SELECT,"Open database error:", err)
		return nil
	}
	if err = db.Ping(); err != nil {
		log.Println(SQL_SELECT,"Ping database error:", err)
		return nil
	}
	return &OwnerList{DB: db, Total: 0, Owners: make([]Owner, 0), Level: level}
}

/*
	说明：得到符合条件的总条数
	入参：s: 查询条件
	出参：参数1：返回符合条件的总条件, 参数2：如果错误返回错误对象
*/

func (r *OwnerList) GetTotal(s Search) (int, error) {
	var where string
	l := time.Now()
	
	
	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}			
	
	
	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}			
	
	
	if s.OwnerType != 0 {
		where += " and owner_type=" + fmt.Sprintf("%d", s.OwnerType)
	}			
	
			
	if s.OwnerNo != "" {
		where += " and owner_no='" + s.OwnerNo + "'"
	}	
	
			
	if s.OwnerName != "" {
		where += " and owner_name='" + s.OwnerName + "'"
	}	
	
	
	if s.Gender != 0 {
		where += " and gender=" + fmt.Sprintf("%d", s.Gender)
	}			
	
			
	if s.Mail != "" {
		where += " and mail='" + s.Mail + "'"
	}	
	
			
	if s.Logo != "" {
		where += " and logo='" + s.Logo + "'"
	}	
	
			
	if s.Url != "" {
		where += " and url='" + s.Url + "'"
	}	
	
			
	if s.OwnerAddr != "" {
		where += " and owner_addr='" + s.OwnerAddr + "'"
	}	
	
			
	if s.OwnerDesc != "" {
		where += " and owner_desc='" + s.OwnerDesc + "'"
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

	qrySql := fmt.Sprintf("Select count(1) as total from tba_account_owners   where 1=1 %s", where)
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

func (r OwnerList) Get(s Search) (*Owner, error) {
	var where string
	l := time.Now()
	
	
	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}			
	
	
	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}			
	
	
	if s.OwnerType != 0 {
		where += " and owner_type=" + fmt.Sprintf("%d", s.OwnerType)
	}			
	
			
	if s.OwnerNo != "" {
		where += " and owner_no='" + s.OwnerNo + "'"
	}	
	
			
	if s.OwnerName != "" {
		where += " and owner_name='" + s.OwnerName + "'"
	}	
	
	
	if s.Gender != 0 {
		where += " and gender=" + fmt.Sprintf("%d", s.Gender)
	}			
	
			
	if s.Mail != "" {
		where += " and mail='" + s.Mail + "'"
	}	
	
			
	if s.Logo != "" {
		where += " and logo='" + s.Logo + "'"
	}	
	
			
	if s.Url != "" {
		where += " and url='" + s.Url + "'"
	}	
	
			
	if s.OwnerAddr != "" {
		where += " and owner_addr='" + s.OwnerAddr + "'"
	}	
	
			
	if s.OwnerDesc != "" {
		where += " and owner_desc='" + s.OwnerDesc + "'"
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
	
	qrySql := fmt.Sprintf("Select id,user_id,owner_type,owner_no,owner_name,gender,mail,logo,url,owner_addr,owner_desc,insert_time,update_time,version from tba_account_owners where 1=1 %s ", where)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, err
	}
	defer rows.Close()

	var p  Owner
	if !rows.Next() {
		return nil, fmt.Errorf("Not Finded Record")
	} else {
		err:=rows.Scan(&p.Id,&p.UserId,&p.OwnerType,&p.OwnerNo,&p.OwnerName,&p.Gender,&p.Mail,&p.Logo,&p.Url,&p.OwnerAddr,&p.OwnerDesc,&p.InsertTime,&p.UpdateTime,&p.Version)
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

func (r *OwnerList) GetList(s Search) ([]Owner, error) {
	var where string
	l := time.Now()
	
	
	
	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}			
	
	
	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}			
	
	
	if s.OwnerType != 0 {
		where += " and owner_type=" + fmt.Sprintf("%d", s.OwnerType)
	}			
	
			
	if s.OwnerNo != "" {
		where += " and owner_no='" + s.OwnerNo + "'"
	}	
	
			
	if s.OwnerName != "" {
		where += " and owner_name='" + s.OwnerName + "'"
	}	
	
	
	if s.Gender != 0 {
		where += " and gender=" + fmt.Sprintf("%d", s.Gender)
	}			
	
			
	if s.Mail != "" {
		where += " and mail='" + s.Mail + "'"
	}	
	
			
	if s.Logo != "" {
		where += " and logo='" + s.Logo + "'"
	}	
	
			
	if s.Url != "" {
		where += " and url='" + s.Url + "'"
	}	
	
			
	if s.OwnerAddr != "" {
		where += " and owner_addr='" + s.OwnerAddr + "'"
	}	
	
			
	if s.OwnerDesc != "" {
		where += " and owner_desc='" + s.OwnerDesc + "'"
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
	if s.PageSize==0 &&s.PageNo==0{
		qrySql = fmt.Sprintf("Select id,user_id,owner_type,owner_no,owner_name,gender,mail,logo,url,owner_addr,owner_desc,insert_time,update_time,version from tba_account_owners where 1=1 %s", where)
	}else{
		qrySql = fmt.Sprintf("Select id,user_id,owner_type,owner_no,owner_name,gender,mail,logo,url,owner_addr,owner_desc,insert_time,update_time,version from tba_account_owners where 1=1 %s Limit %d offset %d", where, s.PageSize, (s.PageNo-1)*s.PageSize)
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

	var p Owner
	for rows.Next() {
		rows.Scan(&p.Id,&p.UserId,&p.OwnerType,&p.OwnerNo,&p.OwnerName,&p.Gender,&p.Mail,&p.Logo,&p.Url,&p.OwnerAddr,&p.OwnerDesc,&p.InsertTime,&p.UpdateTime,&p.Version)
		r.Owners = append(r.Owners, p)
	}
	log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return r.Owners, nil
}

/*
	说明：根据条件查询复核条件对象列表，支持分页查询
	入参：s: 查询条件
	出参：参数1：返回符合条件的对象列表, 参数2：如果错误返回错误对象
*/

func (r *OwnerList) GetListExt(s Search, fList []string) ([][]pubtype.Data, error) {
	var where string
	l := time.Now()
	
	
	
	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}			
	
	
	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}			
	
	
	if s.OwnerType != 0 {
		where += " and owner_type=" + fmt.Sprintf("%d", s.OwnerType)
	}			
	
			
	if s.OwnerNo != "" {
		where += " and owner_no='" + s.OwnerNo + "'"
	}	
	
			
	if s.OwnerName != "" {
		where += " and owner_name='" + s.OwnerName + "'"
	}	
	
	
	if s.Gender != 0 {
		where += " and gender=" + fmt.Sprintf("%d", s.Gender)
	}			
	
			
	if s.Mail != "" {
		where += " and mail='" + s.Mail + "'"
	}	
	
			
	if s.Logo != "" {
		where += " and logo='" + s.Logo + "'"
	}	
	
			
	if s.Url != "" {
		where += " and url='" + s.Url + "'"
	}	
	
			
	if s.OwnerAddr != "" {
		where += " and owner_addr='" + s.OwnerAddr + "'"
	}	
	
			
	if s.OwnerDesc != "" {
		where += " and owner_desc='" + s.OwnerDesc + "'"
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
	if s.PageSize==0 &&s.PageNo==0{
		qrySql = fmt.Sprintf("Select %s from tba_account_owners where 1=1 %s", colNames,where)
	}else{
		qrySql = fmt.Sprintf("Select %s from tba_account_owners where 1=1 %s Limit %d offset %d", colNames,where, s.PageSize, (s.PageNo-1)*s.PageSize)
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

func (r *OwnerList) GetExt(s Search) (map[string]string, error) {
	var where string
	l := time.Now()

	
	
	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}			
	
	
	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}			
	
	
	if s.OwnerType != 0 {
		where += " and owner_type=" + fmt.Sprintf("%d", s.OwnerType)
	}			
	
			
	if s.OwnerNo != "" {
		where += " and owner_no='" + s.OwnerNo + "'"
	}	
	
			
	if s.OwnerName != "" {
		where += " and owner_name='" + s.OwnerName + "'"
	}	
	
	
	if s.Gender != 0 {
		where += " and gender=" + fmt.Sprintf("%d", s.Gender)
	}			
	
			
	if s.Mail != "" {
		where += " and mail='" + s.Mail + "'"
	}	
	
			
	if s.Logo != "" {
		where += " and logo='" + s.Logo + "'"
	}	
	
			
	if s.Url != "" {
		where += " and url='" + s.Url + "'"
	}	
	
			
	if s.OwnerAddr != "" {
		where += " and owner_addr='" + s.OwnerAddr + "'"
	}	
	
			
	if s.OwnerDesc != "" {
		where += " and owner_desc='" + s.OwnerDesc + "'"
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
	

	qrySql := fmt.Sprintf("Select id,user_id,owner_type,owner_no,owner_name,gender,mail,logo,url,owner_addr,owner_desc,insert_time,update_time,version from tba_account_owners where 1=1 %s ", where)
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

func (r OwnerList) Insert(p Owner) error {
	l := time.Now()
	exeSql := fmt.Sprintf("Insert into  tba_account_owners(user_id,owner_type,owner_no,owner_name,gender,mail,logo,url,owner_addr,owner_desc,version)  values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if r.Level == DEBUG {
		log.Println(SQL_INSERT, exeSql)
	}
	_, err := r.DB.Exec(exeSql, p.UserId,p.OwnerType,p.OwnerNo,p.OwnerName,p.Gender,p.Mail,p.Logo,p.Url,p.OwnerAddr,p.OwnerDesc,p.Version)
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


func (r OwnerList) InsertEntity(p Owner, tr *sql.Tx) error {
	l := time.Now()
	var colNames, colTags string
	valSlice := make([]interface{}, 0)
	
	
	if p.UserId != 0 {
		colNames += "user_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.UserId)
	}				
	
	if p.OwnerType != 0 {
		colNames += "owner_type,"
		colTags += "?,"
		valSlice = append(valSlice, p.OwnerType)
	}				
		
	if p.OwnerNo != "" {
		colNames += "owner_no,"
		colTags += "?,"
		valSlice = append(valSlice, p.OwnerNo)
	}			
		
	if p.OwnerName != "" {
		colNames += "owner_name,"
		colTags += "?,"
		valSlice = append(valSlice, p.OwnerName)
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
		
	if p.Logo != "" {
		colNames += "logo,"
		colTags += "?,"
		valSlice = append(valSlice, p.Logo)
	}			
		
	if p.Url != "" {
		colNames += "url,"
		colTags += "?,"
		valSlice = append(valSlice, p.Url)
	}			
		
	if p.OwnerAddr != "" {
		colNames += "owner_addr,"
		colTags += "?,"
		valSlice = append(valSlice, p.OwnerAddr)
	}			
		
	if p.OwnerDesc != "" {
		colNames += "owner_desc,"
		colTags += "?,"
		valSlice = append(valSlice, p.OwnerDesc)
	}			
	
	if p.Version != 0 {
		colNames += "version,"
		colTags += "?,"
		valSlice = append(valSlice, p.Version)
	}				
	
	colNames = strings.TrimRight(colNames, ",")
	colTags = strings.TrimRight(colTags, ",")
	exeSql := fmt.Sprintf("Insert into  tba_account_owners(%s)  values(%s)", colNames, colTags)
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

func (r OwnerList) InsertMap(m map[string]interface{},tr *sql.Tx) error {
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

	exeSql := fmt.Sprintf("Insert into  tba_account_owners(%s)  values(%s)", colNames, colTags)
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


func (r OwnerList) UpdataEntity(keyNo string,p Owner,tr *sql.Tx) error {
	l := time.Now()
	var colNames string
	valSlice := make([]interface{}, 0)
	
	
	if p.Id != 0 {
		colNames += "id=?,"
		valSlice = append(valSlice, p.Id)
	}				
	
	if p.UserId != 0 {
		colNames += "user_id=?,"
		valSlice = append(valSlice, p.UserId)
	}				
	
	if p.OwnerType != 0 {
		colNames += "owner_type=?,"
		valSlice = append(valSlice, p.OwnerType)
	}				
		
	if p.OwnerNo != "" {
		colNames += "owner_no=?,"
		
		valSlice = append(valSlice, p.OwnerNo)
	}			
		
	if p.OwnerName != "" {
		colNames += "owner_name=?,"
		
		valSlice = append(valSlice, p.OwnerName)
	}			
	
	if p.Gender != 0 {
		colNames += "gender=?,"
		valSlice = append(valSlice, p.Gender)
	}				
		
	if p.Mail != "" {
		colNames += "mail=?,"
		
		valSlice = append(valSlice, p.Mail)
	}			
		
	if p.Logo != "" {
		colNames += "logo=?,"
		
		valSlice = append(valSlice, p.Logo)
	}			
		
	if p.Url != "" {
		colNames += "url=?,"
		
		valSlice = append(valSlice, p.Url)
	}			
		
	if p.OwnerAddr != "" {
		colNames += "owner_addr=?,"
		
		valSlice = append(valSlice, p.OwnerAddr)
	}			
		
	if p.OwnerDesc != "" {
		colNames += "owner_desc=?,"
		
		valSlice = append(valSlice, p.OwnerDesc)
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

	exeSql := fmt.Sprintf("update  tba_account_owners  set %s  where id=? ", colNames)
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

func (r OwnerList) UpdateMap(keyNo string, m map[string]interface{},tr *sql.Tx) error {
	l := time.Now()

	var colNames string
	valSlice := make([]interface{}, 0)
	for k, v := range m {
		colNames += k + "=?,"
		valSlice = append(valSlice, v)
	}
	valSlice = append(valSlice, keyNo)
	colNames = strings.TrimRight(colNames, ",")
	updateSql := fmt.Sprintf("Update tba_account_owners set %s where id=?", colNames)
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

func (r OwnerList) Delete(keyNo string,tr *sql.Tx) error {
	l := time.Now()
	delSql := fmt.Sprintf("Delete from  tba_account_owners  where id=?")
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

