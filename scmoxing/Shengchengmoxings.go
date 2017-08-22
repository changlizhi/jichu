package scmoxing

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"io/ioutil"
	"os"
	"gongju"
	"log"
)

func Shengchengmoxings() {
	_, biaos, _ := gongju.Biaolies()
	for bk, _ := range biaos {
		buffer := bytes.Buffer{}
		// model生成时大多数都是一样的，所以提供一个公用的package,import之类的
		// package xxx \n
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pac)

		// type Juese struct{\n
		typestr := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + bk + zfzhi.Zhi.Kgf() +
			zf.Zfs.Struct(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
		buffer.WriteString(typestr)

		for _, lk := range gongju.Biao(bk) {
			field := lk + zfzhi.Zhi.Kgf() + gongju.Lieleixing(lk) + zfzhi.Zhi.Hhf()
			buffer.WriteString(field)
		}
		//左大括号在头里有了
		buffer.WriteString(zfzhi.Zhi.Dkhy()) // }
		// hanfuxin/appmodels/Juese.go
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + gongju.Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi +
			zfzhi.Zhi.Xx() + zf.Zfs.Moxings(true)
		path := dir + zfzhi.Zhi.Xx() + bk + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		log.Println(buffer.String())
	}
}
