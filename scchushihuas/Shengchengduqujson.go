package scchushihuas

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
	"gongju"
	"io/ioutil"
	"os"
	"strings"
)

func Shengchengduqujson() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		mkk := mks[mkvo].Bianma
		buffer := &bytes.Buffer{}
		//package chushihuas
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pac)

		//import
		buffer.WriteString(zf.Zfs.Import(true))
		//(\n
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
		//"changliang/zf"
		clstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
			zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(clstr)
		//"gongju"
		gjstr := zfzhi.Zhi.Syh() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Syh()
		buffer.WriteString(gjstr)

		//\n)\n
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

		for _, jl := range gongju.Jsonliesarr {
			jlko := gongju.Jsonlies[jl]
			jlk := jlko.Bianma
			//func Shezhipath()string
			funpath := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + jlk +
				zf.Zfs.Path(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
			buffer.WriteString(funpath)
			buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

			//return gongju.Getwenjianmulu
			retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Dh() + zf.Zfs.Getwenjianmulu(false)
			buffer.WriteString(retstr)
			buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

			//zf.Zfs.Mokuai(true),
			mkstr := zh.Zhs.Zfszhtrue(mkk) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(mkstr)
			//zf.Zfs.Peizhi(true),
			pzstr := zh.Zhs.Zfszhtrue(zf.Zfs.Peizhi(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(pzstr)

			if gongju.Jsonliezhiding(jlk) == zf.Zfs.Xitong(false) {
				// zf.Zfs.Shezhi(false),
				szstr := zh.Zhs.Zfszhfalse(jlk) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
				buffer.WriteString(szstr)
			}
			if gongju.Jsonliezhiding(jlk) == zf.Zfs.Yuyan(false) {
				//Chushihuas[zf.Zfs.Yuyan(false)].Zhi,
				cshstr := zf.Zfs.Chushihuas(false) + zfzhi.Zhi.Zkhz() + zh.Zhs.Zfszhfalse(gongju.Jsonliezhiding(jlk)) +
					zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
				buffer.WriteString(cshstr)
			}
			//zf.Zfs.Json(true),
			jsstr := zh.Zhs.Zfszhtrue(zf.Zfs.Json(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(jsstr)
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

			//func Xxxjson()*Xxx
			funjsonstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + jlk + zf.Zfs.Json(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xh() + jlk
			buffer.WriteString(funjsonstr)
			buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

			jlx := strings.ToLower(jlk)
			//xxx := Xxx{}
			jlo := jlx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + jlk + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(jlo)
			//ret := gongju.Jiexi
			ostr := zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
				zf.Zfs.Gongju(true) + zfzhi.Zhi.Dh() + zf.Zfs.Jiexi(false)
			buffer.WriteString(ostr)

			buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
			// Shezhipath(),
			pstr := jlk + zf.Zfs.Path(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(pstr)
			// &shezhi,
			jstr := zfzhi.Zhi.Qh() + jlx + zfzhi.Zhi.Dou()
			buffer.WriteString(jstr)
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
			//return ret.(*Xxx)
			restr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ret(true) + zfzhi.Zhi.Dh() + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xh() + jlk + zfzhi.Zhi.Xkhy()
			buffer.WriteString(restr)
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
		}

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() +
			mkv + zfzhi.Zhi.Xx() + zf.Zfs.Chushihuas(true)
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Duqujson(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
