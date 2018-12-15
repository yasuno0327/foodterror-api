package serializer

import "sandbox-api/model"

type Foods struct {
	Foods []model.Food `json:"foods"`
}
