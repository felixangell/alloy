package parser

import (
	"container/list"

	"github.com/ark-lang/ark-go/util"
)

type Type interface {
	TypeName() string
	RawType() Type            // type disregarding pointers
	LevelsOfIndirection() int // number of pointers you have to go through to get to the actual type
	IsIntegerType() bool
	IsFloatingType() bool
}

//go:generate stringer -type=PrimitiveType
type PrimitiveType int

const (
	PRIMITIVE_i8 PrimitiveType = iota
	PRIMITIVE_i16
	PRIMITIVE_i32
	PRIMITIVE_i64
	PRIMITIVE_i128

	PRIMITIVE_u8
	PRIMITIVE_u16
	PRIMITIVE_u32
	PRIMITIVE_u64
	PRIMITIVE_u128

	PRIMITIVE_f32
	PRIMITIVE_f64
	PRIMITIVE_f128

	PRIMITIVE_str
	PRIMITIVE_rune

	PRIMITIVE_int
	PRIMITIVE_uint

	PRIMITIVE_bool
)

func (v PrimitiveType) IsIntegerType() bool {
	switch v {
	case PRIMITIVE_i8, PRIMITIVE_i16, PRIMITIVE_i32, PRIMITIVE_i64, PRIMITIVE_i128,
		PRIMITIVE_u8, PRIMITIVE_u16, PRIMITIVE_u32, PRIMITIVE_u64, PRIMITIVE_u128,
		PRIMITIVE_int, PRIMITIVE_uint:
		return true
	default:
		return false
	}
}

func (v PrimitiveType) IsFloatingType() bool {
	switch v {
	case PRIMITIVE_f32, PRIMITIVE_f64, PRIMITIVE_f128:
		return true
	default:
		return false
	}
}

func (v PrimitiveType) TypeName() string {
	return v.String()[10:]
}

func (v PrimitiveType) RawType() Type {
	return v
}

func (v PrimitiveType) LevelsOfIndirection() int {
	return 0
}

// StructType

type StructType struct {
	Name  string
	Items list.List
	Attrs []*Attr
}

func (v *StructType) String() string {
	result := "(" + util.Blue("StructType") + ": "
	for _, attr := range v.Attrs {
		result += attr.String() + " "
	}
	result += v.Name + "\n"
	for item := v.Items.Front(); item != nil; item = item.Next() {
		result += "\t" + item.Value.(*VariableDecl).String() + "\n"
	}
	return result + ")"
}

func (v *StructType) TypeName() string {
	return v.Name
}

func (v *StructType) RawType() Type {
	return v
}

func (v *StructType) LevelsOfIndirection() int {
	return 0
}

func (v *StructType) IsIntegerType() bool {
	return false
}

func (v *StructType) IsFloatingType() bool {
	return false
}

// PointerType

type PointerType struct {
	Addressee Type
}

func pointerTo(t Type) PointerType {
	return PointerType{Addressee: t}
}

func (v PointerType) TypeName() string {
	return "^" + v.Addressee.TypeName()
}

func (v PointerType) RawType() Type {
	return v.Addressee.RawType()
}

func (v PointerType) LevelsOfIndirection() int {
	return v.Addressee.LevelsOfIndirection() + 1
}

func (v PointerType) IsIntegerType() bool {
	return false
}

func (v PointerType) IsFloatingType() bool {
	return false
}
