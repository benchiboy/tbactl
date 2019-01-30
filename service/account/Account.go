package account

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	PartnerUserId	string	`json:"partner_user_id"`
	ParentUserId	int64	`json:"parent_user_id"`
	UserRole	int64	`json:"user_role"`
	UserStatus	int64	`json:"user_status"`
	AvatarUrl	string	`json:"avatar_url"`
	LoginMode	int64	`json:"login_mode"`
	LoginName	string	`json:"login_name"`
	LoginPassword	string	`json:"login_password"`
	ErrorCount	int64	`json:"error_count"`
	LastLoginTime	string	`json:"last_login_time"`
	LastDeviceId	string	`json:"last_device_id"`
	AccountBal	float64	`json:"account_bal"`
	GoodsCount	float64	`json:"goods_count"`
	DeviceIp	string	`json:"device_ip"`
	DeviceType	int64	`json:"device_type"`
	DeviceOs	string	`json:"device_os"`
	DeviceOsVer	string	`json:"device_os_ver"`
	DeviceId	string	`json:"device_id"`
	Latitude	string	`json:"latitude"`
	Longitude	string	`json:"longitude"`
	Market	string	`json:"market"`
	UserChannel	string	`json:"user_channel"`
	RandomNo	int64	`json:"random_no"`
	RegionNo	string	`json:"region_no"`
	CustomerId	int64	`json:"customer_id"`
	CreatedTime	string	`json:"created_time"`
	UpdatedTime	string	`json:"updated_time"`
	Memo	string	`json:"memo"`
	InsertUser	string	`json:"insert_user"`
	UpdateUser	string	`json:"update_user"`
	Version	int64	`json:"version"`
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
	SortFld  string `json:"sort_fld"`
}

type AccountList struct {
	DB      *sql.DB
	Level   int
	Total   int      `json:"total"`
	Accounts []Account `json:"Account"`
}

type Account struct {
	
	Id	int64	`json:"id"`
	UserId	int64	`json:"user_id"`
	PartnerUserId	string	`json:"partner_user_id"`
	ParentUserId	int64	`json:"parent_user_id"`
	UserRole	int64	`json:"user_role"`
	UserStatus	int64	`json:"user_status"`
	AvatarUrl	string	`json:"avatar_url"`
	LoginMode	int64	`json:"login_mode"`
	LoginName	string	`json:"login_name"`
	LoginPassword	string	`json:"login_password"`
	ErrorCount	int64	`json:"error_count"`
	LastLoginTime	string	`json:"last_login_time"`
	LastDeviceId	string	`json:"last_device_id"`
	AccountBal	float64	`json:"account_bal"`
	GoodsCount	float64	`json:"goods_count"`
	DeviceIp	string	`json:"device_ip"`
	DeviceType	int64	`json:"device_type"`
	DeviceOs	string	`json:"device_os"`
	DeviceOsVer	string	`json:"device_os_ver"`
	DeviceId	string	`json:"device_id"`
	Latitude	string	`json:"latitude"`
	Longitude	string	`json:"longitude"`
	Market	string	`json:"market"`
	UserChannel	string	`json:"user_channel"`
	RandomNo	int64	`json:"random_no"`
	RegionNo	string	`json:"region_no"`
	CustomerId	int64	`json:"customer_id"`
	CreatedTime	string	`json:"created_time"`
	UpdatedTime	string	`json:"updated_time"`
	Memo	string	`json:"memo"`
	InsertUser	string	`json:"insert_user"`
	UpdateUser	string	`json:"update_user"`
	Version	int64	`json:"version"`
}


type Form struct {
	Form   Account `json:"Account"`
}

/*
	说明：创建实例对象
	入参：db:数据库sql.DB, 数据库已经连接, level:日志级别
	出参：实例对象
*/

func New(db *sql.DB, level int) *AccountList {
	if db==nil{
		log.Println(SQL_SELECT,"Database is nil")
		return nil
	}
	return &AccountList{DB: db, Total: 0, Accounts: make([]Account, 0), Level: level}
}

/*
	说明：创建实例对象
	入参：url:连接数据的url, 数据库还没有CONNECTED, level:日志级别
	出参：实例对象
*/

func NewUrl(url string, level int) *AccountList {
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
	return &AccountList{DB: db, Total: 0, Accounts: make([]Account, 0), Level: level}
}

/*
	说明：得到符合条件的总条数
	入参：s: 查询条件
	出参：参数1：返回符合条件的总条件, 参数2：如果错误返回错误对象
*/

func (r *AccountList) GetTotal(s Search) (int, error) {
	var where string
	l := time.Now()
	
	
	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}			
	
	
	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}			
	
			
	if s.PartnerUserId != "" {
		where += " and partner_user_id='" + s.PartnerUserId + "'"
	}	
	
	
	if s.ParentUserId != 0 {
		where += " and parent_user_id=" + fmt.Sprintf("%d", s.ParentUserId)
	}			
	
	
	if s.UserRole != 0 {
		where += " and user_role=" + fmt.Sprintf("%d", s.UserRole)
	}			
	
	
	if s.UserStatus != 0 {
		where += " and user_status=" + fmt.Sprintf("%d", s.UserStatus)
	}			
	
			
	if s.AvatarUrl != "" {
		where += " and avatar_url='" + s.AvatarUrl + "'"
	}	
	
	
	if s.LoginMode != 0 {
		where += " and login_mode=" + fmt.Sprintf("%d", s.LoginMode)
	}			
	
			
	if s.LoginName != "" {
		where += " and login_name='" + s.LoginName + "'"
	}	
	
			
	if s.LoginPassword != "" {
		where += " and login_password='" + s.LoginPassword + "'"
	}	
	
	
	if s.ErrorCount != 0 {
		where += " and error_count=" + fmt.Sprintf("%d", s.ErrorCount)
	}			
	
			
	if s.LastLoginTime != "" {
		where += " and last_login_time='" + s.LastLoginTime + "'"
	}	
	
			
	if s.LastDeviceId != "" {
		where += " and last_device_id='" + s.LastDeviceId + "'"
	}	
	
		
	if s.AccountBal != 0 {
		where += " and account_bal=" + fmt.Sprintf("%f", s.AccountBal)
	}		
	
		
	if s.GoodsCount != 0 {
		where += " and goods_count=" + fmt.Sprintf("%f", s.GoodsCount)
	}		
	
			
	if s.DeviceIp != "" {
		where += " and device_ip='" + s.DeviceIp + "'"
	}	
	
	
	if s.DeviceType != 0 {
		where += " and device_type=" + fmt.Sprintf("%d", s.DeviceType)
	}			
	
			
	if s.DeviceOs != "" {
		where += " and device_os='" + s.DeviceOs + "'"
	}	
	
			
	if s.DeviceOsVer != "" {
		where += " and device_os_ver='" + s.DeviceOsVer + "'"
	}	
	
			
	if s.DeviceId != "" {
		where += " and device_id='" + s.DeviceId + "'"
	}	
	
			
	if s.Latitude != "" {
		where += " and latitude='" + s.Latitude + "'"
	}	
	
			
	if s.Longitude != "" {
		where += " and longitude='" + s.Longitude + "'"
	}	
	
			
	if s.Market != "" {
		where += " and market='" + s.Market + "'"
	}	
	
			
	if s.UserChannel != "" {
		where += " and user_channel='" + s.UserChannel + "'"
	}	
	
	
	if s.RandomNo != 0 {
		where += " and random_no=" + fmt.Sprintf("%d", s.RandomNo)
	}			
	
			
	if s.RegionNo != "" {
		where += " and region_no='" + s.RegionNo + "'"
	}	
	
	
	if s.CustomerId != 0 {
		where += " and customer_id=" + fmt.Sprintf("%d", s.CustomerId)
	}			
	
			
	if s.CreatedTime != "" {
		where += " and created_time='" + s.CreatedTime + "'"
	}	
	
			
	if s.UpdatedTime != "" {
		where += " and updated_time='" + s.UpdatedTime + "'"
	}	
	
			
	if s.Memo != "" {
		where += " and memo='" + s.Memo + "'"
	}	
	
			
	if s.InsertUser != "" {
		where += " and insert_user='" + s.InsertUser + "'"
	}	
	
			
	if s.UpdateUser != "" {
		where += " and update_user='" + s.UpdateUser + "'"
	}	
	
	
	if s.Version != 0 {
		where += " and version=" + fmt.Sprintf("%d", s.Version)
	}			
	
	qrySql := fmt.Sprintf("Select count(1) as total from tba_accounts   where 1=1 %s", where)
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

func (r AccountList) Get(s Search) (*Account, error) {
	var where string
	l := time.Now()
	
	
	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}			
	
			
	if s.PartnerUserId != "" {
		where += " and partner_user_id='" + s.PartnerUserId + "'"
	}	
	
	
	if s.ParentUserId != 0 {
		where += " and parent_user_id=" + fmt.Sprintf("%d", s.ParentUserId)
	}			
	
	
	if s.UserRole != 0 {
		where += " and user_role=" + fmt.Sprintf("%d", s.UserRole)
	}			
	
	
	if s.UserStatus != 0 {
		where += " and user_status=" + fmt.Sprintf("%d", s.UserStatus)
	}			
	
			
	if s.AvatarUrl != "" {
		where += " and avatar_url='" + s.AvatarUrl + "'"
	}	
	
	
	if s.LoginMode != 0 {
		where += " and login_mode=" + fmt.Sprintf("%d", s.LoginMode)
	}			
	
			
	if s.LoginName != "" {
		where += " and login_name='" + s.LoginName + "'"
	}	
	
			
	if s.LoginPassword != "" {
		where += " and login_password='" + s.LoginPassword + "'"
	}	
	
	
	if s.ErrorCount != 0 {
		where += " and error_count=" + fmt.Sprintf("%d", s.ErrorCount)
	}			
	
			
	if s.LastLoginTime != "" {
		where += " and last_login_time='" + s.LastLoginTime + "'"
	}	
	
			
	if s.LastDeviceId != "" {
		where += " and last_device_id='" + s.LastDeviceId + "'"
	}	
	
		
	if s.AccountBal != 0 {
		where += " and account_bal=" + fmt.Sprintf("%f", s.AccountBal)
	}		
	
		
	if s.GoodsCount != 0 {
		where += " and goods_count=" + fmt.Sprintf("%f", s.GoodsCount)
	}		
	
			
	if s.DeviceIp != "" {
		where += " and device_ip='" + s.DeviceIp + "'"
	}	
	
	
	if s.DeviceType != 0 {
		where += " and device_type=" + fmt.Sprintf("%d", s.DeviceType)
	}			
	
			
	if s.DeviceOs != "" {
		where += " and device_os='" + s.DeviceOs + "'"
	}	
	
			
	if s.DeviceOsVer != "" {
		where += " and device_os_ver='" + s.DeviceOsVer + "'"
	}	
	
			
	if s.DeviceId != "" {
		where += " and device_id='" + s.DeviceId + "'"
	}	
	
			
	if s.Latitude != "" {
		where += " and latitude='" + s.Latitude + "'"
	}	
	
			
	if s.Longitude != "" {
		where += " and longitude='" + s.Longitude + "'"
	}	
	
			
	if s.Market != "" {
		where += " and market='" + s.Market + "'"
	}	
	
			
	if s.UserChannel != "" {
		where += " and user_channel='" + s.UserChannel + "'"
	}	
	
	
	if s.RandomNo != 0 {
		where += " and random_no=" + fmt.Sprintf("%d", s.RandomNo)
	}			
	
			
	if s.RegionNo != "" {
		where += " and region_no='" + s.RegionNo + "'"
	}	
	
	
	if s.CustomerId != 0 {
		where += " and customer_id=" + fmt.Sprintf("%d", s.CustomerId)
	}			
	
			
	if s.CreatedTime != "" {
		where += " and created_time='" + s.CreatedTime + "'"
	}	
	
			
	if s.UpdatedTime != "" {
		where += " and updated_time='" + s.UpdatedTime + "'"
	}	
	
			
	if s.Memo != "" {
		where += " and memo='" + s.Memo + "'"
	}	
	
			
	if s.InsertUser != "" {
		where += " and insert_user='" + s.InsertUser + "'"
	}	
	
			
	if s.UpdateUser != "" {
		where += " and update_user='" + s.UpdateUser + "'"
	}	
	
	
	if s.Version != 0 {
		where += " and version=" + fmt.Sprintf("%d", s.Version)
	}			
	
	
	qrySql := fmt.Sprintf("Select user_id,partner_user_id,parent_user_id,user_role,user_status,avatar_url,login_mode,login_name,login_password,error_count,last_login_time,last_device_id,account_bal,goods_count,device_ip,device_type,device_os,device_os_ver,device_id,latitude,longitude,market,user_channel,random_no,region_no,customer_id,created_time,updated_time,memo,insert_user,update_user,version from tba_accounts where 1=1 %s ", where)
	if r.Level == DEBUG {
		log.Println(SQL_SELECT, qrySql)
	}
	rows, err := r.DB.Query(qrySql)
	if err != nil {
		log.Println(SQL_ERROR, err.Error())
		return nil, nil
	}
	defer rows.Close()

	var p  Account
	for rows.Next() {
		rows.Scan(&p.UserId,&p.PartnerUserId,&p.ParentUserId,&p.UserRole,&p.UserStatus,&p.AvatarUrl,&p.LoginMode,&p.LoginName,&p.LoginPassword,&p.ErrorCount,&p.LastLoginTime,&p.LastDeviceId,&p.AccountBal,&p.GoodsCount,&p.DeviceIp,&p.DeviceType,&p.DeviceOs,&p.DeviceOsVer,&p.DeviceId,&p.Latitude,&p.Longitude,&p.Market,&p.UserChannel,&p.RandomNo,&p.RegionNo,&p.CustomerId,&p.CreatedTime,&p.UpdatedTime,&p.Memo,&p.InsertUser,&p.UpdateUser,&p.Version)
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

func (r *AccountList) GetList(s Search) ([]Account, error) {
	var where string
	l := time.Now()
	
	
	
	if s.Id != 0 {
		where += " and id=" + fmt.Sprintf("%d", s.Id)
	}			
	
	
	if s.UserId != 0 {
		where += " and user_id=" + fmt.Sprintf("%d", s.UserId)
	}			
	
			
	if s.PartnerUserId != "" {
		where += " and partner_user_id='" + s.PartnerUserId + "'"
	}	
	
	
	if s.ParentUserId != 0 {
		where += " and parent_user_id=" + fmt.Sprintf("%d", s.ParentUserId)
	}			
	
	
	if s.UserRole != 0 {
		where += " and user_role=" + fmt.Sprintf("%d", s.UserRole)
	}			
	
	
	if s.UserStatus != 0 {
		where += " and user_status=" + fmt.Sprintf("%d", s.UserStatus)
	}			
	
			
	if s.AvatarUrl != "" {
		where += " and avatar_url='" + s.AvatarUrl + "'"
	}	
	
	
	if s.LoginMode != 0 {
		where += " and login_mode=" + fmt.Sprintf("%d", s.LoginMode)
	}			
	
			
	if s.LoginName != "" {
		where += " and login_name='" + s.LoginName + "'"
	}	
	
			
	if s.LoginPassword != "" {
		where += " and login_password='" + s.LoginPassword + "'"
	}	
	
	
	if s.ErrorCount != 0 {
		where += " and error_count=" + fmt.Sprintf("%d", s.ErrorCount)
	}			
	
			
	if s.LastLoginTime != "" {
		where += " and last_login_time='" + s.LastLoginTime + "'"
	}	
	
			
	if s.LastDeviceId != "" {
		where += " and last_device_id='" + s.LastDeviceId + "'"
	}	
	
		
	if s.AccountBal != 0 {
		where += " and account_bal=" + fmt.Sprintf("%f", s.AccountBal)
	}		
	
		
	if s.GoodsCount != 0 {
		where += " and goods_count=" + fmt.Sprintf("%f", s.GoodsCount)
	}		
	
			
	if s.DeviceIp != "" {
		where += " and device_ip='" + s.DeviceIp + "'"
	}	
	
	
	if s.DeviceType != 0 {
		where += " and device_type=" + fmt.Sprintf("%d", s.DeviceType)
	}			
	
			
	if s.DeviceOs != "" {
		where += " and device_os='" + s.DeviceOs + "'"
	}	
	
			
	if s.DeviceOsVer != "" {
		where += " and device_os_ver='" + s.DeviceOsVer + "'"
	}	
	
			
	if s.DeviceId != "" {
		where += " and device_id='" + s.DeviceId + "'"
	}	
	
			
	if s.Latitude != "" {
		where += " and latitude='" + s.Latitude + "'"
	}	
	
			
	if s.Longitude != "" {
		where += " and longitude='" + s.Longitude + "'"
	}	
	
			
	if s.Market != "" {
		where += " and market='" + s.Market + "'"
	}	
	
			
	if s.UserChannel != "" {
		where += " and user_channel='" + s.UserChannel + "'"
	}	
	
	
	if s.RandomNo != 0 {
		where += " and random_no=" + fmt.Sprintf("%d", s.RandomNo)
	}			
	
			
	if s.RegionNo != "" {
		where += " and region_no='" + s.RegionNo + "'"
	}	
	
	
	if s.CustomerId != 0 {
		where += " and customer_id=" + fmt.Sprintf("%d", s.CustomerId)
	}			
	
			
	if s.CreatedTime != "" {
		where += " and created_time='" + s.CreatedTime + "'"
	}	
	
			
	if s.UpdatedTime != "" {
		where += " and updated_time='" + s.UpdatedTime + "'"
	}	
	
			
	if s.Memo != "" {
		where += " and memo='" + s.Memo + "'"
	}	
	
			
	if s.InsertUser != "" {
		where += " and insert_user='" + s.InsertUser + "'"
	}	
	
			
	if s.UpdateUser != "" {
		where += " and update_user='" + s.UpdateUser + "'"
	}	
	
	
	if s.Version != 0 {
		where += " and version=" + fmt.Sprintf("%d", s.Version)
	}			
	
	
	var qrySql string
	if s.PageSize==0 &&s.PageNo==0{
		qrySql = fmt.Sprintf("Select id,user_id,partner_user_id,parent_user_id,user_role,user_status,avatar_url,login_mode,login_name,login_password,error_count,last_login_time,last_device_id,account_bal,goods_count,device_ip,device_type,device_os,device_os_ver,device_id,latitude,longitude,market,user_channel,random_no,region_no,customer_id,created_time,updated_time,memo,insert_user,update_user,version from tba_accounts where 1=1 %s", where)
	}else{
		qrySql = fmt.Sprintf("Select id,user_id,partner_user_id,parent_user_id,user_role,user_status,avatar_url,login_mode,login_name,login_password,error_count,last_login_time,last_device_id,account_bal,goods_count,device_ip,device_type,device_os,device_os_ver,device_id,latitude,longitude,market,user_channel,random_no,region_no,customer_id,created_time,updated_time,memo,insert_user,update_user,version from tba_accounts where 1=1 %s Limit %d offset %d", where, s.PageSize, (s.PageNo-1)*s.PageSize)
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

	var p Account
	for rows.Next() {
		rows.Scan(&p.Id,&p.UserId,&p.PartnerUserId,&p.ParentUserId,&p.UserRole,&p.UserStatus,&p.AvatarUrl,&p.LoginMode,&p.LoginName,&p.LoginPassword,&p.ErrorCount,&p.LastLoginTime,&p.LastDeviceId,&p.AccountBal,&p.GoodsCount,&p.DeviceIp,&p.DeviceType,&p.DeviceOs,&p.DeviceOsVer,&p.DeviceId,&p.Latitude,&p.Longitude,&p.Market,&p.UserChannel,&p.RandomNo,&p.RegionNo,&p.CustomerId,&p.CreatedTime,&p.UpdatedTime,&p.Memo,&p.InsertUser,&p.UpdateUser,&p.Version)
		r.Accounts = append(r.Accounts, p)
	}
	log.Println(SQL_ELAPSED, r)
	if r.Level == DEBUG {
		log.Println(SQL_ELAPSED, time.Since(l))
	}
	return r.Accounts, nil
}


/*
	说明：插入对象到数据表中，这个方法要求对象的各个属性必须赋值
	入参：p:插入的对象
	出参：参数1：如果出错，返回错误对象；成功返回nil
*/

func (r AccountList) Insert(p Account) error {
	l := time.Now()
	exeSql := fmt.Sprintf("Insert into  tba_accounts(user_id,partner_user_id,parent_user_id,user_role,user_status,avatar_url,login_mode,login_name,login_password,error_count,last_login_time,last_device_id,account_bal,goods_count,device_ip,device_type,device_os,device_os_ver,device_id,latitude,longitude,market,user_channel,random_no,region_no,customer_id,created_time,updated_time,memo,insert_user,update_user,version)  values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if r.Level == DEBUG {
		log.Println(SQL_INSERT, exeSql)
	}
	_, err := r.DB.Exec(exeSql, p.UserId,p.PartnerUserId,p.ParentUserId,p.UserRole,p.UserStatus,p.AvatarUrl,p.LoginMode,p.LoginName,p.LoginPassword,p.ErrorCount,p.LastLoginTime,p.LastDeviceId,p.AccountBal,p.GoodsCount,p.DeviceIp,p.DeviceType,p.DeviceOs,p.DeviceOsVer,p.DeviceId,p.Latitude,p.Longitude,p.Market,p.UserChannel,p.RandomNo,p.RegionNo,p.CustomerId,p.CreatedTime,p.UpdatedTime,p.Memo,p.InsertUser,p.UpdateUser,p.Version)
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


func (r AccountList) InsertEntity(p Account) error {
	l := time.Now()
	var colNames, colTags string
	valSlice := make([]interface{}, 0)
	
	
	if p.Id != 0 {
		colNames += "id,"
		colTags += "?,"
		valSlice = append(valSlice, p.Id)
	}				
	
	if p.UserId != 0 {
		colNames += "user_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.UserId)
	}				
		
	if p.PartnerUserId != "" {
		colNames += "partner_user_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.PartnerUserId)
	}			
	
	if p.ParentUserId != 0 {
		colNames += "parent_user_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.ParentUserId)
	}				
	
	if p.UserRole != 0 {
		colNames += "user_role,"
		colTags += "?,"
		valSlice = append(valSlice, p.UserRole)
	}				
	
	if p.UserStatus != 0 {
		colNames += "user_status,"
		colTags += "?,"
		valSlice = append(valSlice, p.UserStatus)
	}				
		
	if p.AvatarUrl != "" {
		colNames += "avatar_url,"
		colTags += "?,"
		valSlice = append(valSlice, p.AvatarUrl)
	}			
	
	if p.LoginMode != 0 {
		colNames += "login_mode,"
		colTags += "?,"
		valSlice = append(valSlice, p.LoginMode)
	}				
		
	if p.LoginName != "" {
		colNames += "login_name,"
		colTags += "?,"
		valSlice = append(valSlice, p.LoginName)
	}			
		
	if p.LoginPassword != "" {
		colNames += "login_password,"
		colTags += "?,"
		valSlice = append(valSlice, p.LoginPassword)
	}			
	
	if p.ErrorCount != 0 {
		colNames += "error_count,"
		colTags += "?,"
		valSlice = append(valSlice, p.ErrorCount)
	}				
		
	if p.LastLoginTime != "" {
		colNames += "last_login_time,"
		colTags += "?,"
		valSlice = append(valSlice, p.LastLoginTime)
	}			
		
	if p.LastDeviceId != "" {
		colNames += "last_device_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.LastDeviceId)
	}			
			
	if p.AccountBal != 0.00 {
		colNames += "account_bal,"
		colTags += "?,"
		valSlice = append(valSlice, p.AccountBal)
	}		
			
	if p.GoodsCount != 0.00 {
		colNames += "goods_count,"
		colTags += "?,"
		valSlice = append(valSlice, p.GoodsCount)
	}		
		
	if p.DeviceIp != "" {
		colNames += "device_ip,"
		colTags += "?,"
		valSlice = append(valSlice, p.DeviceIp)
	}			
	
	if p.DeviceType != 0 {
		colNames += "device_type,"
		colTags += "?,"
		valSlice = append(valSlice, p.DeviceType)
	}				
		
	if p.DeviceOs != "" {
		colNames += "device_os,"
		colTags += "?,"
		valSlice = append(valSlice, p.DeviceOs)
	}			
		
	if p.DeviceOsVer != "" {
		colNames += "device_os_ver,"
		colTags += "?,"
		valSlice = append(valSlice, p.DeviceOsVer)
	}			
		
	if p.DeviceId != "" {
		colNames += "device_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.DeviceId)
	}			
		
	if p.Latitude != "" {
		colNames += "latitude,"
		colTags += "?,"
		valSlice = append(valSlice, p.Latitude)
	}			
		
	if p.Longitude != "" {
		colNames += "longitude,"
		colTags += "?,"
		valSlice = append(valSlice, p.Longitude)
	}			
		
	if p.Market != "" {
		colNames += "market,"
		colTags += "?,"
		valSlice = append(valSlice, p.Market)
	}			
		
	if p.UserChannel != "" {
		colNames += "user_channel,"
		colTags += "?,"
		valSlice = append(valSlice, p.UserChannel)
	}			
	
	if p.RandomNo != 0 {
		colNames += "random_no,"
		colTags += "?,"
		valSlice = append(valSlice, p.RandomNo)
	}				
		
	if p.RegionNo != "" {
		colNames += "region_no,"
		colTags += "?,"
		valSlice = append(valSlice, p.RegionNo)
	}			
	
	if p.CustomerId != 0 {
		colNames += "customer_id,"
		colTags += "?,"
		valSlice = append(valSlice, p.CustomerId)
	}				
		
	if p.CreatedTime != "" {
		colNames += "created_time,"
		colTags += "?,"
		valSlice = append(valSlice, p.CreatedTime)
	}			
		
	if p.UpdatedTime != "" {
		colNames += "updated_time,"
		colTags += "?,"
		valSlice = append(valSlice, p.UpdatedTime)
	}			
		
	if p.Memo != "" {
		colNames += "memo,"
		colTags += "?,"
		valSlice = append(valSlice, p.Memo)
	}			
		
	if p.InsertUser != "" {
		colNames += "insert_user,"
		colTags += "?,"
		valSlice = append(valSlice, p.InsertUser)
	}			
		
	if p.UpdateUser != "" {
		colNames += "update_user,"
		colTags += "?,"
		valSlice = append(valSlice, p.UpdateUser)
	}			
	
	if p.Version != 0 {
		colNames += "version,"
		colTags += "?,"
		valSlice = append(valSlice, p.Version)
	}				
	
	colNames = strings.TrimRight(colNames, ",")
	colTags = strings.TrimRight(colTags, ",")
	exeSql := fmt.Sprintf("Insert into  tba_accounts(%s)  values(%s)", colNames, colTags)
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

func (r AccountList) InsertMap(m map[string]interface{}) error {
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

	exeSql := fmt.Sprintf("Insert into  tba_accounts(%s)  values(%s)", colNames, colTags)
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

func (r AccountList) UpdateMap(keyNo string, m map[string]interface{}) error {
	l := time.Now()

	var colNames string
	valSlice := make([]interface{}, 0)
	for k, v := range m {
		colNames += k + "=?,"
		valSlice = append(valSlice, v)
	}
	valSlice = append(valSlice, keyNo)
	colNames = strings.TrimRight(colNames, ",")
	updateSql := fmt.Sprintf("Update tba_accounts set %s where user_id=?", colNames)

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

func (r AccountList) Delete(keyNo string) error {
	l := time.Now()
	delSql := fmt.Sprintf("Delete from  tba_accounts  where user_id=?")
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

