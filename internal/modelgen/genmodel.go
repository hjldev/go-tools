package modelgen

import (
	"fmt"

	"github.com/hjldev/go-tools/internal/modelgen/config"
	"github.com/hjldev/go-tools/internal/modelgen/generate"
	"github.com/hjldev/go-tools/internal/modelgen/mysql"
)

func DbToGoStruct(configPath string) {

	conf, err := config.GetConf(configPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	tables, err := mysql.GetTables(conf.Db.User, conf.Db.Password,
		conf.Db.Host, conf.Db.Port, conf.Db.Database)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, tableName := range tables {
		table, err := mysql.GetTableInfo(conf.Db.User, conf.Db.Password,
			conf.Db.Host, conf.Db.Port, conf.Db.Database, tableName)
		if err != nil {
			fmt.Println("Error in selecting column data information from mysql information schema")
			return
		}

		generate.GenerateFile(table, conf.Db.PackageName,
			conf)

	}
	fmt.Println("model文件生成成功！")

}
