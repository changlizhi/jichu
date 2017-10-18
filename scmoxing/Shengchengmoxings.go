package scmoxing

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"io/ioutil"
	"os"
)

func Shengchengmoxings() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Fanshebiaolies(mkv)
		tmf := false
		for bk, _ := range biaos {
			for _, lk := range gongju.Fanshebiao(mkv, bk) {
				if gongju.Lieleixing(lk) == zf.Zfs.Time(true) {
					tmf = true
					break
				}
			}
		}
		for bk, _ := range biaos {
			buffer := bytes.Buffer{}
			// model生成时大多数都是一样的，所以提供一个公用的package,import之类的
			// package xxx \n
			pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Hhf()
			buffer.WriteString(pac)
			if tmf {
				// import
				buffer.WriteString(zf.Zfs.Import(true))
				buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
				// "time"
				ts := zfzhi.Zhi.Syh() + zf.Zfs.Time(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
				buffer.WriteString(ts)
				buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
			}

			// type Juese struct{\n
			typestr := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + bk + zfzhi.Zhi.Kgf() +
				zf.Zfs.Struct(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
			buffer.WriteString(typestr)

			for _, lk := range gongju.Fanshebiao(mkv, bk) {
				leixing := gongju.Lieleixing(lk)
				field := lk + zfzhi.Zhi.Kgf() + gongju.Lieleixing(lk) + zfzhi.Zhi.Hhf()
				if leixing == zf.Zfs.Time(true) {
					//time.Time
					field = lk + zfzhi.Zhi.Kgf() + zf.Zfs.Time(true) + zfzhi.Zhi.Dh() + zf.Zfs.Time(false) + zfzhi.Zhi.Hhf()
				}
				buffer.WriteString(field)
			}
			//左大括号在头里有了
			buffer.WriteString(zfzhi.Zhi.Dkhy()) // }
			// xxx/moxings/Juese.go
			dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
				zfzhi.Zhi.Xx() + zf.Zfs.Moxings(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}

	}
}
