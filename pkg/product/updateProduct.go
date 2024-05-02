package product

type UpdateProductRequest struct {
	BrandName        string `json:"brand_name"`
	GenericName      string `json:"generic_name"`
	Strength         string `json:"strength"`
	Dosage           string `json:"dosage"`
	DispenceMode     string `json:"dispence_mode"`
	InsurancePlan    string `json:"insurance_plan"`
	PackageSize      string `json:"package_size"`
	ManufacturerName string `json:"manufacturer_name"`
}
