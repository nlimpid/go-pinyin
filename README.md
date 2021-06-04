# go-pinyin

[![Build Status](https://travis-ci.org/longbridgeapp/go-pinyin.svg?branch=master)](https://travis-ci.org/mozillazg/go-pinyin)
[![Coverage Status](https://coveralls.io/repos/mozillazg/go-pinyin/badge.svg?branch=master)](https://coveralls.io/r/mozillazg/go-pinyin?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/longbridgeapp/go-pinyin)](https://goreportcard.com/report/github.com/longbridgeapp/go-pinyin)
[![GoDoc](https://godoc.org/github.com/longbridgeapp/go-pinyin?status.svg)](https://pkg.go.dev/github.com/longbridgeapp/go-pinyin)

汉语拼音转换工具 Go 版。

这是一个 [go-pinyin](https://github.com/mozillazg/go-pinyin) fork 版本，用来支持段落、短语、多一字准确转换的支持。

## Installation

```
go get -u github.com/longbridgeapp/go-pinyin
```

install CLI tool:

```
go get -u github.com/longbridgeapp/go-pinyin/cmd/pinyin
$ pinyin 中国人
zhōng guó rén
```

## Documentation

API documentation can be found here:
https://pkg.go.dev/github.com/longbridgeapp/go-pinyin

## Usage

```go
package main

import (
	"fmt"
	"github.com/longbridgeapp/go-pinyin"
)

func main() {
	hans := "中国人"

	// 默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong] [guo] [ren]]

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhōng] [guó] [rén]]

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng] [guo2] [re2n]]

	// 开启多音字模式
	a = pinyin.NewArgs()
	a.Heteronym = true
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong zhong] [guo] [ren]]
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng zho4ng] [guo2] [re2n]]

	fmt.Println(pinyin.LazyPinyin(hans, pinyin.NewArgs()))
	// [zhong guo ren]

	fmt.Println(pinyin.Convert(hans, nil))
	// [[zhong] [guo] [ren]]

	fmt.Println(pinyin.LazyConvert(hans, nil))
	// [zhong guo ren]

	// 段落转换，支持完整支持多音字，保留符号
	fmt.Println(pinyin.Paragraph("交给团长，告诉他我们给予期望。前线的供给一定要能自给自足！"))
	// jiao gei tuan zhang, gao su ta wo men ji yu qi wang. qian xian de gong ji yi ding yao neng zi ji zi zu!
}
```

## Related Projects

- [hotoo/pinyin](https://github.com/hotoo/pinyin): 汉语拼音转换工具 Node.js/JavaScript 版。
- [mozillazg/python-pinyin](https://github.com/mozillazg/python-pinyin): 汉语拼音转换工具 Python 版。
- [mozillazg/rust-pinyin](https://github.com/mozillazg/rust-pinyin): 汉语拼音转换工具 Rust 版。

## pinyin data

- 使用 [pinyin-data](https://github.com/mozillazg/pinyin-data) 的拼音数据

## License

Under the MIT License.
