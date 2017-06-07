package rtda

// stack frame
type Frame struct {
	lower        *Frame // stack is implemented as linked list
	thread       *Thread
	localVars    LocalVars
	operandStack *OperandStack
	maxLocals    uint
	maxStack     uint
	nextPC       int // the next instruction after the call
	onPopAction  func()
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getters & setters
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) TopFrame() *Frame {
	return self.lower
}
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
func (self *Frame) SetOnPopAction(f func()) {
	self.onPopAction = f
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}
