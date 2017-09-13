package scchushihuas

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
	"gongju"
	"io/ioutil"
	"os"
)

func Shengchengchushihuatest() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		buffer := &bytes.Buffer{}
		pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tests(true) + zfzhi.Zhi.Hhf()
		buffer.WriteString(pac)
		buffer.WriteString(zf.Zfs.Import(true))
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
		//"github.com/astaxie/beego/orm"
		ormstr := zh.Zhs.Beegoormbao() + zfzhi.Zhi.Hhf()
		buffer.WriteString(ormstr)
		//"xxx/chushihuas"
		cshstr := zfzhi.Zhi.Syh() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(cshstr)

		//"log"
		lstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(lstr)
		//"testing"
		tstr := zfzhi.Zhi.Syh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(tstr)
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

		//func TestDayin
		funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Test(false) + zf.Zfs.Dayin(false)
		buffer.WriteString(funstr)

		//(t *testing.T)
		csstr := zfzhi.Zhi.Xkhz() + zf.Zfs.T(true) + zfzhi.Zhi.Kgf() +
			zfzhi.Zhi.Xh() + zf.Zfs.Testing(true) + zfzhi.Zhi.Dh() +
			zf.Zfs.T(false) + zfzhi.Zhi.Xkhy()
		buffer.WriteString(csstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

		//log.Println("chushihuas.Defaultormer()------------ ", chushihuas.Defaultormer())
		dstr := zh.Zhs.Logszh(zf.Zfs.Chushihuas(true)+
			zfzhi.Zhi.Dh()+
			zf.Zfs.Defaultormer(false)+
			zfzhi.Zhi.Xkhz()+
			zfzhi.Zhi.Xkhy()) +
			zfzhi.Zhi.Hhf()
		buffer.WriteString(dstr)
		//log.Println("chushihuas.Chushihuas------------", chushihuas.Chushihuas)
		cstr := zh.Zhs.Logszh(zf.Zfs.Chushihuas(true)+
			zfzhi.Zhi.Dh()+
			zf.Zfs.Chushihuas(false)) +
			zfzhi.Zhi.Hhf()
		buffer.WriteString(cstr)
		//log.Println("chushihuas.Shujukus------------", chushihuas.Shujukus)
		sstr := zh.Zhs.Logszh(zf.Zfs.Chushihuas(true)+
			zfzhi.Zhi.Dh()+
			zf.Zfs.Shujukus(false)) +
			zfzhi.Zhi.Hhf()
		buffer.WriteString(sstr)
		//log.Println("chushihuas.Tishis------------", chushihuas.Tishis)
		tsstr := zh.Zhs.Logszh(zf.Zfs.Chushihuas(true)+
			zfzhi.Zhi.Dh()+
			zf.Zfs.Tishis(false)) +
			zfzhi.Zhi.Hhf()
		buffer.WriteString(tsstr)
		//log.Println("chushihuas.Cuowus------------", chushihuas.Cuowus)
		cwstr := zh.Zhs.Logszh(zf.Zfs.Chushihuas(true)+
			zfzhi.Zhi.Dh()+
			zf.Zfs.Cuowus(false)) +
			zfzhi.Zhi.Hhf()
		buffer.WriteString(cwstr)
		//log.Println("orm.Debug------------", orm.Debug)
		dbstr := zh.Zhs.Logszh(zf.Zfs.Orm(true)+
			zfzhi.Zhi.Dh()+
			zf.Zfs.Debug(false)) +
			zfzhi.Zhi.Hhf()
		buffer.WriteString(dbstr)
		//log.Println("chushihuas.Getapppath():====",chushihuas.Getapppath())
		apppathstr := zh.Zhs.Logszh(zf.Zfs.Chushihuas(true)+
			zfzhi.Zhi.Dh()+
			zf.Zfs.Getapppath(false)+zfzhi.Zhi.Xkhz()+zfzhi.Zhi.Xkhy()) +
			zfzhi.Zhi.Hhf()
		buffer.WriteString(apppathstr)

		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv + zfzhi.Zhi.Xx() + zf.Zfs.Tests(true)
		path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Chushihua(false) +
			zfzhi.Zhi.Xhx() + zf.Zfs.Test(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}
}
