package scluyous

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

func getdelete(m string, bk string, buffer *bytes.Buffer) {
	//func TestBkm
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + bk + strings.ToLower(m)
	buffer.WriteString(funstr)
	//(t *testing.T)
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//r, _ := http.NewRequest
	hnr := zf.Zfs.R(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Xhx() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Http(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewRequest(false)
	buffer.WriteString(hnr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.GET(false),
	cs1 := zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + strings.ToUpper(m) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cs1)
	//zfzhi.Zhi.Xx() + zf.Zfs.Bk(true) + zfzhi.Zhi.Xx() + zfzhi.Zhi.Shuzi1w(),
	cs2 := zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + bk + zfzhi.Zhi.Xkhz() + zf.Zfs.True(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi1w(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cs2)
	//nil,
	buffer.WriteString(zf.Zfs.Nil(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//w := httptest.NewRecorder()
	wstr := zf.Zfs.W(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Httptest(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewRecorder(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(wstr)
	//beego.BeeApp.Handlers.ServeHTTP(w, r)
	bbhstr := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.BeeApp(false) + zfzhi.Zhi.Dh() + zf.Zfs.Handlers(false) + zfzhi.Zhi.Dh() + zf.Zfs.ServeHTTP(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.W(true) + zfzhi.Zhi.Dou() + zf.Zfs.R(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bbhstr)
	//log.Println(w.Body.String())
	lstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.W(true) + zfzhi.Zhi.Dh() + zf.Zfs.Body(false) + zfzhi.Zhi.Dh() + zf.Zfs.String(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func postpatch(canshu string, liebiao string, url string, m string, bk string, buffer *bytes.Buffer) {
	//func TestBkm
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + bk + liebiao + strings.ToLower(m)
	buffer.WriteString(funstr)
	//(t *testing.T)
	csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() + zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(csstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	buffer.WriteString(canshu)

	//r, _ := http.NewRequest
	hnr := zf.Zfs.R(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Xhx() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Http(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewRequest(false)
	buffer.WriteString(hnr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//zf.Zfs.M(false),
	mstr := zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + strings.ToUpper(m) + zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(mstr)
	//url,\n
	lystr := url + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
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

	//w := httptest.NewRecorder()
	hnrs := zf.Zfs.W(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Httptest(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewRecorder(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(hnrs)
	//beego.BeeApp.Handlers.ServeHTTP(w, r)
	bhss := zf.Zfs.Beego(true) + zfzhi.Zhi.Dh() + zf.Zfs.BeeApp(false) + zfzhi.Zhi.Dh() + zf.Zfs.Handlers(false) + zfzhi.Zhi.Dh() + zf.Zfs.ServeHTTP(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.W(true) + zfzhi.Zhi.Dou() + zf.Zfs.R(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bhss)
	//log.Println(w.Body.String())
	lstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.W(true) + zfzhi.Zhi.Dh() + zf.Zfs.Body(false) + zfzhi.Zhi.Dh() + zf.Zfs.String(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func importstr(mokuai string, buffer *bytes.Buffer) {
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//
	//"net/http"
	buffer.WriteString(zh.Zhs.Httpbao() + zfzhi.Zhi.Hhf())
	//"changliang/zf"
	zfbao := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfbao)
	//"changliang/zfzhi"
	zzstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zzstr)
	//"changliang/fanshe"
	fsstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() + zf.Zfs.Fanshe(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fsstr)
	//"net/http/httptest"
	buffer.WriteString(zh.Zhs.Httptestbao() + zfzhi.Zhi.Hhf())
	//"github.com/astaxie/beego"
	buffer.WriteString(zh.Zhs.Beegobao() + zfzhi.Zhi.Hhf())
	//"bytes"
	bstr := zfzhi.Zhi.Syh() + zf.Zfs.Bytes(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bstr)
	//"testing"
	tstr := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tstr)
	//"log"
	lstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lstr)

	//"xxx/fortests"
	ftstr := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() + zf.Zfs.Fortests(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ftstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

}

func Shengchengluyoustests() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Fanshebiaolies(mkv)
		for bk, bv := range biaos {
			buffer := &bytes.Buffer{}
			buffer.WriteString(zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf())
			importstr(mkv, buffer)

			getdelete(zf.Zfs.Get(false), bk, buffer)
			getdelete(zf.Zfs.Delete(false), bk, buffer)
			url := zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() + zh.Zhs.Zfszhtrue(bk)
			//canshu := fortests.Zuzhuangdtziyuanyigestring
			fzstr := zf.Zfs.Canshu(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
				zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zuzhuang(false) +
				bv + zf.Zfs.Yige(true) + zf.Zfs.String(true) + zfzhi.Zhi.Xkhz() +
				zfzhi.Zhi.Hhf() + zh.Zhs.Zfszhtrue(zf.Zfs.Test(false)) + zfzhi.Zhi.Dou() +
				zfzhi.Zhi.Hhf() + zf.Zfs.Fanshe(true) + zfzhi.Zhi.Dh() + zf.Zfs.Fangfaming(false) +
				zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() +
				zfzhi.Zhi.Hhf() + zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()

			postpatch(fzstr, zfzhi.Zhi.Kzf(), url, zf.Zfs.Post(false), bk, buffer)
			postpatch(fzstr, zfzhi.Zhi.Kzf(), url, zf.Zfs.Patch(false), bk, buffer)
			//列表的post查询
			urllie := zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
				zh.Zhs.Zfszhtrue(bk) + zfzhi.Zhi.Jia() +
				zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
				zh.Zhs.Zfszhtrue(zf.Zfs.Quanbu(false))
			postpatch(fzstr, zf.Zfs.Liebiao(true), urllie, zf.Zfs.Post(false), bk, buffer)

			//canshu := fortests.Zuzhuangxxxbianmastring(zf.Zfs.Test(true) + zf.Zfs.Bianma(false) + zf.Zfs.TestDtziyuanpost(false) + zfzhi.Zhi.Shuzi1w())
			cstiaojian := zf.Zfs.Canshu(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
				zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zuzhuang(false) + bk +
				zf.Zfs.Bianma(true) + zf.Zfs.String(true) + zfzhi.Zhi.Xkhz() +
				zh.Zhs.Zfszhtrue(zf.Zfs.Test(false)) + zfzhi.Zhi.Jia() +
				zh.Zhs.Zfszhfalse(zf.Zfs.Bianma(false)) + zfzhi.Zhi.Jia() +
				zh.Zhs.Zfszhfalse(zf.Zfs.TestDtziyuanpost(false)) + zfzhi.Zhi.Jia() +
				zh.Zhs.Zhiszh(zf.Zfs.Shuzi1w(false)) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			urltiaojian := zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
				zh.Zhs.Zfszhtrue(bk) + zfzhi.Zhi.Jia() +
				zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() +
				zh.Zhs.Zfszhtrue(zf.Zfs.Bianma(false))
			postpatch(cstiaojian, zf.Zfs.Tiaojian(true), urltiaojian, zf.Zfs.Post(false), bk, buffer)

			dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Luyou(true) + zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}
	}
}
