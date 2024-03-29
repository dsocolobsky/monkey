package evaluator

import (
	"fmt"
	"github.com/dsocolobsky/monkey/lexer"
	"github.com/dsocolobsky/monkey/object"
	"github.com/dsocolobsky/monkey/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"--5", 5},
		{"---5", -5},
		{"10 + 10 - 2", 18},
		{"8 - 9", -1},
		{"3 + 2 * 4", 11},
		{"8 * 4 / 2", 16},
		{"3 * (2 + 1) * -1", -9},
		{"(2 + (9 * 2) - -1*(13 * 1 * 5) + 1 + 3 + 1) / 9", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	assert.True(t, ok)
	assert.Equal(t, expected, result.Value)
	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true}, {"false", false},
		{"true == true", true}, {"true == false", false}, {"false == false", true},
		{"true != false", true}, {"true != true", false},
		{"1 < 2", true}, {"1 > 2", false}, {"2 > 1", true},
		{"2 < 1", false}, {"-3 > 1", false}, {"-5 > -7", true},
		{"9 > -15", true},
		{"1 == 1", true}, {"1 != 1", false}, {"1 != 2", true},
		{"-4 == -4", true}, {"-4 == 4", false}, {"-4 != 4", true},
		{"(2 > 4) == false", true}, {"(5 > 1) == true", true}, {"(3 > 1) == (9 > 2)", true},
	}

	for _, tt := range tests {
		fmt.Println(tt.input)
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false}, {"!false", true}, {"!5", false},
		{"!!true", true}, {"!!false", false}, {"!!5", true},
		{"!!!true", false}, {"!!!false", true}, {"!!!5", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	assert.True(t, ok)
	assert.Equal(t, expected, result.Value)
	return true
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}
