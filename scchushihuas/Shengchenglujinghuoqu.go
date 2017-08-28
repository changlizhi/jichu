package scchushihuas

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
	"gongju"
	"io/ioutil"
	"os"
)

func Shengchenglujinghuoqu() {
	buffer := &bytes.Buffer{}
	//package chushihuas
	pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Hhf()
	buffer.WriteString(pac)
	//import
	buffer.WriteString(zf.Zfs.Import(true))
	//(\n
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//"gongju"
	gj := zfzhi.Zhi.Syh() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(gj)
	//"changliang/zf"
	cl := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh()
	buffer.WriteString(cl)
	//\n)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	// func Getapppath() string
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Getapppath(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	// return gongju.Getpath(zfzhi.Zhi.Kzf())
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Gongju(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Getpath(false) + zfzhi.Zhi.Xkhz() +
		zh.Zhs.Zhiszh(zf.Zfs.Kzf(false)) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(retstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy())
	dir := gongju.Getgopath() + zfzhi.Zhi.Xx() +
		gongju.Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi +
		zfzhi.Zhi.Xx() + zf.Zfs.Chushihuas(true)
	path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Lujinghuoqu(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
	os.MkdirAll(dir, os.ModePerm)
	ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
}
