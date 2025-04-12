package operator

import (
	log "github.com/sirupsen/logrus"
)

func (op *Operator) ShowUsers(info Info) SQLResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	rows, err := conn.Query(SqlShowUsers)
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowUserDetail(info Info, name, host string) SQLResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := SqlShowUserDetail(name, host)
	log.Debug(sql)
	rows, err := conn.Query(sql)
	res := scan(rows, err)
	return res
}

func (op *Operator) AddUser(info Info, name string, password string, host string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := AddUser(name, password, host)
	log.Debug(sql)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) DropUser(info Info, name string, host string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	res, err := conn.Exec(DropUser(name, host))
	return result(res, err)
}

func (op *Operator) UpdateName(info Info, oldName string, newName string, host string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := UpdateName(oldName, newName, host)
	log.Debug(sql)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) UpdatePassword(info Info, name string, password string, host string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := UpdatePassword(name, password, host)
	log.Debug(sql)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) LockUser(info Info, name string, host string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := LockUser(name, host)
	log.Debug(sql)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) UnLockUser(info Info, name string, host string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	res, err := conn.Exec(UnLockUser(name, host))
	return result(res, err)
}

func (op *Operator) GrantUserDatabase(info Info, name string, host string, privilege string, db string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := SqlGrantUserDatabase(name, host, privilege, db)
	log.Debug(sql)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) RevokeUserDatabase(info Info, name string, host string, privilege string, db string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := SqlRevokeUserDatabase(name, host, privilege, db)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) ShowDatabasePrivilege(info Info, name string, host string) SQLResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := SqlShowDatabasePrivilege(name, host)
	log.Debug(sql)
	rows, err := conn.Query(sql)
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowTablePrivilege(info Info, name string, host string) SQLResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := SqlShowTablePrivilege(name, host)
	log.Debug(sql)
	rows, err := conn.Query(sql)
	res := scan(rows, err)
	return res
}

func (op *Operator) GrantUserTable(info Info, name string, host string, privilege string, db string, table string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := SqlGrantUserTable(name, host, privilege, db, table)
	log.Debug(sql)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) RevokeUserTable(info Info, name string, host string, privilege string, db string, table string) ExecResult {
	conn := op.getConnection(info, MYSQL_DB_MYSQL)
	sql := SqlRevokeUserTable(name, host, privilege, db, table)
	res, err := conn.Exec(sql)
	return result(res, err)
}
