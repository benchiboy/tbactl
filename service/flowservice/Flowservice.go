package flowservice

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type FlowService struct {
}

/*
	得到服务节点
*/
func (r *FlowService) GetCurrPoint() ([]ServicePoint, error) {

	l := time.Now()

	return
}
