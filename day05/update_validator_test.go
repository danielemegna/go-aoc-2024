package day05

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestUpdateIsAlwaysValidWithoutRules(t *testing.T) {
	var update = PagesToProduceInTheUpdate{61, 13, 29}
	var rules = PageOrderingRules{}

	var isValid = update.IsValidFor(rules)

	assert.True(t, isValid, fmt.Sprintf("%v should be valid!", update))
}

func TestUpdateIsValidWithRulesUnrelatedToUpdatePages(t *testing.T) {
	var update = PagesToProduceInTheUpdate{61, 13, 29}
	var rules = PageOrderingRules{
		{before: 8, after: 19},
		{before: 10, after: 15},
	}

	var isValid = update.IsValidFor(rules)

	assert.True(t, isValid, fmt.Sprintf("%v should be valid!", update))
}

func TestThreeNumbersInvalidUpdate(t *testing.T) {
	var update = PagesToProduceInTheUpdate{19, 46, 90}
	var setOfPageRules = []PageOrderingRules{
		{{before: 46, after: 19}},
		{{before: 90, after: 46}},
		{{before: 90, after: 19}},
		{
			{before: 1, after: 99},
			{before: 90, after: 46},
			{before: 2, after: 99},
		},
		{
			{before: 1, after: 99},
			{before: 2, after: 99},
			{before: 46, after: 19},
		},
		{
			{before: 19, after: 99},
			{before: 46, after: 99},
			{before: 1, after: 19},
			{before: 1, after: 46},
			{before: 46, after: 19},
		},
	}

	for index, rules := range setOfPageRules {
		t.Run("Rules #"+strconv.Itoa(index+1), func(t *testing.T) {
			var isValid = update.IsValidFor(rules)
			assert.False(t, isValid, fmt.Sprintf("%v should not be valid for %v !", update, rules))
		})
	}
}

func TestThreeNumbersValidUpdate(t *testing.T) {
	var update = PagesToProduceInTheUpdate{19, 46, 90}
	var setOfPageRules = []PageOrderingRules{
		{{before: 19, after: 46}},
		{{before: 46, after: 90}},
		{{before: 19, after: 90}},
		{
			{before: 19, after: 46},
			{before: 46, after: 90},
			{before: 19, after: 90},
		},
	}

	for index, rules := range setOfPageRules {
		t.Run("Rules #"+strconv.Itoa(index+1), func(t *testing.T) {
			var isValid = update.IsValidFor(rules)
			assert.True(t, isValid, fmt.Sprintf("%v should be valid for %v !", update, rules))
		})
	}
}

func TestFiveNumbersInvalidUpdate(t *testing.T) {
	var update = PagesToProduceInTheUpdate{75, 97, 47, 61, 53}
	var setOfPageRules = []PageOrderingRules{
		{{before: 97, after: 75}},
		{{before: 53, after: 61}},
		{
			{before: 47, after: 61},
			{before: 53, after: 61},
		},
	}

	for index, rules := range setOfPageRules {
		t.Run("Rules #"+strconv.Itoa(index+1), func(t *testing.T) {
			var isValid = update.IsValidFor(rules)
			assert.False(t, isValid, fmt.Sprintf("%v should not be valid for %v !", update, rules))
		})
	}
}

func TestFiveNumbersValidUpdate(t *testing.T) {
	var update = PagesToProduceInTheUpdate{75, 97, 47, 61, 53}
	var setOfPageRules = []PageOrderingRules{
		{{before: 75, after: 97}},
		{{before: 75, after: 47}},
		{{before: 75, after: 53}},
		{{before: 61, after: 53}},
		{
			{before: 75, after: 97},
			{before: 75, after: 47},
			{before: 75, after: 53},
			{before: 97, after: 53},
			{before: 47, after: 61},
			{before: 47, after: 53},
		},
	}

	for index, rules := range setOfPageRules {
		t.Run("Rules #"+strconv.Itoa(index+1), func(t *testing.T) {
			var isValid = update.IsValidFor(rules)
			assert.True(t, isValid, fmt.Sprintf("%v should be valid for %v !", update, rules))
		})
	}
}

func TestFixThreeNumbersInvalidUpdateWithCloseSwap(t *testing.T) {
	var update = PagesToProduceInTheUpdate{1, 3, 2}
	var rules = PageOrderingRules{
		{before: 2, after: 3},
	}

	var fixed = update.FixWith(rules)

	assert.Equal(t, PagesToProduceInTheUpdate{1, 2, 3}, fixed)
	assert.Equal(t, PagesToProduceInTheUpdate{1, 3, 2}, update) // original update should not change
}

func TestFixFiveNumbersInvalidUpdateWithCloseSwap(t *testing.T) {
	var update = PagesToProduceInTheUpdate{4, 3, 5, 6, 7}
	var rules = PageOrderingRules{
		{before: 3, after: 4},
	}

	var fixed = update.FixWith(rules)

	assert.Equal(t, PagesToProduceInTheUpdate{3, 4, 5, 6, 7}, fixed)
	assert.Equal(t, PagesToProduceInTheUpdate{4, 3, 5, 6, 7}, update) // original update should not change
}

func TestFixThreeNumbersInvalidUpdateWithFarSwap(t *testing.T) {
	var update = PagesToProduceInTheUpdate{3, 2, 1}
	var rules = PageOrderingRules{
		{before: 1, after: 3},
	}

	var fixed = update.FixWith(rules)

	assert.Equal(t, PagesToProduceInTheUpdate{1, 2, 3}, fixed)
}
