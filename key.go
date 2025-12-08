package translate

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	minWord      int64 = 3
	prefixLength int   = 30
)

func Key(input string) string {
	var replaced string
	if len(input) > prefixLength {
		replaced = regexp.MustCompile(fmt.Sprintf(`\b\w{0,%d}\b|\s`, minWord)).ReplaceAllString(input, " ")
	} else {
		replaced = input
	}

	re := regexp.MustCompile(`[A-Za-z]+`)
	long := len(re.FindAllString(input, -1)) > 3

	replaced = strings.ToLower(regexp.MustCompile(`\W+`).ReplaceAllString(strings.TrimSpace(replaced), "_"))
	if len(replaced) > prefixLength {
		replaced = strings.ToLower(strings.Trim(replaced[0:prefixLength], "_"))
	}

	replaced = strings.Trim(replaced, "_")

	hash := md5.Sum([]byte(input))
	md5Hash := fmt.Sprintf("%x", hash)

	if long {
		replaced += "_" + md5Hash[0:6]
		replaced += "_" + strconv.FormatInt(int64(len(input)), 10)
	} else {
		replaced += "_" + md5Hash[0:4]
	}
	return replaced
}
