package flowservice

import (
	"log"
	"testing"
)

func TestGetproduct(t *testing.T) {
	r := New("root:123456@tcp(10.89.4.213:3306)/likemi", DEBUG)
	s := Search{PageNo: 1, PageSize: 1, ProductId: "", ProductName: ""}
	l, _ := r.GetList(s)
	log.Println(l)

}

//func TestInsert(t *testing.T) {
//	r := New("root:123456@tcp(10.89.4.213:3306)/likemi")
//	p := Product{ProductId: "50000", ProductName: "三鲜水饺", Price: 222.11}
//	e := r.Insert(p)
//	log.Println(e)
//}

func TestUpdate(t *testing.T) {
	r := New("root:123456@tcp(10.89.4.213:3306)/likemi", DEBUG)
	p := Product{ProductId: "50000", ProductName: "三鲜水饺111", Price: 222.11}
	e := r.UpdateAll(p)
	log.Println(e)
}

func TestUpdateMap(t *testing.T) {
	r := New("root:123456@tcp(10.89.4.213:3306)/likemi", DEBUG)
	m1 := make(map[string]string)
	m1["Product_Id"] = "2111"
	m1["Product_Name"] = "ProductName"

	e := r.UpdateByMap("50000", m1)
	log.Println(e)
}

func TestDelete(t *testing.T) {
	r := New("root:123456@tcp(10.89.4.213:3306)/likemi", DEBUG)
	e := r.Delete("5000")
	log.Println(e)
}

func TestGetTotal(t *testing.T) {
	r := New("root:123456@tcp(10.89.4.213:3306)/likemi", DEBUG)
	s := Search{PageNo: 1, PageSize: 1, ProductId: "11", ProductName: "jieke"}
	l, _ := r.GetTotal(s)
	log.Println(l)
}
