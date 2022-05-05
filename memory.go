package memoryhttp

import (
	"fmt"
	"reflect"
	"unsafe"
)

type MemoryInfo struct {
	VarName     string      `json:"var_name"`
	VarType     string      `json:"var_type"`
	VarTypeKind string      `json:"var_type_kind"`
	VarSize     string      `json:"var_size"`
	Var         interface{} `json:"var"`
}

var (
	varMap   = make(map[string]*MemoryInfo, 0)
	varNames = make([]string, 0)
)

func Watch(name string, arg interface{}) {
	if arg == nil || name == "" {
		return
	}

	t := reflect.TypeOf(arg)

	varMap[name] = &MemoryInfo{
		VarName:     name,
		VarType:     t.String(),
		VarTypeKind: t.Kind().String(),
		VarSize:     fmt.Sprint(unsafe.Sizeof(arg)),
		Var:         arg,
	}

	fmt.Printf("string:%s, name:%s, pkgPath:%s", t.String(), t.Name(), t.PkgPath())
}
