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

	// "xxx/chushihuas"
	bai := zfzhi.Zhi.Syh() + mokuai + zfzhi.Zhi.Xx() +
		zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(bai)

	//"changliang/zf"
	zfbao := zfzhi.Zhi.Syh() + zf.Zfs.Changliang(true) + zfzhi.Zhi.Xx() +
		zf.Zfs.Zf(true) + zfzhi.Zhi.Syh()
	buffer.WriteString(zfbao)

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

func chaxunquanbu(biao string, buffer *bytes.Buffer) {
	//func Chaxunquanbu() []*moxings.Juese
	fustr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chaxunquanbu(false) + zf.Zfs.Kus(true) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() +
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
	//queryseter.All(&ret)
	alstr := zf.Zfs.QuerySeter(true) + zfzhi.Zhi.Dh() + zf.Zfs.All(false) +
		zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Qh() + zf.Zfs.Ret(true) +
		zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
	buffer.WriteString(alstr)
	//return ret
	retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Ret(true)
	buffer.WriteString(retstr)
	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func Shengchengkus() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Biaolies(mkv)
		for biao, _ := range biaos {
			buffer := bytes.Buffer{}
			buffer.WriteString(zf.Zfs.Package(true)) //package
			buffer.WriteString(zfzhi.Zhi.Konggefu()) // zfzhi.Zhi.Kgf()
			lujing := zf.Zfs.Zd(true) + strings.ToLower(biao) + zf.Zfs.Kus(true)
			buffer.WriteString(lujing) // zdjuesedaos

			imports(mkv, &buffer)       //import all
			chaxunyige(biao, &buffer)   // Chaxunyige
			tianjiayige(biao, &buffer)  // Tianjiayige
			tianjiaduoge(biao, &buffer) // Tianjiaduoge
			xiugaiyige(biao, &buffer)   // Xiugaiyige
			shanchuyige(&buffer)  // Shanchuyige
			chaxunquanbu(biao, &buffer) // Shanchuyige

			dir := gongju.Getgopath() + zfzhi.Zhi.Xx() +
				mkv + zfzhi.Zhi.Xx() + lujing
			path := dir + zfzhi.Zhi.Xx() + biao + zf.Zfs.Kus(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			os.MkdirAll(dir, os.ModePerm)
			ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
		}
	}

}
