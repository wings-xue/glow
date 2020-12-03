package ac

import "fmt"

func ExampleAC_Extract() {

	grow := New()
	grow.AddWord("阿里巴巴")
	grow.AddWord("基金")
	grow.Build()
	words := grow.ExtractSet("阿里巴巴阿里巴巴基金基金")
	fmt.Println(words)

	// Output:
	// [阿里巴巴 基金]
}

func ExampleAC_Extract2() {

	grow := New()
	grow.AddWord("中国")
	grow.AddWord("深圳")
	grow.AddWord("广州")
	grow.AddWord("厦门")
	grow.AddWord("珠海")
	grow.AddWord("江苏")
	grow.AddWord("盐城")
	grow.AddWord("杨超越")
	grow.AddWord("吴亦凡")
	grow.AddWord("杨幂")
	grow.AddWord("美国")
	grow.AddWord("baby")
	grow.Build()
	words := grow.ExtractSet("中国baby和杨超越在厦门看到了吴亦凡")
	fmt.Println(words)

	// Output:
	// [中国 baby 杨超越 厦门 吴亦凡]
}
