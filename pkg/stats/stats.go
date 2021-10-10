package stats

import (
	"github.com/dima-5050/bank/pkg/types"
)

func Avg(payments []types.Payment) types.Money  {
	var avg types.Money=0
	count:=0
	for _, payment := range payments {
		avg+=payment.Amount
		count++
	}
	return avg/types.Money(count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money  {
	var total types.Money = 0

	for _, payment := range payments {
		if !(payment.Category== category) {
			continue
		}
		total+=payment.Amount
	}
	return total
}