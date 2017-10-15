package scguojihua

import (
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"io/ioutil"
	"os"
)

func Shengchengguojihua() {
	mkarr := gongju.Mokuaimingsarr
	mks := gongju.Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		szpath := gongju.Getjichupath() + zfzhi.Zhi.Xx() + zf.Zfs.Sc(true) +
			zf.Zfs.Guojihua(true) + zfzhi.Zhi.Xx() + zf.Zfs.A(false) + zfzhi.Zhi.Dh() + zf.Zfs.Goconf(true)
		ywpath := gongju.Getjichupath() + zfzhi.Zhi.Xx() + zf.Zfs.Sc(true) +
			zf.Zfs.Guojihua(true) + zfzhi.Zhi.Xx() + zf.Zfs.Cuowu(false) + zfzhi.Zhi.Dh() + zf.Zfs.Goconf(true)
		zwpath := gongju.Getjichupath() + zfzhi.Zhi.Xx() + zf.Zfs.Sc(true) +
			zf.Zfs.Guojihua(true) + zfzhi.Zhi.Xx() + zf.Zfs.Tishi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Goconf(true)
		szbs, _ := ioutil.ReadFile(szpath)
		ywbs, _ := ioutil.ReadFile(ywpath)
		zwbs, _ := ioutil.ReadFile(zwpath)

		dir := gongju.Getgopath() + zfzhi.Zhi.Xx() + mkv +
			zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Guojihua(true)
		scszpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.A(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		scywpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Cuowu(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		sczwpath := dir + zfzhi.Zhi.Xx() + zf.Zfs.Tishi(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)

		os.MkdirAll(dir, os.ModePerm)
		ioutil.WriteFile(scszpath, szbs, os.ModePerm)
		ioutil.WriteFile(scywpath, ywbs, os.ModePerm)
		ioutil.WriteFile(sczwpath, zwbs, os.ModePerm)
	}
}
