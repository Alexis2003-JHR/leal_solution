package request

type CreateUser struct {
	Name  string `json:"nombre"`
	Email string `json:"correo"`
}

type CreateCampaign struct {
	BranchID           int     `json:"id_sucursal"`
	StartDate          string  `json:"fecha_inicio"`
	EndDate            string  `json:"end_date"`
	PointsMultiplier   float64 `json:"multiplicador_puntos"`
	CashbackMultiplier float64 `json:"multiplicador_cashback"`
	MinPurchaseAmount  float64 `json:"compra_minima"`
}

type CreateBusiness struct {
	RazonSocial string `json:"razon_social"`
	NIT         int    `json:"nit"`
	Telefono    int    `json:"telefono"`
	Correo      string `json:"correo"`
}

type CreateBranch struct {
	NITEmpresa     int    `json:"nit_empresa"`
	NombreSucursal string `json:"nombre_sucursal"`
}
