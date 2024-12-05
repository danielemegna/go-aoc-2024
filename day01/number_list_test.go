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

func TestSimilarityScoreIsZeroWithoutCommonNumbers(t *testing.T) {
	var firstList = NumberList{[]int{19}}
	var secondList = NumberList{[]int{42}}
	
	assert.Equal(t, 0, firstList.SimilarityScoreWith(secondList));
	assert.Equal(t, 0, secondList.SimilarityScoreWith(firstList));
}

func TestSimilarityScoreWithSameSingleNumberIsTheNumberValue(t *testing.T) {
	var firstList = NumberList{[]int{19}}
	var secondList = NumberList{[]int{19}}
	
	assert.Equal(t, 19, firstList.SimilarityScoreWith(secondList));
	assert.Equal(t, 19, secondList.SimilarityScoreWith(firstList));
}

func TestSimilarityScoreWithSingleCommonNumberIsTheNumberValue(t *testing.T) {
	var firstList = NumberList{[]int{19, 42, 11}}
	var secondList = NumberList{[]int{63, 84, 19}}
	
	assert.Equal(t, 19, firstList.SimilarityScoreWith(secondList));
	assert.Equal(t, 19, secondList.SimilarityScoreWith(firstList));
}

func TestSimilarityScoreIsTheSumOfCommonElements(t *testing.T) {
	var firstList = NumberList{[]int{19, 42, 63, 11, 69}}
	var secondList = NumberList{[]int{63, 84, 19, 72, 88}}
	
	assert.Equal(t, 19 + 63, firstList.SimilarityScoreWith(secondList));
	assert.Equal(t, 19 + 63, secondList.SimilarityScoreWith(firstList));
}

func TestSimilarityScoreShouldCountMultipleCommonElements(t *testing.T) {
	var firstList = NumberList{[]int{19, 42, 63, 11, 69}}
	var secondList = NumberList{[]int{63, 84, 19, 72, 19}}
	
	assert.Equal(t, (19 * 2) + (63 * 1), firstList.SimilarityScoreWith(secondList));
	assert.Equal(t, (19 * 2) + (63 * 1), secondList.SimilarityScoreWith(firstList));
}
