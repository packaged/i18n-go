package i18n

import (
	"fmt"
	"strings"
)

func Replacements(original string, args map[string]interface{}) string {
	if args == nil {
		return original
	}
	for k, v := range args {
		original = strings.ReplaceAll(original, "{"+k+"}", fmt.Sprintf("%v", v))
	}
	return original
}
