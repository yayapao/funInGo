package main

import "fmt"

type People struct {
	Age, Height int
}

// 声明一个空 map（指定类型）
var mappings map[string]People

// 声明一个指定类型的 map，并为其赋值
var otherMappings = map[string]People{
	"a": {1, 2},
	"b": {2, 3},
}

// 为 map 赋值，注意需要通过 `make` 方法来初始化 map，如果直接给 mappings 赋值会报错
func setMapping() {
	if mappings == nil {
		fmt.Println("m is nil!")
	}
	mappings = make(map[string]People)
	mappings["xiaoA"] = People{12, 30}
	fmt.Println(mappings)
}

// 通过第二个返回值为 True 或者 False，判断元素是否在 map 内
// 通过 `delete(v, k)` 删除 map 内指定元素
func muateMapping() {
	v, ok := otherMappings["b"]
	fmt.Println(v, ok)
	delete(otherMappings, "b")
	vv, okk := otherMappings["b"]
	fmt.Println(vv, okk)
}

func main() {
	//setMapping()
	muateMapping()
}
