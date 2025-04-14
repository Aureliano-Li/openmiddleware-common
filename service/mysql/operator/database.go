package operator

func (op *Operator) Login(info Info) ExecResult {
	res, err := op.testConnection(info)
	execResult := ExecResult{
		Success: res,
		Err:     err,
	}
	return execResult
}

// ExecQuery mysql database by SCHEMATA table
func (op *Operator) ShowDatabases(info Info) SQLResult {
	conn := op.getConnection(info, MYSQL_DB_INFORMATION)
	rows, err := conn.Query(SqlShowDatabases)
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowMasterStatus(info Info) SQLResultMap {
	conn := op.getConnection(info, MYSQL_DB_NULL)
	rows, err := conn.Query(SqlShowMasterStatus)
	res := scanSingleMap(rows, err)
	return res
}

func (op *Operator) ShowSlaveStatus(info Info) SQLResultMap {
	conn := op.getConnection(info, MYSQL_DB_NULL)
	rows, err := conn.Query(SqlShowSlaveStatus)
	res := scanSingleMap(rows, err)
	return res
}

func (op *Operator) GrantDevopsUser(info Info, password string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_NULL)
	sql := GrantDevopsUser(password)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) GrantDevopsUserAll(info Info, password string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_NULL)
	sql := GrantDevopsUserAll(password)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) GrantReplUser(info Info, password string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_NULL)
	res, err := conn.Exec(GrantReplUser(password))
	return result(res, err)
}

func (op *Operator) CreateDatabase(info Info, db, options string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_INFORMATION)
	res, err := conn.Exec(SqlCreateDatabase(db, options))
	return result(res, err)
}

func (op *Operator) AlterDatabase(info Info, db, options string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_INFORMATION)
	res, err := conn.Exec(SqlAlterDatabase(db, options))
	return result(res, err)
}

func (op *Operator) DropDatabase(info Info, db string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_INFORMATION)
	res, err := conn.Exec(SqlDropDatabase(db))
	return result(res, err)
}

func (op *Operator) ShowDatabaseDetail(info Info, db string) SQLResult {
	conn := op.getConnection(info, MYSQL_DB_INFORMATION)
	rows, err := conn.Query(SqlShowDatabaseDetail(db))
	res := scan(rows, err)
	return res
}

func (op *Operator) ExecSql(info Info, db, sql string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) ExecQuery(info Info, db, sql string) SQLResult {
	conn := op.getConnection(info, db)
	rows, err := conn.Query(sql)
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowCharset(info Info) SQLResult {
	conn := op.getConnection(info, MYSQL_DB_INFORMATION)
	rows, err := conn.Query(SqlShowCharset)
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowCollation(info Info, charset string) SQLResult {
	conn := op.getConnection(info, MYSQL_DB_INFORMATION)
	rows, err := conn.Query(SqlShowCollation(charset))
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowVariables(info Info, param string) SQLResult {
	conn := op.getConnection(info, MYSQL_DB_INFORMATION)
	rows, err := conn.Query(SqlShowVariables(param))
	res := scan(rows, err)
	return res
}
