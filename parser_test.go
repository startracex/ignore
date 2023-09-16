package ignore

import (
	"testing"
)

func TestGetRules(t *testing.T) {
	content := `
a #
# b
c
#
`
	var rules []string = GetRules(content)
	t.Log(rules)
}
