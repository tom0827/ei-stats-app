package models

type StatsInput struct {
	SoulEggs     float64 `json:"soulEggs"`
	ProphecyEggs uint64  `json:"prophecyEggs"`
	Prestiges    uint64  `json:"prestiges"`
}
