package requests

type CreateMachineRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateMachineRequest struct {
	MachineId uint   `json:"machine_id" validate:"required,number"`
	Name      string `json:"name" validate:"required"`
}
