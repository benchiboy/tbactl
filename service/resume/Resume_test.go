package resume

import (
	"log"
	"testing"
)

//func TestGetproduct(t *testing.T) {
//	r := NewUrl("root:123456@tcp(10.89.4.203:3306)/tba2_db", DEBUG)
//	s := Search{PageNo: 1, PageSize: 10}
//	l, _ := r.GetList(s)
//	log.Println(l)

//}

func TestInsert(t *testing.T) {
	r := NewUrl("root:123456@tcp(10.89.4.203:3306)/tba2_db", DEBUG)

	m := map[string]interface{}{"post_no": "2222222",
		"user_id": 2000}
	e := r.InsertMap(m)
	e = r.InsertMap(m)
	e = r.InsertMap(m)
	e = r.InsertMap(m)
	e = r.InsertMap(m)
	e = r.InsertMap(m)

	log.Println(e)

}

//func TestUpdate(t *testing.T) {
//	r := NewUrl("root:123456@tcp(10.89.4.203:3306)/tba2_db", DEBUG)

//	m := map[string]interface{}{"publish_no": "2222222",
//		"user_id":       1000111,
//		"salary_min":    4000.54,
//		"position_desc": "这是一个测试'222'"}

//	e := r.UpdateMap("2222222", m)
//	log.Println(e)
//}
