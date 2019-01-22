package code

import (
	"log"
	"testing"
)

func TestGetproduct(t *testing.T) {
	r := New("root:123456@tcp(10.89.4.203:3306)/tbar_db", DEBUG)
	s := Search{Table_name: ""}
	l, _ := r.GetCodeMap(s)

	e, ok := l.Load("FLWSTS")

	tvConf, ok := e.([]Code)

	log.Println("======", tvConf, ok)

}
