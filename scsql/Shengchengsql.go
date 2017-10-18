package scsql

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"io/ioutil"
	"os"
	"strconv"
)

func shenginsertsql() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		zwpath := gongju.Getjichupath() + zfzhi.Zhi.Xx() + zf.Zfs.Sc(true) +
			zf.Zfs.Sql(true) + zfzhi.Zhi.Xx() + zf.Zfs.Insert(false) +
			zf.Zfs.Ziyuan(true) + zfzhi.Zhi.Dh() + zf.Zfs.Goconf(true)
		zwbs, _ := ioutil.ReadFile(zwpath)

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Sql(true)
		sczwpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Insert(false) +
			zf.Zfs.Ziyuan(true) + zfzhi.Zhi.Dh() + zf.Zfs.Sql(true)

		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(sczwpath, zwbs, os.ModePerm)
	}
}

func Shengchengsql() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Fanshebiaolies(mkv)
		buffer := &bytes.Buffer{}
		for b, _ := range biaos {
			//CREATE TABLE Shijian (
			crestr := zf.Zfs.Create(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Table(true) +
				zfzhi.Zhi.Kgf() + zfzhi.Zhi.Yzb() + b + zfzhi.Zhi.Yzb()
			buffer.WriteString(crestr)
			buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
			lies := gongju.Fanshebiao(mkv, b)
			for _, l := range lies {
				cd := gongju.Liechangdu(l)
				cdzw := strconv.Itoa(cd)
				lx := gongju.Lieleixing(l)
				if l == zf.Zfs.Id(false) {
					//Id int(10) auto_increment comment '主键',
					vc := zfzhi.Zhi.Yzb() + l + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) +
						zfzhi.Zhi.Xkhz() + cdzw + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf() +
						zf.Zfs.Auto(true) + zfzhi.Zhi.Xhx() + zf.Zfs.Increment(true) +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(l) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vc)
				}
				if l != zf.Zfs.Id(false) && lx == zf.Zfs.Int(true) {
					//paixu int(10) comment '排序',
					vc := zfzhi.Zhi.Yzb() + l + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) +
						zfzhi.Zhi.Xkhz() + cdzw + zfzhi.Zhi.Xkhy() +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(l) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vc)
				}
				if lx == zf.Zfs.String(true) {
					//Mingcheng varchar(250) comment '名称',
					vc := zfzhi.Zhi.Yzb() + l + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Varchar(true) +
						zfzhi.Zhi.Xkhz() + cdzw + zfzhi.Zhi.Xkhy() +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(l) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vc)
				}
				if lx == zf.Zfs.Time(true) {
					//Chuangjianriqi timestamp  comment '创建日期',
					vc := zfzhi.Zhi.Yzb() + l + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Timestamp(true) +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(l) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vc)
				}
			}
			//	PRIMARY KEY (`id`),
			pkstr := zf.Zfs.Primary(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Key(true) +
				zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Yzb() + zf.Zfs.Id(false) + zfzhi.Zhi.Yzb() +
				zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(pkstr)
			//	UNIQUE INDEX `Paixu` (`Paixu`),
			uip := zf.Zfs.Unique(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Index(true) +
				zfzhi.Zhi.Kgf() + zfzhi.Zhi.Yzb() + zf.Zfs.Paixu(false) + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Yzb() +
				zf.Zfs.Paixu(false) + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(uip)
			//	UNIQUE INDEX `Bianma` (`Bianma`)
			uib := zf.Zfs.Unique(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Index(true) +
				zfzhi.Zhi.Kgf() + zfzhi.Zhi.Yzb() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Yzb() +
				zf.Zfs.Bianma(false) + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(uib)

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
			//ENGINE=InnoDB
			ei := zf.Zfs.Engine(true) + zfzhi.Zhi.Dyh() + zf.Zfs.InnoDB(true) + zfzhi.Zhi.Fh() + zfzhi.Zhi.Hhf()
			buffer.WriteString(ei)
		}
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Sql(true)
		scpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Chuangjianbiao(false) + zfzhi.Zhi.Dh() + zf.Zfs.Sql(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(scpath, buffer.Bytes(), os.ModePerm)
	}
	shenginsertsql()
}
