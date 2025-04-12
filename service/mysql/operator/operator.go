package operator

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

type Operator struct {
	mysqlClients map[string]*Client
}

var operatorInstance *Operator

func GetOperatorInstance() *Operator {
	if operatorInstance == nil {
		operatorInstance = &Operator{
			mysqlClients: map[string]*Client{},
		}
	}
	return operatorInstance
}

func (op *Operator) getConnection(info Info, dbName string) *sql.DB {
	var conn *sql.DB
	// Add username as mysql client key,in order to prevent different authorities user use same connection.
	name := info.Path + "/" + info.Port + "/" + info.Username
	client := op.mysqlClients[name]
	if client == nil {
		res, _ := op.testConnection(info)
		if res {
			logrus.Info("添加mysql连接:" + name)
			op.addConnection(info)
			client = op.mysqlClients[name]
		}
	}
	conn, _ = client.GetConnection(dbName)
	return conn
}

func (op *Operator) createConnection(info Info) (*sql.DB, error) {
	client := &Client{
		Mysql: info,
		DBs:   map[string]*sql.DB{},
	}
	return client.GetConnection(MYSQL_DB_INFORMATION)
}

func (op *Operator) testConnection(info Info) (bool, error) {
	client := &Client{
		Mysql: info,
		DBs:   map[string]*sql.DB{},
	}
	con, err := client.GetConnection(MYSQL_DB_INFORMATION)
	if err != nil {
		return false, err
	}
	err = con.Ping()
	if err != nil {
		return false, err
	}
	op.addConnection(info)
	return true, nil
}

func (op *Operator) addConnection(info Info) {
	name := info.Path + "/" + info.Port + "/" + info.Username
	client := &Client{
		Mysql: info,
		DBs:   map[string]*sql.DB{},
	}
	op.mysqlClients[name] = client
}

func (op *Operator) RemoveConnection(info Info) {
	name := info.Path + "/" + info.Port + "/" + info.Username
	client := op.mysqlClients[name]
	if client != nil {
		client.CloseDb()
	}
	logrus.Info("移除mysql连接:" + name)
	delete(op.mysqlClients, name)
}

type SQLResult struct {
	Column []string   `json:"column"`
	Data   [][]string `json:"data"`
	Err    error      `json:"err"`
}

type SQLResultMap struct {
	Column    []string          `json:"column"`
	Data      [][]string        `json:"data"`
	ResultMap map[string]string `json:"resultMap"`
	Err       error             `json:"err"`
}

func scan(r *sql.Rows, err error) SQLResult {
	sqlRes := SQLResult{}
	var values []interface{}

	if err != nil {
		logrus.Error("ExecQuery error", err)
		return SQLResult{Err: err}
	}

	sqlRes.Column, _ = r.Columns()

	for r.Next() {
		values = make([]interface{}, len(sqlRes.Column))
		for i := range values {
			tmp := sql.NullString{}
			values[i] = &tmp
		}
		var line []string
		_ = r.Scan(values...)
		for _, value := range values {
			v := value.(*sql.NullString)
			line = append(line, v.String)
		}
		sqlRes.Data = append(sqlRes.Data, line)
	}

	return sqlRes
}

func scanSingleMap(r *sql.Rows, err error) SQLResultMap {
	sqlResMap := SQLResultMap{}
	var values []interface{}

	if err != nil {
		logrus.Error("ExecQuery error", err)
		return SQLResultMap{Err: err}
	}

	sqlResMap.Column, _ = r.Columns()

	hasData := false
	for r.Next() {
		hasData = true
		values = make([]interface{}, len(sqlResMap.Column))
		for i := range values {
			tmp := sql.NullString{}
			values[i] = &tmp
		}
		var line []string
		_ = r.Scan(values...)
		for _, value := range values {
			v := value.(*sql.NullString)
			line = append(line, v.String)
		}
		sqlResMap.Data = append(sqlResMap.Data, line)
	}
	if !hasData {
		return sqlResMap
	}
	data := sqlResMap.Data
	resMap := make(map[string]string)
	for index, columnName := range sqlResMap.Column {
		resMap[columnName] = data[0][index]
	}
	sqlResMap.ResultMap = resMap
	return sqlResMap
}

type ExecResult struct {
	*SQLResult
	LastInsertId int64       `json:"lastInsertId"`
	RowsAffected int64       `json:"rowsAffected"`
	Err          error       `json:"err"`
	ExecTime     int64       `json:"execTime"`
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
}

func result(res sql.Result, err error) ExecResult {
	execResult := ExecResult{}

	if err != nil {
		execResult.Err = err
	} else {
		execResult.LastInsertId, _ = res.LastInsertId()
		execResult.RowsAffected, _ = res.RowsAffected()
	}
	return execResult
}

type CommonResult struct {
	Success  bool   `json:"success"`
	ErrorMsg string `json:"errorMsg"`
	Data     string `json:"data"`
	Err      error  `json:"err"`
}
