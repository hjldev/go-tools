package mysql

import (
	"bytes"
	"fmt"

	"github.com/hjldev/go-tools/internal/modelgen/model"
)

func generateGormAnnotation(col model.Column) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("column:%s", col.Name))
	if col.PrimaryKey {
		buf.WriteString(";primaryKey")
	}
	if col.Unique {
		buf.WriteString(";unique")
	}
	if col.NotNull {
		buf.WriteString(";not null")
	}
	if col.AutoInc {
		buf.WriteString(";autoIncrement")
	}
	return fmt.Sprintf(`gorm:"%s"`, buf.String())
}
func generateJSONAnnotation(col model.Column) string {
	return fmt.Sprintf(`json:"%s"`, col.Name)
}
func generateDBAnnotation(col model.Column) string {
	return fmt.Sprintf(`db:"%s"`, col.Name)
}
func generateXMLAnnotation(col model.Column) string {
	return fmt.Sprintf(`xml:"%s"`, col.Name)
}
func generateXormAnnotation(col model.Column) string {
	return fmt.Sprintf(`xorm:"%s"`, col.Name)
}
func generateFakerAnnotation(col model.Column) string {
	return fmt.Sprintf(`faker:"%s"`, col.Name)
}
