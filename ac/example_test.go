package ac

import "fmt"

func ExampleAC_Extract() {

	grow := New()
	grow.AddWord("阿里巴巴")
	grow.AddWord("基金")
	grow.Build()
	words := grow.ExtractSet("阿里巴巴asdfasd阿里巴巴I love Big Apple and Bay Area.基金基金")
	fmt.Println(words)
	// Output:
	// [阿里巴巴 基金]
}
