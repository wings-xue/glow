## 部署
go get https://github.com/wings-xue/glow


## 快速开始
参考example_test.go
```
grow := ac.New()
grow.AddWord("阿里巴巴")
grow.AddWord("基金")
grow.Build()
words := grow.ExtractSet("阿里巴巴阿里巴巴基金基金")
fmt.Println(words)

// output:
// [阿里巴巴，基金]
```

