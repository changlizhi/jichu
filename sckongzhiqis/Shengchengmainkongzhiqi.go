package sckongzhiqis

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
	"gongju"
	"io/ioutil"
	"os"
)

func Shengchengmainkongzhiqi() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		buffer := &bytes.Buffer{}
		//package zdkongzhiqis
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pac)
		buffer.WriteString(zf.Zfs.Import(true))
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

		//"changliang/zf"
		ztr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(ztr)

		//"github.com/astaxie/beego"
		buffer.WriteString(zh.Zhs.Beegobao() + zfzhi.Zhi.Hhf())

		//"gongju"
		gjstr := zfzhi.Zhi.Syh() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(gjstr)

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

		//type Mainkongzhiqi struct
		tstr := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Main(false) + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Struct(true)
		buffer.WriteString(tstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//beego.Controller
		bcstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Controller(false)
		buffer.WriteString(bcstr)
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

		//func (c *Mainkongzhiqi) Get()
		gfstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() + zf.Zfs.C(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Main(false) + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Xkhy() + zf.Zfs.Get(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
		buffer.WriteString(gfstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//c.Data[zf.Zfs.Json(true)] = gongju.Mokuaimings
		jstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Data(false) + zfzhi.Zhi.Zkhz() + zh.Zhs.Zfszhtrue(zf.Zfs.Json(false)) + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dyh() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Dh() + zf.Zfs.Mokuaimings(false) + zfzhi.Zhi.Hhf()
		buffer.WriteString(jstr)
		//c.ServeJSON()
		csjstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.ServeJSON(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(csjstr)

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true)
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Main(false) + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
