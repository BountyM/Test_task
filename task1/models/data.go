package models

type Data struct {
	Data []int64 `json:"data" binding:"required"`
}
