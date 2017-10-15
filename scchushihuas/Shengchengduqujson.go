package scchushihuas

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
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
		buffer := &bytes.Buffer{}
		//package chushihuas
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pac)

		//import
		buffer.WriteString(zf.Zfs.Import(true))
		//(\n
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
		//"gongju"
		gjstr := zfzhi.Zhi.Syh() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(gjstr)

		//"mhsydata/zdguojihua"
		mzgstr := zfzhi.Zhi.Syh() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) +
			zf.Zfs.Guojihua(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(mzgstr)
		//"mhsydata/zdshezhi"
		mzpstr := zfzhi.Zhi.Syh() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) +
			zf.Zfs.Shezhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(mzpstr)

		//\n)\n
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

		for _, jl := range gongju.Jsonliesarr {
			jlko := gongju.Jsonlies[jl]
			jlk := jlko.Bianma

			//func Xxxjson()*Xxx
			funjsonstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + jlk + zf.Zfs.Json(true) +
				zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xh() + jlk
			buffer.WriteString(funjsonstr)
			buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

			jlx := strings.ToLower(jlk)
			//xxx := Xxx{}
			jlo := jlx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + jlk +
				zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(jlo)
			//ret := gongju.Fanshejiexi
			ostr := zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
				zf.Zfs.Gongju(true) + zfzhi.Zhi.Dh() + zf.Zfs.Fanshejiexi(false)
			buffer.WriteString(ostr)

			buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
			// &jlx,
			pstr := zfzhi.Zhi.Qh() + jlx + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(pstr)
			// zdguojihua.Guojihua{},
			jstr := zf.Zfs.Zd(true) + jlx + zfzhi.Zhi.Dh() + jlk + zfzhi.Zhi.Dkhz() +
				zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(jstr)
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
			//return ret.(*Xxx)
			restr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ret(true) +
				zfzhi.Zhi.Dh() + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xh() + jlk + zfzhi.Zhi.Xkhy()
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
