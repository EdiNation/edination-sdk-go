/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package configuration_pkg



type CONFIGURATION interface {
        OcpApimSubscriptionKey() string
        SetOcpApimSubscriptionKey(ocpApimSubscriptionKey   string)
}

/*
 * Factory for the CONFIGURATION interface returning CONFIGURATION_IMPL
 */
func NewCONFIGURATION() CONFIGURATION{
    configuration := new(CONFIGURATION_IMPL)
    return configuration
}