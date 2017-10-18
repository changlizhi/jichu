package sckongzhiqis

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

func importskongzhiqitest(main bool, mokuai string, buffer *bytes.Buffer) {
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//"xxx/kongzhiqis"
	constr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(constr)
	if !main {
		//"xxx/moxings"
		mxstr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(mxstr)
		//"xxx/fortests"
		ftstr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(ftstr)

		// "changliang/fanshe"
		fsstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
			zf.Zfs.Fanshe(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(fsstr)

	}

	//"gongju"
	gjstr := zfzhi.Zhi.Syh() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(gjstr)

	//"changliang/zf"
	zfstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfstr)

	// "changliang/zfzhi"
	zfzhistr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfzhistr)

	//"log"
	logstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)
	//"strconv"
	strcstr := zfzhi.Zhi.Syh() + zf.Zfs.Strconv(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(strcstr)
	//testing
	testingstr := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(testingstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
}

func postpatchkongzhiqitest(fangfa string, bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	//func TestFangfajuese(t *testing.T)
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		fangfa + bmx + zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//c := zdkongzhiqis.Juesekongzhiqi{}
	costr := zf.Zfs.C(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() +
		bianma + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Dkhz() +
		zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(costr)
	//gongju.Kongzhiqi
	gkstr := zf.Zfs.Gongju(true) + zfzhi.Zhi.Dh() + zf.Zfs.Kongzhiqi(false)
	buffer.WriteString(gkstr)

	// (&c.Controller)
	gkcs := zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Qh() + zf.Zfs.C(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Controller(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(gkcs)

	// reqjson := fortests.Zuzhuangdtziyuanyigestring
	jsonstr := zf.Zfs.Req(true) + zf.Zfs.Json(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Zuzhuang(false) + bmx + zf.Zfs.Yige(true) + zf.Zfs.String(true)
	buffer.WriteString(jsonstr)

	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	//zf.Zfs.Test(true),
	ztstr := zh.Zhs.Zfszhtrue(zf.Zfs.Test(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ztstr)
	//fanshe.Fangfaming(false),
	fsstr := zf.Zfs.Fanshe(true) + zfzhi.Zhi.Dh() + zf.Zfs.Fangfaming(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fsstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	// c.Ctx.Input.RequestBody = []byte(reqjson)
	cinstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ctx(false) +
		zfzhi.Zhi.Dh() + zf.Zfs.Input(false) + zfzhi.Zhi.Dh() +
		zf.Zfs.RequestBody(false) + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Zkhz() +
		zfzhi.Zhi.Zkhy() + zf.Zfs.Byte(true) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Req(true) + zf.Zfs.Json(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cinstr)

	//c.Fangfa()
	pstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + fangfa +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(pstr)

	//log.Println(c.Data)
	logstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.C(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Data(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func postkongzhiqitest(bianma string, buffer *bytes.Buffer) {
	postpatchkongzhiqitest(zf.Zfs.Post(false), bianma, buffer)
}
func patchkongzhiqitest(bianma string, buffer *bytes.Buffer) {
	postpatchkongzhiqitest(zf.Zfs.Patch(false), bianma, buffer)
}
func deletegetkongzhiqitest(fangfa string, bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	// func TestDeletezfzhi.Zhi.Xx()x(t * testing.T)
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		fangfa + bmx + zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//paramid := strconv.Itoa(zfzhi.Shuzi1zhi())
	parstr := zf.Zfs.Param(true) + zf.Zfs.Id(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.Strconv(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Itoa(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) +
		zfzhi.Zhi.Dh() + zf.Zfs.Shuzi1(false) + zfzhi.Zhi.Xkhz() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(parstr)

	//c := zdkongzhiqis.Juesekongzhiqi{}
	costr := zf.Zfs.C(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() +
		bianma + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Dkhz() +
		zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(costr)
	//gongju.Kongzhiqi
	gkstr := zf.Zfs.Gongju(true) + zfzhi.Zhi.Dh() + zf.Zfs.Kongzhiqi(false)
	buffer.WriteString(gkstr)

	// (&c.Controller)
	gkcs := zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Qh() + zf.Zfs.C(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Controller(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(gkcs)

	//c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false),paramid)
	cinstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ctx(false) +
		zfzhi.Zhi.Dh() + zf.Zfs.Input(false) + zfzhi.Zhi.Dh() + zf.Zfs.SetParam(false)
	buffer.WriteString(cinstr)

	//(zfzhi.Zhi.Mh() + zf.Zfs.Id(false),paramid)
	cinparam := zfzhi.Zhi.Xkhz() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Mh(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Jia() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) +
		zfzhi.Zhi.Dh() + zf.Zfs.Id(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() +
		zf.Zfs.Param(true) + zf.Zfs.Id(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cinparam)

	// c.Get()
	callstr := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + fangfa + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(callstr)

	//log.Println(c.Data[zf.Zfs.Json(true)])
	logstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.C(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Data(false) + zfzhi.Zhi.Zkhz() + zf.Zfs.Zf(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.True(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func deletekongzhiqitest(bianma string, buffer *bytes.Buffer) {
	deletegetkongzhiqitest(zf.Zfs.Delete(false), bianma, buffer)
}
func postliebiaokongzhiqitest(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	//func TestPostquanbuxxx
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		zf.Zfs.Post(false) + zf.Zfs.Quanbu(true) + bmx
	buffer.WriteString(funstr)
	//(t *testing.T)
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//c := zdkongzhiqis.Xxxliebiaokongzhiqi{}
	ckstr := zf.Zfs.C(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() +
		bianma + zf.Zfs.Liebiao(true) + zf.Zfs.Kongzhiqi(true) +
		zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ckstr)
	//gongju.Kongzhiqi
	gkstr := zf.Zfs.Gongju(true) + zfzhi.Zhi.Dh() + zf.Zfs.Kongzhiqi(false)
	buffer.WriteString(gkstr)
	//(&c.Controller)
	gkcs := zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Qh() + zf.Zfs.C(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Controller(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(gkcs)
	//c.Post()
	cp := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Post(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cp)
	//zhis := c.Data[zf.Zfs.Json(true)]
	zstr := zf.Zfs.Zhi(true) + zf.Zfs.S(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.C(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Data(false) + zfzhi.Zhi.Zkhz() + zh.Zhs.Zfszhtrue(zf.Zfs.Json(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zstr)
	//if zhi, ok := zhis.([]*moxings.Juese); ok
	ifstr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zhi(true) +
		zfzhi.Zhi.Dou() + zf.Zfs.Ok(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.Zhi(true) + zf.Zfs.S(true) +
		zfzhi.Zhi.Dh() + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Zkhz() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xh() + zf.Zfs.Moxings(true) +
		zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Fh() + zf.Zfs.Ok(true)
	buffer.WriteString(ifstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//for _, z := range zhi
	forstr := zf.Zfs.For(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() +
		zf.Zfs.Z(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Range(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zhi(true)
	buffer.WriteString(forstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//log.Println(z)
	lstr := zh.Zhs.Logszh(zf.Zfs.Z(true))
	buffer.WriteString(lstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func getkongzhiqitest(bianma string, buffer *bytes.Buffer) {
	deletegetkongzhiqitest(zf.Zfs.Get(false), bianma, buffer)
}
func posttiaojiankongzhiqitest(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	//func TestPosttiaojianshanchu
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Post(false) +
		zf.Zfs.Tiaojian(true) + zf.Zfs.Shanchu(true) + bmx
	buffer.WriteString(funstr)
	//(t *testing.T)
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//c := zdkongzhiqis.Dtziyuantiaojiankongzhiqi{}
	czd := zf.Zfs.C(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() +
		bianma + zf.Zfs.Tiaojian(true) + zf.Zfs.Kongzhiqi(true) +
		zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(czd)
	//gongju.Kongzhiqi(&c.Controller)
	gkc := zf.Zfs.Gongju(true) + zfzhi.Zhi.Dh() + zf.Zfs.Kongzhiqi(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Qh() + zf.Zfs.C(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Controller(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(gkc)
	//reqjson := fortests.Zuzhuangbianmastring
	rfz := zf.Zfs.Req(true) + zf.Zfs.Json(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Zuzhuang(false) + bianma + zf.Zfs.Bianma(true) + zf.Zfs.String(true)
	buffer.WriteString(rfz)
	// (zf.Zfs.Test(true) + zf.Zfs.Bianma(false) +
	// zf.Zfs.TestPostdtziyuan(false) + zfzhi.Zhi.Shuzi1w())
	zzz := zfzhi.Zhi.Xkhz() + zh.Zhs.Zfszhtrue(zf.Zfs.Test(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhfalse(zf.Zfs.Bianma(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhfalse(zf.Zfs.TestPostdtziyuan(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi1w(false)) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zzz)
	//c.Ctx.Input.RequestBody = []byte(reqjson)
	ccir := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ctx(false) + zfzhi.Zhi.Dh() +
		zf.Zfs.Input(false) + zfzhi.Zhi.Dh() + zf.Zfs.RequestBody(false) + zfzhi.Zhi.Dyh() +
		zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zf.Zfs.Byte(true) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Req(true) + zf.Zfs.Json(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ccir)
	//c.Post()
	cp := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Post(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cp)
	//log.Println(c.Data[zf.Zfs.Json(true)])
	lpcd := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Data(false) + zfzhi.Zhi.Zkhz() +
		zh.Zhs.Zfszhtrue(zf.Zfs.Json(false)) + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lpcd)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func Shengchengkongzhiqitest() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Fanshebiaolies(mkv)
		for bk, _ := range biaos {
			buffer := bytes.Buffer{}
			buffer.WriteString(zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf())
			importskongzhiqitest(false, mkv, &buffer)
			postkongzhiqitest(bk, &buffer)
			patchkongzhiqitest(bk, &buffer)
			deletekongzhiqitest(bk, &buffer)
			getkongzhiqitest(bk, &buffer)
			postliebiaokongzhiqitest(bk, &buffer)
			posttiaojiankongzhiqitest(bk, &buffer)

			//hanfuxn/tesets
			dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
			//xxx/tests/Xxx_test.go
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Kongzhiqi(true) +
				zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}
	}
}
