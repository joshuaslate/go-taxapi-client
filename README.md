# TaxiAPI.io Go Client
This is a Go-based wrapper for the free (but rate-limited) [TaxAPI.io](https://taxapi.io) API that allows users to get up-to-date U.S. sales tax and EU country VAT tax rates.

## Usage
```go
package main

import "github.com/joshuaslate/go-taxapi-client"

func main() {
  // To enable caching (the argument passed here is a boolean for whether or not you would like the results to be cached in-memory)
  taxClient := taxapi.NewClient(true)

  // Verify a VAT number
  taxClient.VerifyVATNumber("GB943684002")

  // Get ALL VAT rates
  taxClient.GetAllVATRates()

  // Get VAT rates for a single country
  taxClient.GetVATByCountryCode("ES")

  // Get sales tax rates by U.S. postal code
  taxClient.GetSalesTaxByZip("80521")
}
```

## Caching
By default, results from these API calls are cached in-memory for one day. This is at the request of TaxAPI.io, due to the API being free. It is currently rate-limited at one request per second. If you would like to handle caching on your own, or just throw caution to the wind and not cache the results, pass false to `taxapi.NewClient`.

### Note
Special thanks to [Abs Farah](https://twitter.com/absfarah) for creating this free API for tax information. Note that I am not the creator or maintainer of the API, just this convenience wrapper to access it with. I will try to keep this repository up to date with the API as I can. If you would like to contribute new features, bug fixes, or improvements, please open a pull request or issue. Please direct questions, comments, or concerns regarding the API itself to the API maintainer.