package requests

type WithdrawRequest struct {
	MachineId uint                  `json:"machine_id" validate:"required,number"`
	Items     []withdrawItemRequest `json:"items" validate:"required"`
}
type withdrawItemRequest struct {
	ID     uint  `json:"config_id" validate:"required,number"`
	Amount int32 `json:"amount" validate:"required,number"`
}
