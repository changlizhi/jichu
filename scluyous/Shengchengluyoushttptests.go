package scluyous

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
	"gongju"
	"strings"
	"os"
	"io/ioutil"
)

func httptestimport(buffer *bytes.Buffer) {
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//"bytes"
	bstr := zfzhi.Zhi.Syh() + zf.Zfs.Bytes(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bstr)
	//"changliang/zf"
	ztr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ztr)
	//"changliang/zfzhi"

	zfzhistr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfzhistr)
	//"io/ioutil"
	iostr := zfzhi.Zhi.Syh() + zf.Zfs.Io(true) + zfzhi.Zhi.Xx() + zf.Zfs.Ioutil(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(iostr)
	//"log"
	lstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lstr)
	//"net/http"
	buffer.WriteString(zh.Zhs.Httpbao() + zfzhi.Zhi.Hhf())
	//"testing"
	tstr := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

}

func httppostpatch(bk string, m string, buffer *bytes.Buffer) {
	//func TestJuesehttppost
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + bk + strings.ToLower(zf.Zfs.Http(false)) + strings.ToLower(m)
	buffer.WriteString(funstr)
	//(t *testing.T)
	cs := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(cs)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//canshu := zfzhi.Zhi.Mjuese()
	reqcs := zf.Zfs.Canshu(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dh() + m + strings.ToLower(bk) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(reqcs)
	//p := zfzhi.Zhi.Shuzi8w() + zfzhi.Zhi.Shuzi0w() + zfzhi.Zhi.Shuzi8w() + zfzhi.Zhi.Shuzi0w()
	pstr := zf.Zfs.P(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi8w(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi0w(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi8w(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi0w(false)) + zfzhi.Zhi.Hhf()
	buffer.WriteString(pstr)

	//r, _ := http.NewRequest
	hnr := zf.Zfs.R(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Xhx() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Http(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewRequest(false)
	buffer.WriteString(hnr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.M(false),
	mstr := zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + strings.ToUpper(m) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(mstr)
	//zf.Zfs.Http(true)+zfzhi.Zhi.Mh()+zfzhi.Zhi.Xx()+zfzhi.Zhi.Xx()+zf.Zfs.Localhost(true)+zfzhi.Zhi.Mh()+p+
	prestr := zh.Zhs.Zfszhtrue(zf.Zfs.Http(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Mh(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhtrue(zf.Zfs.Localhost(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Mh(false)) + zfzhi.Zhi.Jia() +
		zf.Zfs.P(true) + zfzhi.Zhi.Jia()
	buffer.WriteString(prestr)

	//zfzhi.Zhi.Xx() + zf.Zfs.Bk(true),
	lystr := zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() + zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + bk + zfzhi.Zhi.Xkhz() + zf.Zfs.True(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lystr)
	//bytes.NewBuffer([]byte(canshu)),
	bfstr := zf.Zfs.Bytes(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewBuffer(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zf.Zfs.Byte(true) + zfzhi.Zhi.Xkhz() + zf.Zfs.Canshu(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bfstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//r.Header.Set
	sstr := zf.Zfs.R(true) + zfzhi.Zhi.Dh() + zf.Zfs.Header(false) + zfzhi.Zhi.Dh() + zf.Zfs.Set(false)
	buffer.WriteString(sstr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.Content(false) + zfzhi.Zhi.Jian() + zf.Zfs.Type(false),
	scs1 := zh.Zhs.Zfszhfalse(zf.Zfs.Content(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Jian(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhfalse(zf.Zfs.Type(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(scs1)
	//zf.Zfs.Application(true) + zfzhi.Zhi.Xx() + zf.Zfs.Json(true),
	scs2 := zh.Zhs.Zfszhtrue(zf.Zfs.Application(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhtrue(zf.Zfs.Json(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(scs2)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	//c := &http.Client{}
	clstr := zf.Zfs.C(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Qh() + zf.Zfs.Http(true) + zfzhi.Zhi.Dh() + zf.Zfs.Client(false) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(clstr)
	//response, _ := c.Do(r)
	restr := zf.Zfs.Response(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Xhx() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.C(true) + zfzhi.Zhi.Dh() + zf.Zfs.Do(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.R(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(restr)

	//body, _ := ioutil.ReadAll(response.Body)
	bstrs := zf.Zfs.Body(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Xhx() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Ioutil(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.ReadAll(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.Response(true) + zfzhi.Zhi.Dh() + zf.Zfs.Body(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bstrs)
	//log.Println(string(body))
	lstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.String(true) + zfzhi.Zhi.Xkhz() + zf.Zfs.Body(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}

func Shengchengluyoushttptests() {
	mks := gongju.Mokuaimings
	for _, mkvo := range mks {
		mkv := mkvo.Zhi
		_, biaos, _ := gongju.Biaolies(mkv)
		for bk, _ := range biaos {
			buffer := &bytes.Buffer{}
			buffer.WriteString(zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf())
			httptestimport(buffer)
			httppostpatch(bk, zf.Zfs.Post(false), buffer)
			httppostpatch(bk, zf.Zfs.Patch(false), buffer)
			dir := gongju.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Luyou(true) + zf.Zfs.Http(true) + zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}
	}
}
