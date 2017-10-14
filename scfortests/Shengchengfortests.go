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

func importsfortests(mkv string, bk string, buffer *bytes.Buffer) {
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	//"xxx/moxings"
	mstr := zfzhi.Zhi.Syh() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(mstr)

	for _, lk := range gongju.Biao(mkv, bk) {
		if gongju.Lieleixing(lk) == zf.Zfs.Time(true) {
			// "time" \n
			timebao := zfzhi.Zhi.Syh() + zf.Zfs.Time(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
			buffer.WriteString(timebao)
			break
		}
	}

	//"strconv"
	sstr := zfzhi.Zhi.Syh() + zf.Zfs.Strconv(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(sstr)
	//"changliang/zf"
	zstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zstr)
	//"changliang/zfzhi"
	zzstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zzstr)
	//"encoding/json"
	jstr := zfzhi.Zhi.Syh() + zf.Zfs.Encoding(true) + zfzhi.Zhi.Xx() + zf.Zfs.Json(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(jstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
}

func zuzhuangjiegoutistr(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	//func Zuzhuangyigexxxstring
	fstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zuzhuang(false) + bmx + zf.Zfs.Yige(true) + zf.Zfs.String(true)
	buffer.WriteString(fstr)

	//(leixing string, fangfa string)
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.Leixing(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.String(true) + zfzhi.Zhi.Dou() + zf.Zfs.Fangfa(true) +
		zfzhi.Zhi.Kgf() + zf.Zfs.String(true) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	buffer.WriteString(zf.Zfs.String(true))
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//ret := Zuzhuangdtziyuan(leixing, zfzhi.Zhi.Shuzi1(), fangfa)
	rstr := zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zuzhuang(false) + bmx + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Leixing(true) + zfzhi.Zhi.Dou() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) +
		zfzhi.Zhi.Dou() + zf.Zfs.Fangfa(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(rstr)

	//jstring, _ := json.Marshal(ret[zfzhi.Zhi.Shuzi0()])
	jmstr := zf.Zfs.J(true) + zf.Zfs.String(true) + zfzhi.Zhi.Dou() +
		zfzhi.Zhi.Xhx() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Json(true) + zfzhi.Zhi.Dh() + zf.Zfs.Marshal(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Ret(true) + zfzhi.Zhi.Zkhz() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(jmstr)
	//return string(jstring)
	rrstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.String(true) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.J(true) + zf.Zfs.String(true) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(rrstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func zuzhuangjiegouti(mokuai string, bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	//func Zuzhuangxx
	fstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zuzhuang(false) + bmx
	buffer.WriteString(fstr)
	//(leixing string, jige int, fangfa string)
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.Leixing(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.String(true) + zfzhi.Zhi.Dou() + zf.Zfs.Jige(true) +
		zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) + zfzhi.Zhi.Dou() +
		zf.Zfs.Fangfa(true) + zfzhi.Zhi.Kgf() + zf.Zfs.String(true) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	// []*moxings.Dtziyuan
	fhstr := zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(fhstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//ret := make
	rmstr := zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Make(true)
	buffer.WriteString(rmstr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//[]*moxings.Dtziyuan,
	cs1 := zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xh() +
		zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cs1)
	//zfzhi.Zhi.Shuzi0(),
	cs2 := zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cs2)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	//for i := zfzhi.Zhi.Shuzi1(); i <= jige; i++
	forstr := zf.Zfs.For(true) + zfzhi.Zhi.Kgf() + zf.Zfs.I(true) +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) +
		zfzhi.Zhi.Fh() + zf.Zfs.I(true) + zfzhi.Zhi.Xy() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Jige(true) + zfzhi.Zhi.Fh() + zf.Zfs.I(true) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Jia()
	buffer.WriteString(forstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//istring := strconv.Itoa(i)
	sstr := zf.Zfs.I(true) + zf.Zfs.String(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Strconv(true) + zfzhi.Zhi.Dh() + zf.Zfs.Itoa(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.I(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(sstr)
	//dtziyuan := &moxings.Dtziyuan
	ostr := bmx + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Qh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(ostr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	gongju.Pinjieziduan(mokuai, bianma, buffer)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	//ret = append(ret, dtziyuan)
	rstr := zf.Zfs.Ret(true) + zfzhi.Zhi.Dyh() + zf.Zfs.Append(true) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Ret(true) + zfzhi.Zhi.Dou() +
		bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(rstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	//return ret
	rrstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ret(true)
	buffer.WriteString(rrstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func zuzhuangbianmajiegoutistring(bianma string, buffer *bytes.Buffer) {
	// func Zuzhuangxxxbianmastring
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zuzhuang(false) +
		bianma + zf.Zfs.Bianma(true) + zf.Zfs.String(true)
	buffer.WriteString(funstr)
	// (bianma string) string
	funcsstr := zfzhi.Zhi.Xkhz() + zf.Zfs.Bianma(true) + zfzhi.Zhi.Kgf() + zf.Zfs.String(true) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf() + zf.Zfs.String(true)
	buffer.WriteString(funcsstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//ret := Zuzhuangxxxbianma(bianma)
	rstr := zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zuzhuang(false) + bianma + zf.Zfs.Bianma(true) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Bianma(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(rstr)
	//jstring, _ := json.Marshal(ret)
	jmstr := zf.Zfs.J(true) + zf.Zfs.String(true) + zfzhi.Zhi.Dou() +
		zfzhi.Zhi.Xhx() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Json(true) + zfzhi.Zhi.Dh() + zf.Zfs.Marshal(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Ret(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(jmstr)
	//return string(jstring)
	rrstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.String(true) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.J(true) + zf.Zfs.String(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(rrstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func zuzhuangbianmajiegou(bianma string, buffer *bytes.Buffer) {
	//func Zuzhuangbianma
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zuzhuang(false) + bianma + zf.Zfs.Bianma(true)
	buffer.WriteString(funstr)
	//(bianma string) *moxings.Dtziyuan
	funcsstr := zfzhi.Zhi.Xkhz() + zf.Zfs.Bianma(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.String(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(funcsstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//ret := &moxings.Dtziyuan
	rstr := zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zfzhi.Zhi.Qh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(rstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//Bianma:bianma,
	bbstr := zf.Zfs.Bianma(false) + zfzhi.Zhi.Mh() + zf.Zfs.Bianma(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bbstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//return ret
	rrstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ret(true)
	buffer.WriteString(rrstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func Shengchengfortests() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Biaolies(mkv)
		for bk, _ := range biaos {
			buffer := &bytes.Buffer{}
			pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Hhf()
			buffer.WriteString(pac)
			importsfortests(mkv, bk, buffer)
			zuzhuangjiegouti(mkv, bk, buffer)
			zuzhuangjiegoutistr(bk, buffer)
			zuzhuangbianmajiegou(bk, buffer)
			zuzhuangbianmajiegoutistring(bk, buffer)

			dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Fortests(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}

	}
}
