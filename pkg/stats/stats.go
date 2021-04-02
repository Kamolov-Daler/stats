package stats

import (
	"github.com/Kamolov-Daler/bank/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	result := types.Money(0)

	for _, papayment := range payments {
		result += papayment.Amount
	}
	result = result / types.Money(len(payments))
	return result
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	categoryResult := types.Money(0)

	for _, payment := range payments {
		if category == payment.Category {
			categoryResult += payment.Amount
		}
	}
	return categoryResult
}
