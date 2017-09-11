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

func scinitfun(buffer *bytes.Buffer) {
	//func init()
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Init(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//ormdebug()
	buffer.WriteString(zf.Zfs.Ormdebug(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//ormermoxing()
	buffer.WriteString(zf.Zfs.Ormermoxing(true) + zfzhi.Zhi.Xkhz() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//ormershujuku()
	buffer.WriteString(zf.Zfs.Ormershujuku(true) + zfzhi.Zhi.Xkhz() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func scormershujuku(buffer *bytes.Buffer) {
	//func ormershujuku()
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ormershujuku(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//url:=
	urstr := zf.Zfs.Url(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(urstr)
	//Shujukus[zf.Zfs.Yonghu(false)].Zhi+
	yhstr := zf.Zfs.Shujukus(false) + zfzhi.Zhi.Zkhz() + zh.Zhs.Zfszhfalse(zf.Zfs.Yonghu(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(yhstr)
	//zfzhi.Zhi.Mh() +
	mhstr := zh.Zhs.Zhiszh(zf.Zfs.Mh(false)) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(mhstr)
	//Shujukus[zf.Zfs.Mima(false)].Zhi +
	mmstr := zf.Zfs.Shujukus(false) + zfzhi.Zhi.Zkhz() + zh.Zhs.Zfszhfalse(zf.Zfs.Mima(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(mmstr)
	//zfzhi.Zhi.Qa() +
	qastr := zh.Zhs.Zhiszh(zf.Zfs.Qa(false)) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(qastr)
	//zf.Zfs.Tcp(true) +
	tcpstr := zh.Zhs.Zfszhtrue(zf.Zfs.Tcp(false)) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(tcpstr)
	//zfzhi.Zhi.Xkhz() +
	xkhzstr := zh.Zhs.Zhiszh(zf.Zfs.Xkhz(false)) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(xkhzstr)
	//Shujukus[zf.Zfs.Ip(false)].Zhi +
	ipstr := zf.Zfs.Shujukus(false) + zfzhi.Zhi.Zkhz() + zh.Zhs.Zfszhfalse(zf.Zfs.Ip(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ipstr)
	//zfzhi.Zhi.Mh() +
	mhsstr := zh.Zhs.Zhiszh(zf.Zfs.Mh(false)) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(mhsstr)
	//Shujukus[zf.Zfs.Duankou(false)].Zhi +
	dkstr := zf.Zfs.Shujukus(false) + zfzhi.Zhi.Zkhz() + zh.Zhs.Zfszhfalse(zf.Zfs.Duankou(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dkstr)
	//zfzhi.Zhi.Xkhy() +
	xkhystr := zh.Zhs.Zhiszh(zf.Zfs.Xkhy(false)) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(xkhystr)
	//zfzhi.Zhi.Xx() +
	xxstr := zh.Zhs.Zhiszh(zf.Zfs.Xx(false)) + zfzhi.Zhi.Jia() + zfzhi.Zhi.Hhf()
	buffer.WriteString(xxstr)
	//Shujukus[zf.Zfs.Mingcheng(false)].Zhi
	mcstr := zf.Zfs.Shujukus(false) + zfzhi.Zhi.Zkhz() + zh.Zhs.Zfszhfalse(zf.Zfs.Mingcheng(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(mcstr)

	//orm.RegisterDataBase
	regstr := zf.Zfs.Orm(true) + zfzhi.Zhi.Dh() + zf.Zfs.RegisterDataBase(false)
	buffer.WriteString(regstr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	//zf.Zfs.Default(true),
	destr := zh.Zhs.Zfszhtrue(zf.Zfs.Default(false)) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(destr)
	//Shujukus[zf.Zfs.Qudong(false)].Zhi,
	qdstr := zf.Zfs.Shujukus(false) + zfzhi.Zhi.Zkhz() + zh.Zhs.Zfszhfalse(zf.Zfs.Qudong(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(qdstr)
	//url,
	urlstrs := zf.Zfs.Url(true) + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
	buffer.WriteString(urlstrs)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func scormdebug(buffer *bytes.Buffer) {
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ormdebug(true) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	buffer.WriteString(zf.Zfs.Setormerdebug(false))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	//Chushihuas[zf.Zfs.Ormdebug(false)].Zhi,
	//TODO 这里要怎么控制一下读取Chushihuas的数据
	csh := zf.Zfs.Chushihuas(false) + zfzhi.Zhi.Zkhz() +
		zh.Zhs.Zfszhfalse(zf.Zfs.Ormdebug(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dou()
	buffer.WriteString(csh)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func scormermoxing(buffer *bytes.Buffer) {
	mks := gongju.Mokuaimings
	for _, mkvo := range mks {
		mkv := mkvo.Zhi
		//func ormermoxing()
		funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ormermoxing(true) +
			zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
		buffer.WriteString(funstr)
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//orm.RegisterModel
		regstr := zf.Zfs.Orm(true) + zfzhi.Zhi.Dh() + zf.Zfs.RegisterModel(false)
		buffer.WriteString(regstr)
		buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
		_, biaos, _ := gongju.Biaolies(mkv)
		for biao, _ := range biaos {
			newstr := zf.Zfs.New(true) + zfzhi.Zhi.Xkhz() + zf.Zfs.Moxings(true) +
				zfzhi.Zhi.Dh() + biao + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(newstr)
		}
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
		buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	}
}

func scsetormerdebug(buffer *bytes.Buffer) {
	//func Setormerdebug(boolstr string)
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Setormerdebug(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Boolstr(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.String(true) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//orm.Debug,_=strconv.ParseBool(boolstr)
	destr := zf.Zfs.Orm(true) + zfzhi.Zhi.Dh() + zf.Zfs.Debug(false) +
		zfzhi.Zhi.Dou() + zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Strconv(true) + zfzhi.Zhi.Dh() + zf.Zfs.ParseBool(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Boolstr(true) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(destr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func scdefaultormer(buffer *bytes.Buffer) {
	//func Defaultormer() orm.Ormer
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Defaultormer(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zf.Zfs.Orm(true) + zfzhi.Zhi.Dh() + zf.Zfs.Ormer(false)
	buffer.WriteString(funstr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//return Ormerbyname(zf.Zfs.Default(true))
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ormerbyname(false) +
		zfzhi.Zhi.Xkhz() + zh.Zhs.Zfszhtrue(zf.Zfs.Default(false)) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(retstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}
func scormerbyname(buffer *bytes.Buffer) {
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ormerbyname(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Name(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.String(true) + zfzhi.Zhi.Xkhy() + zf.Zfs.Orm(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Ormer(false)
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//ret:=orm.NewOrm()
	rstr := zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Orm(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewOrm(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(rstr)
	//ret.Using(name)
	useret := zf.Zfs.Ret(true) + zfzhi.Zhi.Dh() + zf.Zfs.Using(false) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Name(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(useret)
	//return ret
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ret(true)
	buffer.WriteString(retstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func Shengchengchushihuaormer() {
	buffer := &bytes.Buffer{}
	//package chushihuas
	pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Hhf()
	buffer.WriteString(pac)
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	//_"github.com/go-sql-driver/mysql"
	buffer.WriteString(zh.Zhs.Mysqlbao())
	//"changliang/zf"
	zfbao := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfbao)
	//"changliang/zfzhi"
	zfzhibao := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zfzhibao)
	//"strconv"
	constr := zfzhi.Zhi.Syh() + zf.Zfs.Strconv(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(constr)

	//"github.com/astaxie/beego/orm"
	buffer.WriteString(zh.Zhs.Beegoormbao())
	//"xxx/moxings"
	mxstr := zfzhi.Zhi.Syh() + gongju.Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi + zfzhi.Zhi.Xx() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Syh()
	buffer.WriteString(mxstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	scinitfun(buffer)
	scormdebug(buffer)
	scormermoxing(buffer)
	scsetormerdebug(buffer)
	scdefaultormer(buffer)
	scormerbyname(buffer)
	scormershujuku(buffer)
	dir := gongju.Getgopath() + zfzhi.Zhi.Xx() +
		gongju.Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi + zfzhi.Zhi.Xx() + zf.Zfs.Chushihuas(true)
	path := dir + zfzhi.Zhi.Xx() + zf.Zfs.Chushihua(false) + zf.Zfs.Ormer(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
	os.MkdirAll(dir, os.ModePerm)
	ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
}
