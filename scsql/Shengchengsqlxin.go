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
			//type Cujinfangan struct{}

			//CREATE TABLE `mkv`.`bshuju` (
			crestr := zf.Zfs.Create(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Table(true) +
				zfzhi.Zhi.Kgf() + zfzhi.Zhi.Yzb() + mkv + zfzhi.Zhi.Yzb() +
				zfzhi.Zhi.Dh() + zfzhi.Zhi.Yzb() + b + zf.Zfs.Shuju(true) + zfzhi.Zhi.Yzb()
			buffer.WriteString(crestr)
			buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

			for _, sjzd := range gongju.Fanshejichushuju() {
				cd := gongju.Liechangdu(sjzd)
				cdzw := strconv.Itoa(cd)
				if sjzd == zf.Zfs.Id(false) {
					//Id int(10) auto_increment comment '主键',
					vc := zfzhi.Zhi.Yzb() + sjzd + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) +
						zfzhi.Zhi.Xkhz() + cdzw + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf() +
						zf.Zfs.Auto(true) + zfzhi.Zhi.Xhx() + zf.Zfs.Increment(true) +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(sjzd) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vc)
				} else {
					vc := zfzhi.Zhi.Yzb() + sjzd + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Varchar(true) +
						zfzhi.Zhi.Xkhz() + cdzw + zfzhi.Zhi.Xkhy() +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(sjzd) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vc)
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

			//create table `mkv`.`bbiao` (
			crestrbiao := zf.Zfs.Create(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Table(true) +
				zfzhi.Zhi.Kgf() + zfzhi.Zhi.Yzb() + mkv + zfzhi.Zhi.Yzb() +
				zfzhi.Zhi.Dh() + zfzhi.Zhi.Yzb() + b + zf.Zfs.Biao(true) + zfzhi.Zhi.Yzb()
			buffer.WriteString(crestrbiao)
			buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

			//type Xxxjiegou struct{}
			for _, sjzd := range gongju.Fanshejichubiao() {
				cdbiao := gongju.Liechangdu(sjzd)
				cdzwbiao := strconv.Itoa(cdbiao)
				if sjzd == zf.Zfs.Id(false) {
					//Id int(10) auto_increment comment '主键',
					vcbiao := zfzhi.Zhi.Yzb() + sjzd + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Int(true) +
						zfzhi.Zhi.Xkhz() + cdzwbiao + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf() +
						zf.Zfs.Auto(true) + zfzhi.Zhi.Xhx() + zf.Zfs.Increment(true) +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(sjzd) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vcbiao)
				} else {
					vcbiao := zfzhi.Zhi.Yzb() + sjzd + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Kgf() + zf.Zfs.Varchar(true) +
						zfzhi.Zhi.Xkhz() + cdzwbiao + zfzhi.Zhi.Xkhy() +
						zfzhi.Zhi.Kgf() + zf.Zfs.Comment(true) + zfzhi.Zhi.Kgf() +
						zfzhi.Zhi.Dyhe() + gongju.Zhongwen(sjzd) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
					buffer.WriteString(vcbiao)
				}
			}
			//primary key (`id`)
			pkstrbiao := zf.Zfs.Primary(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Key(true) +
				zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Yzb() + zf.Zfs.Id(false) + zfzhi.Zhi.Yzb() +
				zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf()
			buffer.WriteString(pkstrbiao)
			buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
			//COMMENT='时间'
			ctbiao := zf.Zfs.Comment(true) + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyhe() +
				gongju.Zhongwen(b) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Hhf()
			buffer.WriteString(ctbiao)
			//COLLATE='utf8_general_ci'
			cugbiao := zf.Zfs.Collate(true) + zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyhe() +
				zf.Zfs.Utf8(true) + zfzhi.Zhi.Xhx() + zf.Zfs.General(true) +
				zfzhi.Zhi.Xhx() + zf.Zfs.Ci(true) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Hhf()
			buffer.WriteString(cugbiao)
			//ENGINE=InnoDB;
			eibiao := zf.Zfs.Engine(true) + zfzhi.Zhi.Dyh() + zf.Zfs.InnoDB(true) +
				zfzhi.Zhi.Fh() + zfzhi.Zhi.Hhf() + zfzhi.Zhi.Hhf()
			buffer.WriteString(eibiao)
		}
		//log.Println(bf.String())
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Sql(true)
		scpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Chuangjianbiao(false) + zf.Zfs.Xin(true) + zfzhi.Zhi.Dh() + zf.Zfs.Sql(true)
		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(scpath, buffer.Bytes(), os.ModePerm)
	}
}
