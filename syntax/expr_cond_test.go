package syntax

import (
	"testing"
)

func TestEvalCond(t *testing.T) {
	t.Parallel()
	AssertCodesEvalToSameValue(t, `1`, `cond {(1 > 0) : 1, (2 > 3): 2, _:1 + 2,}`)
	AssertCodesEvalToSameValue(t, `1`, `cond {1 > 0 : 1, 2 > 3: 2, _:1 + 2}`)
	AssertCodesEvalToSameValue(t, `2`, `cond {1 > 2 : 1, 2 < 3: 2, _:1 + 2}`)
	AssertCodesEvalToSameValue(t, `3`, `cond {1 > 2 : 1, 2 > 3: 2, _:1 + 2}`)
	AssertCodesEvalToSameValue(t, `1`, `cond {(1 < 2) : 1,}`)
	AssertCodesEvalToSameValue(t, `1`, `cond {(1 < 2) : 1}`)
	AssertCodesEvalToSameValue(t, `2`, `cond {1 > 2 : 1, 2 < 3: 2}`)
	AssertCodesEvalToSameValue(t, `3`, `cond {_ : 1 + 2}`)
	AssertCodesEvalToSameValue(t, `1`, `cond {1 < 2: 1, _ : 1 + 2}`)
	AssertCodesEvalToSameValue(t, `3`, `cond {1 > 2: 1, _ : 1 + 2}`)
	AssertCodesEvalToSameValue(t, `3`, `let a = cond {1 > 2: 1, _ : 1 + 2};a`)
	AssertCodesEvalToSameValue(t, `3`, `let a = cond {1 > 2: 1, _ : 1 + 2,};a`)
	AssertCodesEvalToSameValue(t, `1`, `let a = cond {1 < 2: 1, _ : 1 + 2};a * 1`)
	// // Multiple true conditions
	AssertCodesEvalToSameValue(t, `1`, `cond {1 > 0 : 1, 2 < 3: 2, _:1 + 2}`)
	AssertCodesEvalToSameValue(t, `2`, `cond {1 > 2 : 1, 2 < 3: 2, _:1 + 2}`)
	AssertCodesEvalToSameValue(t, `3`, `cond {1 > 2 : 1, 2 > 3: 2, 3 < 4 :3, _:2 + 2}`)
	AssertCodesEvalToSameValue(t, `3`, `cond {1 > 2 : 1, 2 > 3: 2, 3 < 4 :3, 4 < 5 : 5, _:2 + 2}`)
	AssertCodesEvalToSameValue(t, `3`, `cond {1 > 2 : 1, 2 > 3: 2, 3 < 4 :3, 4 < 5 : 5, 5 > 6 : 6, _:2 + 2}`)
	AssertCodesEvalToSameValue(t, `2`, `cond {1 > 2 : 1, 2 < 3: 2, 3 < 4 :3, 4 < 5 : 5, 5 > 6 : 6, _:2 + 2}`)

	// // Nested call
	AssertCodesEvalToSameValue(t, `1`, `cond {cond {1 > 0 : 1} > 0 : 1, 2 < 3: 2, _:1 + 2}`)
	AssertCodesEvalToSameValue(t, `2`, `cond {cond {1 > 2 : 1, _ : 11} < 2 : 1, 2 < 3: 2, _:1 + 2}`)
	AssertCodesEvalToSameValue(t, `20`, `let a = cond {cond {1 > 2 : 1, _ : 11} < 2 : 1, 2 < 3: 2, _:1 + 2};a * 10`)

	AssertCodesEvalToSameValue(t, `{}`, `cond {1 < 0 : 1, 2 > 3: 2}`)
	AssertCodesEvalToSameValue(t, `{}`, `cond {1 < 0 : 1}`)
	AssertCodesEvalToSameValue(t, `{}`, `cond {}`)
}

// TestEvalCondMulti executes the cases whose condition has multiple expressions.
func TestEvalCondMulti(t *testing.T) {
	t.Parallel()
	AssertCodesEvalToSameValue(t, `1`, `cond {1 > 0 || 3 > 2: 1, 2 > 3: 2, _:1 + 2,}`)
	AssertCodesEvalToSameValue(t, `1`, `cond {0 > 1 || 3 > 2: 1, 2 > 3: 2, _:1 + 2,}`)
	AssertCodesEvalToSameValue(t, `3`, `cond {0 > 1 || 3 > 4: 1, 2 > 3: 2, _:1 + 2,}`)
	AssertCodesEvalToSameValue(t, `1`, `cond {1 > 0 && 3 > 2: 1, 2 > 3: 2, _:1 + 2,}`)
	AssertCodesEvalToSameValue(t, `1`, `cond {(1 > 0 && 3 > 2): 1, (2 > 3) || (1 < 0): 2, _:1 + 2,}`)
	AssertCodesEvalToSameValue(t, `3`, `let a = cond {1 > 2 && 2 > 1: 1, _ : 1 + 2};a`)
	// Multiple true conditions
	AssertCodesEvalToSameValue(t, `1`, `cond {1 > 0 && 3 > 2: 1, 2 > 1: 2, _:1 + 2,}`)
	AssertCodesEvalToSameValue(t, `2`, `cond {(1 > 0 && 3 < 2): 1, (2 > 1) || (1 > 0): 2, _:1 + 2,}`)
	AssertCodesEvalToSameValue(t, `2`, `let a = cond {1 > 2 && 2 > 1: 1, (2 > 1) : 2, _ : 1 + 2};a`)

	AssertCodesEvalToSameValue(t, `{}`, `cond {1 < 0 || 2 > 3 : 1, 2 > 3: 2}`)
	AssertCodesEvalToSameValue(t, `{}`, `cond {1 < 0 || 3 > 4 : 1}`)
}

func TestEvalCondStr(t *testing.T) {
	t.Parallel()
	AssertEvalExprString(t, "{(1>0):1,(2>3):2,_:(1+2)}", "cond {1 > 0 : 1, 2 > 3: 2, _:1 + 2,}")
	AssertEvalExprString(t, "{(1>0):1,(2>3):2,_:(1+2)}", "cond {1 > 0 : 1, 2 > 3: 2, _:1 + 2}")
	AssertEvalExprString(t, "{(1<2):1}", "cond {1 < 2 : 1}")
	AssertEvalExprString(t, "{(1>2):1,(2<3):2}", "cond {1 > 2 : 1, 2 < 3: 2}")
	AssertEvalExprString(t, "{_:(1+2)}", "cond {_: 1 + 2}")
	AssertEvalExprString(t, "{(1<2):1,_:(1+2)}", "cond {1 < 2: 1, _ : 1 + 2}")
}

func TestEvalCondWithControlVar(t *testing.T) {
	t.Parallel()

	// AssertCodesEvalToSameValue(t, `{}`, `let a = 1; a cond {(1 + 2) :1, (2) :2}`)
	// // Control var conditions
	// AssertCodesEvalToSameValue(t, `1`, `let a = 1; a cond {(1) :1}`)
	// AssertCodesEvalToSameValue(t, `1`, `let a = 1; a cond {(1) :1, (2) :2}`)
	// AssertCodesEvalToSameValue(t, `1`, `let a = 1; a cond {(1) :1, (2) :2, _:1 + 2}`)
	// AssertCodesEvalToSameValue(t, `11`, `let a = 1; a cond {(1) :1 + 10, (2) : 2, _:1 + 2}`)
	// AssertCodesEvalToSameValue(t, `3`, `let a = 1; a cond {(2) :2, _:1 + 2}`)
	// AssertCodesEvalToSameValue(t, `3`, `let a = 1; a cond {(2) :2, _:1 + 2}`)
	// AssertCodesEvalToSameValue(t, `13`, `let a = 1; a cond {(2) :2, _:11 + 2}`)
	// AssertCodesEvalToSameValue(t, `5`, `let a = 3; a cond {_:5}`)
	// AssertCodesEvalToSameValue(t, `3`, `let a = 3; a cond {_:1 + 2}`)

	// AssertCodesEvalToSameValue(t, `1`, `let a = 1; let b = a cond {(1) :1, (2) :2, _:1 + 2}; b`)
	// AssertCodesEvalToSameValue(t, `100`, `let a = 1; let b = a cond {(1) :1, (2) :2, _:1 + 2}; b * 100`)
	// // //
	// AssertCodesEvalToSameValue(t, `2`, `let a = 1; (a + 1) cond {(1) :1, (2) :2, _:1 + 2}`)
	// AssertCodesEvalToSameValue(t, `3`, `let a = 1; (a + 10) cond {(1) :1, (2) :2, _:1 + 2}`)
	// AssertCodesEvalToSameValue(t, `2`, `let a = 1; let b = (a + 1) cond {(1) :1, (2) :2, _:1 + 2}; b`)
	// AssertCodesEvalToSameValue(t, `300`, `let a = 1; let b = (a + 10) cond {(1) :1, (2) :2, _:1 + 2}; b * 100`)
	// Nested call
	// AssertCodesEvalToSameValue(t, "B", `let a = 2; a cond {(a cond {(1,2) : 1}): "A", (2, 3): "B", _: "C"}`)
	// AssertCodesEvalToSameValue(t, "A", `let a = 1; a cond { (cond {(2 > 1) : 1}): "A", (2, 3): "B", _: "C"}`)
	// AssertCodesEvalToSameValue(t, "A", `let a = 1; cond { (a cond {(1) : 1}) : "A", (2): "B", _: "C"}`)

	AssertCodesEvalToSameValue(t, `{}`, `let a = 3; a cond {(1) :1, (2) :2 + 1}`)
	AssertCodesEvalToSameValue(t, `{}`, `let a = 3; let b = a cond {(1) :1, (2) :2 + 1}; b`)
	AssertCodesEvalToSameValue(t, `{}`, `let a = 3; let b = (a + 10) cond {(1) :1, (2) :2 + 1}; b`)
}

func TestEvalCondWithControlVarMulti(t *testing.T) {
	AssertCodesEvalToSameValue(t, `1`, `let a = 1; a cond {(1,2) :1}`)
	AssertCodesEvalToSameValue(t, `1`, `let a = 2; a cond {(1,2,3) :1, (2):2}`)
	AssertCodesEvalToSameValue(t, `1`, `let a = 3; a cond {(1,2,3) :1, (2) :2, _:1 + 2}`)
	AssertCodesEvalToSameValue(t, `2`, `let a = 2; a cond {(1) :1 + 10, (2,3) : 2, _:1 + 2}`)

	// AssertCodesEvalToSameValue(t, `med`, `let a = 2;
	// a cond (
	// 	1:"lo",
	// 	(2,3): "med",
	// 	*: "hi")`)

	// var sb strings.Builder
	// assert.Error(t, evalImpl(`let a = 1; a cond ((2,3)) : 2, 3: 3)`, &sb))
	// assert.Error(t, evalImpl(`let a = 1; a cond ((2,3)) : 2, (3,5): 3)`, &sb))
}

// TestEvalCondMultiStr executes the cases whose condition has multiple expressions.
func TestEvalCondMultiStr(t *testing.T) {
	t.Parallel()
	AssertEvalExprString(t, "((control_var:1),{((1>0))&&((2>1)):1})", "(1) cond {(1 > 0 && 2 > 1) : 1}")
	AssertEvalExprString(t, "((control_var:1),{((1>0))||((2>1)):1})", "(1) cond {(1 > 0 || 2 > 1) : 1}")
	AssertEvalExprString(t, "((control_var:1),{((1>0))||((2>1)):1,_:11})", "(1) cond {(1 > 0 || 2 > 1) : 1, _ : 11}")
}

func TestEvalCondWithControlVarStr(t *testing.T) {
	t.Parallel()
	AssertCodesEvalToSameValue(t, "((control_var:1),(1:1))", "(1) cond {(1) : 1}")
	AssertCodesEvalToSameValue(t, "((control_var:1),(1:1))", "(1) cond {(1) : 1,}")
	AssertCodesEvalToSameValue(t, "((control_var:1),(1:1,(2+1):3))", "(1) cond {(1) : 1, (2 + 1) : 3}")
	AssertCodesEvalToSameValue(t, "((control_var:1),(1:1,(2+1):3))", "(1) cond {(1) : 1, (2 + 1) : 3,}")
	AssertCodesEvalToSameValue(t, "((control_var:1),(1:1,(2+1):3,*:4))", "(1) cond {(1) : 1, (2 + 1) : 3, _ : 4}")
	AssertCodesEvalToSameValue(t, "((control_var:1),(1:1,(2+1):3,*:4))", "(1) cond {(1) : 1, (2 + 1) : 3, _ : 4,}")

	AssertCodesEvalToSameValue(t, "(1->(\\a((control_var:a),(1:1))))",
		"let a = 1; a cond (1 : 1)")
	AssertCodesEvalToSameValue(t, "(1->(\\a((control_var:a),((1+2):1,*:(1+2)))))",
		"let a = 1; a cond (1 + 2: 1, * : 1 + 2)")
	AssertCodesEvalToSameValue(t, "(2->(\\a(((control_var:a),((1+2):1,*:(1+2)))->(\\b(b*1)))))",
		"let a = 2; let b = a cond (1 + 2: 1, * : 1 + 2); b * 1")
	AssertCodesEvalToSameValue(t, "(3->(\\a((control_var:(a+2)),((1+2):1,*:(1+2)))))",
		"let a = 3; (a + 2) cond (1 + 2: 1, * : 1 + 2)")
}

// func TestEvalCondWithControlVarMultiStr(t *testing.T) {
// 	t.Parallel()
// 	AssertCodesEvalToSameValue(t, "((control_var:1),([1,2]:1))", "(1) cond ((1,2) :1)")
// 	AssertCodesEvalToSameValue(t, "((control_var:2),(1:(1+10),[2,3]:2,*:(1+2)))", "(2) cond (1 :1 + 10, (2,3) : 2, *:1 + 2)")
// }

func TestEvalCondPatternMatchingWithControlVar(t *testing.T) {
	// ./arrai e "let a = 1; a cond {(type:'hello'):1}"
}