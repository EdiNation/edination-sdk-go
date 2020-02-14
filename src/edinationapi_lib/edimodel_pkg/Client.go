/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package edimodel_pkg


import(
	"encoding/json"
	"github.com/apimatic/unirest-go"
	"edinationapi_lib/apihelper_pkg"
	"edinationapi_lib/configuration_pkg"
	"edinationapi_lib/models_pkg"
)
/*
 * Client structure as interface implementation
 */
type EDIMODEL_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Uploads a model file to a subscription. It must be a valid DOT NET assembly.
 * @param    []byte        fileName     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *EDIMODEL_IMPL) Upload (
            fileName []byte) (error) {
    //the endpoint path uri
    _pathUrl := "/models"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "Ocp-Apim-Subscription-Key" : edinationapi_lib.config.OcpApimSubscriptionKey,
    }

    //form parameters
    parameters := map[string]interface{} {

        "fileName" : fileName,

    }


    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Bad Request", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Server Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return err
    }

    //returning the response
    return nil

}

/**
 * Retrieves all models for a subscription.
 * @param    *bool        system     parameter: Optional
 * @return	Returns the []*models_pkg.EdiModel response from the API call
 */
func (me *EDIMODEL_IMPL) GetModels (
            system *bool) ([]*models_pkg.EdiModel, error) {
    //the endpoint path uri
    _pathUrl := "/models"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "system" : apihelper_pkg.ToString(*system, true),
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "Ocp-Apim-Subscription-Key" : edinationapi_lib.config.OcpApimSubscriptionKey,
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Bad Request", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Server Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal []*models_pkg.EdiModel
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Retrieve a particular model file as a stream.
 * @param    string        id         parameter: Required
 * @param    *bool         system     parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *EDIMODEL_IMPL) GetModel (
            id string,
            system *bool) ([]byte, error) {
    //the endpoint path uri
    _pathUrl := "/models/{id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "system" : apihelper_pkg.ToString(*system, true),
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "Ocp-Apim-Subscription-Key" : edinationapi_lib.config.OcpApimSubscriptionKey,
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Bad Request", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Server Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    return _response.RawBody, nil

}

/**
 * Deletes a particular model from the custom models.
 * @param    string        id     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *EDIMODEL_IMPL) Delete (
            id string) (error) {
    //the endpoint path uri
    _pathUrl := "/models/{id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "Ocp-Apim-Subscription-Key" : edinationapi_lib.config.OcpApimSubscriptionKey,
    }

    //prepare API request
    _request := unirest.Delete(_queryBuilder, headers, nil)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Bad Request", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Server Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return err
    }

    //returning the response
    return nil

}

