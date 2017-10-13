package scsql

import (
	"gongju"
	"strconv"
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"os"
	"io/ioutil"
)

func Shengchengsql() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := gongju.Biaolies(mkv)
		buffer := &bytes.Buffer{}
		for b, _ := range biaos {
			//CREATE TABLE Shijian (
			crestr := zf.Zfs.Create(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Table(true) + zfzhi.Zhi.Kgf() + b
			buffer.WriteString(crestr)
			buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())
			lies := gongju.Biao(mkv, b)
			for _, l := range lies {
				cd := gongju.Liechangdu(l)
				cdzw := strconv.Itoa(cd)
				lx := gongju.Lieleixing(l)
				if l == zf.Zfs.Id(true) {
					//Id int(10) not null auto_increment comment '主键',
					vc := l + zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) +
						zfzhi.Zhi.Xkhz() + cdzw + zfzhi.Zhi.Xkhy() +
						zfzhi.Zhi.Kgf() + zf.Zfs.Not(true) + zf.Zfs.Null(true) +
						zf.Zfs.Auto(true) + zfzhi.Zhi.Xhx() + zf.Zfs.Increment(true) +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(l) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vc)
				}
				if l != zf.Zfs.Id(true) && lx == zf.Zfs.Int(true) {
					//paixu int(10) default null comment '排序',
					vc := l + zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) +
						zfzhi.Zhi.Xkhz() + cdzw + zfzhi.Zhi.Xkhy() +
						zfzhi.Zhi.Kgf() + zf.Zfs.Default(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Null(true) +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(l) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vc)
				}
				if lx == zf.Zfs.String(true) {
					//Mingcheng varchar(250) default null comment '名称',
					vc := l + zfzhi.Zhi.Kgf() + zf.Zfs.Varchar(true) +
						zfzhi.Zhi.Xkhz() + cdzw + zfzhi.Zhi.Xkhy() +
						zfzhi.Zhi.Kgf() + zf.Zfs.Default(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Null(true) +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(l) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vc)
				}
				if lx == zf.Zfs.Time(true) {
					//Chuangjianriqi timestamp default null comment '创建日期',
					vc := l + zfzhi.Zhi.Kgf() + zf.Zfs.Timestamp(true) +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(l) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vc)
				}
			}
			//	PRIMARY KEY (id),
			pkstr := zf.Zfs.Primary(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Key(true) +
				zfzhi.Zhi.Xkhz() + zf.Zfs.Id(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(pkstr)
			//	UNIQUE INDEX Paixu (Paixu),
			uip := zf.Zfs.Unique(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Index(true) +
				zfzhi.Zhi.Kgf() + zf.Zfs.Paixu(false) + zfzhi.Zhi.Xkhz() +
				zf.Zfs.Paixu(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
			buffer.WriteString(uip)
			//	UNIQUE INDEX Bianma (Bianma)
			uib := zf.Zfs.Unique(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Index(true) +
				zfzhi.Zhi.Kgf() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Xkhz() +
				zf.Zfs.Bianma(false) + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
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
			ei := zf.Zfs.Engine(true) + zfzhi.Zhi.Dyh() + zf.Zfs.InnoDB(true) + zfzhi.Zhi.Hhf()
			buffer.WriteString(ei)

			//AUTO_INCREMENT=200;
			ai := zf.Zfs.Auto(true) + zfzhi.Zhi.Xhx() + zf.Zfs.Increment(true) +
				zfzhi.Zhi.Dyh() + zfzhi.Zhi.Shuzi2w() + zfzhi.Zhi.Shuzi0w() +
				zfzhi.Zhi.Shuzi0w() + zfzhi.Zhi.Fh() + zfzhi.Zhi.Hhf()
			buffer.WriteString(ai)
		}
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Sql(true)
		scpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Chuangjianbiao(false) + zfzhi.Zhi.Dh() + zf.Zfs.Sql(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(scpath, buffer.Bytes(), os.ModePerm)
	}
}
