package classfile

type DeprecatedAttribute struct { MarkerAttribute }
type SyntheticAttrbute struct { MarkerAttribute }

type MarkerAttribute struct {}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}