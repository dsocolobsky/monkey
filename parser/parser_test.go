package parser

import (
	"testing"

	"github.com/dysoco/monkey/ast"
	"github.com/dysoco/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 81921024;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.parseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements but got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		st := program.Statements[i]
		if !testLetStatement(t, st, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Tokenliteral not 'let' got=%q", s.TokenLiteral())
	}

	st, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}

	if st.Name.Value != name {
		t.Errorf("st name value not %s got=%s", name, st.Name.Value)
		return false
	}

	return true
}
