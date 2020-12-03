package ac

import (
	"testing"
)

func TestAC_build(t *testing.T) {
	ac := New()
	// ac.AddWord("我爱北京")
	// ac.AddWord("我爱北京")
	// ac.AddWord("我爱河北")
	// ac.AddWord("我爱河南")
	// ac.AddWord("我爱河南北")

	ac.AddWord("hers")
	ac.AddWord("his")
	ac.AddWord("she")

	tests := []struct {
		name string
		ac   *AC
	}{
		// TODO: Add test cases.
		{
			ac: ac,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ac.Build()
			for key, val := range tt.ac.FailPoint {
				t.Logf("node: %s, failNode: %s", key.Word, val.Word)
			}

		})
	}
}

func compare(x, y []string) bool {

	if len(x) != len(y) {
		return false
	}
	xMap := make(map[string]int)
	yMap := make(map[string]int)

	for _, xElem := range x {
		xMap[xElem]++
	}
	for _, yElem := range y {
		yMap[yElem]++
	}

	for xMapKey, xMapVal := range xMap {
		if yMap[xMapKey] != xMapVal {
			return false
		}
	}
	return true
}

func TestAC_Extract(t *testing.T) {
	type args struct {
		text string
	}
	ac := New()
	// ac.AddWord("我爱北京")
	// ac.AddWord("我爱北京")
	// ac.AddWord("我爱河北")
	// ac.AddWord("我爱河南")
	// ac.AddWord("我爱河南北")

	ac.AddWord("hers")
	ac.AddWord("his")
	ac.AddWord("she")
	ac.AddWord("he")
	ac.Build()
	tests := []struct {
		name string
		ac   *AC
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			ac: ac,
			args: args{
				text: "ushers",
			},
			want: []string{"he", "she", "hers"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ac.Extract(tt.args.text); !compare(got, tt.want) {
				t.Errorf("AC.Extract() = %v, want %v", got, tt.want)
			}

		})
	}
}
