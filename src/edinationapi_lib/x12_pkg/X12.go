/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package x12_pkg

import "edinationapi_lib/configuration_pkg"
import "edinationapi_lib/models_pkg"

/*
 * Interface for the X12_IMPL
 */
type X12 interface {
    Read ([]byte, *bool, *bool, *string, *string) ([]*models_pkg.X12Interchange, error)

    Write (*bool, *string, *string, *string, *models_pkg.X12Interchange) ([]byte, error)

    Validate (*bool, *string, *bool, *bool, *string, *models_pkg.X12Interchange) (*models_pkg.OperationResult, error)

    Ack (*bool, *string, *bool, *int64, *int64, *bool, *bool, *string, *string, *bool, *string, *models_pkg.X12Interchange) ([]*models_pkg.X12Interchange, error)
}

/*
 * Factory for the X12 interaface returning X12_IMPL
 */
func NewX12(config configuration_pkg.CONFIGURATION) *X12_IMPL {
    client := new(X12_IMPL)
    client.config = config
    return client
}