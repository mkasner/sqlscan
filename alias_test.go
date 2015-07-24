package sqlscan

import "testing"

type aliasTestStruct struct {
	fields []string
	alias  string
	result []string
}

var (
	aliasTestData = []aliasTestStruct{
		aliasTestStruct{
			fields: []string{"HELLO", "WORLD"},
			alias:  "dt",
			result: []string{"dt.HELLO", "dt.WORLD"},
		},
	}
)

func TestAddAlias(t *testing.T) {
	for _, td := range aliasTestData {
		result := AddAlias(td.alias, td.fields)
		for i, _ := range td.result {
			if result[i] != td.result[i] {
				t.Errorf("aliased fileds don't match: expected: %s result: %s", td.result[i], result[i])
				t.Fail()
			}
		}
	}
}
