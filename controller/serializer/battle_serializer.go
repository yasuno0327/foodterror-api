package serializer

type BattleResult struct {
	Result string `json:"result"`
}

type BattleRequest struct {
	OpponenEmail string `json:"oponent_email"`
	BattleType   string `json:"battle_type"`
}
