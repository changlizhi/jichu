package scchushihuas

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
	"log"
	"gongju"
)

func scinitfun(buffer *bytes.Buffer) {
	//func init()
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Init(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//ormerdebug()
	buffer.WriteString(zf.Zfs.Ormerdebug(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//ormermoxing()
	buffer.WriteString(zf.Zfs.Ormermoxing(true) + zfzhi.Zhi.Xkhz() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	//ormershujuku()
	buffer.WriteString(zf.Zfs.Ormershujuku(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func scormerdebug(buffer *bytes.Buffer) {
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ormdebug(true) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	buffer.WriteString(zf.Zfs.Setormerdebug(false))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	//Chushihuas[zf.Zfs.Ormdebug(false)].Zhi,
	//TODO 这里要怎么控制一下读取Chushihuas的数据
	csh := zf.Zfs.Chushihuas(false) + zfzhi.Zhi.Zkhz() +
		zh.Zhs.Zfszh(zf.Zfs.Ormdebug(false)) +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Dou()
	buffer.WriteString(csh)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func scormermoxing(buffer *bytes.Buffer) {
	//func ormermoxing()
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ormermoxing(true) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//ormer.RegisterModel
	regstr := zf.Zfs.Ormer(true) + zfzhi.Zhi.Dh() + zf.Zfs.RegisterModel(false)
	buffer.WriteString(regstr)
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
	_, biaos, _ := gongju.Biaolies()
	for biao, _ := range biaos {
		newstr := zf.Zfs.New(true) + zfzhi.Zhi.Xkhz() + zf.Zfs.Moxings(true) +
			zfzhi.Zhi.Dh() + biao + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		buffer.WriteString(newstr)
	}
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
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
		zfzhi.Zhi.Xkhz() + zh.Zhs.Zfszh(zf.Zfs.Default(false)) + zfzhi.Zhi.Xkhy()
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
		zf.Zfs.Orm(true) + zfzhi.Zhi.Dh() + zf.Zfs.NewOrm(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(rstr)
	//ret.Using(name)
	useret := zf.Zfs.Ret(true) + zfzhi.Zhi.Dh() +zf.Zfs.Using(false) + zfzhi.Zhi.Xkhz() +
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

	//"xxx/moxings"
	mxstr := zfzhi.Zhi.Syh() + gongju.Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi + zfzhi.Zhi.Xx() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Syh()
	buffer.WriteString(mxstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
	scinitfun(buffer)
	scormerdebug(buffer)
	scormermoxing(buffer)
	scsetormerdebug(buffer)
	scdefaultormer(buffer)
	scormerbyname(buffer)
	log.Println(buffer.String())
}