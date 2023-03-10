package requests

type ReceiveRequest struct {
	MachineId uint                 `json:"machine_id" validate:"required,number"`
	Price     float32              `json:"price" validate:"required"`
	Items     []receiveItemRequest `json:"items" validate:"required"`
}

type receiveItemRequest struct {
	ID     uint  `json:"config_id" validate:"required,number"`
	Amount int32 `json:"amount" validate:"required,number"`
}
