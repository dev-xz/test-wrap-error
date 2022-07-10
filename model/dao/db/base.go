package db

import (
	"database/sql"
	"github.com/pkg/errors"
	"test-wrap-error/library/resource"
)

type BaseDbDao struct {
	ClusterName string
	DB          *sql.DB
}

func newDbDao(cluster string) *BaseDbDao {
	return &BaseDbDao{
		ClusterName: cluster,
		DB:          resource.MySQLClient(cluster),
	}
}

func (dao *BaseDbDao) BaseSelect(sqlStr string, args ...any) ([]map[string]interface{}, error) {
	rows, err := dao.DB.Query(sqlStr, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "db query err, cluster: %s, sql: %s, args: %v", dao.ClusterName, sqlStr, args)
	}
	return rowsToMap(rows), nil
}

func (dao *BaseDbDao) BaseSelectOne(sqlStr string, args ...any) (map[string]interface{}, error) {
	result, err := dao.BaseSelect(sqlStr, args...)
	if err == nil {
		if len(result) == 0 {
			return nil, errors.Wrapf(sql.ErrNoRows, "db query one no data, cluster: %s, sql: %s, args: %v", dao.ClusterName, sqlStr, args)
		}
		return result[0], err
	}
	return nil, err
}

func rowsToMap(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	columnLength := len(columns)
	// 临时存储每行数据
	cache := make([]interface{}, columnLength)
	// 为每一 列初始化一个指针
	for index, _ := range cache {
		var a interface{}
		cache[index] = &a
	}

	// 返回的结果
	var list []map[string]interface{}
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			// 取实际类型
			item[columns[i]] = *data.(*interface{})
		}
		list = append(list, item)
	}
	_ = rows.Close()
	return list
}
