/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package EdiNationAPIClient

import(
	"edinationapi_lib/configuration_pkg"
	"edinationapi_lib/edifact_pkg"
	"edinationapi_lib/edimodel_pkg"
	"edinationapi_lib/x12_pkg"
)


/*
 * Interface for the EDINATIONAPI_IMPL
 */
type EDINATIONAPI interface {
    Edifact()               edifact_pkg.EDIFACT
    EdiModel()              edimodel_pkg.EDIMODEL
    X12()                   x12_pkg.X12
    Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the EDINATIONAPI interface returning EDINATIONAPI_IMPL
 */
func NewEDINATIONAPI() EDINATIONAPI {
    ediNationAPIClient := new(EDINATIONAPI_IMPL)
    ediNationAPIClient.config = configuration_pkg.NewCONFIGURATION()

    return ediNationAPIClient
}
