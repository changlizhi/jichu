package scpeizhi

import (
	"gongju"
	"changliang/zfzhi"
	"changliang/zf"
	"os"
	"io/ioutil"
	"bytes"
	"log"
)

func Shengchengxitongpeizhi() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		buffer := &bytes.Buffer{}
		mkv := mks[mkvo].Zhi
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Peizhi(true)
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Shezhi(false)+zf.Zfs.S(true) + zfzhi.Zhi.Dh() + zf.Zfs.Json(true)
		log.Println(buffer.String())
		log.Println("path----------", path)

		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
