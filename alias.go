package sqlscan

import "fmt"

var (
	tplALIAS = "%s.%s"
)

func AddAlias(alias string, fields []string) []string {
	for i, _ := range fields {
		fields[i] = fmt.Sprintf(tplALIAS, alias, fields[i])
	}
	return fields
}
