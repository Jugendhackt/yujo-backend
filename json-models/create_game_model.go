package jsonmodels

type CreateGame struct {
	ID      string `json:"uuid"`
	GamePin uint64 `json:"gamePin"`
}
