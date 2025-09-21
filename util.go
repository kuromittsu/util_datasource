package util_datasource

import (
	"fmt"
	"strings"
)

func queryReplacePlaceholder(rawQuery, replaceWith string) string {

	if len(replaceWith) == 0 {
		return rawQuery
	}

	count := strings.Count(rawQuery, "?")
	for i := 1; i <= count; i++ {
		placeholder := fmt.Sprintf("%s%d", replaceWith, i)
		rawQuery = strings.Replace(rawQuery, "?", placeholder, 1)
	}
	return rawQuery
}
