package applicant

import (
	"log"
	"testing"
)

func TestGetproduct(t *testing.T) {
	r := New("root:123456@tcp(10.89.4.203:3306)/tbar_db", DEBUG)
	s := Search{Table_name: ""}
	l, _ := r.GetFieldMap(s)

	e, ok := l.Load("CHANNEL_NAME")

	tvConf, _ := e.(ApplicantFields)

	log.Println("======", tvConf.Field_desc, ok)

}
