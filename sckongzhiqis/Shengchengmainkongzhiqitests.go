package sckongzhiqis

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"os"
	"io/ioutil"
)

func Shengchengmainkongzhiqitest() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		buffer := bytes.Buffer{}
		buffer.WriteString(zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf())
		importskongzhiqitest(mkv, &buffer)
		getkongzhiqitest(zf.Zfs.Main(false), &buffer)

		//hanfuxn/tesets
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
		//xxx/tests/Xxx_test.go
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Main(false) + zf.Zfs.Kongzhiqi(true) +
			zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
