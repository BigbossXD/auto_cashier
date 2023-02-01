package responses

type ChangeRequest struct {
	TotalChange float32             `json:"totalChange" `
	Items       []ChangeItemRequest `json:"items" `
}
type ChangeItemRequest struct {
	ID         uint    `json:"config_id" `
	MoneyValue float32 `json:"money_value"`
	Amount     int32   `json:"amount" `
}
