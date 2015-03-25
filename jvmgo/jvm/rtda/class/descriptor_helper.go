package class

import (
	"github.com/zxh0/jvm.go/jvmgo/jtype"
	"strings"
)

func calcArgSlotCount(descriptor string) uint {
	md := parseMethodDescriptor(descriptor)
	slotCount := md.argCount()
	for _, paramType := range md.ParameterTypes() {
		if paramType.IsLongOrDouble() {
			slotCount++
		}
	}
	return slotCount
}

// [XXX -> [XXX
// LXXX; -> XXX
// I -> int ...
func getClassName(descriptor string) string {
	switch descriptor[0] {
	case '[':
		return descriptor // array
	case 'L':
		return descriptor[1 : len(descriptor)-1] // object
	default:
		return jtype.GetPrimitiveType(descriptor) // primirive types
	}
}

func GetReturnDescriptor(methodDescriptor string) string {
	start := strings.Index(methodDescriptor, ")") + 1
	return methodDescriptor[start:]
}
