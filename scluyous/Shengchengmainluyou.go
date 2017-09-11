package scluyous

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"changliang/zh"
	"os"
	"io/ioutil"
)

func Shengchengmainluyou() {
	mks := gongju.Mokuaimings
	for _, mkvo := range mks {
		mkv := mkvo.Zhi
		buffer := &bytes.Buffer{}
		//package luyous
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Luyous(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pac)
		buffer.WriteString(zf.Zfs.Import(true))
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
		//"xxx/zdkongzhiqis"
		kzqstr := zfzhi.Zhi.Syh() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(kzqstr)

		//"changliang/zfzhi"
		zhistr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(zhistr)
		//"changliang/zh"
		zhbao := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zh(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(zhbao)
		//"github.com/astaxie/beego"
		beestr := zh.Zhs.Beegobao() + zfzhi.Zhi.Hhf()
		buffer.WriteString(beestr)
		//"github.com/astaxie/beego/plugins/cors"
		beecorsstr := zh.Zhs.Beegocorsbao() + zfzhi.Zhi.Hhf()
		buffer.WriteString(beecorsstr)

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

		// func init()
		funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Init(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
		buffer.WriteString(funstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

		//beego.InsertFilter
		bistr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.InsertFilter(false)
		buffer.WriteString(bistr)
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
		//zfzhi.Zhi.Xh(),
		xhstr := zh.Zhs.Zhiszh(zf.Zfs.Xh(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		buffer.WriteString(xhstr)
		//beego.BeforeRouter,
		bbstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.BeforeRouter(false) + zfzhi.Zhi.Dou()
		buffer.WriteString(bbstr)
		//cors.Allow
		castr := zf.Zfs.Cors(true) + zfzhi.Zhi.Dh() + zf.Zfs.Allow(false)
		buffer.WriteString(castr)
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
		//&cors.Options
		costr := zfzhi.Zhi.Qh() + zf.Zfs.Cors(true) + zfzhi.Zhi.Dh() + zf.Zfs.Options(false)
		buffer.WriteString(costr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

		//AllowAllOrigins:  true,
		aastr := zf.Zfs.AllowAllOrigins(false) + zfzhi.Zhi.Mh() + zf.Zfs.True(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		buffer.WriteString(aastr)

		//	AllowMethods:     zh.Zhs.AllowMethods(),
		amstr := zf.Zfs.AllowMethods(false) + zfzhi.Zhi.Mh() + zh.Zhszh(zf.Zfs.AllowMethods(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		buffer.WriteString(amstr)

		//	AllowHeaders:     zh.Zhs.AllowHeaders(),
		ahstr := zf.Zfs.AllowHeaders(false) + zfzhi.Zhi.Mh() + zh.Zhszh(zf.Zfs.AllowHeaders(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		buffer.WriteString(ahstr)
		//	ExposeHeaders:    zh.Zhs.ExposeHeaders(),
		ehstr := zf.Zfs.ExposeHeaders(false) + zfzhi.Zhi.Mh() + zh.Zhszh(zf.Zfs.ExposeHeaders(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		buffer.WriteString(ehstr)
		//	AllowCredentials: true,
		acrstr := zf.Zfs.AllowCredentials(false) + zfzhi.Zhi.Mh() + zf.Zfs.True(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		buffer.WriteString(acrstr)

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

		//beego.Router
		brstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Router(false)
		buffer.WriteString(brstr)
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
		//zfzhi.Zhi.Xx(),
		xxstr := zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		buffer.WriteString(xxstr)
		//	&zdkongzhiqis.Mainkongzhiqi{},
		zmstr := zfzhi.Zhi.Qh() + zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() + zf.Zfs.Main(false) + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		buffer.WriteString(zmstr)

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

		dir := gongju.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Luyous(true)
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Main(false) + zf.Zfs.Luyou(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}