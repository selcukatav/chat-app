package model

type Error struct {
	ID             uint   `gorm:"primaryKey" json:"log_id"`
	ResponseCode      int16  `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}

func (Error) TableName() string {
	return "errors"
}
