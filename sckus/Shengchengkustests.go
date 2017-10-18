package sckus

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

// tests的import
func testsimports(mokuai string, bianma string, buffer *bytes.Buffer) {
	pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf()
	buffer.WriteString(pac)

	// import(\n
	impstr := zf.Zfs.Import(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(impstr)

	//"xxx/fortests"
	ftstr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ftstr)

	// "testing" \n
	tbao := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tbao)

	// "log" \n
	logbao := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logbao)

	//"changliang/zf"
	zstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zstr)

	//"changliang/zfzhi"
	zzstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zzstr)

	//"changliang/fanshe"
	fsstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Fanshe(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fsstr)

	// "xxx/zdjueseshujukuduqus"
	daobao := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Kus(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(daobao)

	buffer.WriteString(zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
}

func testchaxunquanbu(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	//func TestChaxunquanbuJuese
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		zf.Zfs.Chaxunquanbu(false) + bianma + zf.Zfs.Kus(true)
	buffer.WriteString(funstr)
	// (t *testing.T)
	cs := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(cs)
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

	//all := zdjuesekus.Chaxunquanbu(dtziyuan[zfzhi.Zhi.Shuzi0()])
	astr := zf.Zfs.All(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zd(true) + bmx + zf.Zfs.Kus(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Chaxunquanbu(false) + zf.Zfs.Kus(true) + zfzhi.Zhi.Xkhz() +
		bmx + zfzhi.Zhi.Zkhz() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(astr)
	//log.Println(all[zfzhi.Zhi.Shuzi0()])
	lstr := zh.Zhs.Logszh(zf.Zfs.All(true) + zfzhi.Zhi.Zkhz() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) + zfzhi.Zhi.Zkhy())
	buffer.WriteString(lstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func testchaxunyige(bianma string, buffer *bytes.Buffer) {

	// func TestChaxunyigeXXX
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		zf.Zfs.Chaxunyige(false) + bianma + zf.Zfs.Kus(true)
	buffer.WriteString(funstr)

	//(t*testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	//juese:=
	//zdjuesedaos.Chaxunyige(1)
	fhstr := strings.ToLower(bianma) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Kus(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Chaxunyige(false) + zf.Zfs.Kus(true) +
		zfzhi.Zhi.Xkhz() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fhstr)

	//log.Println(juese)
	dayin := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) + zfzhi.Zhi.Xkhz() + strings.ToLower(bianma) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dayin)

	// \n } \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testtianjiayige(bianma string, buffer *bytes.Buffer) {
	tianjiahuoxiugai(bianma, buffer, zf.Zfs.Tianjiayige(false))
}
func testxiugaiyige(bianma string, buffer *bytes.Buffer) {
	tianjiahuoxiugai(bianma, buffer, zf.Zfs.Xiugaiyige(false))
}

func tianjiahuoxiugai(bianma string, buffer *bytes.Buffer, fangfa string) {
	bmx := strings.ToLower(bianma)
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + fangfa + bianma + zf.Zfs.Kus(true)
	buffer.WriteString(funstr)

	// (t *testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() +
		zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//dtziyuan := fortests.Zuzhuangdtziyuan
	ostr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zuzhuang(false) + bmx
	buffer.WriteString(ostr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.Kus(false),
	kstr := zh.Zhs.Zfszhfalse(zf.Zfs.Kus(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(kstr)
	//zfzhi.Zhi.Shuzi1(),
	szstr := zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(szstr)
	//fanshe.Fangfaming(false),
	fsstr := zf.Zfs.Fanshe(true) + zfzhi.Zhi.Dh() + zf.Zfs.Fangfaming(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fsstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//zddtziyuankus.Xxx(dtziyuan[zfzhi.Zhi.Shuzi0()])
	dystr := zf.Zfs.Zd(true) + bmx + zf.Zfs.Kus(true) + zfzhi.Zhi.Dh() +
		fangfa + zf.Zfs.Kus(true) + zfzhi.Zhi.Xkhz() + bmx +
		zfzhi.Zhi.Zkhz() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dystr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testshanchuyige(bianma string, buffer *bytes.Buffer) {

	// func TestShanchuyige
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Shanchuyige(false) + bianma
	buffer.WriteString(funstr)

	// (t *testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	// zdjuesedaos.Shanchuyige(1)
	baoming := zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Kus(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Shanchuyige(false) + zf.Zfs.Kus(true) + zfzhi.Zhi.Xkhz() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(baoming)
	// \n } \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testtianjiaduoge(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		zf.Zfs.Tianjiaduoge(false) + bianma + zf.Zfs.Kus(true)
	buffer.WriteString(funstr)

	// (t *testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() +
		zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//dtziyuan := fortests.Zuzhuangdtziyuan
	ostr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zuzhuang(false) + bmx
	buffer.WriteString(ostr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.Kus(false),
	kstr := zh.Zhs.Zfszhfalse(zf.Zfs.Kus(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(kstr)
	//zfzhi.Zhi.Shuzi3(),
	szstr := zh.Zhs.Zhiszh(zf.Zfs.Shuzi3(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(szstr)
	//fanshe.Fangfaming(false),
	fsstr := zf.Zfs.Fanshe(true) + zfzhi.Zhi.Dh() + zf.Zfs.Fangfaming(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fsstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//zddtziyuankus.Xxx(dtziyuan)
	dystr := zf.Zfs.Zd(true) + bmx + zf.Zfs.Kus(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Tianjiaduoge(false) + zf.Zfs.Kus(true) + zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dystr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func testshanchutiaojian(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	//func TestShanchutiaojianxxxkus(t *testing.T)
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		zf.Zfs.Shanchu(false) + zf.Zfs.Tiaojian(true) + bianma + zf.Zfs.Kus(true)
	buffer.WriteString(funstr)
	// (t *testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() +
		zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//dtziyuan := fortests.Zuzhuangxxxbianma
	dfz := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zuzhuang(false) +
		bianma + zf.Zfs.Bianma(true)
	buffer.WriteString(dfz)
	// (zf.Zfs.Kus(false) + zf.Zfs.Bianma(false) + "TestTianjiayigeDtziyuankus1")
	dfzcs := zfzhi.Zhi.Xkhz() + zh.Zhs.Zfszhfalse(zf.Zfs.Kus(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhfalse(zf.Zfs.Bianma(false)) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf() +
		zh.Zhs.Zfszhfalse(zf.Zfs.TestTianjiayigeDtziyuankus(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi1w(false)) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dfzcs)
	//zdxxxkus.Shanchutiaojiankus(dtziyuan)
	zs := zf.Zfs.Zd(true) + bmx + zf.Zfs.Kus(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Shanchu(false) + zf.Zfs.Tiaojian(true) + zf.Zfs.Kus(true) +
		zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zs)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func Shengchengkustests() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Fanshebiaolies(mkv)
		for bk, _ := range biaos {
			buffer := bytes.Buffer{}
			testsimports(mkv, bk, &buffer)
			testtianjiaduoge(bk, &buffer)
			testtianjiayige(bk, &buffer)
			testxiugaiyige(bk, &buffer)
			testchaxunyige(bk, &buffer)
			testshanchuyige(bk, &buffer)
			testchaxunquanbu(bk, &buffer)
			testshanchutiaojian(bk, &buffer)
			dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Ku(true) + zfzhi.Zhi.Xhx() +
				zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}
	}
}
