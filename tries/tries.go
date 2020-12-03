package tries

import (
	"grow/str"
)

// 初始化
// addKeyword
// 创建succeed表
// 创建fail表：最大后缀
// abc
// abcd
// cde

// 每次都重新生成和清空
type Tries struct {
	Word  string
	Child []*Tries
	// Fail存入一个hash表
	IsMatchWord bool
}

func New(word string) *Tries {
	n := make([]*Tries, 0)
	root := &Tries{
		Word:        word,
		Child:       n,
		IsMatchWord: false,
	}
	return root
}

func (root *Tries) AddWord(word string) {
	words := str.Suffix(word)
	root.addWord(words)

}

func (root *Tries) Build(m map[string]*Tries) {
	if root != nil {
		if root.Word != "" {
			m[root.Word] = root
		}

		for _, node := range root.Child {
			node.Build(m)
		}
	}
}

func (root *Tries) addWord(word []string) {
	if root != nil && len(word) >= 1 {

		w := word[0]
		isNewNode := true
		for _, node := range root.Child {
			if node.Word == w {
				isNewNode = false
				if len(word) == 1 {
					node.IsMatchWord = true
				}
				node.addWord(word[1:])
			}
		}
		if isNewNode {
			newNode := New(w)
			if len(word) == 1 {
				newNode.IsMatchWord = true
			}

			root.Child = append(root.Child, newNode)
			newNode.addWord(word[1:])

		}
	}
}
