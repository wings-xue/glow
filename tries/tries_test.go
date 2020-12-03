package tries

import (
	"fmt"
	"grow/str"
	"reflect"
	"testing"
)

// 我爱北京
// 我爱河北
// 我爱河南
var root = New("")
var word1 = []string{"我", "我爱", "我爱北", "我爱北京"}
var word2 = []string{"我", "我爱", "我爱河", "我爱河北"}
var word3 = []string{"我", "我爱", "我爱河", "我爱河南"}
var word4 = []string{"我", "我爱", "我爱河", "我爱河南", "我爱河南北"}

func TestStringSuffix(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			args: args{
				s: "我爱中国",
			},
			want: []string{
				"我", "我爱", "我爱中", "我爱中国",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := str.Suffix(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTries_AddWord(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		root *Tries
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.root.AddWord(tt.args.word)
		})
	}
}

func TestTries_addWord(t *testing.T) {
	type args struct {
		word []string
	}

	tests := []struct {
		name string
		root *Tries
		args args
	}{
		// TODO: Add test cases.
		{
			root: root,
			args: args{
				word: word1,
			},
		},
		{
			root: root,
			args: args{
				word: word2,
			},
		},
		{
			root: root,
			args: args{
				word: word3,
			},
		},
		{
			root: root,
			args: args{
				word: word4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.root.addWord(tt.args.word)
			fmt.Println(tt.root)
		})
	}
	fmt.Println("")
}
