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

func Shengchengsql() {
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
		_, biaos, _ := gongju.Fanshebiaolies(mkv)
		for b, _ := range biaos {
			//CREATE TABLE `mkv`.`b` (
			crestr := zf.Zfs.Create(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Table(true) +
				zfzhi.Zhi.Kgf() + zfzhi.Zhi.Yzb() + mkv + zfzhi.Zhi.Yzb() +
				zfzhi.Zhi.Dh() + zfzhi.Zhi.Yzb() + b + zfzhi.Zhi.Yzb()
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
	Shengchenginsertsql()
	Shengchengdropsql()

}
func Shengchenginsertsql() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		bjg := gongju.Fanshebiaojiegou(mkv)
		buffer := &bytes.Buffer{}
		for b, _ := range bjg {
			ls := gongju.Fanshejichulie()//所有的表只会自动插入标准字段的数据

			for i := zfzhi.Zhi.Shuzi2() * zfzhi.Zhi.Shuzi10() * zfzhi.Zhi.Shuzi10() + zfzhi.Zhi.Shuzi1();
				i < zfzhi.Zhi.Shuzi4() * zfzhi.Zhi.Shuzi10() * zfzhi.Zhi.Shuzi10() + zfzhi.Zhi.Shuzi1(); i++ {
				bufferb := &bytes.Buffer{}
				bufferv := &bytes.Buffer{}
				//INSERT INTO `mkv`.`b`
				iib := zf.Zfs.Insert(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Into(true) +
					zfzhi.Zhi.Kgf() + zfzhi.Zhi.Yzb() + mkv + zfzhi.Zhi.Yzb() +
					zfzhi.Zhi.Dh() + zfzhi.Zhi.Yzb() + b + zfzhi.Zhi.Yzb()
				buffer.WriteString(iib)
				buffer.WriteString(zfzhi.Zhi.Xkhz())
				for _, l := range ls {
					//`caozuoren`,
					lstr := zfzhi.Zhi.Yzb() + l + zfzhi.Zhi.Yzb() + zfzhi.Zhi.Dou()
					bufferb.WriteString(lstr)
					bj := gongju.Biaojishuoming(l)
					if bj == zf.Zfs.Zhujian(false) {
						bufferv.WriteString(strconv.Itoa(i) + zfzhi.Zhi.Dou())
					} else if bj == zf.Zfs.Sql(false) + zf.Zfs.Bianma(false) {
						//'DT_i',
						val := zfzhi.Zhi.Dyhe() + zf.Zfs.DT(false) + zfzhi.Zhi.Xhx() +
							strconv.Itoa(i) + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou()
						bufferv.WriteString(val)
					} else {
						//'1',
						val := zfzhi.Zhi.Dyhe() + bj + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou()
						bufferv.WriteString(val)
					}
				}
				bbstr := bufferb.String()[zfzhi.Zhi.Shuzi0():len(bufferb.String()) - zfzhi.Zhi.Shuzi1()]
				bvstr := bufferv.String()[zfzhi.Zhi.Shuzi0():len(bufferv.String()) - zfzhi.Zhi.Shuzi1()]
				buffer.WriteString(bbstr)
				buffer.WriteString(zfzhi.Zhi.Xkhy() + zf.Zfs.Values(true))
				buffer.WriteString(zfzhi.Zhi.Xkhz())

				buffer.WriteString(bvstr)

				buffer.WriteString(zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Fh())
			}
		}
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Sql(true)
		scpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Insert(false) +
			zf.Zfs.Tables(true) + zfzhi.Zhi.Dh() + zf.Zfs.Sql(true)

		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(scpath, buffer.Bytes(), os.ModePerm)
	}
}
func Shengchengdropsql() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		bjg := gongju.Fanshebiaojiegou(mkv)
		buffer := &bytes.Buffer{}
		for b, _ := range bjg {
			//drop table `b`;
			dtstr := zf.Zfs.Drop(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Table(true) +
				zfzhi.Zhi.Kgf() + zfzhi.Zhi.Yzb() + mkv + zfzhi.Zhi.Yzb() +
				zfzhi.Zhi.Dh() + zfzhi.Zhi.Yzb() + b + zfzhi.Zhi.Yzb() +
				zfzhi.Zhi.Fh() + zfzhi.Zhi.Hhf()
			buffer.WriteString(dtstr)
		}
		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Sql(true)
		scpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Drop(false) +
			zf.Zfs.Tables(true) + zfzhi.Zhi.Dh() + zf.Zfs.Sql(true)

		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(scpath, buffer.Bytes(), os.ModePerm)
	}
}