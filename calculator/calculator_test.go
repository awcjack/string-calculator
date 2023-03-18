package calculator_test

import (
	"calculator/calculator"
	"testing"
)

func TestAddEmptyReturn0(t *testing.T) {
	result, err := calculator.StringAdd("")
	if err != nil {
		t.Fatal(err)
	}
	if result != 0 {
		t.Errorf("Expected 0 but returned %d", result)
	}
}

func TestAdd4Return4(t *testing.T) {
	result, err := calculator.StringAdd("4")
	if err != nil {
		t.Fatal(err)
	}
	if result != 4 {
		t.Errorf("Expected 4 but returned %d", result)
	}
}
func TestAdd14Return14(t *testing.T) {
	result, err := calculator.StringAdd("14")
	if err != nil {
		t.Fatal(err)
	}
	if result != 14 {
		t.Errorf("Expected 14 but returned %d", result)
	}
}

func TestAdd1And2Return3(t *testing.T) {
	result, err := calculator.StringAdd("1,2")
	if err != nil {
		t.Fatal(err)
	}
	if result != 3 {
		t.Errorf("Expected 3 but returned %d", result)
	}
}

func TestAdd1And3Return5(t *testing.T) {
	result, err := calculator.StringAdd("1,4")
	if err != nil {
		t.Fatal(err)
	}
	if result != 5 {
		t.Errorf("Expected 5 but returned %d", result)
	}
}

func TestAdd1To9Return45(t *testing.T) {
	result, err := calculator.StringAdd("1,2,3,4,5,6,7,8,9")
	if err != nil {
		t.Fatal(err)
	}
	if result != 45 {
		t.Errorf("Expected 45 but returned %d", result)
	}
}

func TestAdd1newline2comma3Return6(t *testing.T) {
	result, err := calculator.StringAdd("1\n2,3")
	if err != nil {
		t.Fatal(err)
	}
	if result != 6 {
		t.Errorf("Expected 6 but returned %d", result)
	}
}

func TestAdd1CustomSeparator2Return3(t *testing.T) {
	result, err := calculator.StringAdd("//;\n1;2")
	if err != nil {
		t.Fatal(err)
	}
	if result != 3 {
		t.Errorf("Expected 3 but returned %d", result)
	}
}

func TestAdd1AddNegative2AddNegative3ReturnNegativeNotAllowed(t *testing.T) {
	_, err := calculator.StringAdd("1,-2,-3")
	if err.Error() != "negatives not allowed: -2 -3" {
		t.Errorf("Expected error: negatives not allowed: -2 -3 but returned %f", err)
	}
}

func TestAdd1001Add2Return2(t *testing.T) {
	result, err := calculator.StringAdd("1001,2")
	if err != nil {
		t.Fatal(err)
	}
	if result != 2 {
		t.Errorf("Expected 2 but returned %d", result)
	}
}

func TestAdd1Add2Add3WithArbitraryLengthSeparatorReturn6(t *testing.T) {
	result, err := calculator.StringAdd("//[***]\n1***2***3")
	if err != nil {
		t.Fatal(err)
	}
	if result != 6 {
		t.Errorf("Expected 6 but returned %d", result)
	}
}

func TestAdd1Add2Add3WithMultipleSingleLengthSeparatorReturn6(t *testing.T) {
	result, err := calculator.StringAdd("//[*][%]\n1*2%3")
	if err != nil {
		t.Fatal(err)
	}
	if result != 6 {
		t.Errorf("Expected 6 but returned %d", result)
	}
}
