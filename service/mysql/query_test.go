package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestLink(t *testing.T) {
	// 数据库连接信息
	dsn := "root:OpenMiddleware@01@tcp(10.10.104.88:30514)/"
	// 数据库连接信息，不指定数据库名

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// 检查连接
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// 执行 SHOW MASTER STATUS 查询
	row := db.QueryRow("SHOW MASTER STATUS")

	// 定义变量来存储结果
	var file, binlogDoDB, binlogIgnoreDB string
	var position uint64
	var executedGtidSet sql.NullString // 可能为空

	// 扫描结果
	err = row.Scan(&file, &position, &binlogDoDB, &binlogIgnoreDB, &executedGtidSet)
	if err != nil {
		log.Fatal("Failed to scan master status:", err)
	}

	// 输出结果
	fmt.Printf("Binary Log File: %s\n", file)
	fmt.Printf("Position: %d\n", position)
	fmt.Printf("Binlog Do DB: %s\n", binlogDoDB)
	fmt.Printf("Binlog Ignore DB: %s\n", binlogIgnoreDB)
	if executedGtidSet.Valid {
		fmt.Printf("Executed GTID Set: %s\n", executedGtidSet.String)
	} else {
		fmt.Println("Executed GTID Set: NULL")
	}
}
