/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package configuration_pkg


/** The base Uri for API calls */
const BASEURI string = "https://api.edination.com/v2"



type CONFIGURATION_IMPL struct {
    /** API key to authenticate requests */
    /** Replace the value of OcpApimSubscriptionKey with an appropriate value */
    ocp-apim-subscription-key string
}

/*
 * Getter function returning ocp-apim-subscription-key
 */
func (me *CONFIGURATION_IMPL) OcpApimSubscriptionKey() string{
    return me.ocp-apim-subscription-key
}

/*
 * Setter function setting up ocp-apim-subscription-key
 */
func (me *CONFIGURATION_IMPL) SetOcpApimSubscriptionKey(ocpApimSubscriptionKey string) {
    me.ocp-apim-subscription-key = ocpApimSubscriptionKey
}

