package sqlscan

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
)

var (
	WRAPPER_DEFAULT = "%s"
	WRAPPER_STRING  = `"%s"`
	WRAPPER_PARAM   = `t.%s`
)

func GenerateFieldList(typeName string, fields []string) string {
	return fmt.Sprintf(fieldList, typeName, typeName, conjoin(",", WRAPPER_DEFAULT, fields))
}

func GenerateFieldListV2(typeName string, fields []string) string {
	return fmt.Sprintf(fieldListV2, typeName, typeName, conjoin(",", WRAPPER_STRING, fields))
}
func GenerateValueList(typeName string, fields []string) string {
	return fmt.Sprintf(valueList, typeName, typeName, conjoin(",", WRAPPER_PARAM, fields))
}

func conjoin(conj string, wrapper string, items []string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return fmt.Sprintf(wrapper, items[0])
	}
	if len(items) == 2 { // "a and b" not "a, and b"
		return fmt.Sprintf(wrapper, items[0]) + conj + fmt.Sprintf(wrapper, items[1])
	}

	pieces := []string{fmt.Sprintf(wrapper, items[0])}
	for _, item := range items[1 : len(items)-1] {
		pieces = append(pieces, conj, fmt.Sprintf(wrapper, item))
	}
	pieces = append(pieces, conj, fmt.Sprintf(wrapper, items[len(items)-1]))

	return strings.Join(pieces, "")
}

const fieldList = `// Returns all field names from %s
func (t *%s) Fields() string {
	return "%s"
}
`

const fieldListV2 = `// Returns all field names from %s
func (t *%s) Fields() []string {
	return []string{%s}
}
`

const valueList = `// Returns all values from %s
func (t *%s) Values() []interface{} {
	return []interface{}{%s}
}
`

func GenerateScanFn(typeName string, fields []string) string {
	funcs := make(map[string]interface{})
	// funcs["conjoin"] = conjoin
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
