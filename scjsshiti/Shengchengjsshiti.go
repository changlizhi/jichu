package scjsshiti

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"io/ioutil"
	"os"
	"changliang/zw"
	"reflect"
)

func Shengchengjsshiti() {
	for _, mkv := range gongju.Mokuaimingsarr {
		zw := &zw.Zw{}
		zwrv := reflect.ValueOf(zw)
		zwrt := reflect.TypeOf(zw)
		buffer := &bytes.Buffer{}
		//export function lieming (bianma) {
		funstr := zf.Zfs.Export(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Function(true) +
			zfzhi.Zhi.Kgf() + zf.Zfs.Lieming(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() +
			zf.Zfs.Bianma(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf()
		buffer.WriteString(funstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//var shiti =
		vstr := zf.Zfs.Var(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Shiti(true) +
			zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Kgf()
		buffer.WriteString(vstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//A: 'A'为了好生成
		astr := zf.Zfs.A(false) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Kgf() +
			zfzhi.Zhi.Dyhe() + zf.Zfs.A(false) + zfzhi.Zhi.Dyhe()
		buffer.WriteString(astr)

		for i := zfzhi.Zhi.Shuzi0(); i < zwrt.NumMethod(); i++ {
			mn := zwrt.Method(i).Name
			m := zwrv.MethodByName(mn)
			val := m.Call(nil)[zfzhi.Zhi.Shuzi0()].String()
			zfff := zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf() + mn + zfzhi.Zhi.Mh() +
				zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dyhe() + val + zfzhi.Zhi.Dyhe()
			buffer.WriteString(zfff)
		}
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
		//return shiti[bianma]
		restr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Shiti(true) +
			zfzhi.Zhi.Zkhz() + zf.Zfs.Bianma(true) + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(restr)
		buffer.WriteString(zfzhi.Zhi.Dkhy())

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Js(true)
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Lie(true) + zfzhi.Zhi.Dh() + zf.Zfs.Js(true)

		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
