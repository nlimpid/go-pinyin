package pinyin

import (
	"strings"

	"github.com/yanyiwu/gojieba"
)

var (
	jieba = &gojieba.Jieba{}
)

func InitJieba(path ...string) {
	jieba = gojieba.NewJieba(path...)
}

func cutWords(s string) []string {
	return jieba.CutAll(s)
}

func pinyinPhrase(s string) string {
	words := cutWords(s)
	for _, word := range words {
		match := phraseDict[word]
		if match == "" {
			match = phraseDictAddition[word]
		}

		match = toFixed(match, paragraphOption)
		if match != "" {
			s = strings.Replace(s, word, " "+match+" ", 1)
		}
	}

	return s
}
