package scjsshiti

import (
	"gongju"
	"log"
)

func Shengchengjsshiti() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		bjg := gongju.Suoyoubiaojiegou(mkv)
		for b, l := range bjg {
			log.Println(b, l)
		}

	}

}
