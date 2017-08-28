package scluyous

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

func routersimports(bianma string, buffer *bytes.Buffer) {
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	//"github.com/astaxie/beego"
	buffer.WriteString(zh.Zhs.Beegobao())

	//"xxx/zdkongzhiqis"
	constr := zfzhi.Zhi.Syh() + gongju.Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi + zfzhi.Zhi.Xx() +
		zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(constr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

}
func routersinit(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	//func init()
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Init(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//beego.Router("/zfzhi.Zhi.Xx()x",&kongzhiqis.Xxxkongzhiqi{})
	rstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Router(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Syh() + zfzhi.Zhi.Xx() +
		bmx + zfzhi.Zhi.Syh() + zfzhi.Zhi.Dou() +
		zfzhi.Zhi.Qh() + zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() +
		bianma + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Dkhz() +
		zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(rstr)

	//beego.Router("/zfzhi.Zhi.Xx()x/:Id",&kongzhiqis.Xxxkongzhiqi{})
	ridstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Router(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Syh() + zfzhi.Zhi.Xx() +
		bmx + zfzhi.Zhi.Xx() + zfzhi.Zhi.Mh() +
		zf.Zfs.Id(false) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Dou() +
		zfzhi.Zhi.Qh() + zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() +
		bianma + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Dkhz() +
		zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ridstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func Shengchengluyous() {
	_, biaos, _ := gongju.Biaolies()
	for bk, _ := range biaos {
		buffer := bytes.Buffer{}
		pacstr := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Luyous(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pacstr)

		routersimports(bk, &buffer)
		routersinit(bk, &buffer)
		dir := gongju.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Luyous(true)
		path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Luyou(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
