package scdir

import (
	"gongju"
	"changliang/zfzhi"
	"changliang/zf"
	"os"
	"strings"
	"io/ioutil"
	"bytes"
)

func Shengchengdir() {
	dirs := []string{
		"Dtyonghu",
		"Dtziyuan",
		"Dtjuese",
		"Dtjiance",
		"Dtxiangmu",
		"Dtshijian",
		"Dtshebei",
		"Dtwenjuan",
		"Dtyinpin",
		"Dtjiankang",
		"Dtliuyan",
		"Dtcaiji",
	}

	for _, dir := range dirs {
		dirx := strings.ToLower(dir)
		buffer := &bytes.Buffer{}
		//<template>
		tzstr := zfzhi.Zhi.Xy() + zf.Zfs.Template(true) + zfzhi.Zhi.Dy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(tzstr)

		//<div class='kongbai'>
		dzstr := zfzhi.Zhi.Xy() + zf.Zfs.Div(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Class(true) +
			zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyhe() + dirx + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(dzstr)
		//</div>
		dystr := zfzhi.Zhi.Xy() + zfzhi.Zhi.Xx() + zf.Zfs.Div(true) + zfzhi.Zhi.Dy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(dystr)
		//</template>
		tystr := zfzhi.Zhi.Xy() + zfzhi.Zhi.Xx() + zf.Zfs.Template(true) + zfzhi.Zhi.Dy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(tystr)
		//<script type='text/ecmascript-6'>
		szstr := zfzhi.Zhi.Xy() + zf.Zfs.Script(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Type(true) +
			zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyhe() + zf.Zfs.Text(true) + zfzhi.Zhi.Xx() +
			zf.Zfs.Ecmascript(true) + zfzhi.Zhi.Jian() + zfzhi.Zhi.Shuzi6w() +
			zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(szstr)
		//	export default{}
		edstr := zf.Zfs.Export(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Default(true) +
			zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(edstr)
		//</script>
		systr := zfzhi.Zhi.Xy() + zfzhi.Zhi.Xx() + zf.Zfs.Script(true) + zfzhi.Zhi.Dy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(systr)
		//<style scope>
		yzstr := zfzhi.Zhi.Xy() + zf.Zfs.Style(true) + zfzhi.Zhi.Kgf() +
			zf.Zfs.Scope(true) + zfzhi.Zhi.Dy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(yzstr)
		//</style>
		yystr := zfzhi.Zhi.Xy() + zfzhi.Zhi.Xx() + zf.Zfs.Style(true) + zfzhi.Zhi.Dy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(yystr)

		d :=  gongju.Getgopath() + zfzhi.Zhi.Xx() + zf.Zfs.Mulu(true) + zfzhi.Zhi.Xx() + dirx
		path := d + zfzhi.Zhi.Xx() + dir + zfzhi.Zhi.Dh() + zf.Zfs.Vue(true)
		os.MkdirAll(d, os.ModePerm)
		ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	}

}
