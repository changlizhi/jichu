package sckongzhiqis

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"gongju"
)

func importskongzhiqitest(buffer *bytes.Buffer) {
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//"github.com/astaxie/beego/context"
	gitstr := zfzhi.Zhi.Syh() + zf.Zfs.Github(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Com(true) + zfzhi.Zhi.Xx() + zf.Zfs.Astaxie(true) +
		zfzhi.Zhi.Xx() + zf.Zfs.Beego(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Context(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(gitstr)

	//"xxx/kongzhiqis"
	constr := zfzhi.Zhi.Syh() + gongju.Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi + zfzhi.Zhi.Xx() +
		zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(constr)

	//"gongju"
	gjstr := zfzhi.Zhi.Syh() + zf.Zfs.Gongju(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(gjstr)

	//"xxx/zf"
	zfstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfstr)

	// "xxx/zfzhi"
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
func bkkongzhiqitest(bianma string, buffer *bytes.Buffer) {
	sz200 := zfzhi.Zhi.Shuzi2() * zfzhi.Zhi.Shuzi10() * zfzhi.Zhi.Shuzi10()
	sz200str := strconv.Itoa(sz200)

	bmx := strings.ToLower(bianma)
	// func Xxkongzhiqi() *kongzhiqis.Xxxkongzhiqi
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + bmx +
		zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Xh() + zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() +
		bianma + zf.Zfs.Kongzhiqi(true)
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//  c := kongzhiqis.Xxxkongzhiqi{}
	objcon := zf.Zfs.C(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zd(true) + zf.Zfs.Kongzhiqis(true) + zfzhi.Zhi.Dh() + bianma +
		zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(objcon)
	// c.Data=make(map[interface{}]interface{})
	newd := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Data(false) +
		zfzhi.Zhi.Dyh() + zf.Zfs.Make(true) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Map(true) + zfzhi.Zhi.Zkhz() + zf.Zfs.Interface(true) +
		zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Zkhy() +
		zf.Zfs.Interface(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(newd)
	// c.Ctx=context.NewContext()
	newcon := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ctx(false) +
		zfzhi.Zhi.Dyh() + zf.Zfs.Context(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.NewContext(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(newcon)
	// c.Ctx.Input =context.NewInput()
	newin := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ctx(false) +
		zfzhi.Zhi.Dh() + zf.Zfs.Input(false) + zfzhi.Zhi.Dyh() +
		zf.Zfs.Context(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewInput(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(newin)
	//c.Ctx.Output=context.NewOutput()
	newout := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ctx(false) +
		zfzhi.Zhi.Dh() + zf.Zfs.Output(false) + zfzhi.Zhi.Dyh() +
		zf.Zfs.Context(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewOutput(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(newout)
	//c.Ctx.Output.Context = context.NewContext()
	outcon := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ctx(false) +
		zfzhi.Zhi.Dh() + zf.Zfs.Output(false) + zfzhi.Zhi.Dh() +
		zf.Zfs.Context(false) + zfzhi.Zhi.Dyh() + zf.Zfs.Context(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.NewContext(false) + zfzhi.Zhi.Xkhz() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(outcon)
	//c.Ctx.Output.Context.ResponseWriter = &context.Response
	conresp := zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ctx(false) +
		zfzhi.Zhi.Dh() + zf.Zfs.Output(false) + zfzhi.Zhi.Dh() +
		zf.Zfs.Context(false) + zfzhi.Zhi.Dh() + zf.Zfs.ResponseWriter(false) +
		zfzhi.Zhi.Dyh() + zfzhi.Zhi.Qh() + zf.Zfs.Context(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Response(false)
	buffer.WriteString(conresp)
	//{new(gongju.Mywriter),true,200}
	respobj := zfzhi.Zhi.Dkhz() + zf.Zfs.New(true) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Gongju(true) + zfzhi.Zhi.Dh() + zf.Zfs.Mywriter(false) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zf.Zfs.True(true) +
		zfzhi.Zhi.Dou() + sz200str + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(respobj)
	//return &c
	retcon := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Qh() + zf.Zfs.C(true)
	buffer.WriteString(retcon)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
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

	//c:= Xxxkongzhiqi()
	contr := zf.Zfs.C(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		bmx + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Xkhz() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(contr)

	// reqjson := zfzhi.Fangfajuesezhi()
	jsonstr := zf.Zfs.Req(true) + zf.Zfs.Json(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + fangfa + bmx +
		zf.Zfs.Zhi(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(jsonstr)

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

	//c := Xxxkongzhiqi()
	cstr := zf.Zfs.C(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		bmx + zf.Zfs.Kongzhiqi(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cstr)

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
func getkongzhiqitest(bianma string, buffer *bytes.Buffer) {
	deletegetkongzhiqitest(zf.Zfs.Get(false), bianma, buffer)
}
func Shengchengkongzhiqi() {
	_, biaos, _ := gongju.Biaolies()
	for bk, _ := range biaos {
		buffer := bytes.Buffer{}
		buffer.WriteString(zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf())
		importskongzhiqitest(&buffer)
		bkkongzhiqitest(bk, &buffer)
		postkongzhiqitest(bk, &buffer)
		patchkongzhiqitest(bk, &buffer)
		deletekongzhiqitest(bk, &buffer)
		getkongzhiqitest(bk, &buffer)

		//hanfuxn/tesets
		dir := gongju.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
		//xxx/tests/Xxx_test.go
		path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Kongzhiqi(true) +
			zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
