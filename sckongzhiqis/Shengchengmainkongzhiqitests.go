package sckongzhiqis

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"log"
	"os"
	"io/ioutil"
)
//
//func importsmainkongzhiqitest(mokuai string, buffer *bytes.Buffer) {
//	buffer.WriteString(zf.Zfs.Import(true))
//	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
//	//"github.com/astaxie/beego/context"
//	gitstr := zfzhi.Zhi.Syh() + zf.Zfs.Github(true) + zfzhi.Zhi.Dh() +
//		zf.Zfs.Com(true) + zfzhi.Zhi.Xx() + zf.Zfs.Astaxie(true) +
//		zfzhi.Zhi.Xx() + zf.Zfs.Beego(true) + zfzhi.Zhi.Xx() +
//		zf.Zfs.Context(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(gitstr)
//
//	//"xxx/kongzhiqis"
//	constr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
//		zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(constr)
//
//	//"gongju"
//	gjstr := zfzhi.Zhi.Syh() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(gjstr)
//
//	//"xxx/zf"
//	zfstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
//		zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(zfstr)
//
//	// "xxx/zfzhi"
//	zfzhistr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
//		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(zfzhistr)
//
//	//"log"
//	logstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(logstr)
//	//"strconv"
//	strcstr := zfzhi.Zhi.Syh() + zf.Zfs.Strconv(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(strcstr)
//	//testing
//	testingstr := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(testingstr)
//
//	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
//}
//func deletegetmainkongzhiqitest(fangfa string, buffer *bytes.Buffer) {
//	// func TestDeletemain(t * testing.T)
//	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
//		fangfa + zf.Zfs.Main(true) + zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
//		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
//	buffer.WriteString(funstr)
//	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
//
//	//paramid := strconv.Itoa(zfzhi.Shuzi1zhi())
//	parstr := zf.Zfs.Param(true) + zf.Zfs.Id(true) + zfzhi.Zhi.Mh() +
//		zfzhi.Zhi.Dyh() + zf.Zfs.Strconv(true) + zfzhi.Zhi.Dh() +
//		zf.Zfs.Itoa(false) + zfzhi.Zhi.Xkhz() +
//		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) +
//		zfzhi.Zhi.Dh() + zf.Zfs.Shuzi1(false) + zfzhi.Zhi.Xkhz() +
//		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(parstr)
//
//	//c := Xxxkongzhiqi()
//	cstr := zf.Zfs.C(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
//		zf.Zfs.Main(false) + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(cstr)
//
//	//c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false),paramid)
//	cinstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ctx(false) +
//		zfzhi.Zhi.Dh() + zf.Zfs.Input(false) + zfzhi.Zhi.Dh() + zf.Zfs.SetParam(false)
//	buffer.WriteString(cinstr)
//
//	//(zfzhi.Zhi.Mh() + zf.Zfs.Id(false),paramid)
//	cinparam := zfzhi.Zhi.Xkhz() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() +
//		zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Mh(false) +
//		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Jia() +
//		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) +
//		zfzhi.Zhi.Dh() + zf.Zfs.Id(false) + zfzhi.Zhi.Xkhz() +
//		zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() +
//		zf.Zfs.Param(true) + zf.Zfs.Id(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(cinparam)
//
//	// c.Get()
//	callstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + fangfa + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(callstr)
//
//	//log.Println(c.Data[zf.Zfs.Json(true)])
//	logstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
//		zfzhi.Zhi.Xkhz() + zf.Zfs.C(true) + zfzhi.Zhi.Dh() +
//		zf.Zfs.Data(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) +
//		zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(false) + zfzhi.Zhi.Xkhz() +
//		zf.Zfs.True(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() +
//		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
//	buffer.WriteString(logstr)
//
//	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
//
//}
//func getmainkongzhiqitest(buffer *bytes.Buffer) {
//	deletegetmainkongzhiqitest(zf.Zfs.Get(false), buffer)
//}
func Shengchengmainkongzhiqitest() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		buffer := bytes.Buffer{}
		buffer.WriteString(zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf())
		importskongzhiqitest(mkv, &buffer)
		bkkongzhiqitest(zf.Zfs.Main(false), &buffer)
		getkongzhiqitest(zf.Zfs.Main(false), &buffer)

		//hanfuxn/tesets
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
		//xxx/tests/Xxx_test.go
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Main(false) + zf.Zfs.Kongzhiqi(true) +
			zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		log.Print("buffer------------", buffer.String())
		log.Print("path------------", path)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
