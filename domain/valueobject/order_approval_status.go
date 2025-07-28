package valueobject

type OrderApprovalStatus string

const (
	OrderApprovalStatusApproved OrderApprovalStatus = "APPROVED"
	OrderApprovalStatusRejected OrderApprovalStatus = "REJECTED"
)

func (o OrderApprovalStatus) String() string {
	return string(o)
}