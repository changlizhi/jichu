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

func bufferwriteinit(jsonlies map[string]gongju.Tongyong, buffer *bytes.Buffer) {
	for jlk, _ := range jsonlies {
		// Shezhilie()\n
		// Guojihualie()
		dystr := jlk + zf.Zfs.Lie(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(dystr)
	}
}

func bufferwritelie(jsonlies map[string]gongju.Tongyong, bufferfun *bytes.Buffer, buffer *bytes.Buffer) {

}

func Shengchengchushihuajsonlie() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		buffer := &bytes.Buffer{}
		//package chushihuas
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pac)

		// func init()
		inistr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Init(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
		buffer.WriteString(inistr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		for _, jl := range gongju.Jsonliesarr {
			jlko := gongju.Jsonlies[jl]
			jlk := jlko.Bianma
			// Shezhilie()\n
			// Guojihualie()
			dystr := jlk + zf.Zfs.Lie(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(dystr)
		}
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

		bufferfun := &bytes.Buffer{}

		for _, jl := range gongju.Jsonliesarr {
			jlko := gongju.Jsonlies[jl]
			jlk := jlko.Bianma
			lies := gongju.Yigejsonlies(jlk)
			funone := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + jlk + zf.Zfs.Lie(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
			bufferfun.WriteString(funone)
			bufferfun.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
			//jlk := Jlkjson()
			jlkstr := strings.ToLower(jlk) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
				jlk + zf.Zfs.Json(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			bufferfun.WriteString(jlkstr)
			for _, lie := range lies {
				//var Xxxs = make(map[string]Tongyong)
				varstr := zf.Zfs.Var(true) + zfzhi.Zhi.Kgf() + lie + zf.Zfs.S(true) + zfzhi.Zhi.Dyh() +
					zf.Zfs.Make(true) + zfzhi.Zhi.Xkhz() + zf.Zfs.Map(true) +
					zfzhi.Zhi.Zkhz() + zf.Zfs.String(true) + zfzhi.Zhi.Zkhy() +
					zf.Zfs.Tongyong(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
				buffer.WriteString(varstr)
				//for _,tolower(lie):=
				forpre := zf.Zfs.For(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() +
					strings.ToLower(lie) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh()
				bufferfun.WriteString(forpre)
				//range jlk.Lie
				foraft := zf.Zfs.Range(true) + zfzhi.Zhi.Kgf() + strings.ToLower(jlk) + zfzhi.Zhi.Dh() + lie
				bufferfun.WriteString(foraft)
				bufferfun.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

				liezhi := lie + zf.Zfs.S(true) + zfzhi.Zhi.Zkhz() +
					strings.ToLower(lie) + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) +
					zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dyh() + strings.ToLower(lie)

				bufferfun.WriteString(liezhi)
				bufferfun.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

			}
			bufferfun.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
		}

		buffer.Write(bufferfun.Bytes())
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() +
			mkv + zfzhi.Zhi.Xx() + zf.Zfs.Chushihuas(true)
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Chushihua(false) + zf.Zfs.Jsonlie(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
