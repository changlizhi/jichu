package scluyous

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"changliang/zh"
	"gongju"
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
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	//func init()
	funini:=zf.Zfs.Func(true )+zfzhi.Zhi.Kgf()+zf.Zfs.Init(true)+zfzhi.Zhi.Xkhz()+zfzhi.Zhi.Xkhy()
	buffer.WriteString(funini)
	buffer.WriteString(zfzhi.Zhi.Dkhz()+zfzhi.Zhi.Hhf())
	//beego.TestBeegoInit
	btbstr := zf.Zfs.Beego(true)+zfzhi.Zhi.Dh()+zf.Zfs.TestBeegoInit(false)
	buffer.WriteString(btbstr)
	//(gongju.Getapppath())
	xkh := zfzhi.Zhi.Xkhz()+zf.Zfs.Gongju(true)+zfzhi.Zhi.Dh()+zf.Zfs.Getapppath(false)+zfzhi.Zhi.Xkhz()+zfzhi.Zhi.Xkhy()+zfzhi.Zhi.Xkhy()
	buffer.WriteString(xkh)

	buffer.WriteString(zfzhi.Zhi.Hhf()+zfzhi.Zhi.Dkhy()+zfzhi.Zhi.Hhf())

	//func TestMainluyou()

	log.Println(buffer.String())
}