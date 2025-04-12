package mysql

import (
	"fmt"
	"github.com/Aureliano-Li/openmiddleware-common/service/mysql/operator"
	"testing"
)

func TestShow(t *testing.T) {
	info := operator.Info{
		Username: "root",
		Password: "OpenMiddleware@01",
		Path:     "10.10.104.88",
		Port:     "30514",
	}
	instance := operator.GetOperatorInstance()
	login := instance.Login(info)
	fmt.Println(login)
	masterStatus := instance.ShowMasterStatus(info)
	fmt.Println(masterStatus.ResultMap["File"])
	fmt.Println(masterStatus.ResultMap["Position"])
	slaveStatus := instance.ShowSlaveStatus(info)
	fmt.Println(slaveStatus)
	sql := instance.ExecSql(info, "default", "GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'devops'@'localhost' IDENTIFIED BY 'OpenMiddleware@01'")
	fmt.Println(sql)
	psd := "OpenMiddleware@01"
	user := instance.GrantDevopsUser(info, psd)
	fmt.Println(user)
}
