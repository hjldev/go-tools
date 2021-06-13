package generate

import (
	"fmt"
	"html/template"
	"os"

	"github.com/hjldev/go-tools/internal/modelgen/config"
	"github.com/hjldev/go-tools/internal/modelgen/model"
	"github.com/hjldev/go-tools/internal/modelgen/mysql"
	"github.com/hjldev/go-tools/internal/modelgen/utils"
)

func GenerateFile(table *model.Table, pkgName string, conf *config.Config) {
	if pkgName == "" {
		pkgName = "models"
	}

	structName := utils.FmtFieldName(table.Name)

	fields := mysql.GenerateModelFields(table, 0, conf)

	structPath := conf.GenPath

	if conf.GenPath == "" {
		structPath = "./extra/modelgen/models"
	}

	os.MkdirAll(structPath, os.ModePerm)

	structFileName := structPath + "/" + structName + ".go"

	structFile, _ := os.Create(structFileName)

	structTemplate, err := template.ParseFiles("./internal/modelgen/templates/struct.tpl")

	if err != nil {
		fmt.Println(err)
		return
	}

	if err := structTemplate.Execute(structFile, map[string]interface{}{
		"Table":       table,
		"PackageName": pkgName,
		"ModelName":   structName,
		"Option":      conf.Db,
		"Fields":      template.HTML(fields), // 防止自动转义
	}); err != nil {
		fmt.Println("execute template error:", err)
	}

}
