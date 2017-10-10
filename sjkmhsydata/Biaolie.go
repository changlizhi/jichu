package sjkmhsydata

import (
	"changliang/biaolie"
)

func (sjkmhsydata *Sjkmhsydata) DtziyuanId(xiaoxie bool) string {
	return biaolie.Bls.Dtziyuan(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) DtziyuanMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Dtziyuan(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) DtziyuanPaixu(xiaoxie bool) string {
	return biaolie.Bls.Dtziyuan(xiaoxie) + biaolie.Bls.Paixu(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) DtziyuanBianma(xiaoxie bool) string {
	return biaolie.Bls.Dtziyuan(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) DtziyuanFubianma(xiaoxie bool) string {
	return biaolie.Bls.Dtziyuan(xiaoxie) + biaolie.Bls.Fubianma(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) DtziyuanLujing(xiaoxie bool) string {
	return biaolie.Bls.Dtziyuan(xiaoxie) + biaolie.Bls.Lujing(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) DtziyuanChuangjianriqi(xiaoxie bool) string {
	return biaolie.Bls.Dtziyuan(xiaoxie) + biaolie.Bls.Xiugairiqi(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) DtziyuanXiugairiqi(xiaoxie bool) string {
	return biaolie.Bls.Dtziyuan(xiaoxie) + biaolie.Bls.Xiugairiqi(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) DtziyuanBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Dtziyuan(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) DtziyuanLeixing(xiaoxie bool) string {
	return biaolie.Bls.Dtziyuan(xiaoxie) + biaolie.Bls.Leixing(xiaoxie)
}
