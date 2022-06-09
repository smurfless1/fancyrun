package fancyrun

import "github.com/sirupsen/logrus"

func CheckFinal(e error) error {
	if e != nil {
		logrus.Fatal(e)
		panic(e)
	}
	return nil
}

func CheckInline(e error) {
	if e != nil {
		logrus.Fatal(e)
		panic(e)
	}
}
