package skjhfxyonghu

import (
	"changliang/biaolie"
)

func (sjk *Sjk) JueseId(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjk *Sjk) JuesequanxianId(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjk *Sjk) QuanxianId(xiaoxie bool) string {
	return biaolie.Bls.Quanxian(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjk *Sjk) XinxiId(xiaoxie bool) string {
	return biaolie.Bls.Xinxi(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjk *Sjk) XinxijueseId(xiaoxie bool) string {
	return biaolie.Bls.Xinxijuese(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjk *Sjk) YanzhengId(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjk *Sjk) YanzhengleixingId(xiaoxie bool) string {
	return biaolie.Bls.Yanzhengleixing(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjk *Sjk) JueseBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjk *Sjk) JuesequanxianBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjk *Sjk) QuanxianBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Quanxian(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjk *Sjk) XinxiBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Xinxi(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjk *Sjk) XinxijueseBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Xinxijuese(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjk *Sjk) YanzhengBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjk *Sjk) YanzhengleixingBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Yanzhengleixing(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjk *Sjk) JueseBianma(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjk *Sjk) QuanxianBianma(xiaoxie bool) string {
	return biaolie.Bls.Quanxian(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjk *Sjk) XinxiBianma(xiaoxie bool) string {
	return biaolie.Bls.Xinxi(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjk *Sjk) YanzhengBianma(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjk *Sjk) YanzhengleixingBianma(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjk *Sjk) JueseMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjk *Sjk) QuanxianMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Quanxian(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjk *Sjk) XinxiMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Xinxi(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjk *Sjk) YanzhengMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjk *Sjk) YanzhengleixingMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Yanzhengleixing(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjk *Sjk) JuesequanxianJuesebianma(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Juesebianma(xiaoxie)
}
func (sjk *Sjk) XinxijueseJuesebianma(xiaoxie bool) string {
	return biaolie.Bls.Xinxijuese(xiaoxie) + biaolie.Bls.Juesebianma(xiaoxie)
}
func (sjk *Sjk) JuesequanxianQuanxianbianma(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Quanxianbianma(xiaoxie)
}
func (sjk *Sjk) XinxijueseXinxibianma(xiaoxie bool) string {
	return biaolie.Bls.Xinxijuese(xiaoxie) + biaolie.Bls.Xinxibianma(xiaoxie)
}
func (sjk *Sjk) QuanxianLujing(xiaoxie bool) string {
	return biaolie.Bls.Quanxian(xiaoxie) + biaolie.Bls.Lujing(xiaoxie)
}
func (sjk *Sjk) XinxiYouxiang(xiaoxie bool) string {
	return biaolie.Bls.Xinxi(xiaoxie) + biaolie.Bls.Youxiang(xiaoxie)
}
func (sjk *Sjk) YanzhengLeixing(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Leixing(xiaoxie)
}
func (sjk *Sjk) YanzhengZhi(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Zhi(xiaoxie)
}
