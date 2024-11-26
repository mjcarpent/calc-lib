package calc

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

var Add = &Addition{}
var Sub = &Subtraction{}
var Mul = &Multiplication{}
var Div = &Division{}
var Mod = &Modulo{}

var maxInt = math.MaxInt
var minInt = math.MinInt

func assertEquals(t *testing.T, got, expected int) {

	t.Helper()
	if got != expected {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestHandler_Add_ZeroAsFirstOperandPositiveOutput(t *testing.T) {
	got := Add.Calculate(0, 1)
	assertEquals(t, got, 1)
}

func TestHandler_Add_ZeroAsSecondOperandPositiveOutput(t *testing.T) {
	got := Add.Calculate(1, 0)
	assertEquals(t, got, 1)
}

func TestHandler_Add_ZeroAsFirstOperandNegativeOutput(t *testing.T) {
	got := Add.Calculate(0, -3)
	assertEquals(t, got, -3)
}

func TestHandler_Add_ZeroAsSecondOperandNegativeOutput(t *testing.T) {
	got := Add.Calculate(-17, 0)
	assertEquals(t, got, -17)
}

func TestHandler_Add_NegativeAsFirstOperandPositiveOutput(t *testing.T) {
	got := Add.Calculate(-1, 5)
	assertEquals(t, got, 4)
}

func TestHandler_Add_NegativeAsSecondOperandPositiveOutput(t *testing.T) {
	got := Add.Calculate(5, -1)
	assertEquals(t, got, 4)
}

func TestHandler_Add_NegativeAsFirstOperandNegativeOutput(t *testing.T) {
	got := Add.Calculate(-9, 5)
	assertEquals(t, got, -4)
}

func TestHandler_Add_NegativeAsSecondOperandNegativeOutput(t *testing.T) {
	got := Add.Calculate(5, -9)
	assertEquals(t, got, -4)
}

func TestHandler_Add_MaxNegativeOperandWithPositive(t *testing.T) {
	got := Add.Calculate(minInt, 5)
	assertEquals(t, got, minInt+5)
}

func TestHandler_Add_MaxNegativeOperandWithNegative(t *testing.T) {
	got := Add.Calculate(minInt, -5)
	assertEquals(t, got, maxInt-4)
}

func TestHandler_Add_MaxPositiveOperandWithPositive(t *testing.T) {
	got := Add.Calculate(5, maxInt)
	assertEquals(t, got, maxInt+5)
}

func TestHandler_Add_MaxPositiveOperandWithNegative(t *testing.T) {
	got := Add.Calculate(maxInt, -5)
	assertEquals(t, got, maxInt-5)
}

func TestHandler_Sub_ZeroAsFirstOperandPositiveOutput(t *testing.T) {
	got := Sub.Calculate(0, -1)
	assertEquals(t, got, 1)
}

func TestHandler_Sub_ZeroAsSecondOperandPositiveOutput(t *testing.T) {
	got := Sub.Calculate(1, 0)
	assertEquals(t, got, 1)
}

func TestHandler_Sub_ZeroAsFirstOperandNegativeOutput(t *testing.T) {
	got := Sub.Calculate(0, 3)
	assertEquals(t, got, -3)
}

func TestHandler_Sub_ZeroAsSecondOperandNegativeOutput(t *testing.T) {
	got := Sub.Calculate(-17, 0)
	assertEquals(t, got, -17)
}

func TestHandler_Sub_NegativeAsFirstOperandPositiveOutput(t *testing.T) {
	got := Sub.Calculate(-1, -5)
	assertEquals(t, got, 4)
}

func TestHandler_Sub_NegativeAsSecondOperandPositiveOutput(t *testing.T) {
	got := Sub.Calculate(5, -1)
	assertEquals(t, got, 6)
}

func TestHandler_Sub_NegativeAsFirstOperandNegativeOutput(t *testing.T) {
	got := Sub.Calculate(-9, 5)
	assertEquals(t, got, -14)
}

func TestHandler_Sub_NegativeAsSecondOperandNegativeOutput(t *testing.T) {
	got := Sub.Calculate(5, 9)
	assertEquals(t, got, -4)
}

func TestHandler_Sub_MaxNegativeOperandWithPositive(t *testing.T) {
	got := Sub.Calculate(minInt, 5)
	assertEquals(t, got, maxInt-4)
}

func TestHandler_Sub_MaxNegativeOperandWithNegative(t *testing.T) {
	got := Sub.Calculate(minInt, -5)
	assertEquals(t, got, minInt+5)
}

func TestHandler_Sub_MaxPositiveWithPositive(t *testing.T) {
	got := Sub.Calculate(maxInt, 5)
	assertEquals(t, got, maxInt-5)
}

func TestHandler_Sub_MaxPositiveWithNegative(t *testing.T) {
	got := Sub.Calculate(maxInt, -5)
	assertEquals(t, got, minInt+4)
}

func TestHandler_Mul_ZeroAsFirstOperandPositiveOutput(t *testing.T) {
	got := Mul.Calculate(0, 1)
	assertEquals(t, got, 0)
}

func TestHandler_Mul_ZeroAsSecondOperandPositiveOutput(t *testing.T) {
	got := Mul.Calculate(1, 0)
	assertEquals(t, got, 0)
}

func TestHandler_Mul_ZeroAsFirstOperandNegativeOutput(t *testing.T) {
	got := Mul.Calculate(0, -3)
	assertEquals(t, got, -0)
}

func TestHandler_Mul_ZeroAsSecondOperandNegativeOutput(t *testing.T) {
	got := Mul.Calculate(-17, 0)
	assertEquals(t, got, 0)
}

func TestHandler_Mul_NegativeAsFirstOperandNegativeOutput(t *testing.T) {
	got := Mul.Calculate(-1, 5)
	assertEquals(t, got, -5)
}

func TestHandler_Mul_NegativeAsSecondOperandNegativeOutput(t *testing.T) {
	got := Mul.Calculate(5, -1)
	assertEquals(t, got, -5)
}

func TestHandler_Mul_NegativeBothOperandsPositiveOutput(t *testing.T) {
	got := Mul.Calculate(-9, -5)
	assertEquals(t, got, 45)
}

func TestHandler_Mul_BothOperandsPositiveWithPositiveOutput(t *testing.T) {
	got := Mul.Calculate(5, 9)
	assertEquals(t, got, 45)
}

func TestHandler_Mul_MaxNegativeOperandWithPositiveOutput(t *testing.T) {
	got := Mul.Calculate(minInt, 1)
	assertEquals(t, got, minInt)
}

func TestHandler_Mul_MaxNegativeOperandWithNegativeOutput(t *testing.T) {
	got := Mul.Calculate(minInt, -1)
	assertEquals(t, got, minInt)
}

func TestHandler_Mul_MaxPositiveOperandWithPositiveOutput(t *testing.T) {
	got := Mul.Calculate(1, maxInt)
	assertEquals(t, got, maxInt)
}

func TestHandler_Mul_MaxPositiveOperandWithNegativeOutput(t *testing.T) {
	got := Mul.Calculate(maxInt, -1)
	assertEquals(t, got, minInt+1)
}

func TestHandler_Div_ZeroAsFirstOperandZeroOutput(t *testing.T) {
	got := Div.Calculate(0, 1)
	assertEquals(t, got, 0)
}

func TestHandler_Div_ZeroAsSecondOperandNanOutput(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			errMsg := fmt.Sprintf("%v", r)
			if strings.Compare(errMsg, "runtime error: integer divide by zero") != 0 {
				t.Errorf("Expected division by zero message and received [%v]", r)
			}
			return
		}
	}()
	_ = Div.Calculate(1, 0)
}

func TestHandler_Div_FirstOperandNegativeOutputNegative(t *testing.T) {
	got := Div.Calculate(-18, 3)
	assertEquals(t, got, -6)
}

func TestHandler_Div_SecondOperandNegativeOutputNegative(t *testing.T) {
	got := Div.Calculate(18, -3)
	assertEquals(t, got, -6)
}

func TestHandler_Div_BothOperandsNegativePositiveOutput(t *testing.T) {
	got := Div.Calculate(-18, -3)
	assertEquals(t, got, 6)
}

func TestHandler_Div_BothOperandsPositivePositiveOutput(t *testing.T) {
	got := Div.Calculate(18, 3)
	assertEquals(t, got, 6)
}

func TestHandler_Div_MaxNegativeOperand(t *testing.T) {
	got := Div.Calculate(minInt, 1000000000000000000)
	assertEquals(t, got, -9)
}

func TestHandler_Div_MaxPositiveOperandWithPositiveOutput(t *testing.T) {
	got := Div.Calculate(maxInt, 1000000000000000000)
	assertEquals(t, got, 9)
}

func TestHandler_Div_OperandsProvideOnlyRemainder(t *testing.T) {
	got := Div.Calculate(2, 9)
	assertEquals(t, got, 0)
}

func TestHandler_Mod_ZeroAsFirstOperandZeroOutput(t *testing.T) {
	got := Mod.Calculate(0, 1)
	assertEquals(t, got, 0)
}

func TestHandler_Mod_ZeroAsSecondOperandNanOutput(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			errMsg := fmt.Sprintf("%v", r)
			if strings.Compare(errMsg, "runtime error: integer divide by zero") != 0 {
				t.Errorf("Expected division by zero message and received [%v]", r)
			}
			return
		}
	}()
	_ = Mod.Calculate(1, 0)
}

func TestHandler_Mod_FirstOperandNegativeOutputNegative(t *testing.T) {
	got := Mod.Calculate(-18, 4)
	assertEquals(t, got, -2)
}

func TestHandler_Mod_SecondOperandNegativeOutputNegative(t *testing.T) {
	got := Mod.Calculate(18, -4)
	assertEquals(t, got, 2)
}

func TestHandler_Mod_BothOperandsNegativePositiveOutput(t *testing.T) {
	got := Mod.Calculate(-18, -4)
	assertEquals(t, got, -2)
}

func TestHandler_Mod_BothOperandsPositivePositiveOutput(t *testing.T) {
	got := Mod.Calculate(18, 4)
	assertEquals(t, got, 2)
}

func TestHandler_Mod_MaxNegativeOperand(t *testing.T) {
	got := Mod.Calculate(minInt, 1000000000000000000)
	assertEquals(t, got, -223372036854775808)
}

func TestHandler_Mod_MaxPositiveOperandWithPositiveOutput(t *testing.T) {
	got := Mod.Calculate(maxInt, 1000000000000000000)
	assertEquals(t, got, 223372036854775807)
}

func TestHandler_Mod_OperandsProvideNoRemainder(t *testing.T) {
	got := Mod.Calculate(18, 3)
	assertEquals(t, got, 0)
}
