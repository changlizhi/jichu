package scconf

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"os"
	"io/ioutil"
)

func Shengchengconf() {
	mks := gongju.Mokuaimings
	for _, mkvo := range mks {
		mkv := mkvo.Zhi
		buffer := &bytes.Buffer{}
		//appname = hfxyonghu
		astr := zf.Zfs.Appname(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dyh() +
			zfzhi.Zhi.Kgf() + mkv + zfzhi.Zhi.Hhf()
		buffer.WriteString(astr)

		//httpport = 8080
		hstr := zf.Zfs.Httpport(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Kgf() +
			zfzhi.Zhi.Shuzi8w() + zfzhi.Zhi.Shuzi0w() +
			zfzhi.Zhi.Shuzi8w() + zfzhi.Zhi.Shuzi0w() + zfzhi.Zhi.Hhf()
		buffer.WriteString(hstr)

		//runmode = dev
		rstr := zf.Zfs.Runmode(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Kgf() + zf.Zfs.Dev(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(rstr)

		//copyrequestbody = true
		crbtr := zf.Zfs.Copyrequestbody(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Kgf() + zf.Zfs.True(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(crbtr)

		dir := gongju.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Conf(true)
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.App(true) + zfzhi.Zhi.Dh() + zf.Zfs.Conf(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}