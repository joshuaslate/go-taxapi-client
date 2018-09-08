package taxapi

import "net/http"

// VATVerificationResponse is the response from vat?vat_number={vatNumber}
type VATVerificationResponse struct {
	Status         string `json:"status"`
	Valid          bool   `json:"valid"`
	FormatValid    bool   `json:"format_valid"`
	Query          bool   `json:"query"`
	CountryCode    string `json:"country_code"`
	VatNumber      string `json:"vat_number"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
}

// VerifyVATNumber will check that a VAT number is valid
func (c *Client) VerifyVATNumber(vatNumber string) (VATVerificationResponse, error) {
	result := new(VATVerificationResponse)
	vatQuery := []string{"vat_number", vatNumber}
	_, err := c.makeRequest("/vat", http.MethodGet, [][]string{vatQuery}, result)

	return *result, err
}
