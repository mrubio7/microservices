package webhooks

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// AdditionalData define un array de arrays de estructuras
type AdditionalData [][]struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Implementación de los métodos para manejar AdditionalData como jsonb
func (a AdditionalData) Value() (driver.Value, error) {
	// Serializar el array de arrays como JSON
	return json.Marshal(a)
}

func (a *AdditionalData) Scan(value interface{}) error {
	if value == nil {
		*a = AdditionalData{} // Inicializa como un array vacío si es nulo
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, a)
}

// Estructura principal
type AllstarClipProcessed struct {
	Event             string         `json:"event"`
	ID                string         `json:"_id"`
	ClipURL           string         `json:"clipUrl"`
	Username          string         `json:"username"`
	DemoURL           string         `json:"demoUrl"`
	RoundNumber       int            `json:"roundNumber"`
	SteamID           string         `json:"steamid"`
	ClipLength        float64        `json:"clipLength"`
	Status            string         `json:"status"`
	ClipTitle         string         `json:"clipTitle"`
	ShareID           string         `json:"shareId"`
	CreatedDate       string         `json:"createdDate"`
	Updated           string         `json:"updated"`
	ClipSnapshotURL   string         `json:"clipSnapshotURL"`
	ClipImageThumbURL string         `json:"clipImageThumbURL"`
	RequestID         string         `json:"requestId"`
	AdditionalData    AdditionalData `json:"additionalData" gorm:"type:jsonb"`
}
