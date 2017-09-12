package sccuowus

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"io/ioutil"
	"os"
)

func Shengchengziduancuowu() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		buffer := &bytes.Buffer{}
		//package suoyoucuowus
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Suoyoucuowus(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pac)
		buffer.WriteString(zf.Zfs.Import(true))
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
		//"encoding/json"
		jstr := zfzhi.Zhi.Syh() + zf.Zfs.Encoding(true) + zfzhi.Zhi.Xx() + zf.Zfs.Json(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(jstr)
		//"time"
		tstr := zfzhi.Zhi.Syh() + zf.Zfs.Time(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(tstr)
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

		//type Ziduancuowu struct
		zdstr := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ziduancuowu(false) + zfzhi.Zhi.Kgf() + zf.Zfs.Struct(true)
		buffer.WriteString(zdstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//Shijian time.Time
		sjstr := zf.Zfs.Shijian(false) + zfzhi.Zhi.Kgf() + zf.Zfs.Time(true) + zfzhi.Zhi.Dh() + zf.Zfs.Time(false) + zfzhi.Zhi.Hhf()
		buffer.WriteString(sjstr)
		//Wenti string
		wtstr := zf.Zfs.Wenti(false) + zfzhi.Zhi.Kgf() + zf.Zfs.String(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(wtstr)
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

		//func (e Ziduancuowu)Error()string{}
		funerr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() +
			zf.Zfs.E(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ziduancuowu(false) + zfzhi.Zhi.Xkhy() +
			zfzhi.Zhi.Kgf() + zf.Zfs.Error(false) + zfzhi.Zhi.Xkhz() +
			zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf() + zf.Zfs.String(true)
		buffer.WriteString(funerr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//ret,_:=json.Marshal(e)
		rets := zf.Zfs.Ret(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Xhx() +
			zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Json(true) +
			zfzhi.Zhi.Dh() + zf.Zfs.Marshal(false) + zfzhi.Zhi.Xkhz() +
			zf.Zfs.E(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(rets)
		//return string(ret)
		retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.String(true) +
			zfzhi.Zhi.Xkhz() + zf.Zfs.Ret(true) + zfzhi.Zhi.Xkhy()
		buffer.WriteString(retstr)

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Suoyoucuowus(true)
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Ziduancuowu(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
