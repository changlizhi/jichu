package sckus

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

// tests的import
func testsimports(mokuai string, bianma string, buffer *bytes.Buffer) {
	pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf()
	buffer.WriteString(pac)

	// import(\n
	impstr := zf.Zfs.Import(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(impstr)

	for _, lk := range gongju.Biao(mokuai, bianma) {
		if gongju.Lieleixing(lk) == zf.Zfs.Time(true) {
			// "time" \n
			timebao := zfzhi.Zhi.Syh() + zf.Zfs.Time(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
			buffer.WriteString(timebao)
			break
		}
	}
	// "xxx/moxings" \n
	apm := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Moxings(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(apm)

	// "testing" \n
	tbao := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tbao)

	// "log" \n
	logbao := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(logbao)
	//"changliang/zfzhi"
	clstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(clstr)

	// "xxx/zdjueseshujukuduqus"
	daobao := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Kus(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(daobao)

	buffer.WriteString(zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
}

func testchaxunyige(bianma string, buffer *bytes.Buffer) {

	// func TestChaxunyigeXXX
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Chaxunyige(false) + bianma
	buffer.WriteString(funstr)

	//(t*testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	//juese:=
	//zdjuesedaos.Chaxunyige(1)
	fhstr := strings.ToLower(bianma) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Kus(true) + zfzhi.Zhi.Dh() + zf.Zfs.Chaxunyige(false) + zfzhi.Zhi.Xkhz() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fhstr)

	//log.Println(juese)
	dayin := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) + zfzhi.Zhi.Xkhz() + strings.ToLower(bianma) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dayin)

	// \n } \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testtianjiayige(mokuai string, bianma string, buffer *bytes.Buffer) {
	tianjiahuoxiugai(mokuai, bianma, buffer, zf.Zfs.Tianjiayige(false))
}
func testxiugaiyige(mokuai string, bianma string, buffer *bytes.Buffer) {
	tianjiahuoxiugai(mokuai, bianma, buffer, zf.Zfs.Xiugaiyige(false))
}

func tianjiahuoxiugai(mokuai string, bianma string, buffer *bytes.Buffer, fangfa string) {
	sz1 := strconv.Itoa(zfzhi.Zhi.Shuzi1())

	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + fangfa + bianma
	buffer.WriteString(funstr)

	// (t *testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	// juese := &moxings.Juese{\n
	dxstr := strings.ToLower(bianma) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Qh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dxstr)

	// 生成字段和值
	pinjieziduan(mokuai, bianma, buffer, fangfa, sz1)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	// zdjuesedaos.Testzfzhi.Zhi.Xx()xJuese(juese)
	baoming := zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Kus(true) + zfzhi.Zhi.Dh() + fangfa + zfzhi.Zhi.Xkhz() + strings.ToLower(bianma) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(baoming)

	// \n } \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func pinjieziduan(mokuai string, bianma string, buffer *bytes.Buffer, fangfa string, houzhui string) {

	for _, lk := range gongju.Biao(mokuai, bianma) {
		buffer.WriteString(lk + zfzhi.Zhi.Mh())
		zhi := shengchengzhi(lk, gongju.Lieleixing(lk), fangfa, houzhui)
		buffer.WriteString(zhi + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())
	}

}
func shengchengzhi(lieming string, leixing string, fangfa string, houzhui string) string {
	if leixing == zf.Zfs.String(true) {
		// "lieZengjiaTest1"
		ret := zfzhi.Zhi.Syh() + lieming + fangfa + zf.Zfs.Test(false) + houzhui + zfzhi.Zhi.Syh()
		return ret
	}
	// 1
	if leixing == zf.Zfs.Int(true) {
		ret := zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) // 1
		return ret
	}
	// time.Now()
	if leixing == zf.Zfs.Time(true) {
		// time.Now()
		ret := zf.Zfs.Time(true) + zfzhi.Zhi.Dh() + zf.Zfs.Now(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
		return ret
	}
	// 1.0
	if leixing == zf.Zfs.Float32(true) || leixing == zf.Zfs.Float64(true) {
		// 1.0
		ret := zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) + zfzhi.Zhi.Dh() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false))
		return ret
	}
	// Null
	return zf.Zfs.Null(false)
}

func testshanchuyige(bianma string, buffer *bytes.Buffer) {

	// func TestShanchuyige
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Shanchuyige(false) + bianma
	buffer.WriteString(funstr)

	// (t *testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	// zdjuesedaos.Shanchuyige(1)
	baoming := zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Kus(true) + zfzhi.Zhi.Dh() + zf.Zfs.Shanchuyige(false) + zfzhi.Zhi.Xkhz() + zh.Zhs.Zhiszh(zf.Zfs.Shuzi1(false)) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(baoming)
	// \n } \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func testtianjiaduoge(mokuai string, bianma string, buffer *bytes.Buffer) {
	//zfzhi.Zhi.Syh() := zfzhi.Shuangyinhaozhi()
	sz2 := strconv.Itoa(zfzhi.Zhi.Shuzi2())
	sz3 := strconv.Itoa(zfzhi.Zhi.Shuzi3())
	//组装方法名
	// func TestTianjiaduogeJuese
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Tianjiaduoge(false) + bianma
	buffer.WriteString(funstr)

	//(t *testing.T){\n
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	// 组装第一个实体
	//juese2 := moxings.Juese
	shiti2 := strings.ToLower(bianma) + sz2 + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(shiti2)

	//{ \n zhi... \n }
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	pinjieziduan(mokuai, bianma, buffer, zf.Zfs.Tianjiaduoge(false), sz2)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	// juese3 := moxings.Juese{\n
	shiti3 := strings.ToLower(bianma) + sz3 + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(shiti3)

	// \n zhi \n
	pinjieziduan(mokuai, bianma, buffer, zf.Zfs.Tianjiaduoge(false), sz3)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	// jueses :=[]moxings.Juese
	duoge := strings.ToLower(bianma) + zf.Zfs.S(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(duoge)

	//{juese1,juese2}
	szs := zfzhi.Zhi.Dkhz() + strings.ToLower(bianma) + sz2 + zfzhi.Zhi.Dou() + strings.ToLower(bianma) + sz3 + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(szs)

	//zdjuesedaos.Tianjiaduoge(jueses)
	baoming := zf.Zfs.Zd(true) + strings.ToLower(bianma) + zf.Zfs.Kus(true) + zfzhi.Zhi.Dh() + zf.Zfs.Tianjiaduoge(false) + zfzhi.Zhi.Xkhz() + strings.ToLower(bianma) + zf.Zfs.S(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(baoming)

	//} \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func Shengchengdaostests() {
	mks := gongju.Mokuaimings
	for _, mkvo := range mks {
		mkv := mkvo.Zhi
		_, biaos, _ := gongju.Biaolies(mkv)
		for bk, _ := range biaos {
			buffer := bytes.Buffer{}
			testsimports(mkv, bk, &buffer)
			testtianjiaduoge(mkv, bk, &buffer)
			testtianjiayige(mkv, bk, &buffer)
			testxiugaiyige(mkv, bk, &buffer)
			testchaxunyige(bk, &buffer)
			testshanchuyige(bk, &buffer)

			path := gongju.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true) +
				zfzhi.Zhi.Xx() + bk + zf.Zfs.Ku(true) + zfzhi.Zhi.Xhx() +
				zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}
	}
}
