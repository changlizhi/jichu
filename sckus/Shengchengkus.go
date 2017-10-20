package sckus

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

func imports(mokuai string, buffer *bytes.Buffer) {
	// \n import (\n
	importstr := zfzhi.Zhi.Hhf() + zf.Zfs.Import(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(importstr)

	// "xxx/moxings" \n
	apm := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Moxings(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(apm)

	//"log"
	lstr := zfzhi.Zhi.Syh() + zf.Zfs.Log(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lstr)

	// "xxx/chushihuas"
	bai := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bai)

	//"changliang/zf"
	zstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(zstr)
	//"changliang/zfzhi"
	zzstr := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Syh()
	buffer.WriteString(zzstr)

	//\n ) \n
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
}
func chaxunyige(bianma string, buffer *bytes.Buffer) {
	// func Chaxunyige (id int)
	idx := zf.Zfs.Id(true)
	idd := zf.Zfs.Id(false)

	bmx := strings.ToLower(bianma)

	//func Chaxunyige(id int)
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chaxunyige(false) + zf.Zfs.Kus(true) +
		zfzhi.Zhi.Xkhz() + idx + zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) + zfzhi.Zhi.Xkhy()
	buffer.WriteString(funstr)

	// *moxings.Juese{ \n
	fanhui := zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Moxings(true) +
		zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(fanhui)

	// juese := &moxings.Juese
	modelstr := bmx + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zfzhi.Zhi.Kgf() + zfzhi.Zhi.Qh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma
	buffer.WriteString(modelstr)

	// {Id:id}
	idstr := zfzhi.Zhi.Dkhz() + idd + zfzhi.Zhi.Mh() + idx + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(idstr)

	// \n err := appinits.Defaultormer().Read(juese)
	read := zf.Zfs.Err(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Dh() + zf.Zfs.Defaultormer(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Dh() + zf.Zfs.Read(false) + zfzhi.Zhi.Xkhz() +
		bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(read)

	// if err := nil
	errstrss := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Kgf() + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(errstrss)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//log.Println(err)
	lgstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Err(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lgstr)

	//return nil
	retstrsss := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Nil(true)
	buffer.WriteString(retstrsss)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	// \n return juese
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + bmx + zfzhi.Zhi.Hhf()
	buffer.WriteString(retstr)

	//} \n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func tianjiayige(bianma string, buffer *bytes.Buffer) {
	// func Tianjiayige

	bmx := strings.ToLower(bianma)

	funcstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tianjiayige(false) + zf.Zfs.Kus(true)
	buffer.WriteString(funcstr)

	//(juese*moxings.Juese){\n
	canshu := zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() +
		zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma +
		zfzhi.Zhi.Xkhy() + zf.Zfs.String(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(canshu)

	// _, err := appinits.Defaultormer().Insert(juese)
	insert := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Defaultormer(false) + zfzhi.Zhi.Xkhz() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Insert(false) +
		zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(insert)

	// if err != nil
	iferr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(iferr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//log.Println(err)
	lgstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Err(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lgstr)

	// return chushihuas.Tishis[zf.Zfs.Tishi04].Bianma
	errret := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi04(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(errret)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	// return chushihuas.Tishis[zf.Zfs.Tishi03].Bianma
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi03(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(retstr)
	//\n}\n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func tianjiaduoge(bianma string, buffer *bytes.Buffer) {
	//func Tianjiaduoge(
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Tianjiaduoge(false) + zf.Zfs.Kus(true) + zfzhi.Zhi.Xkhz()
	buffer.WriteString(funstr)

	//jueseshuzu
	canshu := strings.ToLower(bianma) + zf.Zfs.Shuzu(true)
	buffer.WriteString(canshu)

	//[]*moxings.Juese{\n
	cslx := zfzhi.Zhi.Kgf() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xh() +
		zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma + zfzhi.Zhi.Xkhy() +
		zf.Zfs.String(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cslx)

	//_err:=appinits.Defaultormer().InsertMulti
	insert := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Defaultormer(false) + zfzhi.Zhi.Xkhz() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dh() + zf.Zfs.InsertMulti(false)
	buffer.WriteString(insert)

	//(len(jueseshuzu+,jueseshuzu+
	lenstr := zfzhi.Zhi.Xkhz() + zf.Zfs.Len(true) + zfzhi.Zhi.Xkhz() +
		canshu + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() +
		canshu + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lenstr)
	// if err != nil
	iferr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(iferr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//log.Println(err)
	lgstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Err(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lgstr)

	// return chushihuas.Tishis[zf.Zfs.Tishi04].Bianma
	errret := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi04(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(errret)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	// return chushihuas.Tishis[zf.Zfs.Tishi03].Bianma
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi03(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(retstr)

	//}\n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func shanchuyige(buffer *bytes.Buffer) {
	//func Shanchuyige(id int+{ \n
	funcstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Shanchuyige(false) + zf.Zfs.Kus(true) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Id(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.Int(true) + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(funcstr)

	//_,err:=chushihuas.Defaultormer().Delete
	delete := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Defaultormer(false) + zfzhi.Zhi.Xkhz() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Delete(false)
	buffer.WriteString(delete)

	//(Chaxunyige(id++
	cxstr := zfzhi.Zhi.Xkhz() + zf.Zfs.Chaxunyige(false) + zf.Zfs.Kus(true) + zfzhi.Zhi.Xkhz() +
		zf.Zfs.Id(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(cxstr)
	// if err != nil
	iferr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(iferr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//log.Println(err)
	lgstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Err(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lgstr)

	// return chushihuas.Tishis[zf.Zfs.Tishi04].Bianma
	errret := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi08(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(errret)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	// return chushihuas.Tishis[zf.Zfs.Tishi03].Bianma
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi07(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(retstr)

	//\n}\n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}
func xiugaiyige(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)

	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Xiugaiyige(false) + zf.Zfs.Kus(true)
	buffer.WriteString(funstr)

	//(juese*moxings.Juese){\n
	csstr := zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Xh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() +
		bianma + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(csstr)

	//_,err := chushihuas.Defaultormer().Update(juese)
	update := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Defaultormer(false) + zfzhi.Zhi.Xkhz() +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Update(false) +
		zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(update)
	// if err != nil
	iferr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)

	buffer.WriteString(iferr)

	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//log.Println(err)
	lgstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Err(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lgstr)

	// return chushihuas.Tishis[zf.Zfs.Tishi04].Bianma
	errret := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi06(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(errret)
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	// return chushihuas.Tishis[zf.Zfs.Tishi03].Bianma
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Dh() + zf.Zfs.Zfs(false) + zfzhi.Zhi.Dh() + zf.Zfs.Tishi05(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.False(true) + zfzhi.Zhi.Xkhy() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Hhf()
	buffer.WriteString(retstr)

	//\n}\n
	buffer.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

}

func chaxunquanbu(mokuai string, biao string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(biao)
	//func Chaxunquanbukus(dtziyuan *moxings.Dtziyuan) []*moxings.Juese
	fustr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chaxunquanbu(false) + zf.Zfs.Kus(true) +
		zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Moxings(true) +
		zfzhi.Zhi.Dh() + biao + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Xh() + zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + biao
	buffer.WriteString(fustr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//ret := []*moxings.Juese{}
	restr := zf.Zfs.Ret(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xh() +
		zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() +
		biao + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(restr)
	//queryseter := chushihuas.Defaultormer().QueryTable(zf.Zfs.Juese(true))
	qsstr := zf.Zfs.QuerySeter(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Dh() + zf.Zfs.Defaultormer(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dh() +
		zf.Zfs.QueryTable(false) + zfzhi.Zhi.Xkhz() + zh.Zhs.Zfszhtrue(biao) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(qsstr)
	//queryseter = queryseter.Filter(zf.Zfs.Youxiaoxing(true), zfzhi.Zhi.Shuzi1w())
	qqfy := zf.Zfs.QuerySeter(true) + zfzhi.Zhi.Dyh() + zf.Zfs.QuerySeter(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Filter(false) + zfzhi.Zhi.Xkhz() +
		zh.Zhs.Zfszhtrue(zf.Zfs.Youxiaoxing(false)) + zfzhi.Zhi.Dou() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi1w(false)) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(qqfy)

	// queryseter = queryseter.Limit(zfzhi.Zhi.Shuzifu1())
	listr := zf.Zfs.QuerySeter(true) + zfzhi.Zhi.Dyh() + zf.Zfs.QuerySeter(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Limit(false) + zfzhi.Zhi.Xkhz() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzifu1(false)) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(listr)
	//queryseter = queryseter.OrderBy(zf.Zfs.Id(true))
	obstr := zf.Zfs.QuerySeter(true) + zfzhi.Zhi.Dyh() + zf.Zfs.QuerySeter(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.OrderBy(false) + zfzhi.Zhi.Xkhz() +
		zh.Zhs.Zfszhtrue(zf.Zfs.Id(false)) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(obstr)
	//houzhui := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Xhx() + zf.Zfs.I(true) + zf.Zfs.Exact(true)
	hzstr := zf.Zfs.Houzhui(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zh.Zhs.Zhiszh(zf.Zfs.Xhx(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zhiszh(zf.Zfs.Xhx(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhtrue(zf.Zfs.I(false)) + zfzhi.Zhi.Jia() +
		zh.Zhs.Zfszhtrue(zf.Zfs.Exact(false)) + zfzhi.Zhi.Hhf()
	buffer.WriteString(hzstr)

	for _, lk := range gongju.Fanshebiao(mokuai, biao) {
		if gongju.Lieleixing(lk) == zf.Zfs.String(true) {
			//if dtziyuan.lk != zfzhi.Zhi.Kzf()
			ifstr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + bmx + zfzhi.Zhi.Dh() + lk + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zh.Zhs.Zhiszh(zf.Zfs.Kzf(false))
			buffer.WriteString(ifstr)
			buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
			//queryseter = queryseter.Filter(zf.Zfs.lk(true) + houzhui, dtziyuan.Mingcheng)
			qsstr := zf.Zfs.QuerySeter(true) + zfzhi.Zhi.Dyh() + zf.Zfs.QuerySeter(true) + zfzhi.Zhi.Dh() +
				zf.Zfs.Filter(false) + zfzhi.Zhi.Xkhz() + zh.Zhs.Zfszhtrue(lk) +
				zfzhi.Zhi.Jia() + zf.Zfs.Houzhui(true) + zfzhi.Zhi.Dou() + bmx + zfzhi.Zhi.Dh() + lk +
				zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(qsstr)
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
		}
	}

	//_,err := queryseter.All(&ret)
	alstr := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Dyh() + zf.Zfs.QuerySeter(true) + zfzhi.Zhi.Dh() + zf.Zfs.All(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Qh() + zf.Zfs.Ret(true) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(alstr)
	//if err != nil
	iestr := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(iestr)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//log.Println(err)
	lpstr := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.Err(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lpstr)
	//return nil
	rnstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Nil(true)
	buffer.WriteString(rnstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//return ret
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ret(true)
	buffer.WriteString(retstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func shanchutiaojiankus(bianma string, buffer *bytes.Buffer) {
	bmx := strings.ToLower(bianma)
	//func Shanchutiaojiankus
	funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Shanchu(false) +
		zf.Zfs.Tiaojian(true) + zf.Zfs.Kus(true)
	buffer.WriteString(funstr)
	//(dtziyuan *moxings.Dtziyuan) string
	funcs := zfzhi.Zhi.Xkhz() + bmx + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() +
		zf.Zfs.Moxings(true) + zfzhi.Zhi.Dh() + bianma +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf() + zf.Zfs.String(true)
	buffer.WriteString(funcs)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//dtziyuans := Chaxunquanbukus(dtziyuan)
	dcs := bmx + zf.Zfs.S(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() +
		zf.Zfs.Chaxunquanbu(false) + zf.Zfs.Kus(true) + zfzhi.Zhi.Xkhz() +
		bmx + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(dcs)
	//_, err := chushihuas.Defaultormer().Delete(dtziyuans[zfzhi.Zhi.Shuzi0()])
	ecd := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Dou() + zf.Zfs.Err(true) +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Defaultormer(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dh() +
		zf.Zfs.Delete(false) + zfzhi.Zhi.Xkhz() + bmx + zf.Zfs.S(true) + zfzhi.Zhi.Zkhz() +
		zh.Zhs.Zhiszh(zf.Zfs.Shuzi0(false)) + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(ecd)

	//if err != nil
	ien := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Err(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Nil(true)
	buffer.WriteString(ien)
	buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//log.Println(err)
	lpe := zf.Zfs.Log(true) + zfzhi.Zhi.Dh() + zf.Zfs.Println(false) +
		zfzhi.Zhi.Xkhz() + zf.Zfs.Err(true) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(lpe)
	//return chushihuas.Tishis[zf.Zfs.Tishi08(false)].Bianma
	rct8 := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zh.Zhs.Zfszhfalse(zf.Zfs.Tishi08(false)) + zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false)
	buffer.WriteString(rct8)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//return chushihuas.Tishis[zf.Zfs.Tishi07(false)].Bianma
	rct7 := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) +
		zfzhi.Zhi.Dh() + zf.Zfs.Tishis(false) + zfzhi.Zhi.Zkhz() +
		zh.Zhs.Zfszhfalse(zf.Zfs.Tishi07(false)) + zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Dh() + zf.Zfs.Bianma(false)
	buffer.WriteString(rct7)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func Shengchengkus() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Fanshebiaolies(mkv)
		for biao, _ := range biaos {
			buffer := bytes.Buffer{}
			buffer.WriteString(zf.Zfs.Package(true)) //package
			buffer.WriteString(zfzhi.Zhi.Konggefu()) // zfzhi.Zhi.Kgf()
			lujing := zf.Zfs.Zd(true) + strings.ToLower(biao) + zf.Zfs.Kus(true)
			buffer.WriteString(lujing) // zdjuesedaos

			imports(mkv, &buffer)             //import all
			chaxunyige(biao, &buffer)         // Chaxunyige
			tianjiayige(biao, &buffer)        // Tianjiayige
			tianjiaduoge(biao, &buffer)       // Tianjiaduoge
			xiugaiyige(biao, &buffer)         // Xiugaiyige
			shanchuyige(&buffer)              // Shanchuyige
			chaxunquanbu(mkv, biao, &buffer)  // Shanchuyige
			shanchutiaojiankus(biao, &buffer) // Shanchuyige

			dir := gongju.Getgopath() + zfzhi.Zhi.Xx() +
				mkv + zfzhi.Zhi.Xx() + lujing
			path := dir + zfzhi.Zhi.Xx() + biao + zf.Zfs.Kus(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}
	}

}
