package requests

type CreateConfigRequest struct {
	MoneyValue    float32 `json:"money_value" validate:"required,number"`
	MaximumAmount int32   `json:"maximum_amount" validate:"required,number"`
}

type UpdateConfigRequest struct {
	ConfigId      uint    `json:"config_id" validate:"required,number"`
	MaximumAmount int32   `json:"maximum_amount" validate:"required,number"`
}
