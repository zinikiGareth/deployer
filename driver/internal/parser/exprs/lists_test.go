package exprs_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

func TestAnEmptyListIsParsed(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := &driverbottom.CoreTools{Reporter: reporter}
	lx := lexicator.NewLineLexicator(tools, "test")
	p := exprs.NewExprParser(tools)

	f := errorsink.InFile("test")
	l := f.AtLine(1, 0, "[]")

	tokens := lx.BlockedLine(l)
	expr, ok := p.Parse(tokens)
	if !ok {
		t.Fatalf("parsing failed")
	}

	le, ok := expr.(*exprs.ListExpr)
	if !ok {
		t.Fatalf("Expr was not a list: %T", expr)
	}

	if le.Length() != 0 {
		t.Fatalf("Expected list of length 0, not %d", le.Length())
	}

}

func TestAnEmptyListIsParsedInParens(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := &driverbottom.CoreTools{Reporter: reporter}
	lx := lexicator.NewLineLexicator(tools, "test")
	p := exprs.NewExprParser(tools)

	f := errorsink.InFile("test")
	l := f.AtLine(1, 0, "([])")

	tokens := lx.BlockedLine(l)
	expr, ok := p.Parse(tokens)
	if !ok {
		t.Fatalf("parsing failed")
	}

	le, ok := expr.(*exprs.ListExpr)
	if !ok {
		t.Fatalf("Expr was not a list: %T", expr)
	}

	if le.Length() != 0 {
		t.Fatalf("Expected list of length 0, not %d", le.Length())
	}
}

func TestAnEmptyListIsParsedAsAnArgument(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	recall = myRecall{things: make(map[string]any)}
	tools := &driverbottom.CoreTools{Reporter: reporter, Recall: recall}
	lx := lexicator.NewLineLexicator(tools, "test")
	p := exprs.NewExprParser(tools)

	recall.things["sum"] = konstFunc

	f := errorsink.InFile("test")
	l := f.AtLine(1, 0, "sum []")

	tokens := lx.BlockedLine(l)
	expr, ok := p.Parse(tokens)
	if !ok {
		t.Fatalf("parsing failed")
	}

	ae, ok := expr.(*exprs.Apply)
	if !ok {
		t.Fatalf("Expr was not a list: %T", expr)
	}

	if len(ae.Args) != 1 {
		t.Fatalf("Expected list of length 0, not %d", len(ae.Args))
	}

	le, ok := ae.Args[0].(*exprs.ListExpr)
	if !ok {
		t.Fatalf("Expr was not a list: %T", expr)
	}

	if le.Length() != 0 {
		t.Fatalf("Expected list of length 0, not %d", le.Length())
	}
}

func TestAnSingletonListIsParsed(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := &driverbottom.CoreTools{Reporter: reporter}
	lx := lexicator.NewLineLexicator(tools, "test")
	p := exprs.NewExprParser(tools)

	f := errorsink.InFile("test")
	l := f.AtLine(1, 0, "[3]")

	tokens := lx.BlockedLine(l)
	expr, ok := p.Parse(tokens)
	if !ok {
		t.Fatalf("parsing failed")
	}

	le, ok := expr.(*exprs.ListExpr)
	if !ok {
		t.Fatalf("Expr was not a list: %T", expr)
	}

	if le.Length() != 1 {
		t.Fatalf("Expected list of length 1, not %d", le.Length())
	}

}

func TestADoubleListIsParsed(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := &driverbottom.CoreTools{Reporter: reporter}
	lx := lexicator.NewLineLexicator(tools, "test")
	p := exprs.NewExprParser(tools)

	f := errorsink.InFile("test")
	l := f.AtLine(1, 0, "[3, 8]")

	tokens := lx.BlockedLine(l)
	expr, ok := p.Parse(tokens)
	if !ok {
		t.Fatalf("parsing failed")
	}

	le, ok := expr.(*exprs.ListExpr)
	if !ok {
		t.Fatalf("Expr was not a list: %T", expr)
	}

	if le.Length() != 2 {
		t.Fatalf("Expected list of length 2, not %d", le.Length())
	}

}
