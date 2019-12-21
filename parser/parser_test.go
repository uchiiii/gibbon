package parser

import (
	"testing"
	"github.com/uchiiii/gibbon/ast"
	"github.com/uchiiii/gibbon/lexer"
)

func TestLetStatements(t *testing.T){
	input := `
	let x = 5;
	let y = 10;
	let footbar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil{
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3{
		t.Fatalf("program.Statements does not constain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct{
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"footbar"},
	}

	
}
