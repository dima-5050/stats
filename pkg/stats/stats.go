package stats

import (
	"fmt"
	"github.com/dima-5050/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var avg types.Money = 0
	count := 0

	for _, payment := range payments {
		if payment.Status==types.StatusFail {
			continue
		}
		avg += payment.Amount
		count++
	}
	return avg / types.Money(count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money = 0

	for _, payment := range payments {
		if !(payment.Category == category) {
			continue
		}

		if payment.Status==types.StatusFail {
			continue
		}

		total += payment.Amount
	}
	return total
}


func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment  {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered=append(filtered, payment)
		}
	}
	return filtered
}


func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money  {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category]+=payment.Amount
	}
	return categories
}

func CategoriesAvg(payments []types.Payment)map[types.Category]types.Money  {
	categories:=map[types.Category]types.Money{}
	count:=map[types.Category]int{}
	coun := 0

	for _, payment := range payments {
		if payment.Status==types.StatusFail {
			continue
		}
		
		categories[payment.Category]+= payment.Amount

		count[payment.Category]+=(coun+1)
		
	}


	for categ := range categories {
		
		categories[categ]/=types.Money(count[categ])

	}
	
	return categories
}


func PeriodsDynamic(
	first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {

	//  count := 0
	 result := map[types.Category]types.Money{}

	 
	for keySecond, valSecond := range second {
		result[keySecond]=valSecond

	}

	for k, v := range first {
		result[k]-=v
	}
	fmt.Println(result)
	return result
}