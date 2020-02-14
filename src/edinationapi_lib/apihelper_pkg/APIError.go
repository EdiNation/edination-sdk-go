/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package apihelper_pkg



// APIError structure for custom error handling in API invocation
type APIError struct {
    ResponseCode    int     //The HTTP response code from the API request
    Reason          string  //The reason for throwing exception
    Response		[]byte
}

// NewAPIError implements initialization constructor
// @param   string  reason  The reason for throwing exception
// @param   int     code    The HTTP response code from the API request
// @return  new APIException object
func NewAPIError(reason string, code int, res []byte) *APIError {
    ex := new(APIError)
    ex.ResponseCode = code
    ex.Reason = reason
    ex.Response = res
    return ex
}

// Error method implements error interface
func (e *APIError) Error() string {
    return e.Reason
}
