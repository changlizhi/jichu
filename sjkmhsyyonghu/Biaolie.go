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
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianJuesebianma(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Juesebianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianQuanxianbianma(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Quanxianbianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Juesequanxian(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseBianma(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
