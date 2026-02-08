package types

import (
	"encoding/json"
	"time"
)

// Bin представляет контейнер для хранения JSON данных
type Bin struct {
	ID        string          `json:"id"`
	Private   bool            `json:"private"`
	CreatedAt time.Time       `json:"created_at"`
	Name      string          `json:"name"`
	Content   json.RawMessage `json:"content,omitempty"` // Для хранения JSON данных
}

// BinList представляет список контейнеров
type BinList struct {
	Bins  []Bin `json:"bins"`
	Total int   `json:"total"`
}

// CreateBinRequest - запрос на создание нового Bin
type CreateBinRequest struct {
	Name    string          `json:"name"`
	Private bool            `json:"private"`
	Content json.RawMessage `json:"content"`
}

// UpdateBinRequest - запрос на обновление Bin
type UpdateBinRequest struct {
	Name    string          `json:"name,omitempty"`
	Private *bool           `json:"private,omitempty"` // Используем указатель для различия false и отсутствия значения
	Content json.RawMessage `json:"content,omitempty"`
}
