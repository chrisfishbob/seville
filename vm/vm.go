package vm

import (
	"fmt"
	"seville/compiler"
	"seville/object"
	"seville/opcode"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions opcode.Instructions

	stack []object.Object
	sp    int // Always points to the next value. Top of the stack is stack[sp - 1]
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,
		stack:        make([]object.Object, StackSize),
		sp:           0,
	}
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := opcode.Opcode(vm.instructions[ip])

		switch op {
		case opcode.OpConstant:
			constIndex := opcode.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		case opcode.OpAdd:
			right := vm.pop()
			left := vm.pop()
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			vm.push(&object.Integer{Value: leftValue + rightValue})
		case opcode.OpSubtract:
			right := vm.pop()
			left := vm.pop()
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			vm.push(&object.Integer{Value: leftValue - rightValue})
		}
	}

	return nil
}

func (vm *VM) push(obj object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow: exceeded stack limit of %d", StackSize)
	}

	vm.stack[vm.sp] = obj
	vm.sp++

	return nil
}

func (vm *VM) pop() object.Object {
	obj := vm.stack[vm.sp-1]
	vm.sp--

	return obj
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}
