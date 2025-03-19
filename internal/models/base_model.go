package models

import "time"

type BaseModel struct {
	ID        uint      `db:"id" json:"id"`
	PartnerID uint      `db:"partner_id" json:"partner_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
