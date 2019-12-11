package mysql

import (
	"context"
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
)

type DaoMysqlSchema interface {
	TableName() string
}

// GetOne gets one record from table user by condition "where"
func GetOne(tableName string,where map[string]interface{},result DaoMysqlSchema) error{
	if nil == Db {
		return  errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildSelect(tableName, where, nil)
	if nil != err {
		return err
	}

	row, err := Db.Query(cond, vals...)
	if nil != err || nil == row {
		return  err
	}
	defer row.Close()
	return scanner.Scan(row, result)
}

// GetMulti gets multiple records from table user by condition "where"
func GetMulti(tableName string,where map[string]interface{},results []DaoMysqlSchema) error {
	if nil == Db {
		return errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildSelect(tableName, where, nil)
	if nil != err {
		return  err
	}
	row, err := Db.Query(cond, vals...)
	if nil != err || nil == row {
		return  err
	}
	defer row.Close()
	return scanner.Scan(row, results)
}

// Insert inserts an array of data into table user
func Insert(tableName string,data []map[string]interface{}) (int64, error) {
	if nil == Db {
		return -1, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildInsert(tableName, data)
	if nil != err {
		return -1, err
	}
	result, err := Db.Exec(cond, vals...)
	if nil != err || nil == result {
		return -1, err
	}
	return result.LastInsertId()
}

// Update updates the table user
func Update(tableName string,where, data map[string]interface{}) (int64, error) {
	if nil == Db {
		return -1, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildUpdate(tableName, where, data)
	if nil != err {
		return -1, err
	}
	result, err := Db.Exec(cond, vals...)
	if nil != err {
		return -1, err
	}
	return result.RowsAffected()
}

// Delete deletes matched records in user
func Delete(tableName string,where map[string]interface{}) (int64, error) {
	if nil == Db {
		return -1, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildDelete(tableName, where)
	if nil != err {
		return -1, err
	}
	result, err := Db.Exec(cond, vals...)
	if nil != err {
		return -1, err
	}
	return result.RowsAffected()
}

// GetCount
func GetCount(tableName string,where map[string]interface{}) (int64, error) {
	if nil == Db {
		return -1, errors.New("sql.DB object couldn't be nil")
	}
	res, err := builder.AggregateQuery(context.TODO(), Db, tableName, where, builder.AggregateCount("*"))
	if nil != err {
		return -1, err
	}

	return res.Int64(), err
}