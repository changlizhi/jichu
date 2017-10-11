package sjkmhsyyonghu

import (
	"changliang/zf"
)

func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseId(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Id(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianId(xiaoxie bool) string {
	return zf.Zfs.Juesequanxian(xiaoxie) + zf.Zfs.Id(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianJuesebianma(xiaoxie bool) string {
	return zf.Zfs.Juesequanxian(xiaoxie) + zf.Zfs.Juesebianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianQuanxianbianma(xiaoxie bool) string {
	return zf.Zfs.Juesequanxian(xiaoxie) + zf.Zfs.Quanxianbianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseBiaoji(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JuesequanxianBiaoji(xiaoxie bool) string {
	return zf.Zfs.Juesequanxian(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseBianma(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Bianma(xiaoxie)
}
func (sjkmhsyyonghu *Sjkmhsyyonghu) JueseMingcheng(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)
}
