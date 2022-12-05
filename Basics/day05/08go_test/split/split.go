package split

import "strings"

// Split 字符串切割
func Split(str, sep string) []string {
	// var ret []string
	var ret = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep) // 获取切割字符串的索引
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):] // 使用len获取sep长度
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}
