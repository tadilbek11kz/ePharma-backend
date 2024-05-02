package product

type CreateProductRequest struct {
	BrandName        string `json:"brand_name" binding:"required"`
	GenericName      string `json:"generic_name" binding:"required"`
	Strength         string `json:"strength" binding:"required"`
	Dosage           string `json:"dosage" binding:"required"`
	DispenceMode     string `json:"dispence_mode" binding:"required"`
	InsurancePlan    string `json:"insurance_plan" binding:"required"`
	PackageSize      string `json:"package_size" binding:"required"`
	ManufacturerName string `json:"manufacturer_name" binding:"required"`
}
