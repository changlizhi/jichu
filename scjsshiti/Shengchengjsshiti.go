package scjsshiti

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"io/ioutil"
	"os"
)

func Shengchengjsshiti() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		mkk := mks[mkvo].Bianma
		bjg := gongju.Suoyoubiaojiegou(mkv)
		zwmkk := gongju.Liezhongwen(mkk)

		buffer := &bytes.Buffer{}
		//export function lieming (bianma) {
		funstr := zf.Zfs.Export(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Function(true) +
			zfzhi.Zhi.Kgf() + zf.Zfs.Lieming(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() +
			zf.Zfs.Bianma(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf()
		buffer.WriteString(funstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//var shiti =
		vstr := zf.Zfs.Var(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Shiti(true) +
			zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Kgf()
		buffer.WriteString(vstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//A: 'A'为了好生成
		astr := zf.Zfs.A(false) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Kgf() +
			zfzhi.Zhi.Dyhe() + zf.Zfs.A(false) + zfzhi.Zhi.Dyhe()
		buffer.WriteString(astr)

		//,\nzwmkk:'zwmkk'
		mkstr := zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf() + mkk + zfzhi.Zhi.Mh() + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dyhe() + zwmkk + zfzhi.Zhi.Dyhe()
		buffer.WriteString(mkstr)

		for b, ls := range bjg {
			zwb := gongju.Liezhongwen(b)
			zwbstr := zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf() + b + zfzhi.Zhi.Mh() + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dyhe() + zwb + zfzhi.Zhi.Dyhe()
			buffer.WriteString(zwbstr)
			for _, l := range ls {
				zwl := gongju.Liezhongwen(l)
				zwlstr := zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf() + l + zfzhi.Zhi.Mh() + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dyhe() + zwl + zfzhi.Zhi.Dyhe()
				buffer.WriteString(zwlstr)
			}
		}
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
		//return shiti[bianma]
		restr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Shiti(true) +
			zfzhi.Zhi.Zkhz() + zf.Zfs.Bianma(true) + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(restr)
		buffer.WriteString(zfzhi.Zhi.Dkhy())

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Js(true)
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Lie(true) + zfzhi.Zhi.Dh() + zf.Zfs.Js(true)

		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}

}
