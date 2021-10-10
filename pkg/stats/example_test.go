package stats

import (
	"fmt"
	"github.com/dima-5050/bank/v2/pkg/types"
)

func ExampleAvg() {
	
	payment := []types.Payment{
		{
			ID: 1,
			Amount: 1_000_00,
			Category: "Auto",
			Status: "INPROGRESS",
		},
		{
			ID: 2,
			Amount: 2_000_00,
			Category: "Restaurant",
			Status: "OK",
		},
		{
			ID: 3,
			Amount: 3_000_00,
			Category: "cafe",
			Status: "FAIL",
		},
	}

	fmt.Println(Avg(payment))
	//Output: 150000
	
}

func ExampleTotalInCategory() {
	
	
	payment := []types.Payment{
		{
			ID: 1,
			Amount: 1_000_00,
			Category: "Auto",
			Status: "FAIL",
		},
		{
			ID: 2,
			Amount: 2_000_00,
			Category: "Restaurant",
			Status: "OK",
		},
		{
			ID: 3,
			Amount: 3_000_00,
			Category: "cafe",
			Status: "FAIL",

		},
	}

	fmt.Println(TotalInCategory(payment,"cafe"))
	//Output: 0
	
}