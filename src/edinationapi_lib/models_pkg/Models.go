/*
 * edinationapi_lib
 *
 * This file was automatically generated for EdiNation by APIMATIC v2.0 ( https://apimatic.io ).
 */

package models_pkg



/*
 * Structure for the custom type UNA
 */
type UNA struct {
    ComponentDataElement *string         `json:"ComponentDataElement,omitempty" form:"ComponentDataElement,omitempty"` //TODO: Write general description for this field
    DataElement          *string         `json:"DataElement,omitempty" form:"DataElement,omitempty"` //TODO: Write general description for this field
    DecimalNotation      *string         `json:"DecimalNotation,omitempty" form:"DecimalNotation,omitempty"` //TODO: Write general description for this field
    ReleaseIndicator     *string         `json:"ReleaseIndicator,omitempty" form:"ReleaseIndicator,omitempty"` //TODO: Write general description for this field
    Reserved             *string         `json:"Reserved,omitempty" form:"Reserved,omitempty"` //TODO: Write general description for this field
    Segment              *string         `json:"Segment,omitempty" form:"Segment,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type S001
 */
type S001 struct {
    SyntaxIdentifier1                       *string         `json:"SyntaxIdentifier_1,omitempty" form:"SyntaxIdentifier_1,omitempty"` //TODO: Write general description for this field
    SyntaxVersionNumber2                    *string         `json:"SyntaxVersionNumber_2,omitempty" form:"SyntaxVersionNumber_2,omitempty"` //TODO: Write general description for this field
    ServiceCodeListDirectoryVersionNumber3  *string         `json:"ServiceCodeListDirectoryVersionNumber_3,omitempty" form:"ServiceCodeListDirectoryVersionNumber_3,omitempty"` //TODO: Write general description for this field
    CharacterEncoding4                      *string         `json:"CharacterEncoding_4,omitempty" form:"CharacterEncoding_4,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type S002
 */
type S002 struct {
    InterchangeSenderIdentification1             *string         `json:"InterchangeSenderIdentification_1,omitempty" form:"InterchangeSenderIdentification_1,omitempty"` //TODO: Write general description for this field
    IdentificationCodeQualifier2                 *string         `json:"IdentificationCodeQualifier_2,omitempty" form:"IdentificationCodeQualifier_2,omitempty"` //TODO: Write general description for this field
    InterchangeSenderInternalIdentification3     *string         `json:"InterchangeSenderInternalIdentification_3,omitempty" form:"InterchangeSenderInternalIdentification_3,omitempty"` //TODO: Write general description for this field
    InterchangeSenderInternalSubIdentification4  *string         `json:"InterchangeSenderInternalSubIdentification_4,omitempty" form:"InterchangeSenderInternalSubIdentification_4,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type S003
 */
type S003 struct {
    InterchangeRecipientIdentification1             *string         `json:"InterchangeRecipientIdentification_1,omitempty" form:"InterchangeRecipientIdentification_1,omitempty"` //TODO: Write general description for this field
    IdentificationCodeQualifier2                    *string         `json:"IdentificationCodeQualifier_2,omitempty" form:"IdentificationCodeQualifier_2,omitempty"` //TODO: Write general description for this field
    InterchangeRecipientInternalIdentification3     *string         `json:"InterchangeRecipientInternalIdentification_3,omitempty" form:"InterchangeRecipientInternalIdentification_3,omitempty"` //TODO: Write general description for this field
    InterchangeRecipientInternalSubIdentification4  *string         `json:"InterchangeRecipientInternalSubIdentification_4,omitempty" form:"InterchangeRecipientInternalSubIdentification_4,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type S004
 */
type S004 struct {
    Date1           *string         `json:"Date_1,omitempty" form:"Date_1,omitempty"` //TODO: Write general description for this field
    Time2           *string         `json:"Time_2,omitempty" form:"Time_2,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type S005
 */
type S005 struct {
    RecipientReferencePassword1           *string         `json:"RecipientReferencePassword_1,omitempty" form:"RecipientReferencePassword_1,omitempty"` //TODO: Write general description for this field
    RecipientReferencePasswordQualifier2  *string         `json:"RecipientReferencePasswordQualifier_2,omitempty" form:"RecipientReferencePasswordQualifier_2,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type UNB
 */
type UNB struct {
    SYNTAXIDENTIFIER1                    *S001           `json:"SYNTAXIDENTIFIER_1,omitempty" form:"SYNTAXIDENTIFIER_1,omitempty"` //TODO: Write general description for this field
    INTERCHANGESENDER2                   *S002           `json:"INTERCHANGESENDER_2,omitempty" form:"INTERCHANGESENDER_2,omitempty"` //TODO: Write general description for this field
    INTERCHANGERECIPIENT3                *S003           `json:"INTERCHANGERECIPIENT_3,omitempty" form:"INTERCHANGERECIPIENT_3,omitempty"` //TODO: Write general description for this field
    DATEANDTIMEOFPREPARATION4            *S004           `json:"DATEANDTIMEOFPREPARATION_4,omitempty" form:"DATEANDTIMEOFPREPARATION_4,omitempty"` //TODO: Write general description for this field
    InterchangeControlReference5         *string         `json:"InterchangeControlReference_5,omitempty" form:"InterchangeControlReference_5,omitempty"` //TODO: Write general description for this field
    RECIPIENTSREFERENCEPASSWORDDETAILS6  *S005           `json:"RECIPIENTSREFERENCEPASSWORDDETAILS_6,omitempty" form:"RECIPIENTSREFERENCEPASSWORDDETAILS_6,omitempty"` //TODO: Write general description for this field
    ApplicationReference7                *string         `json:"ApplicationReference_7,omitempty" form:"ApplicationReference_7,omitempty"` //TODO: Write general description for this field
    ProcessingPriorityCode8              *string         `json:"ProcessingPriorityCode_8,omitempty" form:"ProcessingPriorityCode_8,omitempty"` //TODO: Write general description for this field
    AcknowledgementRequest9              *string         `json:"AcknowledgementRequest_9,omitempty" form:"AcknowledgementRequest_9,omitempty"` //TODO: Write general description for this field
    InterchangeAgreementIdentifier10     *string         `json:"InterchangeAgreementIdentifier_10,omitempty" form:"InterchangeAgreementIdentifier_10,omitempty"` //TODO: Write general description for this field
    TestIndicator11                      *string         `json:"TestIndicator_11,omitempty" form:"TestIndicator_11,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type S006
 */
type S006 struct {
    ApplicationSenderIdentification1  *string         `json:"ApplicationSenderIdentification_1,omitempty" form:"ApplicationSenderIdentification_1,omitempty"` //TODO: Write general description for this field
    IdentificationCodeQualifier2      *string         `json:"IdentificationCodeQualifier_2,omitempty" form:"IdentificationCodeQualifier_2,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type S007
 */
type S007 struct {
    ApplicationRecipientIdentification1  *string         `json:"ApplicationRecipientIdentification_1,omitempty" form:"ApplicationRecipientIdentification_1,omitempty"` //TODO: Write general description for this field
    IdentificationCodeQualifier2         *string         `json:"IdentificationCodeQualifier_2,omitempty" form:"IdentificationCodeQualifier_2,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type S008
 */
type S008 struct {
    MessageVersionNumber1     *string         `json:"MessageVersionNumber_1,omitempty" form:"MessageVersionNumber_1,omitempty"` //TODO: Write general description for this field
    MessageReleaseNumber2     *string         `json:"MessageReleaseNumber_2,omitempty" form:"MessageReleaseNumber_2,omitempty"` //TODO: Write general description for this field
    AssociationAssignedCode3  *string         `json:"AssociationAssignedCode_3,omitempty" form:"AssociationAssignedCode_3,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type UNG
 */
type UNG struct {
    MessageGroupIdentification1          *string         `json:"MessageGroupIdentification_1,omitempty" form:"MessageGroupIdentification_1,omitempty"` //TODO: Write general description for this field
    APPLICATIONSENDERIDENTIFICATION2     *S006           `json:"APPLICATIONSENDERIDENTIFICATION_2,omitempty" form:"APPLICATIONSENDERIDENTIFICATION_2,omitempty"` //TODO: Write general description for this field
    APPLICATIONRECIPIENTIDENTIFICATION3  *S007           `json:"APPLICATIONRECIPIENTIDENTIFICATION_3,omitempty" form:"APPLICATIONRECIPIENTIDENTIFICATION_3,omitempty"` //TODO: Write general description for this field
    DATEANDTIMEOFPREPARATION4            *S004           `json:"DATEANDTIMEOFPREPARATION_4,omitempty" form:"DATEANDTIMEOFPREPARATION_4,omitempty"` //TODO: Write general description for this field
    GroupReferenceNumber5                *string         `json:"GroupReferenceNumber_5,omitempty" form:"GroupReferenceNumber_5,omitempty"` //TODO: Write general description for this field
    ControllingAgency6                   *string         `json:"ControllingAgency_6,omitempty" form:"ControllingAgency_6,omitempty"` //TODO: Write general description for this field
    MESSAGEVERSION7                      *S008           `json:"MESSAGEVERSION_7,omitempty" form:"MESSAGEVERSION_7,omitempty"` //TODO: Write general description for this field
    ApplicationPassword8                 *string         `json:"ApplicationPassword_8,omitempty" form:"ApplicationPassword_8,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type UNE
 */
type UNE struct {
    GroupControlCount1     *string         `json:"GroupControlCount_1,omitempty" form:"GroupControlCount_1,omitempty"` //TODO: Write general description for this field
    GroupReferenceNumber2  *string         `json:"GroupReferenceNumber_2,omitempty" form:"GroupReferenceNumber_2,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type EdifactGroup
 */
type EdifactGroup struct {
    UNG             *UNG            `json:"UNG,omitempty" form:"UNG,omitempty"` //TODO: Write general description for this field
    Transactions    []interface{}   `json:"Transactions" form:"Transactions"` //TODO: Write general description for this field
    UNETrailers     []*UNE          `json:"UNETrailers,omitempty" form:"UNETrailers,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type UNZ
 */
type UNZ struct {
    InterchangeControlCount1      *string         `json:"InterchangeControlCount_1,omitempty" form:"InterchangeControlCount_1,omitempty"` //TODO: Write general description for this field
    InterchangeControlReference2  *string         `json:"InterchangeControlReference_2,omitempty" form:"InterchangeControlReference_2,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type OperationDetail
 */
type OperationDetail struct {
    Index            *int64          `json:"Index,omitempty" form:"Index,omitempty"` //TODO: Write general description for this field
    TransactionIndex *int64          `json:"TransactionIndex,omitempty" form:"TransactionIndex,omitempty"` //TODO: Write general description for this field
    TransactionRef   *string         `json:"TransactionRef,omitempty" form:"TransactionRef,omitempty"` //TODO: Write general description for this field
    SegmentId        *string         `json:"SegmentId,omitempty" form:"SegmentId,omitempty"` //TODO: Write general description for this field
    Value            *string         `json:"Value,omitempty" form:"Value,omitempty"` //TODO: Write general description for this field
    Message          *string         `json:"Message,omitempty" form:"Message,omitempty"` //TODO: Write general description for this field
    Status           *string         `json:"Status,omitempty" form:"Status,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type OperationResult
 */
type OperationResult struct {
    LastIndex       *int64          `json:"LastIndex,omitempty" form:"LastIndex,omitempty"` //TODO: Write general description for this field
    Details         []*OperationDetail `json:"Details,omitempty" form:"Details,omitempty"` //TODO: Write general description for this field
    Status          *string         `json:"Status,omitempty" form:"Status,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type EdifactInterchange
 */
type EdifactInterchange struct {
    UNA             *UNA            `json:"UNA,omitempty" form:"UNA,omitempty"` //TODO: Write general description for this field
    UNB             UNB             `json:"UNB" form:"UNB"` //TODO: Write general description for this field
    Groups          []*EdifactGroup `json:"Groups" form:"Groups"` //TODO: Write general description for this field
    UNZTrailers     []*UNZ          `json:"UNZTrailers,omitempty" form:"UNZTrailers,omitempty"` //TODO: Write general description for this field
    Result          *OperationResult `json:"Result,omitempty" form:"Result,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type Link
 */
type Link struct {
    Rel             *string         `json:"Rel,omitempty" form:"Rel,omitempty"` //TODO: Write general description for this field
    Href            *string         `json:"Href,omitempty" form:"Href,omitempty"` //TODO: Write general description for this field
    Action          *string         `json:"Action,omitempty" form:"Action,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type EdiModel
 */
type EdiModel struct {
    Name            *string         `json:"Name,omitempty" form:"Name,omitempty"` //TODO: Write general description for this field
    DateCreated     *string         `json:"DateCreated,omitempty" form:"DateCreated,omitempty"` //TODO: Write general description for this field
    Links           []*Link         `json:"Links,omitempty" form:"Links,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ISA
 */
type ISA struct {
    AuthorizationInformationQualifier1       *string         `json:"AuthorizationInformationQualifier_1,omitempty" form:"AuthorizationInformationQualifier_1,omitempty"` //TODO: Write general description for this field
    AuthorizationInformation2                *string         `json:"AuthorizationInformation_2,omitempty" form:"AuthorizationInformation_2,omitempty"` //TODO: Write general description for this field
    SecurityInformationQualifier3            *string         `json:"SecurityInformationQualifier_3,omitempty" form:"SecurityInformationQualifier_3,omitempty"` //TODO: Write general description for this field
    SecurityInformation4                     *string         `json:"SecurityInformation_4,omitempty" form:"SecurityInformation_4,omitempty"` //TODO: Write general description for this field
    SenderIDQualifier5                       *string         `json:"SenderIDQualifier_5,omitempty" form:"SenderIDQualifier_5,omitempty"` //TODO: Write general description for this field
    InterchangeSenderID6                     *string         `json:"InterchangeSenderID_6,omitempty" form:"InterchangeSenderID_6,omitempty"` //TODO: Write general description for this field
    ReceiverIDQualifier7                     *string         `json:"ReceiverIDQualifier_7,omitempty" form:"ReceiverIDQualifier_7,omitempty"` //TODO: Write general description for this field
    InterchangeReceiverID8                   *string         `json:"InterchangeReceiverID_8,omitempty" form:"InterchangeReceiverID_8,omitempty"` //TODO: Write general description for this field
    InterchangeDate9                         *string         `json:"InterchangeDate_9,omitempty" form:"InterchangeDate_9,omitempty"` //TODO: Write general description for this field
    InterchangeTime10                        *string         `json:"InterchangeTime_10,omitempty" form:"InterchangeTime_10,omitempty"` //TODO: Write general description for this field
    InterchangeControlStandardsIdentifier11  *string         `json:"InterchangeControlStandardsIdentifier_11,omitempty" form:"InterchangeControlStandardsIdentifier_11,omitempty"` //TODO: Write general description for this field
    InterchangeControlVersionNumber12        *string         `json:"InterchangeControlVersionNumber_12,omitempty" form:"InterchangeControlVersionNumber_12,omitempty"` //TODO: Write general description for this field
    InterchangeControlNumber13               *string         `json:"InterchangeControlNumber_13,omitempty" form:"InterchangeControlNumber_13,omitempty"` //TODO: Write general description for this field
    AcknowledgementRequested14               *string         `json:"AcknowledgementRequested_14,omitempty" form:"AcknowledgementRequested_14,omitempty"` //TODO: Write general description for this field
    UsageIndicator15                         *string         `json:"UsageIndicator_15,omitempty" form:"UsageIndicator_15,omitempty"` //TODO: Write general description for this field
    ComponentElementSeparator16              *string         `json:"ComponentElementSeparator_16,omitempty" form:"ComponentElementSeparator_16,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type TA1
 */
type TA1 struct {
    InterchangeControlNumber1       *string         `json:"InterchangeControlNumber_1,omitempty" form:"InterchangeControlNumber_1,omitempty"` //TODO: Write general description for this field
    InterchangeDate2                *string         `json:"InterchangeDate_2,omitempty" form:"InterchangeDate_2,omitempty"` //TODO: Write general description for this field
    InterchangeTime3                *string         `json:"InterchangeTime_3,omitempty" form:"InterchangeTime_3,omitempty"` //TODO: Write general description for this field
    InterchangeAcknowledgmentCode4  *string         `json:"InterchangeAcknowledgmentCode_4,omitempty" form:"InterchangeAcknowledgmentCode_4,omitempty"` //TODO: Write general description for this field
    InterchangeNoteCode5            *string         `json:"InterchangeNoteCode_5,omitempty" form:"InterchangeNoteCode_5,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type GS
 */
type GS struct {
    CodeIdentifyingInformationType1  *string         `json:"CodeIdentifyingInformationType_1,omitempty" form:"CodeIdentifyingInformationType_1,omitempty"` //TODO: Write general description for this field
    SenderIDCode2                    *string         `json:"SenderIDCode_2,omitempty" form:"SenderIDCode_2,omitempty"` //TODO: Write general description for this field
    ReceiverIDCode3                  *string         `json:"ReceiverIDCode_3,omitempty" form:"ReceiverIDCode_3,omitempty"` //TODO: Write general description for this field
    Date4                            *string         `json:"Date_4,omitempty" form:"Date_4,omitempty"` //TODO: Write general description for this field
    Time5                            *string         `json:"Time_5,omitempty" form:"Time_5,omitempty"` //TODO: Write general description for this field
    GroupControlNumber6              *string         `json:"GroupControlNumber_6,omitempty" form:"GroupControlNumber_6,omitempty"` //TODO: Write general description for this field
    TransactionTypeCode7             *string         `json:"TransactionTypeCode_7,omitempty" form:"TransactionTypeCode_7,omitempty"` //TODO: Write general description for this field
    VersionAndRelease8               *string         `json:"VersionAndRelease_8,omitempty" form:"VersionAndRelease_8,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type GE
 */
type GE struct {
    NumberOfIncludedSets1  *string         `json:"NumberOfIncludedSets_1,omitempty" form:"NumberOfIncludedSets_1,omitempty"` //TODO: Write general description for this field
    GroupControlNumber2    *string         `json:"GroupControlNumber_2,omitempty" form:"GroupControlNumber_2,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type X12Group
 */
type X12Group struct {
    GS              GS              `json:"GS" form:"GS"` //TODO: Write general description for this field
    Transactions    []interface{}   `json:"Transactions" form:"Transactions"` //TODO: Write general description for this field
    GETrailers      []*GE           `json:"GETrailers,omitempty" form:"GETrailers,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type IEA
 */
type IEA struct {
    NumberOfIncludedGroups1    *string         `json:"NumberOfIncludedGroups_1,omitempty" form:"NumberOfIncludedGroups_1,omitempty"` //TODO: Write general description for this field
    InterchangeControlNumber2  *string         `json:"InterchangeControlNumber_2,omitempty" form:"InterchangeControlNumber_2,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type X12Interchange
 */
type X12Interchange struct {
    SegmentDelimiter     string          `json:"SegmentDelimiter" form:"SegmentDelimiter"` //TODO: Write general description for this field
    DataElementDelimiter string          `json:"DataElementDelimiter" form:"DataElementDelimiter"` //TODO: Write general description for this field
    ISA                  ISA             `json:"ISA" form:"ISA"` //TODO: Write general description for this field
    TA1                  *TA1            `json:"TA1,omitempty" form:"TA1,omitempty"` //TODO: Write general description for this field
    Groups               []*X12Group     `json:"Groups" form:"Groups"` //TODO: Write general description for this field
    IEATrailers          []*IEA          `json:"IEATrailers,omitempty" form:"IEATrailers,omitempty"` //TODO: Write general description for this field
    Result               *OperationResult `json:"Result,omitempty" form:"Result,omitempty"` //TODO: Write general description for this field
}
