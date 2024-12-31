package day05

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
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
		{{before: 90, after: 19}},
	}

	for index, rules := range setOfPageRules {
		t.Run("Rules #"+strconv.Itoa(index+1), func(t *testing.T) {
			var isValid = update.IsValidFor(rules)
			assert.False(t, isValid, fmt.Sprintf("%v should not be valid for %v !", update, rules))
		})
	}

}
