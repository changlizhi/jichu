package sjkmhsydata

import "changliang/zf"

type Chujinfangan struct{}

func (cjfa *Chujinfangan) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
func (cjfa *Chujinfangan) Cankao() string {
	return zf.Zfs.Cankao(false)
}
func (cjfa *Chujinfangan) Neirong() string {
	return zf.Zfs.Neirong(false)
}
func (cjfa *Chujinfangan) Zhidingren() string {
	return zf.Zfs.Zhidingren(false)
}
func (cjfa *Chujinfangan) Weiyibianma() string {
	return zf.Zfs.Weiyibianma(false)
}
