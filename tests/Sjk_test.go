package tests

import (
	"jichu/sjkhfxyonghu"
	"log"
	"testing"
)

func TestPrintsome(t *testing.T) {
	sjkhfxyonghu.Sjkhfxyonghus.JueseBianma(false)
	log.Println(sjkhfxyonghu.Sjkhfxyonghus.JueseBianma(false))
}
