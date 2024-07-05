# hashplate-go

English | [简体中文](README-cn.md)

A tiny and fast lib to generate human-readable hash from a string in the style of license plates.

## Use

```go
result := Hashplate("Hello World!", defaultOptions)
// returns "🍢 渝G·VGUA1 🪣"
```

## Format

>(CN option enabled by default)

The hash is generated in the following format:
```
<emoji> <2 uppercase letters>-<3 digits>-<2 uppercase letters> <emoji>
```

If the CN option is enabled in the options, the format is:
```
<emoji> <emoji> <1 Chinese character provincial abbreviation, 1 uppercase letter>·<5 uppercase letters> <emoji> <emoji>
```

You can remove emoji by passing the second parameter.
```go
Hashplate("Hello World!", Options{Emoji: false})
```

To exclude uppercase 'O' and 'I', you can enable ExcludeOI in the options:
```go
Hashplate("Hello World!", Options{ExcludeOI: true})
```

## Why?

> Hugo ATTAL : I needed a way to generate a hash that was **readable** and **easy to recognize** to **anonymize** data. This is the result.

## Credit

forked from Project [hashplate](https://github.com/hugoattal/hashplate)

forked from Project [hashplate-cn](https://github.com/cunzaizhuyi/hashplate-cn)
