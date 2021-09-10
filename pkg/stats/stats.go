package stats

import (
	"github.com/ikrom97/bank/v2/pkg/types"
)
// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	var totalAmount types.Money

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			totalAmount += payment.Amount
		}
	}

	avgAmount := totalAmount / types.Money(len(payments))

	return avgAmount
}
// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var totalAmount types.Money

	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			totalAmount += payment.Amount
		}
	}

	return totalAmount
}

