package mysql

import (
	"fmt"
	"github.com/Aureliano-Li/openmiddleware-common/service/mysql/operator"
	"github.com/Aureliano-Li/openmiddleware-common/service/mysql/service"
	"testing"
)

func TestLogin(t *testing.T) {
	info := operator.Info{
		Username: "devops",
		Password: "OpenMiddleware@01",
		Path:     "10.10.104.88",
		Port:     "30514",
	}
	exists := service.CheckUserExists(info)
	fmt.Println(exists)

}
