package requests

type DepositRequest struct {
	MachineId uint                 `json:"machine_id" validate:"required,number"`
	Items     []depositItemRequest `json:"items" validate:"required"`
}
type depositItemRequest struct {
	ID     uint  `json:"config_id" validate:"required,number"`
	Amount int32 `json:"amount" validate:"required,number"`
}
