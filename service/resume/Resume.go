package resume

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
	Id             int     `json:"id"`
	PostNo         int     `json:"post_no"`
	UserId         int     `json:"user_id"`
	WantPositionNo string  `json:"want_position_no"`
	WantSalary     float64 `json:"want_salary"`
	WantArea       string  `json:"want_area"`
	ToUserId       int     `json:"to_user_id"`
	PostionNo      string  `json:"postion_no"`
	InsertTime     string  `json:"insert_time"`
	UpdateTime     string  `json:"update_time"`
	Version        int     `json:"version"`
	PageNo         int     `json:"page_no"`
	PageSize       int     `json:"page_size"`
	SortFld        string  `json:"sort_fld"`
}

type PostResumeList struct {
	DB          *sql.DB
	Level       int
	Total       int          `json:"total"`
	PostResumes []PostResume `json:"PostResume"`
}

type PostResume struct {
	Id             int     `json:"id"`
	PostNo         int     `json:"post_no"`
	UserId         int     `json:"user_id"`
	WantPositionNo string  `json:"want_position_no"`
	WantSalary     float64 `json:"want_salary"`
	WantArea       string  `json:"want_area"`
	ToUserId       int     `json:"to_user_id"`
	PostionNo      string  `json:"postion_no"`
	InsertTime     string  `json:"insert_time"`
	UpdateTime     string  `json:"update_time"`
	Version        int     `json:"version"`
}

/*
	说明：创建实例对象
	入参：db:数据库sql.DB, 数据库已经连接, level:日志级别
	出参：实例对象
*/

func New(db *sql.DB, level int) *PostResumeList {
	if db == nil {
		log.Println(SQL_SELECT, "Database is nil")
		return nil
	}
	return &PostResumeList{DB: db, Total: 0, PostResumes: make([]PostResume, 0), Level: level}
}

/*
	说明：创建实例对象
	入参：url:连接数据的url, 数据库还没有CONNECTED, level:日志级别
	出参：实例对象
*/

func NewUrl(url string, level int) *PostResumeList {
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
	return &PostResumeList{DB: db, Total: 0, PostResumes: make([]PostResume, 0), Level: level}
}

/*
	说明：得到符合条件的总条数
	入参：s: 查询条件
	出参：参数1：返回符合条件的总条件, 参数2：如果错误返回错误对象
*/

func (r *PostResumeList) GetTotal(s Search) (int, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.PostNo != 0 {
		where += " and post_no=" + fmt.Sprintf("%d", s.PostNo)
	}

	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}

	if s.WantPositionNo != "" {
		where += " and want_position_no='" + s.WantPositionNo + "'"
	}

	if s.WantSalary != 0 {
		where += " and want_salary=" + fmt.Sprintf("%f", s.WantSalary)
	}

	if s.WantArea != "" {
		where += " and want_area='" + s.WantArea + "'"
	}

	if s.ToUserId != 0 {
		where += " and to_user_id=" + fmt.Sprintf("%d", s.ToUserId)
	}

	if s.PostionNo != "" {
		where += " and postion_no='" + s.PostionNo + "'"
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

	qrySql := fmt.Sprintf("Select count(1) as total from tba_account_post_resumes   where 1=1 %s", where)
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

func (r PostResumeList) Get(s Search) (*PostResume, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.PostNo != 0 {
		where += " and post_no=" + fmt.Sprintf("%d", s.PostNo)
	}

	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}

	if s.WantPositionNo != "" {
		where += " and want_position_no='" + s.WantPositionNo + "'"
	}

	if s.WantSalary != 0 {
		where += " and want_salary=" + fmt.Sprintf("%f", s.WantSalary)
	}

	if s.WantArea != "" {
		where += " and want_area='" + s.WantArea + "'"
	}

	if s.ToUserId != 0 {
		where += " and to_user_id=" + fmt.Sprintf("%d", s.ToUserId)
	}

	if s.PostionNo != "" {
		where += " and postion_no='" + s.PostionNo + "'"
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

	qrySql := fmt.Sprintf("Select id,post_no,user_id,want_position_no,want_salary,want_area,to_user_id,postion_no,insert_time,update_time,version from tba_account_post_resumes where 1=1 %s Limit %d offset %d", where, s.PageSize, (s.PageNo-1)*s.PageSize)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, nil
	}
	defer rows.Close()

	var p PostResume
	for rows.Next() {
		rows.Scan(&p.Id, &p.PostNo, &p.UserId, &p.WantPositionNo, &p.WantSalary, &p.WantArea, &p.ToUserId, &p.PostionNo, &p.InsertTime, &p.UpdateTime, &p.Version)
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

func (r *PostResumeList) GetList(s Search) ([]PostResume, error) {
	var where string
	l := time.Now()

	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}

	if s.PostNo != 0 {
		where += " and post_no=" + fmt.Sprintf("%d", s.PostNo)
	}

	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}

	if s.WantPositionNo != "" {
		where += " and want_position_no='" + s.WantPositionNo + "'"
	}

	if s.WantSalary != 0 {
		where += " and want_salary=" + fmt.Sprintf("%f", s.WantSalary)
	}

	if s.WantArea != "" {
		where += " and want_area='" + s.WantArea + "'"
	}

	if s.ToUserId != 0 {
		where += " and to_user_id=" + fmt.Sprintf("%d", s.ToUserId)
	}

	if s.PostionNo != "" {
		where += " and postion_no='" + s.PostionNo + "'"
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
		qrySql = fmt.Sprintf("Select id,post_no,user_id,want_position_no,want_salary,want_area,to_user_id,postion_no,insert_time,update_time,version from tba_account_post_resumes where 1=1 %s", where)
	} else {
		qrySql = fmt.Sprintf("Select id,post_no,user_id,want_position_no,want_salary,want_area,to_user_id,postion_no,insert_time,update_time,version from tba_account_post_resumes where 1=1 %s Limit %d offset %d", where, s.PageSize, (s.PageNo-1)*s.PageSize)
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

	var p PostResume
	for rows.Next() {
		rows.Scan(&p.Id, &p.PostNo, &p.UserId, &p.WantPositionNo, &p.WantSalary, &p.WantArea, &p.ToUserId, &p.PostionNo, &p.InsertTime, &p.UpdateTime, &p.Version)
		r.PostResumes = append(r.PostResumes, p)
	}
	log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return r.PostResumes, nil
}

/*
	说明：插入对象到数据表中，这个方法要求对象的各个属性必须赋值
	入参：p:插入的对象
	出参：参数1：如果出错，返回错误对象；成功返回nil
*/

func (r PostResumeList) Insert(p PostResume) error {
	l := time.Now()
	exeSql := fmt.Sprintf("Insert into  tba_account_post_resumes(post_no,user_id,want_position_no,want_salary,want_area,to_user_id,postion_no,version)  values(?,?,?,?,?,?,?,?,?,?,?)")
	if r.Level == DEBUG {
		log.Println(SQL_INSERT, exeSql)
	}
	_, err := r.DB.Exec(exeSql, p.PostNo, p.UserId, p.WantPositionNo, p.WantSalary, p.WantArea, p.ToUserId, p.PostionNo, p.Version)
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

func (r PostResumeList) InsertEntity(p PostResume) error {
	l := time.Now()
	var colNames, colTags string
	valSlice := make([]interface{}, 0)

	if p.Id != 0 {
		colNames += "id,"
		colTags += "?,"
		valSlice = append(valSlice, p.Id)
	}

	if p.PostNo != 0 {
		colNames += "post_no,"
		colTags += "?,"
		valSlice = append(valSlice, p.PostNo)
	}

	if p.UserId != 0 {
		colNames += "user_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.UserId)
	}

	if p.WantPositionNo != "" {
		colNames += "want_position_no,"
		colTags += "?,"
		valSlice = append(valSlice, p.WantPositionNo)
	}

	if p.WantSalary != 0.00 {
		colNames += "want_salary,"
		colTags += "?,"
		valSlice = append(valSlice, p.WantSalary)
	}

	if p.WantArea != "" {
		colNames += "want_area,"
		colTags += "?,"
		valSlice = append(valSlice, p.WantArea)
	}

	if p.ToUserId != 0 {
		colNames += "to_user_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.ToUserId)
	}

	if p.PostionNo != "" {
		colNames += "postion_no,"
		colTags += "?,"
		valSlice = append(valSlice, p.PostionNo)
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

func (r PostResumeList) InsertMap(m map[string]interface{}) error {
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

	exeSql := fmt.Sprintf("Insert into  tba_account_post_resumes(%s)  values(%s)", colNames, colTags)
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

func (r PostResumeList) UpdateMap(keyNo string, m map[string]interface{}) error {
	l := time.Now()

	var colNames string
	valSlice := make([]interface{}, 0)
	for k, v := range m {
		colNames += k + "=?,"
		valSlice = append(valSlice, v)
	}
	valSlice = append(valSlice, keyNo)
	colNames = strings.TrimRight(colNames, ",")
	updateSql := fmt.Sprintf("Update tba_account_post_resumes set %s where post_no=?", colNames)

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

func (r PostResumeList) Delete(keyNo string) error {
	l := time.Now()
	delSql := fmt.Sprintf("Delete from  tba_account_post_resumes  where post_no=?")
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
