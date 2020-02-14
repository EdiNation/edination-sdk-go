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
 * Client structure as interface implementation
 */
type EDINATIONAPI_IMPL struct {
     edifact edifact_pkg.EDIFACT
     edimodel edimodel_pkg.EDIMODEL
     x12 x12_pkg.X12
     config  configuration_pkg.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *EDINATIONAPI_IMPL) Configuration() configuration_pkg.CONFIGURATION {
    return me.config
}
/**
     * Access to Edifact controller
     * @return Returns the Edifact() instance
*/
func (me *EDINATIONAPI_IMPL) Edifact() edifact_pkg.EDIFACT {
    if(me.edifact) == nil {
        me.edifact = edifact_pkg.NewEDIFACT(me.config)
    }
    return me.edifact
}
/**
     * Access to EdiModel controller
     * @return Returns the EdiModel() instance
*/
func (me *EDINATIONAPI_IMPL) EdiModel() edimodel_pkg.EDIMODEL {
    if(me.edimodel) == nil {
        me.edimodel = edimodel_pkg.NewEDIMODEL(me.config)
    }
    return me.edimodel
}
/**
     * Access to X12 controller
     * @return Returns the X12() instance
*/
func (me *EDINATIONAPI_IMPL) X12() x12_pkg.X12 {
    if(me.x12) == nil {
        me.x12 = x12_pkg.NewX12(me.config)
    }
    return me.x12
}

