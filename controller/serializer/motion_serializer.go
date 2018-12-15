package serializer

import "sandbox-api/model"

type Motions struct {
	Motions []model.Motion `json:"motions"`
}
