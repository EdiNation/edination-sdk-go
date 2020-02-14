# Getting started

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=EdiNation%20API-GoLang&projectName=edinationapi_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=edinationapi_lib)

## How to Use

The following section explains how to use the EdinationapiLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=edinationapi_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=edinationapi_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=edinationapi_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=edinationapi_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=edinationapi_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=EdiNation%20API-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=edinationapi_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=edinationapi_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| ocpApimSubscriptionKey | API key to authenticate requests |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [edifact_pkg](#edifact_pkg)
* [edimodel_pkg](#edimodel_pkg)
* [x12_pkg](#x12_pkg)

## <a name="edifact_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".edifact_pkg") edifact_pkg

### Get instance

Factory for the ``` EDIFACT ``` interface can be accessed from the package edifact_pkg.

```go
edifact := edifact_pkg.NewEDIFACT()
```

### <a name="read"></a>![Method: ](https://apidocs.io/img/method.png ".edifact_pkg.Read") Read

> Reads an EDIFACT file and returns its contents translated to an array of EdifactInterchange objects.


```go
func (me *EDIFACT_IMPL) Read(
            fileName []byte,
            ignoreNullValues *bool,
            continueOnError *bool,
            charSet *string,
            model *string,
            eancomS3 *bool)([]*models_pkg.EdifactInterchange,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| fileName |  ``` Required ```  | Upload File |
| ignoreNullValues |  ``` Optional ```  ``` DefaultValue ```  | Whether to ignore all null values in the response. The default is false. |
| continueOnError |  ``` Optional ```  ``` DefaultValue ```  | Whether to continue reading if a corrupt interchange is encountered. The default is false. |
| charSet |  ``` Optional ```  ``` DefaultValue ```  | The encoding of the file contents. The default is utf-8. The available values are: unicodeFFFE, utf-32, utf-32BE, us-ascii, iso-8859-1, utf-7, utf-8, utf-16. |
| model |  ``` Optional ```  | The model to use. By default, the API will infer the model based on the transaction and version identifiers. |
| eancomS3 |  ``` Optional ```  ``` DefaultValue ```  | The default syntax for EANCOM transactions. By default all EANCOM transactions will be translated and validated according to the rules of Syntax 4. Set this flag to true if you need Syntax 3 to be used. |


#### Example Usage

```go
fileName :=  []byte("")
ignoreNullValues := false
continueOnError := false
charSet := "utf-8"
model := "model"
eancomS3 := false

var result []*models_pkg.EdifactInterchange
result,_ = edifact.Read(fileName, ignoreNullValues, continueOnError, charSet, model, eancomS3)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



### <a name="write"></a>![Method: ](https://apidocs.io/img/method.png ".edifact_pkg.Write") Write

> Translates an EdifactInterchange object to a raw EDIFACT interchange and returns it as a stream.


```go
func (me *EDIFACT_IMPL) Write(
            preserveWhitespace *bool,
            charSet *string,
            postfix *string,
            sameRepetionAndDataElement *bool,
            eancomS3 *bool,
            contentType *string,
            body *models_pkg.EdifactInterchange)([]byte,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| preserveWhitespace |  ``` Optional ```  ``` DefaultValue ```  | Whether to preserve blank data elements so the output contains multiple delimiters instead of omitting any excess delimiters. The default is false. |
| charSet |  ``` Optional ```  ``` DefaultValue ```  | The encoding of the file contents. The default is utf-8. The available values are: unicodeFFFE, utf-32, utf-32BE, us-ascii, iso-8859-1, utf-7, utf-8, utf-16. |
| postfix |  ``` Optional ```  | The postfix to be applied at the end of each segment, just after the segment separator. This is usually a carriage return (CR), line feed (LF) or both. By default, there is no postfix. |
| sameRepetionAndDataElement |  ``` Optional ```  ``` DefaultValue ```  | Sometimes the same delimiter is used to denote data element separator and repetition separator as in IATA transactions. By default, this is false. |
| eancomS3 |  ``` Optional ```  ``` DefaultValue ```  | The default syntax for EANCOM transactions. By default all EANCOM transactions will be translated and validated according to the rules of Syntax 4. Set this flag to true if you need Syntax 3 to be used. |
| contentType |  ``` Optional ```  ``` DefaultValue ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | The EdifactInterchange object to translate to raw EDIFACT. |


#### Example Usage

```go
preserveWhitespace := false
charSet := "utf-8"
postfix := "postfix"
sameRepetionAndDataElement := false
eancomS3 := false
contentType := "application/json"
var body *models_pkg.EdifactInterchange

var result []byte
result,_ = edifact.Write(preserveWhitespace, charSet, postfix, sameRepetionAndDataElement, eancomS3, contentType, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



### <a name="validate"></a>![Method: ](https://apidocs.io/img/method.png ".edifact_pkg.Validate") Validate

> Validates an EdifactInterchange object according to the EDIFACT standard rules for each version and transaction.


```go
func (me *EDIFACT_IMPL) Validate(
            syntaxSet *string,
            skipTrailer *bool,
            structureOnly *bool,
            decimalPoint *string,
            eancomS3 *bool,
            contentType *string,
            body *models_pkg.EdifactInterchange)(*models_pkg.OperationResult,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| syntaxSet |  ``` Optional ```  | In case you need to validate against a syntax set, different than UNOA and UNOB, populate this filed with all of the allowed symbols, url-escaped. All data elements with alpha (A) or alphanumeric (AN) data types are validated against a syntax set of allowed characters. The supported syntax sets are UNOA and UNOB. The validator infers the correct syntax set from EdifactInterchange.UNB.SYNTAXIDENTIFIER_1.SyntaxIdentifier_1. The symbols added to this field will override the corresponding sets in both UNOA and UNOB. |
| skipTrailer |  ``` Optional ```  ``` DefaultValue ```  | You are allowed to validate an EdifactInterchange with missing interchange, functional group or transaction trailers (UNZ, UNE, UNT). This is because these will be automatically applied during the Write oprtaion so you don't have to worry about counting the items. By default it is expected that all trailers are present when you validate the EdifactInterchange and by default, this is set to false. To skip all trailer validation, set this to true. |
| structureOnly |  ``` Optional ```  ``` DefaultValue ```  | This is equivalent to HIPAA Snip level 1, where only the structure and control segments are validated. By default, this is set to false, however if you want to not validate things such as data types, number of repeteitions or dates, set this to true. |
| decimalPoint |  ``` Optional ```  ``` DefaultValue ```  | This could be either point . (default) or comma ,. |
| eancomS3 |  ``` Optional ```  ``` DefaultValue ```  | The default syntax for EANCOM transactions. By default all EANCOM transactions will be validated according to the rules of Syntax 4. Set this flag to true if you need Syntax 3 to be used. |
| contentType |  ``` Optional ```  ``` DefaultValue ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | The EdifactInterchange object to validate. |


#### Example Usage

```go
syntaxSet := "syntaxSet"
skipTrailer := false
structureOnly := false
decimalPoint := "."
eancomS3 := false
contentType := "application/json"
var body *models_pkg.EdifactInterchange

var result *models_pkg.OperationResult
result,_ = edifact.Validate(syntaxSet, skipTrailer, structureOnly, decimalPoint, eancomS3, contentType, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



### <a name="ack"></a>![Method: ](https://apidocs.io/img/method.png ".edifact_pkg.Ack") Ack

> Generates functional and/or technical CONTRL acknowledment(s) for the requested EdifactInterchange.


```go
func (me *EDIFACT_IMPL) Ack(
            syntaxSet *string,
            detectDuplicates *bool,
            tranRefNumber *int64,
            interchangeRefNumber *int64,
            ackForValidTrans *bool,
            batchAcks *bool,
            technicalAck *string,
            eancomS3 *bool,
            contentType *string,
            body *models_pkg.EdifactInterchange)([]*models_pkg.EdifactInterchange,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| syntaxSet |  ``` Optional ```  | In case you need to validate against a syntax set, different than UNOA and UNOB, populate this filed with all of the allowed symbols, url-escaped. All data elements with alpha (A) or alphanumeric (AN) data types are validated against a syntax set of allowed characters. The supported syntax sets are UNOA and UNOB. The validator infers the correct syntax set from EdifactInterchange.UNB.SYNTAXIDENTIFIER_1.SyntaxIdentifier_1. The symbols added to this field will override the corresponding sets in both UNOA and UNOB. |
| detectDuplicates |  ``` Optional ```  ``` DefaultValue ```  | If you need to detect duplicates as in functional groups or transactions with the same reference number, set this flag to true. The default is false. |
| tranRefNumber |  ``` Optional ```  ``` DefaultValue ```  | The default is 1. Set this to whatever the CONTRL UNH.MessageReferenceNumber_01 needs to be. |
| interchangeRefNumber |  ``` Optional ```  ``` DefaultValue ```  | The default is 1. Set this to whatever the CONTRL EdifactInterchange.UNB.InterchangeControlReference_5 needs to be. |
| ackForValidTrans |  ``` Optional ```  ``` DefaultValue ```  | The default is false. Set this to true if you need UCM loops included for all valid transaction as well. By default UCM loops are generated only for invalid transactions. |
| batchAcks |  ``` Optional ```  ``` DefaultValue ```  | The default is true. Set this to false if you need to generate separate EdifactInterchange for each acknowledgment. By default all acknowledgments are batched in the same EdifactInterchange. |
| technicalAck |  ``` Optional ```  | The default technical acknowledgment CONTRL is generated according to EdifactInterchange.UNB.AcknowledgementRequest_9. You can either enforce it to always generate technical CONTRLs or supress it to never generate any technical CONTRLs. This will override the flag in EdifactInterchange.UNB.AcknowledgementRequest_9. The available values are: default, enforce, suppress. |
| eancomS3 |  ``` Optional ```  ``` DefaultValue ```  | The default syntax for EANCOM transactions. By default all EANCOM transactions will be validated according to the rules of Syntax 4. Set this flag to true if you need Syntax 3 to be used. |
| contentType |  ``` Optional ```  ``` DefaultValue ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | The EdifactInterchange object to acknowledge. |


#### Example Usage

```go
syntaxSet := "syntaxSet"
detectDuplicates := false
tranRefNumber,_ := strconv.ParseInt("1", 10, 8)
interchangeRefNumber,_ := strconv.ParseInt("1", 10, 8)
ackForValidTrans := false
batchAcks := true
technicalAck := "technicalAck"
eancomS3 := false
contentType := "application/json"
var body *models_pkg.EdifactInterchange

var result []*models_pkg.EdifactInterchange
result,_ = edifact.Ack(syntaxSet, detectDuplicates, tranRefNumber, interchangeRefNumber, ackForValidTrans, batchAcks, technicalAck, eancomS3, contentType, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



[Back to List of Controllers](#list_of_controllers)

## <a name="edimodel_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".edimodel_pkg") edimodel_pkg

### Get instance

Factory for the ``` EDIMODEL ``` interface can be accessed from the package edimodel_pkg.

```go
ediModel := edimodel_pkg.NewEDIMODEL()
```

### <a name="upload"></a>![Method: ](https://apidocs.io/img/method.png ".edimodel_pkg.Upload") Upload

> Uploads a model file to a subscription. It must be a valid DOT NET assembly.


```go
func (me *EDIMODEL_IMPL) Upload(fileName []byte)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| fileName |  ``` Required ```  | Upload File |


#### Example Usage

```go
fileName :=  []byte("")

var result 
result,_ = ediModel.Upload(fileName)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



### <a name="get_models"></a>![Method: ](https://apidocs.io/img/method.png ".edimodel_pkg.GetModels") GetModels

> Retrieves all models for a subscription.


```go
func (me *EDIMODEL_IMPL) GetModels(system *bool)([]*models_pkg.EdiModel,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| system |  ``` Optional ```  ``` DefaultValue ```  | Whether to retrieve EdiNation's or custom uploaded models. |


#### Example Usage

```go
system := true

var result []*models_pkg.EdiModel
result,_ = ediModel.GetModels(system)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



### <a name="get_model"></a>![Method: ](https://apidocs.io/img/method.png ".edimodel_pkg.GetModel") GetModel

> Retrieve a particular model file as a stream.


```go
func (me *EDIMODEL_IMPL) GetModel(
            id string,
            system *bool)([]byte,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | The model name. |
| system |  ``` Optional ```  ``` DefaultValue ```  | Whether to search in EdiNation's or custom uploaded models. |


#### Example Usage

```go
id := "id"
system := true

var result []byte
result,_ = ediModel.GetModel(id, system)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



### <a name="delete"></a>![Method: ](https://apidocs.io/img/method.png ".edimodel_pkg.Delete") Delete

> Deletes a particular model from the custom models.


```go
func (me *EDIMODEL_IMPL) Delete(id string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | The model name. |


#### Example Usage

```go
id := "id"

var result 
result,_ = ediModel.Delete(id)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



[Back to List of Controllers](#list_of_controllers)

## <a name="x12_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".x12_pkg") x12_pkg

### Get instance

Factory for the ``` X12 ``` interface can be accessed from the package x12_pkg.

```go
x12 := x12_pkg.NewX12()
```

### <a name="read"></a>![Method: ](https://apidocs.io/img/method.png ".x12_pkg.Read") Read

> Reads an X12 file and returns its contents translated to an array of X12Interchange objects.


```go
func (me *X12_IMPL) Read(
            fileName []byte,
            ignoreNullValues *bool,
            continueOnError *bool,
            charSet *string,
            model *string)([]*models_pkg.X12Interchange,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| fileName |  ``` Required ```  | Upload File |
| ignoreNullValues |  ``` Optional ```  ``` DefaultValue ```  | Whether to ignore all null values in the response. The default is false. |
| continueOnError |  ``` Optional ```  ``` DefaultValue ```  | Whether to continue reading if a corrupt interchange is encountered. The default is false. |
| charSet |  ``` Optional ```  ``` DefaultValue ```  | The encoding of the file contents. The default is utf-8. The available values are: unicodeFFFE, utf-32, utf-32BE, us-ascii, iso-8859-1, utf-7, utf-8, utf-16. |
| model |  ``` Optional ```  | The model to use. By default, the API will infer the model based on the transaction and version identifiers. |


#### Example Usage

```go
fileName :=  []byte("")
ignoreNullValues := false
continueOnError := false
charSet := "utf-8"
model := "model"

var result []*models_pkg.X12Interchange
result,_ = x12.Read(fileName, ignoreNullValues, continueOnError, charSet, model)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



### <a name="write"></a>![Method: ](https://apidocs.io/img/method.png ".x12_pkg.Write") Write

> Translates an X12Interchange object to a raw X12 interchange and returns it as a stream.


```go
func (me *X12_IMPL) Write(
            preserveWhitespace *bool,
            charSet *string,
            postfix *string,
            contentType *string,
            body *models_pkg.X12Interchange)([]byte,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| preserveWhitespace |  ``` Optional ```  ``` DefaultValue ```  | Whether to preserve blank data elements so the output contains multiple delimiters instead of omitting any excess delimiters. The default is false. |
| charSet |  ``` Optional ```  ``` DefaultValue ```  | The encoding of the file contents. The default is utf-8. The available values are: unicodeFFFE, utf-32, utf-32BE, us-ascii, iso-8859-1, utf-7, utf-8, utf-16. |
| postfix |  ``` Optional ```  | The postfix to be applied at the end of each segment, just after the segment separator. This is usually a carriage return (CR), line feed (LF) or both. By default, there is no postfix. |
| contentType |  ``` Optional ```  ``` DefaultValue ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | The X12Interchange object to translate to raw X12. |


#### Example Usage

```go
preserveWhitespace := false
charSet := "utf-8"
postfix := "postfix"
contentType := "application/json"
var body *models_pkg.X12Interchange

var result []byte
result,_ = x12.Write(preserveWhitespace, charSet, postfix, contentType, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



### <a name="validate"></a>![Method: ](https://apidocs.io/img/method.png ".x12_pkg.Validate") Validate

> Validates an X12Interchange object according to the X12 standard rules for each version and transaction.


```go
func (me *X12_IMPL) Validate(
            basicSyntax *bool,
            syntaxSet *string,
            skipTrailer *bool,
            structureOnly *bool,
            contentType *string,
            body *models_pkg.X12Interchange)(*models_pkg.OperationResult,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| basicSyntax |  ``` Optional ```  ``` DefaultValue ```  | All data elements with alpha (A) or alphanumeric (AN) data types are validated against a syntax set of allowed characters. The default syntax set is the Extended, hence the default for this parameter is false. By setting this to true, validation will use the Basic syntax set. |
| syntaxSet |  ``` Optional ```  | In case you need to validate against a syntax set, different than Basic and Extended, populate this filed with all of the allowed symbols, url-escaped. |
| skipTrailer |  ``` Optional ```  ``` DefaultValue ```  | You are allowed to validate an X12Interchange with missing interchange, functional group or transaction trailers (IEA, GE, SE). This is because these will be automatically applied during the Write oprtaion so you don't have to worry about counting the items. By default it is expected that all trailers are present when you validate the X12Interchange and by default, this is set to false. To skip all trailer validation, set this to true. |
| structureOnly |  ``` Optional ```  ``` DefaultValue ```  | This is equivalent to HIPAA Snip level 1, where only the structure and control segments are validated. By default, this is set to false, however if you want to not validate things such as data types, number of repeteitions or dates, set this to true. |
| contentType |  ``` Optional ```  ``` DefaultValue ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | The X12Interchange object to validate. |


#### Example Usage

```go
basicSyntax := false
syntaxSet := "syntaxSet"
skipTrailer := false
structureOnly := false
contentType := "application/json"
var body *models_pkg.X12Interchange

var result *models_pkg.OperationResult
result,_ = x12.Validate(basicSyntax, syntaxSet, skipTrailer, structureOnly, contentType, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



### <a name="ack"></a>![Method: ](https://apidocs.io/img/method.png ".x12_pkg.Ack") Ack

> Generates functional/implementation and/or technical acknowledment(s) for the requested X12Interchange.


```go
func (me *X12_IMPL) Ack(
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
            body *models_pkg.X12Interchange)([]*models_pkg.X12Interchange,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| basicSyntax |  ``` Optional ```  ``` DefaultValue ```  | All data elements with alpha (A) or alphanumeric (AN) data types are validated against a syntax set of allowed characters. The default syntax set is the Extended, hence the default for this parameter is false. By setting this to true, validation will use the Basic syntax set. |
| syntaxSet |  ``` Optional ```  | In case you need to validate against a syntax set, different than Basic and Extended, populate this filed with all of the allowed symbols, url-escaped. |
| detectDuplicates |  ``` Optional ```  ``` DefaultValue ```  | If you need to detect duplicates as in functional groups or transactions with the same reference number, set this flag to true. The default is false. |
| tranRefNumber |  ``` Optional ```  ``` DefaultValue ```  | The default is 1. Set this to whatever the 997 or 999 X12Interchange.ST.TransactionSetControlNumber_02 needs to be. In case there are multiple acknowledgments (for multiple functional groups), this will be starting reference number and every subsequent acknowledgment will have the previous reference number incremented with 1. |
| interchangeRefNumber |  ``` Optional ```  ``` DefaultValue ```  | The default is 1. Set this to whatever the 997 or 999 X12Interchange.ISA.InterchangeControlNumber_13 needs to be. |
| ackForValidTrans |  ``` Optional ```  ``` DefaultValue ```  | The default is false. Set this to true if you need AK2 loops included for all valid transaction as well. By default AK2 loops are generated only for invalid transactions. |
| batchAcks |  ``` Optional ```  ``` DefaultValue ```  | The default is true. Set this to false if you need to generate separate X12Interchange for each acknowledgment. By default all acknowledgments are batched in the same X12Interchange. |
| technicalAck |  ``` Optional ```  | The default technical acknowledgment TA1 is generated according to X12Interchange.ISA.AcknowledgementRequested_14. You can either enforce it to always generate TA1s or supress it to never generate any TA1s. This will override the flag in X12Interchange.ISA.AcknowledgementRequested_14. The available values are: default, enforce, suppress. |
| ack |  ``` Optional ```  ``` DefaultValue ```  | The default value is 997. The type of acknowledgment being generated. Set this to 999 if you need to generate an implementation instead of functional acknowledgment. The available values are: 997, 999. |
| ak901isP |  ``` Optional ```  ``` DefaultValue ```  | The value of the AK9's first element. By default it is "E". Set this to true if you want this value to be "P" instead. |
| contentType |  ``` Optional ```  ``` DefaultValue ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | The X12Interchange object to acknowledge. |


#### Example Usage

```go
basicSyntax := false
syntaxSet := "syntaxSet"
detectDuplicates := false
tranRefNumber,_ := strconv.ParseInt("1", 10, 8)
interchangeRefNumber,_ := strconv.ParseInt("1", 10, 8)
ackForValidTrans := false
batchAcks := true
technicalAck := "technicalAck"
ack := "997"
ak901isP := false
contentType := "application/json"
var body *models_pkg.X12Interchange

var result []*models_pkg.X12Interchange
result,_ = x12.Ack(basicSyntax, syntaxSet, detectDuplicates, tranRefNumber, interchangeRefNumber, ackForValidTrans, batchAcks, technicalAck, ack, ak901isP, contentType, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |
| 500 | Server Error |



[Back to List of Controllers](#list_of_controllers)



