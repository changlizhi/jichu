package scmain

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
	"gongju"
	"io/ioutil"
	"os"
)

func Shengchengmain() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		buffer := &bytes.Buffer{}
		//package main
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Main(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pac)
		buffer.WriteString(zf.Zfs.Import(true))
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
		//_ "xxx/luyous"
		lystr := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Syh() +
			mkv +
			zfzhi.Zhi.Xx() +
			zf.Zfs.Luyous(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(lystr)
		//"github.com/astaxie/beego"
		bbao := zh.Zhs.Beegobao() + zfzhi.Zhi.Hhf()
		buffer.WriteString(bbao)
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

		//func main()
		funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Main(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
		buffer.WriteString(funstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//beego.Run()
		rstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Run(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(rstr)

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Main(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)

	}
}
