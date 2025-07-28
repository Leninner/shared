package valueobject

type OrderStatus string

const (
	OrderStatusPending OrderStatus = "PENDING"
	OrderStatusPaid    OrderStatus = "PAID"
	OrderStatusApproved OrderStatus = "APPROVED"
	OrderStatusCancelling OrderStatus = "CANCELLING"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)
