/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package x12_pkg


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
type X12_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Reads an X12 file and returns its contents translated to an array of X12Interchange objects.
 * @param    []byte         fileName             parameter: Required
 * @param    *bool          ignoreNullValues     parameter: Optional
 * @param    *bool          continueOnError      parameter: Optional
 * @param    *string        charSet              parameter: Optional
 * @param    *string        model                parameter: Optional
 * @return	Returns the []*models_pkg.X12Interchange response from the API call
 */
func (me *X12_IMPL) Read (
            fileName []byte,
            ignoreNullValues *bool,
            continueOnError *bool,
            charSet *string,
            model *string) ([]*models_pkg.X12Interchange, error) {
    //the endpoint path uri
    _pathUrl := "/x12/read"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "ignoreNullValues" : apihelper_pkg.ToString(*ignoreNullValues, false),
        "continueOnError" : apihelper_pkg.ToString(*continueOnError, false),
        "charSet" : apihelper_pkg.ToString(*charSet, "utf-8"),
        "model" : model,
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
    var retVal []*models_pkg.X12Interchange
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Translates an X12Interchange object to a raw X12 interchange and returns it as a stream.
 * @param    *bool                             preserveWhitespace     parameter: Optional
 * @param    *string                           charSet                parameter: Optional
 * @param    *string                           postfix                parameter: Optional
 * @param    *string                           contentType            parameter: Optional
 * @param    *models_pkg.X12Interchange        body                   parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *X12_IMPL) Write (
            preserveWhitespace *bool,
            charSet *string,
            postfix *string,
            contentType *string,
            body *models_pkg.X12Interchange) ([]byte, error) {
    //the endpoint path uri
    _pathUrl := "/x12/write"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "preserveWhitespace" : apihelper_pkg.ToString(*preserveWhitespace, false),
        "charSet" : apihelper_pkg.ToString(*charSet, "utf-8"),
        "postfix" : postfix,
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
        "content-type" : "application/json; charset=utf-8",
        "Content-Type" : apihelper_pkg.ToString(*contentType, "application/json"),
        "Ocp-Apim-Subscription-Key" : edinationapi_lib.config.OcpApimSubscriptionKey,
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
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
 * Validates an X12Interchange object according to the X12 standard rules for each version and transaction.
 * @param    *bool                             basicSyntax       parameter: Optional
 * @param    *string                           syntaxSet         parameter: Optional
 * @param    *bool                             skipTrailer       parameter: Optional
 * @param    *bool                             structureOnly     parameter: Optional
 * @param    *string                           contentType       parameter: Optional
 * @param    *models_pkg.X12Interchange        body              parameter: Optional
 * @return	Returns the *models_pkg.OperationResult response from the API call
 */
func (me *X12_IMPL) Validate (
            basicSyntax *bool,
            syntaxSet *string,
            skipTrailer *bool,
            structureOnly *bool,
            contentType *string,
            body *models_pkg.X12Interchange) (*models_pkg.OperationResult, error) {
    //the endpoint path uri
    _pathUrl := "/x12/validate"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "basicSyntax" : apihelper_pkg.ToString(*basicSyntax, false),
        "syntaxSet" : syntaxSet,
        "skipTrailer" : apihelper_pkg.ToString(*skipTrailer, false),
        "structureOnly" : apihelper_pkg.ToString(*structureOnly, false),
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
        "content-type" : "application/json; charset=utf-8",
        "Content-Type" : apihelper_pkg.ToString(*contentType, "application/json"),
        "Ocp-Apim-Subscription-Key" : edinationapi_lib.config.OcpApimSubscriptionKey,
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
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
    var retVal *models_pkg.OperationResult = &models_pkg.OperationResult{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Generates functional/implementation and/or technical acknowledment(s) for the requested X12Interchange.
 * @param    *bool                             basicSyntax              parameter: Optional
 * @param    *string                           syntaxSet                parameter: Optional
 * @param    *bool                             detectDuplicates         parameter: Optional
 * @param    *int64                            tranRefNumber            parameter: Optional
 * @param    *int64                            interchangeRefNumber     parameter: Optional
 * @param    *bool                             ackForValidTrans         parameter: Optional
 * @param    *bool                             batchAcks                parameter: Optional
 * @param    *string                           technicalAck             parameter: Optional
 * @param    *string                           ack                      parameter: Optional
 * @param    *bool                             ak901isP                 parameter: Optional
 * @param    *string                           contentType              parameter: Optional
 * @param    *models_pkg.X12Interchange        body                     parameter: Optional
 * @return	Returns the []*models_pkg.X12Interchange response from the API call
 */
func (me *X12_IMPL) Ack (
            basicSyntax *bool,
            syntaxSet *string,
            detectDuplicates *bool,
            tranRefNumber *int64,
            interchangeRefNumber *int64,
            ackForValidTrans *bool,
            batchAcks *bool,
            technicalAck *string,
            ack *string,
            ak901isP *bool,
            contentType *string,
            body *models_pkg.X12Interchange) ([]*models_pkg.X12Interchange, error) {
    //the endpoint path uri
    _pathUrl := "/x12/ack"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "basicSyntax" : apihelper_pkg.ToString(*basicSyntax, false),
        "syntaxSet" : syntaxSet,
        "detectDuplicates" : apihelper_pkg.ToString(*detectDuplicates, false),
        "tranRefNumber" : apihelper_pkg.ToString(*tranRefNumber, "1"),
        "interchangeRefNumber" : apihelper_pkg.ToString(*interchangeRefNumber, "1"),
        "ackForValidTrans" : apihelper_pkg.ToString(*ackForValidTrans, false),
        "batchAcks" : apihelper_pkg.ToString(*batchAcks, true),
        "technicalAck" : technicalAck,
        "ack" : apihelper_pkg.ToString(*ack, "997"),
        "ak901isP" : apihelper_pkg.ToString(*ak901isP, false),
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
        "content-type" : "application/json; charset=utf-8",
        "Content-Type" : apihelper_pkg.ToString(*contentType, "application/json"),
        "Ocp-Apim-Subscription-Key" : edinationapi_lib.config.OcpApimSubscriptionKey,
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
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
    var retVal []*models_pkg.X12Interchange
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

