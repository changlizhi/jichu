package scchushihuas

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"io/ioutil"
	"os"
)

func Shengchengshezhis() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Hhf()
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() +
			mkv + zfzhi.Zhi.Xx() + zf.Zfs.Chushihuas(true)
		for _, jl := range gongju.Jsonliesarr {
			jlko := gongju.Jsonlies[jl]
			jlk := jlko.Bianma
			buffer := &bytes.Buffer{}
			// package chushihuas
			buffer.WriteString(pac)
			// type Shezhi struct
			typestrlie := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + jlk +
				zfzhi.Zhi.Kgf() + zf.Zfs.Struct(true)
			buffer.WriteString(typestrlie)
			// { \n
			buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
			//从jsonlie包下读取里面的方法解析出所有的*Jl的方法
			lies := gongju.Yigejsonlies(jlk)
			for _, lie := range lies {
				// Xxx []Tongyong
				l := lie + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zf.Zfs.Tongyong(false) + zfzhi.Zhi.Hhf()
				buffer.WriteString(l)
			}
			//\n }
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy())
			pathlie := dir + zfzhi.Zhi.Xx() + jlk + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			ioutil.WriteFile(pathlie, buffer.Bytes(), os.ModePerm)
		}
	}
}
