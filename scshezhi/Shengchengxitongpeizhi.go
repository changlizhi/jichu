package scshezhi

import (
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"io/ioutil"
	"os"
)

func Shengchengxitongpeizhi() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		szpath := gongju.Getjichupath() + zfzhi.Zhi.Xx() + zf.Zfs.Sc(true) + zf.Zfs.Peizhi(true) + zfzhi.Zhi.Xx() + zf.Zfs.Shezhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(true)
		ywpath := gongju.Getjichupath() + zfzhi.Zhi.Xx() + zf.Zfs.Sc(true) + zf.Zfs.Peizhi(true) + zfzhi.Zhi.Xx() + zf.Zfs.Yingwen(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(true)
		zwpath := gongju.Getjichupath() + zfzhi.Zhi.Xx() + zf.Zfs.Sc(true) + zf.Zfs.Peizhi(true) + zfzhi.Zhi.Xx() + zf.Zfs.Zhongwen(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(true)
		szbs, _ := ioutil.ReadFile(szpath)
		ywbs, _ := ioutil.ReadFile(ywpath)
		zwbs, _ := ioutil.ReadFile(zwpath)

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Peizhi(true)
		scszpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Shezhi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(true)
		scywpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Yingwen(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(true)
		sczwpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Zhongwen(false) + zfzhi.Zhi.Dh() + zf.Zfs.Json(true)

		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(scszpath, szbs, os.ModePerm)
		ioutil.WriteFile(scywpath, ywbs, os.ModePerm)
		ioutil.WriteFile(sczwpath, zwbs, os.ModePerm)
	}
}

func Shengchenggoconfpeizhi() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		szpath := gongju.Getjichupath() + zfzhi.Zhi.Xx() + zf.Zfs.Sc(true) +
			zf.Zfs.Shezhi(true) + zfzhi.Zhi.Xx() + zf.Zfs.A(false) + zfzhi.Zhi.Dh() + zf.Zfs.Goconf(true)
		ywpath := gongju.Getjichupath() + zfzhi.Zhi.Xx() + zf.Zfs.Sc(true) +
			zf.Zfs.Shezhi(true) + zfzhi.Zhi.Xx() + zf.Zfs.Chushihua(false) + zfzhi.Zhi.Dh() + zf.Zfs.Goconf(true)
		zwpath := gongju.Getjichupath() + zfzhi.Zhi.Xx() + zf.Zfs.Sc(true) +
			zf.Zfs.Shezhi(true) + zfzhi.Zhi.Xx() + zf.Zfs.Shujuku(false) + zfzhi.Zhi.Dh() + zf.Zfs.Goconf(true)
		szbs, _ := ioutil.ReadFile(szpath)
		ywbs, _ := ioutil.ReadFile(ywpath)
		zwbs, _ := ioutil.ReadFile(zwpath)

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Shezhi(true)
		scszpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.A(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		scywpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Chushihua(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		sczwpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Shujuku(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)

		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(scszpath, szbs, os.ModePerm)
		ioutil.WriteFile(scywpath, ywbs, os.ModePerm)
		ioutil.WriteFile(sczwpath, zwbs, os.ModePerm)
	}
}
