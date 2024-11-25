package handler

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/mcarpenter/calc-lib"
)

type Handler struct {
	Writer io.Writer
	Calc   calc.Addition
}

func (this *Handler) Handle(args []string) error {

	if len(args) != 2 {
		return insufficientArgs
	}

	param1, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %w", invalidArg, err)
	}

	param2, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: %w", invalidArg, err)
	}

	param3 := this.Calc.Calculate(param1, param2)
	_, err = fmt.Fprintf(this.Writer, "%d", param3)
	if err != nil {
		return err
	}

	return nil
}

var (
	insufficientArgs = errors.New("Usage: Addition requires 2 operands")
	invalidArg       = errors.New("invalid argument")
)
