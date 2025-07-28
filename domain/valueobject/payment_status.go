package valueobject

type PaymentStatus string

const (
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusFailed    PaymentStatus = "FAILED"
	PaymentStatusCancelled PaymentStatus = "CANCELLED"
	PaymentStatusPending   PaymentStatus = "PENDING"
)

func (p PaymentStatus) String() string {
	return string(p)
}