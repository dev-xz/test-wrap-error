package resource

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"test-wrap-error/library/conf"
)

const (
	MySQLTestDB1 string = "test_db"
)

var MySQLClientPool = map[string]*sql.DB{}

func InitMysql() {
	// 注册MySQL集群
	MySQLClientPool[MySQLTestDB1] = mustInitOneMySQL(MySQLTestDB1)
}

func mustInitOneMySQL(name string) *sql.DB {
	if _, exist := conf.DBConfs[name]; !exist {
		panic("load mysql conf " + name + " not exist")
	}
	dbConf := conf.DBConfs[name]
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbConf.Username, dbConf.Passwd, dbConf.Host, dbConf.Port, dbConf.DbName)
	client, err := sql.Open(dbConf.Type, dataSource)
	// 理论上，只有配置错误，才有可能返回err
	// 如配置不存在，配置内容错误
	if err != nil {
		panic(err.Error())
	}
	return client
}

// MySQLClient MySQLClient
func MySQLClient(cluster string) *sql.DB {
	if client, have := MySQLClientPool[cluster]; !have {
		panic(fmt.Errorf("MySQL cluster %s doesn't register", cluster))
	} else {
		return client
	}
}
