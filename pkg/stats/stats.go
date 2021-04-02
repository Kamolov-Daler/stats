package stats

import (
	"github.com/Kamolov-Daler/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	result := types.Money(0)
	iter := types.Money(0)

	for _, papayment := range payments {
		if papayment.Status != types.StatusFail {
			iter++
			result += papayment.Amount
		}
	}
	result = result / iter
	return result
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	categoryResult := types.Money(0)

	for _, payment := range payments {
		if category == payment.Category {
			if payment.Status != types.StatusFail {
				categoryResult += payment.Amount
			}
		}
	}
	return categoryResult
}
