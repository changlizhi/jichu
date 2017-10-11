package sjkhfxyonghu

import (
	"changliang/zf"
)

func (sjkhfxyonghu *Sjkhfxyonghu) JueseId(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Id(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) JuesequanxianId(xiaoxie bool) string {
	return zf.Zfs.Juesequanxian(xiaoxie) + zf.Zfs.Id(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) QuanxianId(xiaoxie bool) string {
	return zf.Zfs.Quanxian(xiaoxie) + zf.Zfs.Id(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) XinxiId(xiaoxie bool) string {
	return zf.Zfs.Xinxi(xiaoxie) + zf.Zfs.Id(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) XinxijueseId(xiaoxie bool) string {
	return zf.Zfs.Xinxijuese(xiaoxie) + zf.Zfs.Id(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) YanzhengId(xiaoxie bool) string {
	return zf.Zfs.Yanzheng(xiaoxie) + zf.Zfs.Id(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) YanzhengleixingId(xiaoxie bool) string {
	return zf.Zfs.Yanzhengleixing(xiaoxie) + zf.Zfs.Id(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) JueseBiaoji(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) JuesequanxianBiaoji(xiaoxie bool) string {
	return zf.Zfs.Juesequanxian(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) QuanxianBiaoji(xiaoxie bool) string {
	return zf.Zfs.Quanxian(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) XinxiBiaoji(xiaoxie bool) string {
	return zf.Zfs.Xinxi(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) XinxijueseBiaoji(xiaoxie bool) string {
	return zf.Zfs.Xinxijuese(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) YanzhengBiaoji(xiaoxie bool) string {
	return zf.Zfs.Yanzheng(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) YanzhengleixingBiaoji(xiaoxie bool) string {
	return zf.Zfs.Yanzhengleixing(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) JueseBianma(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Bianma(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) QuanxianBianma(xiaoxie bool) string {
	return zf.Zfs.Quanxian(xiaoxie) + zf.Zfs.Bianma(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) XinxiBianma(xiaoxie bool) string {
	return zf.Zfs.Xinxi(xiaoxie) + zf.Zfs.Bianma(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) YanzhengBianma(xiaoxie bool) string {
	return zf.Zfs.Yanzheng(xiaoxie) + zf.Zfs.Bianma(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) YanzhengleixingBianma(xiaoxie bool) string {
	return zf.Zfs.Yanzheng(xiaoxie) + zf.Zfs.Bianma(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) JueseMingcheng(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) QuanxianMingcheng(xiaoxie bool) string {
	return zf.Zfs.Quanxian(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) XinxiMingcheng(xiaoxie bool) string {
	return zf.Zfs.Xinxi(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) YanzhengMingcheng(xiaoxie bool) string {
	return zf.Zfs.Yanzheng(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) YanzhengleixingMingcheng(xiaoxie bool) string {
	return zf.Zfs.Yanzhengleixing(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) JuesequanxianJuesebianma(xiaoxie bool) string {
	return zf.Zfs.Juesequanxian(xiaoxie) + zf.Zfs.Juesebianma(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) XinxijueseJuesebianma(xiaoxie bool) string {
	return zf.Zfs.Xinxijuese(xiaoxie) + zf.Zfs.Juesebianma(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) JuesequanxianQuanxianbianma(xiaoxie bool) string {
	return zf.Zfs.Juesequanxian(xiaoxie) + zf.Zfs.Quanxianbianma(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) XinxijueseXinxibianma(xiaoxie bool) string {
	return zf.Zfs.Xinxijuese(xiaoxie) + zf.Zfs.Xinxibianma(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) QuanxianLujing(xiaoxie bool) string {
	return zf.Zfs.Quanxian(xiaoxie) + zf.Zfs.Lujing(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) XinxiYouxiang(xiaoxie bool) string {
	return zf.Zfs.Xinxi(xiaoxie) + zf.Zfs.Youxiang(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) YanzhengLeixing(xiaoxie bool) string {
	return zf.Zfs.Yanzheng(xiaoxie) + zf.Zfs.Leixing(xiaoxie)
}
func (sjkhfxyonghu *Sjkhfxyonghu) YanzhengZhi(xiaoxie bool) string {
	return zf.Zfs.Yanzheng(xiaoxie) + zf.Zfs.Zhi(xiaoxie)
}
