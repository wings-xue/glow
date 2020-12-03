package ac

import (
	"grow/str"
	"grow/tries"
	"unicode/utf8"
)

// 构建fail表

type AC struct {
	Dictionary map[string]string
	// 失败跳转你的树
	FailPoint map[*tries.Tries]*tries.Tries
	// word和树的hash表，一个词一定只有一个节点，这里的词指的是词的前缀
	WordTries map[string]*tries.Tries
	Tries     *tries.Tries

	Collect []string
}

func New() *AC {
	Dictionary := make(map[string]string)
	FailPoint := make(map[*tries.Tries]*tries.Tries)
	WordTries := make(map[string]*tries.Tries)
	Collect := make([]string, 0)
	return &AC{
		Dictionary: Dictionary,
		FailPoint:  FailPoint,
		WordTries:  WordTries,
		Tries:      tries.New(""),
		Collect:    Collect,
	}
}

func (ac *AC) AddWord(word string) {
	// 构建tres树
	// 如果word之前没有添加过才执行
	if _, ok := ac.Dictionary[word]; !ok {
		ac.Tries.AddWord(word)
		ac.Dictionary[word] = word
	}
}

func (ac *AC) Build() {
	// tries构建成功时才能做以下事
	// 1. 构建WordTries
	// 2. 构建FailPoint
	ac.Tries.Build(ac.WordTries)
	// 循环wordTries, 基于后配匹配节点
	for key, node := range ac.WordTries {
		isFindFail := false
		if utf8.RuneCountInString(key) == 1 {
			ac.FailPoint[node] = ac.Tries
		} else {
			for _, p := range str.Prefix(key) {
				if failNode, ok := ac.WordTries[p]; ok {
					ac.FailPoint[node] = failNode
					isFindFail = true
					break
				}
			}
			if !isFindFail {
				ac.FailPoint[node] = ac.Tries
			}
		}
	}
}

// 状态转移，输入一个节点，输出一个节点
func (ac *AC) Transfer(c rune, t *tries.Tries) *tries.Tries {
	if node, ok := Match(c, t); ok {
		if node.IsMatchWord {
			ac.Collect = append(ac.Collect, node.Word)
		}
		return node
	} else {
		for node, ok := ac.FailPoint[t]; ok; t = node {
			// 如果跳转到头结点
			if node.Word == "" {
				return node
			}
			// 不是头结点，再一次进行计算
			if node.IsMatchWord {
				ac.Collect = append(ac.Collect, node.Word)
			}
			return ac.Transfer(c, node)
		}
	}
	return t

}

// 匹配是否c在node的子节点
func Match(c rune, t *tries.Tries) (*tries.Tries, bool) {
	for _, node := range t.Child {
		word := []rune(node.Word)
		if word[len(word)-1] == c {
			return node, true
		}
	}
	return nil, false
}

func (ac *AC) Extract(text string) []string {
	// 遍历文本
	// 如果char匹配到了子节点中的一个，当前tries变成子节点，否则还是根节点
	// 如果下次匹配失败，跳转失败节点
	ac.Collect = ac.Collect[:0]
	t := []rune(text)

	current := ac.Tries
	for i := 0; i < len(t); i++ {
		current = ac.Transfer(t[i], current)
	}
	return ac.Collect
}

func removeDup(words []string) []string {
	out := make([]string, 0)
	temp := make(map[string]string, 0)
	for _, item := range words {
		if _, ok := temp[item]; !ok {
			temp[item] = item
			out = append(out, item)
		}
	}
	return out
}

func (ac *AC) ExtractSet(text string) []string {
	// 遍历文本
	// 如果char匹配到了子节点中的一个，当前tries变成子节点，否则还是根节点
	// 如果下次匹配失败，跳转失败节点
	ac.Collect = ac.Collect[:0]
	t := []rune(text)

	current := ac.Tries
	for i := 0; i < len(t); i++ {
		current = ac.Transfer(t[i], current)
	}

	return removeDup(ac.Collect)
}
