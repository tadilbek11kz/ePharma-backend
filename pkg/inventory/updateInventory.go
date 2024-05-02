package inventory

type UpdateInventoryRequest struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
