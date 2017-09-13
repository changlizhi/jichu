package sjkmhsydata

import (
	"changliang/biaolie"
)

func (sjkmhsydata *Sjkmhsydata) JueseId(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Id(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) JueseBiaoji(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Biaoji(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) JueseBianma(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Bianma(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) JueseMingcheng(xiaoxie bool) string {
	return biaolie.Bls.Juese(xiaoxie) + biaolie.Bls.Mingcheng(xiaoxie)
}
