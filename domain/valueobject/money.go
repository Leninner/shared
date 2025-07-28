package valueobject

import (
	"fmt"
	"math"
)

type Money struct {
	Amount   float64
}

func NewMoney(amount float64) *Money {
	return &Money{Amount: amount}
}

func (m *Money) IsGreaterThanZero() bool {
	return m.Amount > 0
}

func (m *Money) IsGreaterThan(other *Money) bool {
	return m.Amount > other.Amount
}

func (m *Money) Add(other *Money) *Money {
	result := m.Amount + other.Amount
	fmt.Println("result", result)
	return &Money{Amount: m.setScale(result)}
}

func (m *Money) setScale(amount float64) float64 {
	return math.Round(amount*100) / 100
}

func (m *Money) Subtract(other *Money) *Money {
	result := m.Amount - other.Amount
	return &Money{Amount: m.setScale(result)}
}

func (m *Money) Multiply(factor int32) *Money {
	result := m.Amount * float64(factor)
	return &Money{Amount: m.setScale(result)}
}

func (m *Money) Equals(other *Money) bool {
	fmt.Println("m.Amount", m.Amount, "other.Amount", other.Amount, "m.Amount == other.Amount", m.Amount == other.Amount)
	return m.Amount == other.Amount
}