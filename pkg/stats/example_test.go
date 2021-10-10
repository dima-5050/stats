package stats

import (
	"fmt"
	"github.com/dima-5050/bank/pkg/types"
)

func ExampleAvg() {
	
	payment := []types.Payment{
		{
			ID: 1,
			Amount: 1_000_00,
			Category: "Auto",
		},
		{
			ID: 2,
			Amount: 2_000_00,
			Category: "Restaurant",
		},
		{
			ID: 3,
			Amount: 3_000_00,
			Category: "cafe",
		},
	}

	fmt.Println(Avg(payment))
	//Output: 200000
	
}

func ExampleTotalInCategory() {
	
	
	payment := []types.Payment{
		{
			ID: 1,
			Amount: 1_000_00,
			Category: "Auto",
		},
		{
			ID: 2,
			Amount: 2_000_00,
			Category: "Restaurant",
		},
		{
			ID: 3,
			Amount: 3_000_00,
			Category: "cafe",
		},
	}

	fmt.Println(TotalInCategory(payment,"cafe"))
	//Output: 300000
	
}