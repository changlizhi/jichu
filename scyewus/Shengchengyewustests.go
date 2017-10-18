package scyewus

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

func yewustestimport(mokuai string, bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	importstr := zf.Zfs.Import(true) + zfzhi.Zhi.Kgf()
	buffer.WriteString(importstr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	// "changliang/zfzhi"
	zzstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zzstr)
	// "changliang/zf"
	zstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zstr)
	// "changliang/fashe"
	fsstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Fanshe(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fsstr)
	// "xxx/zdxxxyewus"
	serstr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Zd(true) + bmx + zf.Zfs.Yewus(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(serstr)
	// "mhsydata/fortests"
	ftstr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Fortests(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ftstr)
	//"log"
	logstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)
	// "testing"
	tstr := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy())
}
func tianjiaxiugai(fangfa string, bianma string, buffer *bytes.Buffer) {

	bmx := strings.ToLower(bianma)
	//func TestXiugaiyigexxxyewus(t *testing.T)
	funcstr := zfzhi.Zhi.Hhf() + zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.Test(false) + fangfa + bianma + zf.Zfs.Yewus(true) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() +
		zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funcstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//dtziyuan := fortests.Zuzhuangdtziyuan
	ostr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zuzhuang(false) + bmx
	buffer.WriteString(ostr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.Yewus(false),
	kstr := zh.Zhs.Zfszhfalse(zf.Zfs.Yewus(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(kstr)
	//zfzhi.Zhi.Shuzi1(),
	szstr := zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(szstr)
	//fanshe.Fangfaming(false),
	fsstr := zf.Zfs.Fanshe(true) + zfzhi.Zhi.Dh() + zf.Zfs.Fangfaming(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fsstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//zddtziyuanyewus.Xxx(dtziyuan[zfzhi.Zhi.Shuzi0()])
	dystr := zf.Zfs.Zd(true) + bmx + zf.Zfs.Yewus(true) + zfzhi.Zhi.Dh() +
		fangfa + zf.Zfs.Yewus(true) + zfzhi.Zhi.Xkhz() + bmx +
		zfzhi.Zhi.Zkhz() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dystr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testyewustianjia(bianma string, buffer *bytes.Buffer) {
	tianjiaxiugai(zf.Zfs.Tianjiayige(false), bianma, buffer)
}
func testyewusxiugai(bianma string, buffer *bytes.Buffer) {
	tianjiaxiugai(zf.Zfs.Xiugaiyige(false), bianma, buffer)
}
func testyewuschaxunquanbu(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	// func TestChaxunquanbuxxxyewus
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Chaxunquanbu(false) + bianma + zf.Zfs.Yewus(true)
	buffer.WriteString(funstr)
	// (t *testing.T)
	tstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//dtziyuan := fortests.Zuzhuangdtziyuan
	ostr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zuzhuang(false) + bmx
	buffer.WriteString(ostr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.Test(true),
	kstr := zh.Zhs.Zfszhtrue(zf.Zfs.Test(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(kstr)
	//zfzhi.Zhi.Shuzi1(),
	szstr := zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(szstr)
	//fanshe.Fangfaming(false),
	fsstr := zf.Zfs.Fanshe(true) + zfzhi.Zhi.Dh() + zf.Zfs.Fangfaming(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fsstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	// quanbu := zdjueseyewus.Chaxunquanbuyewus(dtziyuan[zfzhi.Zhi.Shuzi0()])
	qbstr := zf.Zfs.Quanbu(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zd(true) +
		bmx + zf.Zfs.Yewus(true) + zfzhi.Zhi.Dh() + zf.Zfs.Chaxunquanbu(false) + zf.Zfs.Yewus(true) +
		zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Zkhz() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(qbstr)

	// log.Println(quanbu[zfzhi.Zhi.Shuzi0()])
	lstr := zh.Zhs.Logszh(zf.Zfs.Quanbu(true) + zfzhi.Zhi.Zkhz() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) + zfzhi.Zhi.Zkhy())
	buffer.WriteString(lstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func testyewuschaxun(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	// func TestChaxunyigexxxyewus
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		zf.Zfs.Chaxunyige(false) + bianma + zf.Zfs.Yewus(true)
	buffer.WriteString(funstr)
	//(t *testing.T)
	tstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	// xxx := zdxxxyewus.Chaxunyigeyewus(zfzhi.Zhi.Shuzi1())
	objstr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zd(true) +
		bmx + zf.Zfs.Yewus(true) + zfzhi.Zhi.Dh() + zf.Zfs.Chaxunyige(false) +
		zf.Zfs.Yewus(true) + zfzhi.Zhi.Xkhz() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(objstr)
	logstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testyewusshanchu(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	// func TestShanchuyigexxxyewus
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		zf.Zfs.Shanchuyige(false) + bianma + zf.Zfs.Yewus(true)
	buffer.WriteString(funstr)
	//(t *testing.T)
	tstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//zdxxxyewuss.Shanchuyigeyewus(zfzhi.Zhi.Shuzi1())
	scstr := zf.Zfs.Zd(true) + bmx + zf.Zfs.Yewus(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Shanchuyige(false) + zf.Zfs.Yewus(true) + zfzhi.Zhi.Xkhz() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(scstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testyewustiaojianshanchu(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	//func TestShanchutiaojianxxxyewus
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		zf.Zfs.Shanchu(false) + zf.Zfs.Tiaojian(true) + bianma + zf.Zfs.Yewus(true)
	buffer.WriteString(funstr)
	//(t *testing.T)
	tstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//dtziyuan := fortests.Zuzhuangxxxbianma
	cs := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Fortests(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Zuzhuang(false) + bianma + zf.Zfs.Bianma(true)
	buffer.WriteString(cs)
	// (zf.Zfs.Yewus(false) + zf.Zfs.Bianma(false) +
	// zf.Zfs.TestTianjiayigeyewus(false) + zfzhi.Zhi.Shuzi1w())
	css := zfzhi.Zhi.Xkhz() + zh.Zhs.Zfszhfalse(zf.Zfs.Yewus(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhfalse(zf.Zfs.Bianma(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhfalse(zf.Zfs.TestTianjiayigeyewus(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi1w(false)) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(css)
	//tishis := zddtziyuanyewus.Shanchutiaojianyewus(dtziyuan)
	tsz := zf.Zfs.Tishis(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zd(true) + bmx + zf.Zfs.Yewus(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Shanchu(false) + zf.Zfs.Tiaojian(true) + zf.Zfs.Yewus(true) +
		zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tsz)
	//log.Println("tishis:====", tishis)
	lstr := zh.Zhs.Logszh(zf.Zfs.Tishis(true)) + zfzhi.Zhi.Hhf()
	buffer.WriteString(lstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func Shengchengyewutest() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Fanshebiaolies(mkv)
		for bk, _ := range biaos {
			buffer := bytes.Buffer{}
			pacstr := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf()
			buffer.WriteString(pacstr)
			yewustestimport(mkv, bk, &buffer)
			testyewustianjia(bk, &buffer)
			testyewusxiugai(bk, &buffer)
			testyewuschaxun(bk, &buffer)
			testyewusshanchu(bk, &buffer)
			testyewuschaxunquanbu(bk, &buffer)
			testyewustiaojianshanchu(bk, &buffer)
			dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Yewus(true) + zfzhi.Zhi.Xhx() +
				zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}
	}
}
