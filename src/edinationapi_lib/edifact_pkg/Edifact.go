/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package edifact_pkg

import "edinationapi_lib/configuration_pkg"
import "edinationapi_lib/models_pkg"

/*
 * Interface for the EDIFACT_IMPL
 */
type EDIFACT interface {
    Read ([]byte, *bool, *bool, *string, *string, *bool) ([]*models_pkg.EdifactInterchange, error)

    Write (*bool, *string, *string, *bool, *bool, *string, *models_pkg.EdifactInterchange) ([]byte, error)

    Validate (*string, *bool, *bool, *string, *bool, *string, *models_pkg.EdifactInterchange) (*models_pkg.OperationResult, error)

    Ack (*string, *bool, *int64, *int64, *bool, *bool, *string, *bool, *string, *models_pkg.EdifactInterchange) ([]*models_pkg.EdifactInterchange, error)
}

/*
 * Factory for the EDIFACT interaface returning EDIFACT_IMPL
 */
func NewEDIFACT(config configuration_pkg.CONFIGURATION) *EDIFACT_IMPL {
    client := new(EDIFACT_IMPL)
    client.config = config
    return client
}