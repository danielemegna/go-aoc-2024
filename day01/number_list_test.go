package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLength(t *testing.T) {
	var numberList = NumberList{[]int{1, 13, 9}}
	assert.Equal(t, 3, numberList.Length());

	numberList = NumberList{[]int{42, 73, 38, 11}}
	assert.Equal(t, 4, numberList.Length());
}

func TestGetElementAt(t *testing.T) {
	var numberList = NumberList{[]int{42, 73, 38, 11}}
	assert.Equal(t, 42, numberList.At(0));
	assert.Equal(t, 73, numberList.At(1));
	assert.Equal(t, 11, numberList.At(3));
}

func TestSort(t *testing.T) {
	var numberList = NumberList{[]int{42, 73, 38, 11}}
	numberList.Sort()
	assert.Equal(t, 11, numberList.At(0));
	assert.Equal(t, 38, numberList.At(1));
	assert.Equal(t, 73, numberList.At(3));
}

func TestGetOccurenceCount(t *testing.T) {
	var numberList = NumberList{[]int{42, 73, 38, 11, 32, 42, 66, 12, 83, 23, 42, 11}}
	numberList.Sort()
	assert.Equal(t, 3, numberList.OccurenceCount(42));
	assert.Equal(t, 2, numberList.OccurenceCount(11));
	assert.Equal(t, 1, numberList.OccurenceCount(66));
	assert.Equal(t, 0, numberList.OccurenceCount(99));
}
