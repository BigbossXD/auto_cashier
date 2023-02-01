package responses

type ErrorMachineResponse struct {
	MachineId uint `json:"machine_id" validate:"required,number"`
}
