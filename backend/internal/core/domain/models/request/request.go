package request

type CreateUser struct {
	Name           string `json:"nombre"`
	DocumentNumber int    `json:"numero_documento"`
	Email          string `json:"correo"`
}

type CreateCampaign struct {
	TaxID              int     `json:"nit_empresa"`
	BranchID           int     `json:"id_sucursal"`
	StartDate          string  `json:"fecha_inicio"`
	EndDate            string  `json:"end_date"`
	PointsMultiplier   float64 `json:"multiplicador_puntos"`
	CashbackMultiplier float64 `json:"multiplicador_cashback"`
	MinPurchaseAmount  float64 `json:"compra_minima"`
}

type CreateBusiness struct {
	RazonSocial      string           `json:"razon_social"`
	NIT              int              `json:"nit"`
	Telefono         int              `json:"telefono"`
	Correo           string           `json:"correo"`
	ConversionFactor ConversionFactor `json:"valor_conversion"`
}

type CreateBranch struct {
	NITEmpresa       int              `json:"nit_empresa"`
	NombreSucursal   string           `json:"nombre_sucursal"`
	ConversionFactor ConversionFactor `json:"valor_conversion"`
}

type ConversionFactor struct {
	MinAmount           float64 `json:"valor_minimo"`
	PointsPerCurrency   float64 `json:"puntos_por_unidad"`
	CashbackPerCurrency float64 `json:"cashback_por_unidad"`
}

type ProcessTransaction struct {
	User     CreateUser `json:"usuario"`
	BranchID int        `json:"id_sucursal"`
	Valor    float64    `json:"valor"`
}
