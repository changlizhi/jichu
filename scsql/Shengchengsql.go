package scsql

import (
	"gongju"
	"log"
	"strconv"
)

func Shengchengsql() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, lies := gongju.Biaolies(mkv)
		for b, _ := range biaos {
			for l, _ := range lies {
				cd := gongju.Liechangdu(l)
				lx := gongju.Lieleixing(l)
				log.Println(b + "------------" + l + "----------" + strconv.Itoa(cd) + "---------" + lx)
			}
		}
	}
}
