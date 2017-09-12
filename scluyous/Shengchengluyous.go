package scluyous

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
	"gongju"
	"io/ioutil"
	"os"
)

func routersimports(mokuai string, buffer *bytes.Buffer) {
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	//"github.com/astaxie/beego"
	buffer.WriteString(zh.Zhs.Beegobao())

	//"changliang/zf"
	zfstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfstr)
	//"changliang/zfzhi"
	zfzstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfzstr)

	//"xxx/zdkongzhiqis"
	constr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(constr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

}
func routersinit(bianma string, buffer *bytes.Buffer) {
	//func init()
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Init(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//beego.Router(zfzhi.Zhi.Xx() + zf.Zfs.Xxx(true),&kongzhiqis.Xxxkongzhiqi{})
	rstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Router(false) +
		zfzhi.Zhi.Xkhz() + zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) +
		zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhz() +
		zf.Zfs.True(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() +
		zfzhi.Zhi.Qh() + zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() +
		bianma + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Dkhz() +
		zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(rstr)

	//beego.Router("/Xxx/:Id",&kongzhiqis.Xxxkongzhiqi{})
	ridstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Router(false) +
		zfzhi.Zhi.Xkhz() + zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) +
		zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhz() +
		zf.Zfs.True(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() + zh.Zhs.Zhiszh(zf.Zfs.Mh(false)) + zfzhi.Zhi.Jia() + zh.Zhs.Zfszhfalse(zf.Zfs.Id(false)) +
		zfzhi.Zhi.Dou() + zfzhi.Zhi.Qh() + zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() +
		bianma + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Dkhz() +
		zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ridstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func Shengchengluyous() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Biaolies(mkv)
		for bk, _ := range biaos {
			buffer := bytes.Buffer{}
			pacstr := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Luyous(true) + zfzhi.Zhi.Hhf()
			buffer.WriteString(pacstr)

			routersimports(mkv, &buffer)
			routersinit(bk, &buffer)
			dir := gongju.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Luyous(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Luyou(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}
	}
}
