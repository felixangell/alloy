package parser

import (
	"container/list"
	"fmt"

	"github.com/ark-lang/ark-go/util"
)

type Node interface {
	String() string
	analyze(*semanticAnalyzer)
}

type Stat interface {
	Node
	statNode()
}

type Expr interface {
	Node
	exprNode()
	GetType() Type
}

type Decl interface {
	Node
	declNode()
}

type Variable struct {
	Type    Type
	Name    string
	Mutable bool
	Attrs   []*Attr
}

func (v *Variable) String() string {
	result := "(" + util.Blue("Variable") + ": "
	if v.Mutable {
		result += util.Green("[mutable] ")
	}
	for _, attr := range v.Attrs {
		result += attr.String() + " "
	}
	return result + v.Name + " " + util.Green(v.Type.TypeName()) + ")"
}

type Function struct {
	Name       string
	Parameters *List
	Type       Type
	Mutable    bool
	Attrs      []*Attr
	Body       *Block
}

func (v *Function) String() string {
	result := "(" + util.Blue("Function") + ": "
	if v.Mutable {
		result += util.Green("[mutable] ")
	}
	for _, attr := range v.Attrs {
		result += attr.String() + " "
	}
	result += v.Name + " " + v.Parameters.String() + ": " + util.Green(v.Type.TypeName()) + " "
	if v.Body != nil {
		result += v.Body.String()
	}
	return result + ")"
}

//
// Nodes
//

type Block struct {
	Nodes []Node
}

func newBlock() *Block {
	return &Block{Nodes: make([]Node, 0)}
}

func (v *Block) String() string {
	if len(v.Nodes) == 0 {
		return "(" + util.Blue("Block") + ": )"
	}

	result := "(" + util.Blue("Block") + ":\n"
	for _, n := range v.Nodes {
		result += "\t" + n.String() + "\n"
	}
	return result + ")"
}

func (v *Block) appendNode(n Node) {
	v.Nodes = append(v.Nodes, n)
}

/**
 * Declarations
 */

// VariableDecl

type VariableDecl struct {
	Variable   *Variable
	Assignment Expr
}

func (v *VariableDecl) declNode() {}

func (v *VariableDecl) String() string {
	if v.Assignment == nil {
		return "(" + util.Blue("VariableDecl") + ": " + v.Variable.String() + ")"
	} else {
		return "(" + util.Blue("VariableDecl") + ": " + v.Variable.String() +
			" = " + v.Assignment.String() + ")"
	}
}

// StructDecl

type StructDecl struct {
	Struct *StructType
}

func (v *StructDecl) declNode() {}

func (v *StructDecl) String() string {
	return "(" + util.Blue("StructDecl") + ": " + v.Struct.String() + ")"
}

// FunctionDecl

type FunctionDecl struct {
	Function *Function
}

func (v *FunctionDecl) declNode() {}

func (v *FunctionDecl) String() string {
	return "(" + util.Blue("FunctionDecl") + ": " + v.Function.String() + ")"
}

/**
 * Statements
 */

type ReturnStat struct {
	Value Expr
}

func (v *ReturnStat) statNode() {}

func (v *ReturnStat) String() string {
	return "(" + util.Blue("ReturnStat") + ": " +
		v.Value.String() + ")"
}

/**
 * Expressions
 */

// RuneLiteral

type RuneLiteral struct {
	Value rune
}

func (v *RuneLiteral) exprNode() {}

func (v *RuneLiteral) String() string {
	return fmt.Sprintf("("+util.Blue("RuneLiteral")+": "+util.Yellow("%c")+")", v.Value)
}

func (v *RuneLiteral) GetType() Type {
	return PRIMITIVE_rune
}

// IntegerLiteral

type IntegerLiteral struct {
	Value uint64
}

func (v *IntegerLiteral) exprNode() {}

func (v *IntegerLiteral) String() string {
	return fmt.Sprintf("("+util.Blue("IntegerLiteral")+": "+util.Yellow("%d")+")", v.Value)
}

func (v *IntegerLiteral) GetType() Type {
	return PRIMITIVE_int
}

// FloatingLiteral

type FloatingLiteral struct {
	Value float64
}

func (v *FloatingLiteral) exprNode() {}

func (v *FloatingLiteral) String() string {
	return fmt.Sprintf("("+util.Blue("FloatingLiteral")+": "+util.Yellow("%f")+")", v.Value)
}

func (v *FloatingLiteral) GetType() Type {
	return PRIMITIVE_f64
}

// StringLiteral

type StringLiteral struct {
	Value string
}

func (v *StringLiteral) exprNode() {}

func (v *StringLiteral) String() string {
	return "(" + util.Blue("StringLiteral") + ": " + util.Yellow(v.Value) + ")"
}

func (v *StringLiteral) GetType() Type {
	return PRIMITIVE_str
}

// BinaryExpr

type BinaryExpr struct {
	Lhand, Rhand Expr
	Op           BinOpType
	Type         Type
}

func (v *BinaryExpr) exprNode() {}

func (v *BinaryExpr) String() string {
	return "(" + util.Blue("BinaryExpr") + ": " + v.Lhand.String() + " " +
		v.Op.String() + " " +
		v.Rhand.String() + ")"
}

func (v *BinaryExpr) GetType() Type {
	return v.Type
}

// UnaryExpr

type UnaryExpr struct {
	Expr Expr
	Op   UnOpType
	Type Type
}

func (v *UnaryExpr) exprNode() {}

func (v *UnaryExpr) String() string {
	return "(" + util.Blue("UnaryExpr") + ": " +
		v.Op.String() + " " + v.Expr.String() + ")"
}

func (v *UnaryExpr) GetType() Type {
	return v.Type
}

/**
 * Other
 */

// List

type List struct {
	Items list.List
}

func (v *List) listNode() {}

func (v *List) String() string {
	var result = "(" + util.Blue("List") + ": "
	for item := v.Items.Front(); item != nil; item = item.Next() {
		result += item.Value.(*VariableDecl).String()
		if item.Next() != nil {
			result += " "
		}
	}
	result += ")"
	return result
}
