package split

import (
	"reflect"
	"testing"
)


func TestGroupSplit(t *testing.T) {
	// 定义一个测试用例类型
	type testGroup struct {
		str  string
		sep  string
		want []string
	}

	// 定义一个存储测试用例的map
	tests := map[string]testGroup{
		"sep_1": {str: "abc", sep: "b", want: []string{"a", "c"}},
		"sep_2": {str: "a:b:c:d:e", sep: ":", want: []string{"a", "b", "c", "d", "e"}},
		"sep_3": {str: "abcdef", sep: "cd", want: []string{"ab", "ef"}},
		"sep_4": {str: "四川省成都市", sep: "省", want: []string{"四川", "成都市"}}, // 这里写一个错误的预期结果
	}

	// 遍历map，通过t.Run()逐一执行测试用例
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}


func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("四川省成都市", "省")
	}
}
