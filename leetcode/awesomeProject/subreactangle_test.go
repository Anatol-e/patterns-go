package awesomeProject

import (
	"reflect"
	"testing"
)

const (
	testFirstElement = 1
	testOtherElement = 3
)

var testMatrix = [][]int{
	{1, 2, 1},
	{4, 3, 4},
	{3, 2, 1},
	{1, 1, 1}}

func TestConstructor(t *testing.T) {
	input := testMatrix

	sq := Constructor(input)
	for i := 0; i < len(input); i++ {
		inputSlice := input[i]
		outputSlice := sq.rectangle[i]
		if !reflect.DeepEqual(inputSlice, outputSlice) {
			t.Error("rectangles are not equals")
		}
	}
}

func TestSubrectangleQueries_GetValue(t *testing.T) {
	input := testMatrix

	sq := Constructor(input)
	value := sq.GetValue(0, 0)
	if value != testFirstElement {
		t.Error("invalid first element")
	}

	value = sq.GetValue(1, 1)
	if value != testOtherElement {
		t.Error("invalid first element")
	}
}

func TestSubrectangleQueries_UpdateSubrectangle(t *testing.T) {
	input := testMatrix

	sq := Constructor(input)
	sq.UpdateSubrectangle(0, 0, 3, 2, 5)
	if sq.GetValue(3, 1) != 5  || sq.GetValue(0, 2) != 5{
		t.Error("should be 5")
	}
	sq.UpdateSubrectangle(3, 0, 3, 2, 10)
	if sq.GetValue(3, 1) != 10  || sq.GetValue(0, 2) != 5{
		t.Error("should be 5")
	}
}
