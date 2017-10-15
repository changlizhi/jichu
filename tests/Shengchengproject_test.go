package tests

import (
	"jichu/scchushihuas"
	"jichu/scconf"
	"jichu/sccuowus"
	"jichu/scfortests"
	"jichu/sckongzhiqis"
	"jichu/sckus"
	"jichu/scluyous"
	"jichu/scmain"
	"jichu/scmoxing"
	"jichu/scpeizhi"
	"jichu/scyewus"
	"testing"
)

func TestShengchengproj(t *testing.T) {

	//var _ = gauge.Step("生成Moxings", func() {
	scmoxing.Shengchengmoxings()
	//})

	//var _ = gauge.Step("生成Conf", func() {
	scconf.Shengchengconf()
	//})

	//var _ = gauge.Step("生成Cuowu", func() {
	sccuowus.Shengchengziduancuowu()
	//})
	//var _ = gauge.Step("生成Chushihuas-jsongo", func() {
	scchushihuas.Shengchengjsongo()
	//})
	//var _ = gauge.Step("生成Chushihuas-shezhigo", func() {
	scchushihuas.Shengchengshezhis()
	//})
	//var _ = gauge.Step("生成Chushihuas-lujinghuoqu", func() {
	scchushihuas.Shengchenglujinghuoqu()
	//})

	//var _ = gauge.Step("生成Chushihuas-jsonlie", func() {
	scchushihuas.Shengchengchushihuajsonlie()
	//})
	//需要Chushihuas这个前提与“生成Chushihuas-jsonlie”互相关联
	//var _ = gauge.Step("生成Chushihuas-duqujson", func() {
	scchushihuas.Shengchengduqujson()
	//})

	//需要moxings先生成
	//var _ = gauge.Step("生成Chushihuas-ormer", func() {
	scchushihuas.Shengchengchushihuaormer()
	//})

	//var _ = gauge.Step("生成Chushihuas-test", func() {
	scchushihuas.Shengchengchushihuatest()
	//})

	//var _ = gauge.Step("生成拼接结构体字段", func() {
	scfortests.Shengchengfortests()
	//})
	//var _ = gauge.Step("生成拼接结构体字段test", func() {
	scfortests.Shengchengforteststest()
	//})

	//var _ = gauge.Step("生成数据库操作kus", func() {
	sckus.Shengchengkus()
	//})
	//var _ = gauge.Step("生成数据库操作kus-test", func() {
	sckus.Shengchengkustests()
	//})
	//var _ = gauge.Step("生成数据库操作yewus", func() {
	scyewus.Shengchengyewu()
	//})
	//var _ = gauge.Step("生成数据库操作yewus-test", func() {
	scyewus.Shengchengyewutest()
	//})
	//var _ = gauge.Step("生成数据库操作kongzhiqis-main", func() {
	sckongzhiqis.Shengchengmainkongzhiqi()
	//})
	//var _ = gauge.Step("生成数据库操作kongzhiqis-main-test", func() {
	sckongzhiqis.Shengchengmainkongzhiqitest()
	//})
	//var _ = gauge.Step("生成数据库操作kongzhiqis", func() {
	sckongzhiqis.Shengchengkongzhiqis()
	//})
	//var _ = gauge.Step("生成数据库操作kongzhiqis-test", func() {
	sckongzhiqis.Shengchengkongzhiqitest()
	//})
	//var _ = gauge.Step("生成数据库操作luyous-main", func() {
	scluyous.Shengchengmainluyou()
	//})
	//var _ = gauge.Step("生成数据库操作luyous-main-test", func() {
	scluyous.Shengchengmainluyoutest()
	//})
	//var _ = gauge.Step("生成数据库操作luyous", func() {
	scluyous.Shengchengluyous()
	//})
	//var _ = gauge.Step("生成数据库操作luyous-test", func() {
	scluyous.Shengchengluyoustests()
	//})

	//var _ = gauge.Step("生成Main", func() {
	scmain.Shengchengmain()
	//})
}
