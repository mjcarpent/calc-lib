package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/mcarpenter/calc-lib"
)

func main() {

	//param1, err := strconv.Atoi(os.Args[1])
	//if err != nil {
	//	panic(err)
	//}
	//
	//param2, err := strconv.Atoi(os.Args[2])
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println(param1, " ", param2)

	var a = Handler{os.Stdout, calc.Addition{}}
	err := a.Handle(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

type Handler struct {
	writer io.Writer
	calc   calc.Addition
}

func (this *Handler) Handle(args []string) error {

	if len(args) != 2 {
		return fmt.Errorf("Usage: Addition requires 2 operands")
	}

	param1, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	param2, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	param3 := this.calc.Calculate(param1, param2)
	_, err = fmt.Fprintf(this.writer, "%d + %d = %d\n", param1, param2, param3)
	if err != nil {
		return err
	}

	return nil
}
