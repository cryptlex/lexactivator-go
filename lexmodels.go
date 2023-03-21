package lexactivator
type OrganizationAddress struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City		 string `json:"city"`
	State		 string `json:"state"`
	Country 	 string `json:"country"`
	PostalCode 	 string `json:"postalCode"`
}
