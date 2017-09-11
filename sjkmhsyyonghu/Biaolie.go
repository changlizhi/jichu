package sjkmhsyyonghu

import (
	"changliang/biaolie"
)

func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseId(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianId(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) QuanxianId(xiaoxie bool) string {
	return biaolie.Bls.Quanxian(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) XinxiId(xiaoxie bool) string {
	return biaolie.Bls.Xinxi(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) XinxijueseId(xiaoxie bool) string {
	return biaolie.Bls.Xinxijuese(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) YanzhengId(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) YanzhengleixingId(xiaoxie bool) string {
	return biaolie.Bls.Yanzhengleixing(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) QuanxianBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Quanxian(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) XinxiBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Xinxi(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) XinxijueseBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Xinxijuese(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) YanzhengBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) YanzhengleixingBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Yanzhengleixing(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseBianma(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) QuanxianBianma(xiaoxie bool) string {
	return biaolie.Bls.Quanxian(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) XinxiBianma(xiaoxie bool) string {
	return biaolie.Bls.Xinxi(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) YanzhengBianma(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) YanzhengleixingBianma(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) QuanxianMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Quanxian(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) XinxiMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Xinxi(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) YanzhengMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) YanzhengleixingMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Yanzhengleixing(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianJuesebianma(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Juesebianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) XinxijueseJuesebianma(xiaoxie bool) string {
	return biaolie.Bls.Xinxijuese(xiaoxie) + biaolie.Bls.Juesebianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianQuanxianbianma(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Quanxianbianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) XinxijueseXinxibianma(xiaoxie bool) string {
	return biaolie.Bls.Xinxijuese(xiaoxie) + biaolie.Bls.Xinxibianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) QuanxianLujing(xiaoxie bool) string {
	return biaolie.Bls.Quanxian(xiaoxie) + biaolie.Bls.Lujing(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) XinxiYouxiang(xiaoxie bool) string {
	return biaolie.Bls.Xinxi(xiaoxie) + biaolie.Bls.Youxiang(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) YanzhengLeixing(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Leixing(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) YanzhengZhi(xiaoxie bool) string {
	return biaolie.Bls.Yanzheng(xiaoxie) + biaolie.Bls.Zhi(xiaoxie)
}
