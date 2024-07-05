# hashplate-go

[English](README.md) | 简体中文

一个小而快的库，用于生成类似车牌风格的可读哈希字符串。

## 使用方法

```go
result := Hashplate("Hello World!", defaultOptions)
// returns "🍢 渝G·VGUA1 🪣"
```

## 格式

>(默认启用 CN 选项)

生成的哈希格式如下：
```
<emoji> <2个大写字母>-<3个数字>-<2个大写字母> <emoji>
```

如果在选项中启用了 CN，那么格式是：
```
<emoji> <1个汉字省份缩写1个大写字母>·<5个大写字母> <emoji>
```

可以通过设置Emoji参数为false来移除表情符号。
```go
Hashplate("Hello World!", Options{Emoji: false})
```

如果需要排除掉大写的O和大写的I，可以在选项中启用 ExcludeOI：
```go
Hashplate("Hello World!", Options{ExcludeOI: true})
```

## 为什么?

> Hugo ATTAL : 我需要一种生成**可读**且**易于识别**的哈希方法，以**匿名化**数据。这就是结果。

## 致谢

forked from Project [hashplate](https://github.com/hugoattal/hashplate)

forked from Project [hashplate-cn](https://github.com/cunzaizhuyi/hashplate-cn)
