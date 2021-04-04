package stats

import (
	"reflect"
	"testing"

	"github.com/Kamolov-Daler/bank/v2/pkg/types"
)

func TestCategoriesAvg_Nil(t *testing.T) {
	var payments []types.Payment
	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestCategoriesAvg_Empty(t *testing.T) {
	payments := []types.Payment{}
	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 500_00},
		{ID: 2, Category: "food", Amount: 1_000_00},
		{ID: 3, Category: "auto", Amount: 500_00},
		{ID: 4, Category: "auto", Amount: 500_00},
		{ID: 5, Category: "fun", Amount: 1_500_00},
	}

	expected := map[types.Category]types.Money{
		"auto": 500_00,
		"food": 1_000_00,
		"fun":  1_500_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}

	second := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}

	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
