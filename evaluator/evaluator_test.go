package evaluator

import (
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
		{"5", 2},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	assert.True(t, ok)
	assert.Equal(t, result.Value, expected)
	return true
}
