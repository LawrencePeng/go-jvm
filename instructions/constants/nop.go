package constants

import (
	"go-jvm/instructions/base"
	"go-jvm/rtda"
)

type NOP struct{ base.NoOperandsInstruction }
func (self *NOP) Execute(frame *rtda.Frame) {
}
