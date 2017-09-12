package scyewus

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
	"gongju"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func servicetestimport(mokuai string, bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	importstr := zf.Zfs.Import(true) + zfzhi.Zhi.Kgf()
	buffer.WriteString(importstr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	// "xxx/moxings"
	apmstr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Moxings(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(apmstr)

	// "changliang/zfzhi"
	zfzstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfzstr)
	// "xxx/zdjueseyewus"
	serstr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Zd(true) + bmx + zf.Zfs.Yewus(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(serstr)

	for _, lk := range gongju.Biao(mokuai, bianma) {
		if gongju.Lieleixing(lk) == zf.Zfs.Time(true) {
			// "time" \n
			timebao := zfzhi.Zhi.Syh() + zf.Zfs.Time(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
			buffer.WriteString(timebao)
			break
		}
	}

	//"log"
	logstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)
	// "testing"
	tstr := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy())
}
func tianjiaxiugai(mokuai string, fangfa string, bianma string, buffer *bytes.Buffer) {

	bmx := strings.ToLower(bianma)
	//func TestJueseserviceTianjia(t *testing.T)
	funcstr := zfzhi.Zhi.Hhf() + zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.Test(false) + bianma + zf.Zfs.Yewus(true) + fangfa +
		zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() +
		zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funcstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//juese := moxings.Juese
	objstr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(objstr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	for _, lk := range gongju.Biao(mokuai, bianma) {
		// Id : 1,
		if zf.Zfs.Int(true) == gongju.Lieleixing(lk) {
			intstr := lk + zfzhi.Zhi.Mh() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(intstr)
			continue
		}
		// Bianma :"",
		if zf.Zfs.String(true) == gongju.Lieleixing(lk) {
			strstr := lk + zfzhi.Zhi.Mh() + zfzhi.Zhi.Syh() +
				bianma + fangfa + zfzhi.Zhi.Syh() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(strstr)
			continue
		}
		// time : timeNow()
		if zf.Zfs.Time(true) == gongju.Lieleixing(lk) {
			timestr := lk + zfzhi.Zhi.Mh() + zf.Zfs.Time(true) +
				zfzhi.Zhi.Dh() + zf.Zfs.Now(false) + zfzhi.Zhi.Xkhz() +
				zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(timestr)
			continue
		}
		// Qian : 1.0
		if zf.Zfs.Float(true) == gongju.Lieleixing(lk) {
			flostr := lk + zfzhi.Zhi.Mh() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) + zfzhi.Zhi.Dh() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(flostr)
			continue
		}
	}
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy())
	//zdzfzhi.Zhi.Xxxyewus.Tianjiazfzhi.Zhi.Xxx(&zfzhi.Zhi.Xxx)
	tjstr := zfzhi.Zhi.Hhf() + zf.Zfs.Zd(true) + bmx + zf.Zfs.Yewus(true) +
		zfzhi.Zhi.Dh() + fangfa + bmx + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Qh() + bmx + zfzhi.Zhi.Xkhy()
	buffer.WriteString(tjstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testservicetianjia(mokuai string, bianma string, buffer *bytes.Buffer) {
	tianjiaxiugai(mokuai, zf.Zfs.Tianjia(false), bianma, buffer)
}
func testservicexiugai(mokuai string, bianma string, buffer *bytes.Buffer) {
	tianjiaxiugai(mokuai, zf.Zfs.Xiugai(false), bianma, buffer)
}
func testservicechaxun(bianma string, buffer *bytes.Buffer) {
	sz1str := strconv.Itoa(zfzhi.Zhi.Shuzi1())
	bmx := strings.ToLower(bianma)
	// func TestXzfzhi.Zhi.XxyeuwChaxun
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		bianma + zf.Zfs.Yewus(true) + zf.Zfs.Chaxun(false)
	buffer.WriteString(funstr)
	//(t *testing.T)
	tstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	// zfzhi.Zhi.Xx()x := zdzfzhi.Zhi.Xxxyewus.Chaxunzfzhi.Zhi.Xx()x(1)
	objstr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zd(true) +
		bmx + zf.Zfs.Yewus(true) + zfzhi.Zhi.Dh() + zf.Zfs.Chaxun(false) +
		bmx + zfzhi.Zhi.Xkhz() + sz1str + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(objstr)
	logstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testserviceshanchu(bianma string, buffer *bytes.Buffer) {
	sz1str := strconv.Itoa(zfzhi.Zhi.Shuzi1())
	bmx := strings.ToLower(bianma)
	// func TestXzfzhi.Zhi.Xx()serviceShanchu
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) +
		bianma + zf.Zfs.Yewus(true) + zf.Zfs.Shanchu(false)
	buffer.WriteString(funstr)
	//(t *testing.T)
	tstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//zdzfzhi.Zhi.Xx()xservices.Shanchuzfzhi.Zhi.Xx()x(1)
	scstr := zf.Zfs.Zd(true) + bmx + zf.Zfs.Yewus(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Shanchu(false) + bmx + zfzhi.Zhi.Xkhz() + sz1str + zfzhi.Zhi.Xkhy()
	buffer.WriteString(scstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func Shengchengyewutest() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Biaolies(mkv)
		for bk, _ := range biaos {
			buffer := bytes.Buffer{}
			pacstr := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf()
			buffer.WriteString(pacstr)
			servicetestimport(mkv, bk, &buffer)
			testservicetianjia(mkv, bk, &buffer)
			testservicexiugai(mkv, bk, &buffer)
			testservicechaxun(bk, &buffer)
			testserviceshanchu(bk, &buffer)
			dir := gongju.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Yewus(true) + zfzhi.Zhi.Xhx() +
				zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}
	}
}
