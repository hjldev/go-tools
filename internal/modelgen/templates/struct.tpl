package {{.PackageName}}

type {{.ModelName}} struct {
	{{.Fields}}
}

{{ if .Option.GormAnnotation }}
	// TableName sets the insert table name for this struct type
    func (model *{{.ModelName}})TableName() string {
    	return "{{.Table.Name}}"
    }
{{ end }}



