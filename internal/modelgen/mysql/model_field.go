package mysql

import (
	"fmt"
	"strings"

	"github.com/hjldev/go-tools/internal/modelgen/config"
	"github.com/hjldev/go-tools/internal/modelgen/model"
	"github.com/hjldev/go-tools/internal/modelgen/utils"
)

func GenerateModelFields(table *model.Table, depth int, option *config.Config) string {

	structure := ""

	for _, column := range table.Columns {

		// Get the corresponding go value type for this mysql type
		valueType := mysqlTypeToGoType(column.Type, !column.NotNull, false)

		fieldName := utils.FmtFieldName(utils.StringifyFirstChar(column.Name))
		var annotations []string
		if option.Db.GormAnnotation {
			annotations = append(annotations, generateGormAnnotation(column))
		}
		if option.Db.JSONAnnotation {
			annotations = append(annotations, generateJSONAnnotation(column))
		}
		if option.Db.DBAnnotation {
			annotations = append(annotations, generateDBAnnotation(column))
		}
		if option.Db.XMLAnnotation {
			annotations = append(annotations, generateXMLAnnotation(column))
		}
		if len(annotations) > 0 {
			structure += fmt.Sprintf("\n%s %s `%s`",
				fieldName,
				valueType,
				strings.Join(annotations, " "))

		} else {
			structure += fmt.Sprintf("\n%s %s",
				fieldName,
				valueType)
		}
	}
	return structure
}
