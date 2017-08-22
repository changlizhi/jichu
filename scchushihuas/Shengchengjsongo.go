package scchushihuas

import (
	"gongju"
	"changliang/zf"
	"changliang/zfzhi"
	"bytes"
	"os"
	"io/ioutil"
	"log"
)

func Shengchengjsongo() {
	dir := gongju.Getgopath() + zfzhi.Zhi.Xx() +
		gongju.Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi + zfzhi.Zhi.Xx() + zf.Zfs.Chushihuas(true)
	buffermoji := &bytes.Buffer{}
	//package chushihuas
	pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Chushihuas(true) + zfzhi.Zhi.Hhf()
	buffermoji.WriteString(pac)
	//type Tongyong struct
	typestr := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + gongju.Chushihuas[zf.Zfs.Jsonmojiming(false)].Zhi +
		zfzhi.Zhi.Kgf() + zf.Zfs.Struct(true)
	buffermoji.WriteString(typestr)
	// {\n
	buffermoji.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	for jmk, _ := range gongju.Jsonmojis {
		//Guilei string
		jm := jmk + zfzhi.Zhi.Kgf() + zf.Zfs.String(true) + zfzhi.Zhi.Hhf()
		buffermoji.WriteString(jm)
	}
	// \n}
	buffermoji.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy())
	for jlk, _ := range gongju.Jsonlies {
		bufferlie := &bytes.Buffer{}
		// package chushihuas
		bufferlie.WriteString(pac)
		// type Shezhi struct
		typestrlie := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + jlk +
			zfzhi.Zhi.Kgf() + zf.Zfs.Struct(true)
		bufferlie.WriteString(typestrlie)
		// { \n
		bufferlie.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
		//从jsonlie包下读取里面的方法解析出所有的*Jl的方法
		lies := gongju.Yigejsonlies(jlk)
		for _, lie := range lies {
			// Xxx []Tongyong
			l := lie + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy() + zf.Zfs.Tongyong(false) + zfzhi.Zhi.Hhf()
			bufferlie.WriteString(l)
		}
		//\n }
		bufferlie.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy())
		pathlie := dir + zfzhi.Zhi.Xx() + jlk + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
		ioutil.WriteFile(pathlie, bufferlie.Bytes(), os.ModePerm)
	log.Println(bufferlie.String())
	}
	pathmoji := dir + zfzhi.Zhi.Xx() + gongju.Chushihuas[zf.Zfs.Jsonmojiming(false)].Zhi + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
	os.MkdirAll(dir, os.ModePerm)
	ioutil.WriteFile(pathmoji, buffermoji.Bytes(), os.ModePerm)
}