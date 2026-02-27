package main

import (
	"fmt"
	"ostadbun/pkg/richerror"
)

func main() {

	a()

}

func a() error {

	err := b()

	errR, _ := err.(richerror.RichError)

	errF := errR.Unwrap()

	errY, _ := errF.(richerror.RichError)

	fmt.Println(errR.Error(), errY.Error())

	return richerror.New("layertest - funca").WithErr(err).WithMessage("this is a")

}

func b() error {

	err := c()
	return richerror.New("").WithErr(err).WithMessage("this is b")

}

func c() richerror.RichError {

	return richerror.New("layertest - funcc").WithMessage("this is c")

}
