package models

type Campaign struct {
	ID                 int     `json:"id"`
	BranchID           int     `json:"sucursal_id"`
	StartDate          string  `json:"fecha_inicio"`
	EndDate            string  `json:"fecha_fin"`
	PointsMultiplier   float64 `json:"multiplicador_puntos"`
	CashbackMultiplier float64 `json:"multiplicador_cashback"`
	MinPurchaseAmount  float64 `json:"valor_minimo"`
}
