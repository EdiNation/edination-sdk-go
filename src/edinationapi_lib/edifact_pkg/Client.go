/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package edifact_pkg


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
type EDIFACT_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Reads an EDIFACT file and returns its contents translated to an array of EdifactInterchange objects.
 * @param    []byte         fileName             parameter: Required
 * @param    *bool          ignoreNullValues     parameter: Optional
 * @param    *bool          continueOnError      parameter: Optional
 * @param    *string        charSet              parameter: Optional
 * @param    *string        model                parameter: Optional
 * @param    *bool          eancomS3             parameter: Optional
 * @return	Returns the []*models_pkg.EdifactInterchange response from the API call
 */
func (me *EDIFACT_IMPL) Read (
            fileName []byte,
            ignoreNullValues *bool,
            continueOnError *bool,
            charSet *string,
            model *string,
            eancomS3 *bool) ([]*models_pkg.EdifactInterchange, error) {
    //the endpoint path uri
    _pathUrl := "/edifact/read"

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
        "eancomS3" : apihelper_pkg.ToString(*eancomS3, false),
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
    var retVal []*models_pkg.EdifactInterchange
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Translates an EdifactInterchange object to a raw EDIFACT interchange and returns it as a stream.
 * @param    *bool                                 preserveWhitespace             parameter: Optional
 * @param    *string                               charSet                        parameter: Optional
 * @param    *string                               postfix                        parameter: Optional
 * @param    *bool                                 sameRepetionAndDataElement     parameter: Optional
 * @param    *bool                                 eancomS3                       parameter: Optional
 * @param    *string                               contentType                    parameter: Optional
 * @param    *models_pkg.EdifactInterchange        body                           parameter: Optional
 * @return	Returns the []byte response from the API call
 */
func (me *EDIFACT_IMPL) Write (
            preserveWhitespace *bool,
            charSet *string,
            postfix *string,
            sameRepetionAndDataElement *bool,
            eancomS3 *bool,
            contentType *string,
            body *models_pkg.EdifactInterchange) ([]byte, error) {
    //the endpoint path uri
    _pathUrl := "/edifact/write"

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
        "sameRepetionAndDataElement" : apihelper_pkg.ToString(*sameRepetionAndDataElement, false),
        "eancomS3" : apihelper_pkg.ToString(*eancomS3, false),
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
 * Validates an EdifactInterchange object according to the EDIFACT standard rules for each version and transaction.
 * @param    *string                               syntaxSet         parameter: Optional
 * @param    *bool                                 skipTrailer       parameter: Optional
 * @param    *bool                                 structureOnly     parameter: Optional
 * @param    *string                               decimalPoint      parameter: Optional
 * @param    *bool                                 eancomS3          parameter: Optional
 * @param    *string                               contentType       parameter: Optional
 * @param    *models_pkg.EdifactInterchange        body              parameter: Optional
 * @return	Returns the *models_pkg.OperationResult response from the API call
 */
func (me *EDIFACT_IMPL) Validate (
            syntaxSet *string,
            skipTrailer *bool,
            structureOnly *bool,
            decimalPoint *string,
            eancomS3 *bool,
            contentType *string,
            body *models_pkg.EdifactInterchange) (*models_pkg.OperationResult, error) {
    //the endpoint path uri
    _pathUrl := "/edifact/validate"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "syntaxSet" : syntaxSet,
        "skipTrailer" : apihelper_pkg.ToString(*skipTrailer, false),
        "structureOnly" : apihelper_pkg.ToString(*structureOnly, false),
        "decimalPoint" : apihelper_pkg.ToString(*decimalPoint, "."),
        "eancomS3" : apihelper_pkg.ToString(*eancomS3, false),
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
 * Generates functional and/or technical CONTRL acknowledment(s) for the requested EdifactInterchange.
 * @param    *string                               syntaxSet                parameter: Optional
 * @param    *bool                                 detectDuplicates         parameter: Optional
 * @param    *int64                                tranRefNumber            parameter: Optional
 * @param    *int64                                interchangeRefNumber     parameter: Optional
 * @param    *bool                                 ackForValidTrans         parameter: Optional
 * @param    *bool                                 batchAcks                parameter: Optional
 * @param    *string                               technicalAck             parameter: Optional
 * @param    *bool                                 eancomS3                 parameter: Optional
 * @param    *string                               contentType              parameter: Optional
 * @param    *models_pkg.EdifactInterchange        body                     parameter: Optional
 * @return	Returns the []*models_pkg.EdifactInterchange response from the API call
 */
func (me *EDIFACT_IMPL) Ack (
            syntaxSet *string,
            detectDuplicates *bool,
            tranRefNumber *int64,
            interchangeRefNumber *int64,
            ackForValidTrans *bool,
            batchAcks *bool,
            technicalAck *string,
            eancomS3 *bool,
            contentType *string,
            body *models_pkg.EdifactInterchange) ([]*models_pkg.EdifactInterchange, error) {
    //the endpoint path uri
    _pathUrl := "/edifact/ack"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "syntaxSet" : syntaxSet,
        "detectDuplicates" : apihelper_pkg.ToString(*detectDuplicates, false),
        "tranRefNumber" : apihelper_pkg.ToString(*tranRefNumber, "1"),
        "interchangeRefNumber" : apihelper_pkg.ToString(*interchangeRefNumber, "1"),
        "ackForValidTrans" : apihelper_pkg.ToString(*ackForValidTrans, false),
        "batchAcks" : apihelper_pkg.ToString(*batchAcks, true),
        "technicalAck" : technicalAck,
        "eancomS3" : apihelper_pkg.ToString(*eancomS3, false),
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
    var retVal []*models_pkg.EdifactInterchange
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

