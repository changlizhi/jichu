package scfortests

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

func importforteststest(mkv string, buffer *bytes.Buffer) {
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//"changliang/fanshe"
	cfstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Fanshe(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cfstr)
	//"changliang/zf"
	zstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zstr)
	//"changliang/zfzhi"
	zzstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zzstr)
	//"log"
	lstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lstr)
	//"xxx/fortests"
	ftstr := zfzhi.Zhi.Syh() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ftstr)
	//"testing"
	ttstr := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ttstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
}
func zzcs(houzhui string, bmx string, buffer *bytes.Buffer) {
	//jiegouti+houzhui := fortests.Zuzhuangdtziyuan...string
	jgtstr := zf.Zfs.Jiegouti(true) + houzhui + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zuzhuang(false) + bmx + houzhui
	buffer.WriteString(jgtstr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.Test(true),
	cs1 := zh.Zhs.Zfszhtrue(zf.Zfs.Test(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cs1)
	if houzhui == zfzhi.Zhi.Kzf() {
		//zfzhi.Zhi.Shuzi4(),
		cs2 := zh.Zhs.Zhiszh(zf.Zfs.Shuzi4(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		buffer.WriteString(cs2)
	}
	//fanshe.Fangfaming(false),
	ffstr := zf.Zfs.Fanshe(true) + zfzhi.Zhi.Dh() + zf.Zfs.Fangfaming(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ffstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	//log.Println(jiegouti)
	lj := zh.Zhs.Logszh(zf.Zfs.Jiegouti(true)+houzhui) + zfzhi.Zhi.Hhf()
	buffer.WriteString(lj)
}
func testzuzhuangjiegou(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	//func TestZuzhuangdtziyuan
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Zuzhuang(false) + bmx
	buffer.WriteString(funstr)
	//(t *testing.T)
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	zzcs(zfzhi.Zhi.Kzf(), bmx, buffer)
	zzcs(zf.Zfs.Yige(true)+zf.Zfs.String(true), bmx, buffer)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}

func testzuzhuangbianmacs(houzhui string, bianma string, buffer *bytes.Buffer) {
	//func TestZuzhuangbianma+houzhui
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		zf.Zfs.Zuzhuang(false) + bianma + zf.Zfs.Bianma(true) + houzhui
	buffer.WriteString(funstr)
	//(t *testing.T)
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//jiegouti := fortests.Zuzhuangxxxbianma+houzhui
	jgt := zf.Zfs.Jiegouti(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zuzhuang(false) +
		bianma + zf.Zfs.Bianma(true) + houzhui
	buffer.WriteString(jgt)
	// (zf.Zfs.Kus(false) + zf.Zfs.Bianma(false) + fanshe.Fangfaming(false))
	jgtcs := zfzhi.Zhi.Xkhz() + zh.Zhs.Zfszhfalse(zf.Zfs.Kus(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhfalse(zf.Zfs.Bianma(false)) + zfzhi.Zhi.Jia() +
		zf.Zfs.Fanshe(true) + zfzhi.Zhi.Dh() + zf.Zfs.Fangfaming(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(jgtcs)
	//log.Println(jiegouti)
	lstr := zh.Zhs.Logszh(zf.Zfs.Jiegouti(true)) + zfzhi.Zhi.Hhf()
	buffer.WriteString(lstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}

func Shengchengforteststest() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Fanshebiaolies(mkv)
		for bk, _ := range biaos {
			buffer := &bytes.Buffer{}
			pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf()
			buffer.WriteString(pac)
			importforteststest(mkv, buffer)
			testzuzhuangjiegou(bk, buffer)
			testzuzhuangbianmacs(zfzhi.Zhi.Kzf(), bk, buffer)
			testzuzhuangbianmacs(zf.Zfs.String(true), bk, buffer)

			dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Fortests(true) +
				zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}

	}
}
