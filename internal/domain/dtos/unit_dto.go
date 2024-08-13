package dtos

type UnitDTO struct {
	UnitId      uint   `json:"unit_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Sequence    int    `json:"sequence"`
}
