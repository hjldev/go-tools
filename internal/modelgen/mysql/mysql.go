package mysql

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hjldev/go-tools/internal/modelgen/model"
)

func GetTableInfo(user string, password string, host string, port int, dbname string, tableName string) (*model.Table, error) {

	var err error
	var db *sql.DB
	if password != "" {
		db, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbname+"?&parseTime=True")
	} else {
		db, err = sql.Open("mysql", user+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbname+"?&parseTime=True")
	}

	// Check for error in db, note this does not check connectivity but does check uri
	if err != nil {
		fmt.Println("Error opening mysql db: " + err.Error())
		return nil, err
	}

	defer db.Close()

	// Select column data from INFORMATION_SCHEMA
	columnDataTypeQuery := "SELECT `COLUMN_NAME`, COLUMN_KEY, DATA_TYPE,COLUMN_DEFAULT, IS_NULLABLE,CHARACTER_MAXIMUM_LENGTH,EXTRA,COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND table_name = ?"

	rows, err := db.Query(columnDataTypeQuery, dbname, tableName)
	if err != nil {
		fmt.Println("Error selecting from db: " + err.Error())
		return nil, err
	}

	table := &model.Table{
		Name:    tableName,
		Columns: make([]model.Column, 0),
	}
	for rows.Next() {
		var columnName string
		var columnKey string
		var dataType string
		var defaultValue sql.NullString
		var nullable string
		var maxLength sql.NullInt64
		var extra string
		var comment string

		if err := rows.Scan(&columnName, &columnKey, &dataType, &defaultValue, &nullable, &maxLength, &extra, &comment); err != nil {
			fmt.Println("scan rows error:", err)
			return nil, err
		}

		table.Columns = append(table.Columns, model.Column{
			Name:       columnName,
			Comment:    comment,
			Type:       dataType,
			Size:       int(maxLength.Int64),
			NotNull:    !(nullable == "YES"),
			Default:    defaultValue.String,
			AutoInc:    isAutoInc(extra),
			Unique:     isUnique(columnKey),
			PrimaryKey: isPrimaryKey(columnKey),
		})
	}

	return table, nil
}

func GetTables(user string, password string, host string, port int, dbName string) ([]string, error) {
	var err error
	var db *sql.DB
	if password != "" {
		db, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbName+"?&parseTime=True")
	} else {
		db, err = sql.Open("mysql", user+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbName+"?&parseTime=True")
	}

	// Check for error in db, note this does not check connectivity but does check uri
	if err != nil {
		fmt.Println("Error opening mysql db: " + err.Error())
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA= ?", dbName)
	if err != nil {
		fmt.Println("query tables error:", err)
		return nil, err
	}

	tables := make([]string, 0)
	for rows.Next() {
		var tbName string
		rows.Scan(&tbName)
		tables = append(tables, tbName)
	}
	return tables, nil
}

func isAutoInc(extra string) bool {
	return strings.Contains(extra, "auto_increment")
}

func isUnique(columnKey string) bool {
	return isPrimaryKey(columnKey) || strings.Contains(columnKey, "UNI")
}

func isPrimaryKey(columnKey string) bool {
	return strings.Contains(columnKey, "PRI")
}
