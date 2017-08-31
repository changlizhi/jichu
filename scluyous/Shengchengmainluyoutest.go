package scluyous

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"changliang/zh"
	"gongju"
	"os"
	"io/ioutil"
)

func Shengchengmainluyoutest() {
	buffer := &bytes.Buffer{}
	//package tests
	pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf()
	buffer.WriteString(pac)
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//"github.com/astaxie/beego"
	buffer.WriteString(zh.Zhs.Beegobao() + zfzhi.Zhi.Hhf())
	//_"xxx/luyous"
	lstr := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Syh() + gongju.Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi + zfzhi.Zhi.Xx() + zf.Zfs.Luyous(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lstr)
	//"net/http"
	buffer.WriteString(zh.Zhs.Httpbao() + zfzhi.Zhi.Hhf())
	//. "github.com/smartystreets/goconvey/convey"
	buffer.WriteString(zh.Zhs.Conveybao() + zfzhi.Zhi.Hhf())
	//"testing"
	tstr := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tstr)
	//"gongju"
	gstr := zfzhi.Zhi.Syh() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(gstr)
	//"net/http/httptest"
	buffer.WriteString(zh.Zhs.Httptestbao() + zfzhi.Zhi.Hhf())
	//"changliang/zf"
	zfstr:=zfzhi.Zhi.Syh()+zf.Zfs.Changliang(true)+zfzhi.Zhi.Xx()+zf.Zfs.Zf(true)+zfzhi.Zhi.Syh()+zfzhi.Zhi.Hhf()
	buffer.WriteString(zfstr)
	//"changliang/zfzhi"
	zfzhistr:=zfzhi.Zhi.Syh()+zf.Zfs.Changliang(true)+zfzhi.Zhi.Xx()+zf.Zfs.Zfzhi(true)+zfzhi.Zhi.Syh()+zfzhi.Zhi.Hhf()
	buffer.WriteString(zfzhistr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	//func init()
	funini := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Init(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funini)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//beego.TestBeegoInit
	btbstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.TestBeegoInit(false)
	buffer.WriteString(btbstr)
	//(gongju.Getapppath())
	xkh := zfzhi.Zhi.Xkhz() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Dh() + zf.Zfs.Getapppath(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(xkh)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//func TestMainluyou()
	funtm := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.TestMainluyou(false)
	buffer.WriteString(funtm)
	//(t *testing.T)
	cst := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(cst)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//r, _ := http.NewRequest
	reqstr := zf.Zfs.R(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Xhx() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Http(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewRequest(false)
	buffer.WriteString(reqstr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.GET(false),
	gsstr := zh.Zhs.Zfszhfalse(zf.Zfs.GET(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(gsstr)
	//zfzhi.Zhi.Xx(),
	xxstr := zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(xxstr)
	//nil,
	nilstr := zf.Zfs.Nil(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(nilstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//w := httptest.NewRecorder()
	nrstr := zf.Zfs.W(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Httptest(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewRecorder(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(nrstr)
	//beego.BeeApp.Handlers.ServeHTTP(w, r)
	bhsstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.BeeApp(false) + zfzhi.Zhi.Dh() + zf.Zfs.Handlers(false) + zfzhi.Zhi.Dh() + zf.Zfs.ServeHTTP(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.W(true) + zfzhi.Zhi.Dou() + zf.Zfs.R(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bhsstr)
	//beego.Trace
	btstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.Trace(false)
	buffer.WriteString(btstr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.Testing(true) + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Bfh() + zf.Zfs.D(true) + zfzhi.Zhi.Zkhy(),
	ts1 := zh.Zhs.Zfszhtrue(zf.Zfs.Testing(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Zkhz(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Bfh(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhtrue(zf.Zfs.D(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Zkhy(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ts1)

	//zf.Zfs.TestMainluyou(false) + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Bfh() + zf.Zfs.D(true) + zfzhi.Zhi.Zkhy(),
	ts2 := zh.Zhs.Zfszhfalse(zf.Zfs.TestMainluyou(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Zkhz(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Bfh(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhtrue(zf.Zfs.D(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Zkhy(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ts2)
	//zf.Zfs.Code(false) + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Bfh() + zf.Zfs.D(true) + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Fxx() + zfzhi.Zhi.Bfh() + zf.Zfs.S(true),
	ts3 := zh.Zhs.Zfszhfalse(zf.Zfs.Code(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Zkhz(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Bfh(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhtrue(zf.Zfs.D(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Zkhy(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Fxx(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Bfh(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhtrue(zf.Zfs.S(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ts3)
	//w.Code,
	wcstr := zf.Zfs.W(true) + zfzhi.Zhi.Dh() + zf.Zfs.Code(false) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(wcstr)
	//	w.Code,
	buffer.WriteString(wcstr)
	//	w.Code,
	buffer.WriteString(wcstr)
	//w.Body.String(),
	wbsstr := zf.Zfs.W(true) + zfzhi.Zhi.Dh() + zf.Zfs.Body(false)+zfzhi.Zhi.Dh() + zf.Zfs.String(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(wbsstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	buffer.WriteString(zf.Zfs.Convey(false))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.Code(false) + zfzhi.Zhi.Kgf() + zf.Zfs.And(false) + zf.Zfs.Bodylen(false),
	con1 := zh.Zhs.Zfszhfalse(zf.Zfs.Code(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Kgf(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhfalse(zf.Zfs.And(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhfalse(zf.Zfs.Bodylen(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(con1)
	//t,
	buffer.WriteString(zf.Zfs.T(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())
	//func()
	buffer.WriteString(zf.Zfs.Func(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy())
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	buffer.WriteString(zf.Zfs.Convey(false))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.Code(false) + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Shuzi2w() + zfzhi.Zhi.Shuzi0w() + zfzhi.Zhi.Shuzi0w(),
	con21 := zh.Zhs.Zfszhfalse(zf.Zfs.Code(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Dyh(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Dyh(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi2w(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi0w(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi0w(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(con21)
	//func()
	buffer.WriteString(zf.Zfs.Func(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy())
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//So
	buffer.WriteString(zf.Zfs.So(false))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//w.Code,
	buffer.WriteString(wcstr)
	//ShouldEqual,
	sestr := zf.Zfs.ShouldEqual(false) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(sestr)
	//zfzhi.Zhi.Shuzi2() * zfzhi.Zhi.Shuzi10() * zfzhi.Zhi.Shuzi10(),
	con22 := zh.Zhs.Zhiszh(zf.Zfs.Shuzi2(false)) + zfzhi.Zhi.Xh() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi10(false)) + zfzhi.Zhi.Xh() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi10(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(con22)
	//)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//},
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())
	//)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	//Convey
	buffer.WriteString(zf.Zfs.Convey(false))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.Bodylen(false) + zfzhi.Zhi.Dy() + zfzhi.Zhi.Shuzi0w(),
	con31 := zh.Zhs.Zfszhfalse(zf.Zfs.Bodylen(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Dy(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi0w(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(con31)
	//func()
	buffer.WriteString(zf.Zfs.Func(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy())
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//So
	buffer.WriteString(zf.Zfs.So(false))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//w.Body.Len(),
	wbl := zf.Zfs.W(true) + zfzhi.Zhi.Dh() + zf.Zfs.Body(false) + zfzhi.Zhi.Dh() + zf.Zfs.Len(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(wbl)
	//ShouldBeGreaterThan,
	buffer.WriteString(zf.Zfs.ShouldBeGreaterThan(false) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())
	//zfzhi.Zhi.Shuzi0(),
	con33 := zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(con33)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//},
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())
	//)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//},
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())
	//)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//}
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	dir := gongju.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
	path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Mainluyou(false) + zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
	os.MkdirAll(dir, os.ModePerm)
	ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	log.Println(buffer.String())
}