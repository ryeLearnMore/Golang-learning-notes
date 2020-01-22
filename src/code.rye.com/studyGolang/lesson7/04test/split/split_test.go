package split

import (
	"reflect"
	"testing"
	"os"
	"fmt"
)

func TestSplit(t *testing.T)  {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到：%v, 实际得到：%v\n", want, got)
	}
}

func TestNoneSplit(t *testing.T)  {
	got := Split("a:b:c", "*")
	want := []string{"a:b:c"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到：%v, 实际得到：%v\n", want, got)
	}
}

// 测试分隔多个字符的
func TestMultiSepSplit(t *testing.T)  {
	got := Split("abcfbcabcd", "bc")
	want := []string{"a", "f", "a", "d"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到：%v, 实际得到：%v\n", want, got)
	}
}

// 将多个测试用例放到一起组成：测试组

func TestMultiCaseSepSplit(t *testing.T)  {
	// 定义一个存放测试数据的结构体
	type test struct {
		str string
		sep string
		want []string
	}

	// 创建一个存放所有测试用例的map
	var tests = map[string]test{
		"normal": test{"a:b:c", ":", []string{"a", "b", "c"}},
		"none": test{"a:b:c", "*", []string{"a:b:c"}},
		"multi": test{"abcfbcabcd", "bc", []string{"a", "f", "a", "d"}},
	}

	for name, tc := range tests{
		// ret := Split(item.str, item.sep)
		// if !reflect.DeepEqual(ret, item.want) {
		// 	t.Errorf("测试用例：%v失败，期望得到：%v, 实际得到：%v\n", name, item.want, ret)
		// }

		t.Run(name, func (t *testing.T)  {
			ret := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(ret, tc.want) {
				t.Errorf("测试用例：%v失败，期望得到：%v, 实际得到：%v\n", name, tc.want, ret)
			}
		})
	}
}

// 基准测试
func BenchmarkSplit(b *testing.B)  {
	b.Log("这是一个基准测试")
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

// 整个测试之前做的事和之后做的事
func TestMain(m *testing.M)  {
	fmt.Println("write setup code here.") // 测试之前做一些设置
	// 如果TestMain使用了flags，这里应该加上flag.Parse()
	retCode := m.Run() // 执行测试
	fmt.Println("write teardown code here") // 测试之后做一些拆卸工作
	os.Exit(retCode)
}