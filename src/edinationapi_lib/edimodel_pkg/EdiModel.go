/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package edimodel_pkg

import "edinationapi_lib/configuration_pkg"
import "edinationapi_lib/models_pkg"

/*
 * Interface for the EDIMODEL_IMPL
 */
type EDIMODEL interface {
    Upload ([]byte) (error)

    GetModels (*bool) ([]*models_pkg.EdiModel, error)

    GetModel (string, *bool) ([]byte, error)

    Delete (string) (error)
}

/*
 * Factory for the EDIMODEL interaface returning EDIMODEL_IMPL
 */
func NewEDIMODEL(config configuration_pkg.CONFIGURATION) *EDIMODEL_IMPL {
    client := new(EDIMODEL_IMPL)
    client.config = config
    return client
}