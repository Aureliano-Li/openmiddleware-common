package operator

import (
	log "github.com/sirupsen/logrus"
)

func (op *Operator) ShowEngines(info Info) SQLResult {
	conn := op.getConnection(info, MYSQL_DB_INFORMATION)
	rows, err := conn.Query(SqlShowEngines)
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowTables(info Info, db string) SQLResult {
	conn := op.getConnection(info, db)
	rows, err := conn.Query(SqlShowTables(db))
	res := scan(rows, err)
	return res
}

func (op *Operator) CreateTable(info Info, db, table, options string) ExecResult {
	conn := op.getConnection(info, db)
	sql := SqlCreateTable(db, table, options)
	log.Debug(sql)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) DropTable(info Info, db, table string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(SqlDropTable(db, table))
	return result(res, err)
}

func (op *Operator) ShowTableOptions(info Info, db, table string) SQLResult {
	conn := op.getConnection(info, db)
	sql := SqlShowTableOptions(db, table)
	log.Info(sql)
	rows, err := conn.Query(sql)
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowTableScript(info Info, db, table string) SQLResult {
	conn := op.getConnection(info, db)
	rows, err := conn.Query(SqlShowTableScript(db, table))
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowTableData(info Info, db, table, order string, offset, pageSize int) SQLResult {
	conn := op.getConnection(info, db)
	sql := SqlShowTableData(db, table, order, offset, pageSize)
	log.Debug(sql)
	rows, err := conn.Query(sql)
	res := scan(rows, err)
	return res
}

func (op *Operator) CountTable(info Info, db, table string) SQLResult {
	conn := op.getConnection(info, db)
	sql := CountTable(db, table)
	log.Debug(sql)
	rows, err := conn.Query(sql)
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowTableColumns(info Info, db, table string) SQLResult {
	conn := op.getConnection(info, db)
	rows, err := conn.Query(SqlShowTableColumns(db, table))
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowTableIndices(info Info, db, table string) SQLResult {
	conn := op.getConnection(info, db)
	rows, err := conn.Query(SqlShowIndices(db, table))
	res := scan(rows, err)
	return res
}

func (op *Operator) ShowTableForeignKeys(info Info, db, table string) SQLResult {
	conn := op.getConnection(info, db)
	rows, err := conn.Query(SqlShowTableForeignKeys(db, table))
	res := scan(rows, err)
	return res
}

func (op *Operator) AlterTableName(info Info, db, table, newTableName string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(SqlAlterTableName(db, table, newTableName))
	return result(res, err)
}

func (op *Operator) AlterTableOptions(info Info, db, table, options string) ExecResult {
	conn := op.getConnection(info, db)
	sql := SqlAlterTableOptions(db, table, options)
	res, err := conn.Exec(sql)
	return result(res, err)
}

func (op *Operator) AddTableColumn(info Info, db, table, column, options string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(SqlAddTableColumn(db, table, column, options))
	return result(res, err)
}

func (op *Operator) ChangeTableColumn(info Info, db, table, column, newColumnName, options string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(SqlChangeTableColumn(db, table, column, newColumnName, options))
	return result(res, err)
}

func (op *Operator) ModifyTableColumn(info Info, db, table, column, options string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(SqlModifyTableColumn(db, table, column, options))
	return result(res, err)
}

func (op *Operator) DropTableColumn(info Info, db, table, column string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(SqlDropTableColumn(db, table, column))
	return result(res, err)
}

func (op *Operator) AddTableIndex(info Info, db, table, column, options string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(SqlAddTableIndex(db, table, column, options))
	return result(res, err)
}

func (op *Operator) DropTableIndex(info Info, db, table, index string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(SqlDropTableIndex(db, table, index))
	return result(res, err)
}

func (op *Operator) AddTableForeignKey(info Info, db, table, index string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(SqlAddTableForeignKey(db, table, index))
	return result(res, err)
}

func (op *Operator) UpdateTableForeignKeys(info Info, db, options string) ExecResult {
	conn := op.getConnection(info, db)
	log.Info(options)
	res, err := conn.Exec(options)
	return result(res, err)
}

func (op *Operator) DropTableForeignKey(info Info, db, table, index string) ExecResult {
	conn := op.getConnection(info, db)
	res, err := conn.Exec(SqlDropTableForeignKey(db, table, index))
	return result(res, err)
}
