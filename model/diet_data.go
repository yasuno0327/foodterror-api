package model

import "time"

type DietData struct {
	UserID    int
	Burned    float32 // Burned calories
	Intake    float32 // Intake calories
	CreatedAt time.Time
}
