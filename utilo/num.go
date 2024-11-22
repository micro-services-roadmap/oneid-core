package utilo

import (
	"github.com/spf13/cast"
	"strings"
)

func ConvertString2Ids(str string) []int64 {
	if len(str) > 0 {
		var ids []int64
		for _, v := range strings.Split(str, ",") {
			ids = append(ids, cast.ToInt64(v))
		}
		return ids
	}
	return []int64{}
}

func ConvertPrt2Ids(str *string) []int64 {
	if str == nil {
		return []int64{}
	}

	if len(*str) > 0 {
		var ids []int64
		for _, v := range strings.Split(*str, ",") {
			ids = append(ids, cast.ToInt64(v))
		}
		return ids
	}
	return []int64{}
}

func JonIds(ids []int64, sp string) string {
	if len(ids) == 0 {
		return ""
	}

	return strings.Join(cast.ToStringSlice(ids), sp)
}
