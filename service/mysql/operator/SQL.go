package operator

import "fmt"

const (
	MYSQL_DB_NULL        = "default"
	MYSQL_DB_MYSQL       = "mysql"
	MYSQL_DB_INFORMATION = "information_schema"
	MYSQL_DB_SYS         = "sys"
	MYSQL_DB_PERFORMANCE = "performance_schema"
)

// user
var SqlShowUsers = "select * from mysql.user where Host !='localhost'"
var SqlShowUserDetail = func(name, host string) string {
	return fmt.Sprintf("select * from mysql.user where User =  '%s' and Host = '%s'", name, host)
}
var AddUser = func(name string, password string, host string) string {
	return fmt.Sprintf("create user '%s'@'%s' identified by '%s'", name, host, password)
}
var DropUser = func(name string, host string) string {
	return fmt.Sprintf("drop user '%s'@'%s'", name, host)
}
var UpdateName = func(oldName string, newName string, host string) string {
	return fmt.Sprintf("rename user '%s'@'%s' to '%s'@'%s'", oldName, host, newName, host)
}
var UpdatePassword = func(name string, password string, host string) string {
	return fmt.Sprintf("alter user '%s'@'%s' identified with mysql_native_password by '%s'", name, host, password)
}
var LockUser = func(name string, host string) string {
	return fmt.Sprintf("alter user '%s'@'%s' account lock", name, host)
}
var UnLockUser = func(name string, host string) string {
	return fmt.Sprintf("alter user '%s'@'%s' account unlock", name, host)
}

// privilege
var SqlGrantUserDatabase = func(name string, host string, privilege string, db string) string {
	return fmt.Sprintf("grant %s on %s.* to '%s'@'%s'", privilege, db, name, host)
}
var SqlRevokeUserDatabase = func(name string, host string, privilege string, db string) string {
	return fmt.Sprintf("revoke %s on %s.* from '%s'@'%s'", privilege, db, name, host)
}
var SqlShowDatabasePrivilege = func(name, host string) string {
	return fmt.Sprintf(" SELECT GRANTEE,TABLE_SCHEMA,GROUP_CONCAT(PRIVILEGE_TYPE) PRIVILEGE, IS_GRANTABLE "+
		"FROM information_schema.SCHEMA_PRIVILEGES WHERE GRANTEE = \"'%s'@'%s'\" GROUP BY TABLE_SCHEMA ", name, host)
}
var SqlShowTablePrivilege = func(name, host string) string {
	return fmt.Sprintf(" SELECT TABLE_SCHEMA,TABLE_NAME,concat(TABLE_SCHEMA,\"_\",TABLE_NAME) "+
		"SCHEMA_TABLE,group_concat(PRIVILEGE_TYPE) PRIVILEGE,IS_GRANTABLE FROM information_schema.TABLE_PRIVILEGES WHERE GRANTEE = \"'%s'@'%s'\" GROUP BY SCHEMA_TABLE ", name, host)
}
var SqlGrantUserTable = func(name string, host string, privilege string, db string, table string) string {
	return fmt.Sprintf("grant %s on %s.%s to '%s'@'%s'", privilege, db, table, name, host)
}
var SqlRevokeUserTable = func(name string, host string, privilege string, db string, table string) string {
	return fmt.Sprintf("revoke %s on %s.%s from '%s'@'%s'", privilege, db, table, name, host)
}

// db
var SqlShowDatabases = "SELECT SCHEMA_NAME,DEFAULT_CHARACTER_SET_NAME,DEFAULT_COLLATION_NAME FROM SCHEMATA"

var SqlShowMasterStatus = "SHOW MASTER STATUS"

var SqlShowSlaveStatus = "SHOW SLAVE STATUS"

var GrantDevopsUser = func(password string) string {
	return fmt.Sprintf("GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'devops'@'%%' IDENTIFIED BY '%s'", password)
}

var GrantReplUser = func(password string) string {
	return fmt.Sprintf("GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%%' IDENTIFIED BY '%s'", password)
}

var SqlShowDatabaseDetail = func(db string) string {
	return fmt.Sprintf("SELECT SCHEMA_NAME,DEFAULT_CHARACTER_SET_NAME,DEFAULT_COLLATION_NAME FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = '%s' ", db)
}

var SqlCreateDatabase = func(db, options string) string {
	return fmt.Sprintf("CREATE DATABASE `%s` %s ", db, options)
}

var SqlAlterDatabase = func(db, options string) string {
	return fmt.Sprintf("ALTER DATABASE `%s` %s ", db, options)
}

var SqlDropDatabase = func(db string) string {
	return fmt.Sprintf("DROP DATABASE `%s` ", db)
}

var SqlShowCharset = "SHOW CHARSET"

var SqlShowCollation = func(charset string) string {
	return fmt.Sprintf("SHOW COLLATION WHERE CHARSET = '%s'", charset)
}

var SqlShowVariables = func(param string) string {
	if param != "" {
		return fmt.Sprintf("SHOW GLOBAL VARIABLES LIKE \"%s\";", param)
	} else {
		return fmt.Sprintf("SHOW GLOBAL VARIABLES;")
	}
}

// table
var SqlShowEngines = "SELECT * FROM information_schema.`ENGINES` WHERE `support` != 'NO'"
var SqlShowTables = func(db string) string {
	return fmt.Sprintf("SELECT * from information_schema.tables WHERE table_schema='%s'", db)
}

var SqlCreateTable = func(db, table, options string) string {
	return fmt.Sprintf("CREATE TABLE `%s`.`%s` %s", db, table, options)
}

var SqlDropTable = func(db, table string) string {
	return fmt.Sprintf("DROP TABLE `%s`.`%s` ", db, table)
}

var SqlTruncateTable = func(db, table string) string {
	return fmt.Sprintf("TRUNCATE TABLE `%s`.`%s` ", db, table)
}

var SqlShowTableOptions = func(db, table string) string {
	return fmt.Sprintf("SELECT TABLE_SCHEMA,TABLE_NAME,AUTO_INCREMENT,CREATE_OPTIONS,TABLE_COMMENT,TABLE_COLLATION,ROW_FORMAT,ENGINE,CREATE_OPTIONS FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s'", db, table)
}

var SqlShowTableScript = func(db, table string) string {
	return fmt.Sprintf("SHOW CREATE TABLE `%s`.`%s` ", db, table)
}

var SqlShowTableData = func(db, table, order string, offset, pageSize int) string {
	return fmt.Sprintf("SELECT * FROM `%s`.`%s` %s LIMIT %d, %d", db, table, order, offset, pageSize)
}

var CountTable = func(db, table string) string {
	return fmt.Sprintf("SELECT COUNT(*) num FROM `%s`.`%s` ", db, table)
}

var SqlShowTableColumns = func(db, table string) string {
	return fmt.Sprintf("SELECT COLUMN_NAME,COLUMN_TYPE,IS_NULLABLE,DATA_TYPE,COLUMN_COMMENT,COLUMN_KEY,CHARACTER_MAXIMUM_LENGTH,COLUMN_DEFAULT,COLUMN_KEY,EXTRA,COLLATION_NAME FROM information_schema.COLUMNS WHERE table_schema='%s' AND TABLE_NAME = '%s'", db, table)
}

var SqlShowIndices = func(db, table string) string {
	return fmt.Sprintf("SHOW INDEX FROM `%s`.`%s`", db, table)
}

var SqlShowTableForeignKeys = func(db, table string) string {
	return fmt.Sprintf("SELECT a.CONSTRAINT_NAME,a.TABLE_NAME,a.COLUMN_NAME,a.ORDINAL_POSITION,a.REFERENCED_TABLE_SCHEMA, a.REFERENCED_TABLE_NAME, "+
		"a.REFERENCED_COLUMN_NAME,b.UPDATE_RULE,b.DELETE_RULE FROM information_schema.KEY_COLUMN_USAGE a LEFT JOIN information_schema.REFERENTIAL_CONSTRAINTS b "+
		"ON a.CONSTRAINT_SCHEMA = b.CONSTRAINT_SCHEMA AND a.TABLE_NAME = b.TABLE_NAME and a.CONSTRAINT_NAME = b.CONSTRAINT_NAME "+
		"WHERE a.CONSTRAINT_SCHEMA='%s' AND a.TABLE_NAME = '%s' AND a.REFERENCED_TABLE_SCHEMA IS NOT NULL", db, table)
}

var SqlAlterTableName = func(db, table, newTableName string) string {
	return fmt.Sprintf("ALTER TABLE `%s`.`%s` RENAME to `%s`.`%s`", db, table, db, newTableName)
}

var SqlAlterTableOptions = func(db, table, options string) string {
	return fmt.Sprintf("ALTER TABLE `%s`.`%s` %s", db, table, options)
}

var SqlChangeTableColumn = func(db, table, column, newColumnName, options string) string {
	return fmt.Sprintf("ALTER TABLE `%s`.`%s` CHANGE %s %s %s ", db, table, column, newColumnName, options)
}

var SqlAddTableColumn = func(db, table, column, options string) string {
	return fmt.Sprintf("ALTER TABLE `%s`.`%s` ADD COLUMN %s %s ", db, table, column, options)
}

var SqlModifyTableColumn = func(db, table, column, options string) string {
	return fmt.Sprintf("ALTER TABLE `%s`.`%s` MODIFY %s %s ", db, table, column, options)
}

var SqlDropTableColumn = func(db, table, column string) string {
	return fmt.Sprintf("ALTER TABLE `%s`.`%s` DROP COLUMN %s ", db, table, column)
}

var SqlAddTableIndex = func(db, table, column, options string) string {
	return fmt.Sprintf("ALTER TABLE `%s`.`%s` ADD INDEX %s ", db, table, options)
}

var SqlDropTableIndex = func(db, table, index string) string {
	return fmt.Sprintf("ALTER TABLE `%s`.`%s` DROP INDEX %s ", db, table, index)
}

var SqlAddTableForeignKey = func(db, table, options string) string {
	return fmt.Sprintf("ALTER TABLE `%s`.`%s` ADD CONSTRAINT %s  ", db, table, options)
}

var SqlDropTableForeignKey = func(db, table, foreignKey string) string {
	return fmt.Sprintf("ALTER TABLE `%s`.%s`` DROP FOREIGN KEY `%s`  ", db, table, foreignKey)
}
