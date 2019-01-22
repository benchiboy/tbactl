package point

import (
	"log"
	"tbactl/service/flowservice/applicant"
	"testing"
)

func TestServicePoint(t *testing.T) {
	r := New("root:123456@tcp(10.89.4.203:3306)/tbar_db", DEBUG)
	m, _ := r.GetServicePointsMap()

	e, ok := m.Load("10")
	tvConf, ok := e.([]applicant.ApplicantFields)

	log.Println("======", tvConf, ok)

}
