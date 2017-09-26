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
	rbu := bytes.Buffer{}
	d := gongju.Getgopath() + zfzhi.Zhi.Xx() + zf.Zfs.Mulu(true) + zfzhi.Zhi.Xx()
	for _, dir := range dirs {
		dirx := strings.ToLower(dir)
		buffer := bytes.Buffer{}
		//<template>
		tzstr := zfzhi.Zhi.Xy() + zf.Zfs.Template(true) + zfzhi.Zhi.Dy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(tzstr)

		//<div class='kongbai'>
		dzstr := zfzhi.Zhi.Xy() + zf.Zfs.Div(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Class(true) +
			zfzhi.Zhi.Dyh() + zfzhi.Zhi.Dyhe() + dirx + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dy() + zfzhi.Zhi.Hhf()
		buffer.WriteString(dzstr)
		buffer.WriteString(dirx)
		buffer.WriteString(dirx)
		buffer.WriteString(dirx)

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
		pathd := d + dirx + zfzhi.Zhi.Xx() + dir + zfzhi.Zhi.Dh() + zf.Zfs.Vue(true)

		os.MkdirAll(pathd, os.ModePerm)
		ioutil.WriteFile(pathd, buffer.Bytes(), os.ModePerm)
		//{
		rbu.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//path: '/dtyonghu',
		pastr := zf.Zfs.Path(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Xx() +
			dirx + zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		rbu.WriteString(pastr)
		//name: 'dtyonghu',
		ndstr := zf.Zfs.Name(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyhe() + dirx +
			zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		rbu.WriteString(ndstr)
		//component: Dtyonghu,
		cdstr := zf.Zfs.Component(true) + zfzhi.Zhi.Mh() + dir + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		rbu.WriteString(cdstr)
		//children: [{
		clstr := zf.Zfs.Children(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
		rbu.WriteString(clstr)
		//path: 'biaodan',
		pbstr := zf.Zfs.Path(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyhe() + zf.Zfs.Biaodan(true) +
			zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		rbu.WriteString(pbstr)
		//component: Biaodan
		cbstr := zf.Zfs.Component(true) + zfzhi.Zhi.Mh() + zf.Zfs.Biaodan(false) + zfzhi.Zhi.Hhf()
		rbu.WriteString(cbstr)
		//}, {
		khstr := zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
		rbu.WriteString(khstr)
		//path: 'quanbu',
		pqstr := zf.Zfs.Path(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyhe() + zf.Zfs.Quanbu(true) +
			zfzhi.Zhi.Dyhe() + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf()
		rbu.WriteString(pqstr)
		//component: Quanbu
		cqstr := zf.Zfs.Component(true) + zfzhi.Zhi.Mh() + zf.Zfs.Quanbu(false) + zfzhi.Zhi.Hhf()
		rbu.WriteString(cqstr)
		//}]
		rbu.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Zkhy())
		//},
		rbu.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Dou())
	}
	rpath := d + zf.Zfs.Router(true) + zfzhi.Zhi.Dh() + zf.Zfs.Js(true)
	ioutil.WriteFile(rpath, rbu.Bytes(), os.ModePerm)

}
