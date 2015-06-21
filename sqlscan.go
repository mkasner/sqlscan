package sqlscan

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
)

func GenerateFieldList(typeName string, fields []string) string {
	return fmt.Sprintf(fieldList, typeName, typeName, conjoin(",", fields))
}

func conjoin(conj string, items []string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return items[0]
	}
	if len(items) == 2 { // "a and b" not "a, and b"
		return items[0] + conj + items[1]
	}

	pieces := []string{items[0]}
	for _, item := range items[1 : len(items)-1] {
		pieces = append(pieces, conj, item)
	}
	pieces = append(pieces, conj, items[len(items)-1])

	return strings.Join(pieces, "")
}

const fieldList = `// Returns all field names from %s
func (t *%s) Fields() string {
	return "%s"
}
`

func GenerateScanFn(typeName string, fields []string) string {
	funcs := make(map[string]interface{})
	funcs["conjoin"] = conjoin
	scanFunctions := []string{scanFn, scanRowFn}
	var buff bytes.Buffer
	for _, sf := range scanFunctions {
		tmpl, err := template.New("").Funcs(funcs).Parse(sf)
		if err != nil {
			log.Println(err.Error())
			return ""
		}
		err = tmpl.Execute(&buff, struct {
			TypeName string
			Fields   []string
		}{
			typeName,
			fields,
		})
		if err != nil {
			log.Println(err.Error())
			return ""
		}
	}
	return buff.String()
}

const scanFn = `// // Scans to {{.TypeName}}
func (t *{{.TypeName}}) Scan(rows *sql.Rows) ({{.TypeName}}, error) {
	var r {{.TypeName}}
	err := rows.Scan(
		{{range .Fields}}&r.{{.}},
		{{end}}
	)
	return r, err
}
`

const scanRowFn = `// // Scans to {{.TypeName}}
func (t *{{.TypeName}}) ScanRow(row *sql.Row) ({{.TypeName}}, error) {
	var r {{.TypeName}}
	err := row.Scan(
		{{range .Fields}}&r.{{.}},
		{{end}}
	)
	return r, err
}
`
