package scsql

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"gongju/sjkmh"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Shengchengchuangjian() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		buffer := &bytes.Buffer{}
		mkv := mks[mkvo].Zhi

		//craete database if not exists `mkv`;
		cdm := zf.Zfs.Create(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Database(true) + zfzhi.Zhi.Kgf() +
			zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Not(true) + zfzhi.Zhi.Kgf() +
			zf.Zfs.Exists(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Yzb() +
			mkv + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Fh() + zfzhi.Zhi.Hhf()
		buffer.WriteString(cdm)
		bjg := gongju.Fanshebiaos(mkv)

		for _, b := range bjg {
			//CREATE TABLE `mkv`.`bshuju` (
			crestr := zf.Zfs.Create(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Table(true) +
				zfzhi.Zhi.Kgf() + zfzhi.Zhi.Yzb() + mkv + zfzhi.Zhi.Yzb() +
				zfzhi.Zhi.Dh() + zfzhi.Zhi.Yzb() + b + zfzhi.Zhi.Yzb()
			buffer.WriteString(crestr)
			buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
			idcd := gongju.Liechangdu(zf.Zfs.Id(false))
			idcdzw := strconv.Itoa(idcd)
			//Id int(10) auto_increment comment '主键',
			vc := zfzhi.Zhi.Yzb() + zf.Zfs.Id(false) + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) +
				zfzhi.Zhi.Xkhz() + idcdzw + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf() +
				zf.Zfs.Auto(true) + zfzhi.Zhi.Xhx() + zf.Zfs.Increment(true) +
				zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
				zfzhi.Zhi.Dyhe() + gongju.Zhongwen(zf.Zfs.Id(false)) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(vc)
			if strings.Contains(b, zf.Zfs.Jiegou(true)) {
				for _, sjzd := range gongju.Fanshejichubiao() {
					if zf.Zfs.Id(false) != sjzd {
						cd := gongju.Liechangdu(sjzd)
						cdzf := strconv.Itoa(cd)
						vc := zfzhi.Zhi.Yzb() + sjzd + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Varchar(true) +
							zfzhi.Zhi.Xkhz() + cdzf + zfzhi.Zhi.Xkhy() +
							zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
							zfzhi.Zhi.Dyhe() + gongju.Zhongwen(sjzd) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
						buffer.WriteString(vc)
					}
				}
			}
			if strings.Contains(b, zf.Zfs.Shuju(true)) {
				for _, sjzd := range gongju.Fanshejichushuju() {
					if zf.Zfs.Id(false) != sjzd {
						cd := gongju.Liechangdu(sjzd)
						cdzf := strconv.Itoa(cd)
						vcbiao := zfzhi.Zhi.Yzb() + sjzd + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Varchar(true) +
							zfzhi.Zhi.Xkhz() + cdzf + zfzhi.Zhi.Xkhy() +
							zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
							zfzhi.Zhi.Dyhe() + gongju.Zhongwen(sjzd) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
						buffer.WriteString(vcbiao)
					}
				}
			}
			//primary key (`id`)
			pkstr := zf.Zfs.Primary(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Key(true) +
				zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Yzb() + zf.Zfs.Id(false) + zfzhi.Zhi.Yzb() +
				zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(pkstr)
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
			//COMMENT='时间'
			ct := zf.Zfs.Comment(true) + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyhe() +
				gongju.Zhongwen(b) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Hhf()
			buffer.WriteString(ct)
			//COLLATE='utf8_general_ci'
			cug := zf.Zfs.Collate(true) + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyhe() +
				zf.Zfs.Utf8(true) + zfzhi.Zhi.Xhx() + zf.Zfs.General(true) +
				zfzhi.Zhi.Xhx() + zf.Zfs.Ci(true) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Hhf()
			buffer.WriteString(cug)
			//ENGINE=InnoDB;
			ei := zf.Zfs.Engine(true) + zfzhi.Zhi.Dyh() + zf.Zfs.InnoDB(true) +
				zfzhi.Zhi.Fh() + zfzhi.Zhi.Hhf() + zfzhi.Zhi.Hhf()
			buffer.WriteString(ei)
		}
		//log.Println(bf.String())
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Sql(true)
		scpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Chuangjianbiao(false) + zf.Zfs.Xin(true) + zfzhi.Zhi.Dh() + zf.Zfs.Sql(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(scpath, buffer.Bytes(), os.ModePerm)
	}
	Shengchengcharu()
}

func Shengchengcharu() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		buffer := &bytes.Buffer{}
		mkv := mks[mkvo].Zhi
		bjg := gongju.Fanshebiaos(mkv)
		for _, b := range bjg {
			if strings.Contains(b, zf.Zfs.Jiegou(true)) {

				//insert into `mhsydata`.`juesejiegou`
				instr := zf.Zfs.Insert(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Into(true) + zfzhi.Zhi.Kgf() +
					zfzhi.Zhi.Yzb() + mkv + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Dh() + zfzhi.Zhi.Yzb() + b + zfzhi.Zhi.Yzb()
				buffer.WriteString(instr)

				jczds := gongju.Fanshejichubiao()

				buffer.WriteString(zfzhi.Zhi.Xkhz())
				buffercol := &bytes.Buffer{}
				for _, jczd := range jczds {
					buffercol.WriteString(zfzhi.Zhi.Yzb() + jczd + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Dou())
				}
				buffercolstrneed := buffercol.String()[zfzhi.Zhi.Shuzi0() : len(buffercol.String())-zfzhi.Zhi.Shuzi1()]
				buffer.WriteString(buffercolstrneed)
				buffer.WriteString(zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())

				buffer.WriteString(zf.Zfs.Values(true) + zfzhi.Zhi.Hhf())
				bufferval := &bytes.Buffer{}
				for jcindex, jcval := range sjkmh.Biaozhunjiegous {
					bufferval.WriteString(zfzhi.Zhi.Xkhz())
					bufferval.WriteString(strconv.Itoa(jcindex+zfzhi.Zhi.Shuzi1()) + zfzhi.Zhi.Dou())
					bufferval.WriteString(zfzhi.Zhi.Dyhe() + jcval + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou())
					bufferval.WriteString(zfzhi.Zhi.Dyhe() + gongju.Lieleixing(jcval) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou())
					bufferval.WriteString(zfzhi.Zhi.Dyhe() + gongju.Zhongwen(jcval) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou())
					bufferval.WriteString(zfzhi.Zhi.Dyhe() + strconv.Itoa(gongju.Liechangdu(jcval)) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou())
					bufferval.WriteString(zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dyhe())
					bufferval.WriteString(zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())
				}
				buffervalstrneed := bufferval.String()[zfzhi.Zhi.Shuzi0() : len(bufferval.String())-zfzhi.Zhi.Shuzi2()]
				buffer.WriteString(buffervalstrneed)
				buffer.WriteString(zfzhi.Zhi.Fh() + zfzhi.Zhi.Hhf())
			}
		}
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Sql(true)
		scpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Charujiegou(false) + zfzhi.Zhi.Dh() + zf.Zfs.Sql(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(scpath, buffer.Bytes(), os.ModePerm)
	}
}
