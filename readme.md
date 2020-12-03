## Trie
wiki:https://zh.wikipedia.org/wiki/Trie
节点不直接保存值，而是由节点在树中的位置决定

## AC自动机算法
wiki: https://zh.wikipedia.org/wiki/AC%E8%87%AA%E5%8A%A8%E6%9C%BA%E7%AE%97%E6%B3%95
1. 字符串的每一个前缀都有一个节点来表示
   1. 字符串bca存在节点
      1. bca
      2. bc
      3. b
      4. 空
2. 如果节点在字典中，则为蓝色，否则灰色
3. 黑色有向边: 增加一个字符形成下一个节点字符串
   1. bc增加a变成bca节点
4. 蓝色有向边（后缀节点）: 指向本节点的最大后缀
5. 绿色有向边（字典后缀节点）: 代表终点是起点经过蓝色有向边到达的第一个蓝色节点
   


https://oi-wiki.org/string/ac-automaton/
fail指针指向当前状态的最长后缀状态

## 相似项目
https://github.com/signalsciences/ac/blob/master/ac.go