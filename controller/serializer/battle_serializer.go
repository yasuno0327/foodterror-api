package serializer

type BattleResult struct {
	Result string `json:"result"`
}

type BattleRequest struct {
	OpponentEmail string `json:"opponent_email"`
	BattleType    string `json:"battle_type"`
}
