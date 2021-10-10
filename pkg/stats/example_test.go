package stats

// import (
// 	"fmt"
// 	"github.com/dima-5050/bank/v2/pkg/types"
// )

// func ExampleAvg() {
	
// 	payment := []types.Payment{
// 		{
// 			ID: 1,
// 			Amount: 1_000_00,
// 			Category: "Auto",
// 			Status: "INPROGRESS",
// 		},
// 		{
// 			ID: 2,
// 			Amount: 2_000_00,
// 			Category: "Restaurant",
// 			Status: "OK",
// 		},
// 		{
// 			ID: 3,
// 			Amount: 3_000_00,
// 			Category: "cafe",
// 			Status: "FAIL",
// 		},
// 	}

// 	fmt.Println(Avg(payment))
// 	//Output: 150000
	
// }

// func ExampleTotalInCategory() {
	
	
// 	payment := []types.Payment{
// 		{
// 			ID: 1,
// 			Amount: 1_000_00,
// 			Category: "Auto",
// 			Status: "FAIL",
// 		},
// 		{
// 			ID: 2,
// 			Amount: 2_000_00,
// 			Category: "Restaurant",
// 			Status: "OK",
// 		},
// 		{
// 			ID: 3,
// 			Amount: 3_000_00,
// 			Category: "cafe",
// 			Status: "FAIL",

// 		},
// 	}

// 	fmt.Println(TotalInCategory(payment,"cafe"))
// 	//Output: 0
	

// }

// func ExampleCategoriesAvg() {
// 	payments := []types.Payment{
// 		{ID: 1, Category: "auto", Amount: 1_000_00},
// 		{ID: 2, Category: "food", Amount: 2_000_00},
// 		{ID: 3, Category: "auto", Amount: 3_000_00},
// 		{ID: 4, Category: "auto", Amount: 4_000_00},
// 		{ID: 5, Category: "fun", Amount: 5_000_00},
		
// 	}
	
// 	result := CategoriesAvg(payments)

// 	fmt.Println(result)
// 	//Output: 0

// }