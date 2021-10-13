package stats

import (

	"testing"

	"github.com/dima-5050/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}

	result := CategoriesAvg(payments)

	if len(result) == 0 {
		t.Error("result len !=0")
	}
}

func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
			"auto": 100,
			"food": 20,
			"note":100,
	}

	second := map[types.Category]types.Money{
		"auto": 50,
		"food": 10,
		"mobile":333,
}
	result := PeriodsDynamic(first, second)
	
	
	if len(result) == 0 {
		t.Error("result len !=0")
	}
}
