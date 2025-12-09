package i18n

import (
	"strings"

	"github.com/packaged/helpers-go"
)

// ENPlural allows an EN conversion of {QTY} bag(s) full
func ENPlural(txt string, quantity int64) string {
	txt = strings.ReplaceAll(txt, "(s)", helpers.If(quantity == 1, "", "s"))
	return txt
}
