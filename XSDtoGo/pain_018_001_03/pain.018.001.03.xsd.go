// Code generated by download. DO NOT EDIT.

package iso20022_pain_018_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Othr,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type AuthenticationChannel1Choice struct {
	Cd    ExternalAuthenticationChannel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type Authorisation1Choice struct {
	Cd    Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

// May be one of AUTH, FDET, FSUM, ILEV
type Authorisation1Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PstlAdr,omitempty"`
}

type CashAccount40 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Id,omitempty"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CtryOfBirth"`
}

type DatePeriod3 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 ToDt,omitempty"`
}

type Document struct {
	MndtSspnsnReq MandateSuspensionRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MndtSspnsnReq"`
}

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

// Must match the pattern [0-9]{2}
type Exact2NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalAccountIdentification1Code string

// May be no more than 4 items long
type ExternalAuthenticationChannel1Code string

// May be no more than 4 items long
type ExternalCashAccountType1Code string

// May be no more than 4 items long
type ExternalCategoryPurpose1Code string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be no more than 35 items long
type ExternalLocalInstrument1Code string

// May be no more than 4 items long
type ExternalMandateSetupReason1Code string

// May be no more than 4 items long
type ExternalMandateSuspensionReason1Code string

// May be no more than 4 items long
type ExternalOrganisationIdentification1Code string

// May be no more than 4 items long
type ExternalPersonIdentification1Code string

// May be no more than 4 items long
type ExternalProxyAccountType1Code string

// May be no more than 4 items long
type ExternalServiceLevel1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Othr,omitempty"`
}

// May be one of NEVR, YEAR, RATE, MIAN, QURT
type Frequency10Code string

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Tp,omitempty"`
	Prd    FrequencyPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prd,omitempty"`
	PtInTm FrequencyAndMoment1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PtInTm,omitempty"`
}

type Frequency37Choice struct {
	Cd    Frequency10Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code string

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Tp"`
	PtInTm Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Tp"`
	CntPerPrd float64        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CntPerPrd"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Issr,omitempty"`
}

type GroupHeader80 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Authstn,omitempty"`
	InitgPty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 InitgPty,omitempty"`
	InstgAgt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 InstgAgt,omitempty"`
	InstdAgt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 InstdAgt,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type Mandate17 struct {
	MndtId        Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MndtId"`
	MndtReqId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MndtReqId,omitempty"`
	Authntcn      MandateAuthentication1                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Authntcn,omitempty"`
	Tp            MandateTypeInformation2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Tp,omitempty"`
	Ocrncs        MandateOccurrences5                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Ocrncs,omitempty"`
	TrckgInd      bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 TrckgInd"`
	FrstColltnAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 FrstColltnAmt,omitempty"`
	ColltnAmt     ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 ColltnAmt,omitempty"`
	MaxAmt        ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MaxAmt,omitempty"`
	Adjstmnt      MandateAdjustment1                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Adjstmnt,omitempty"`
	Rsn           MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Rsn,omitempty"`
	CdtrSchmeId   PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CdtrSchmeId,omitempty"`
	Cdtr          PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cdtr"`
	CdtrAcct      CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CdtrAcct,omitempty"`
	CdtrAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CdtrAgt,omitempty"`
	UltmtCdtr     PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 UltmtCdtr,omitempty"`
	Dbtr          PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Dbtr"`
	DbtrAcct      CashAccount40                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 DbtrAcct,omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 DbtrAgt"`
	UltmtDbtr     PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 UltmtDbtr,omitempty"`
	MndtRef       Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MndtRef,omitempty"`
	RfrdDoc       []ReferredMandateDocument1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 RfrdDoc,omitempty"`
}

type MandateAdjustment1 struct {
	DtAdjstmntRuleInd bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 DtAdjstmntRuleInd"`
	Ctgy              Frequency37Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Ctgy,omitempty"`
	Amt               ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Amt,omitempty"`
	Rate              float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Rate,omitempty"`
}

type MandateAuthentication1 struct {
	MsgAuthntcnCd Max16Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MsgAuthntcnCd,omitempty"`
	Dt            ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Dt,omitempty"`
	Chanl         AuthenticationChannel1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Chanl,omitempty"`
}

type MandateClassification1Choice struct {
	Cd    MandateClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

// May be one of FIXE, USGB, VARI
type MandateClassification1Code string

type MandateOccurrences5 struct {
	SeqTp        SequenceType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SeqTp"`
	Frqcy        Frequency36Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Frqcy,omitempty"`
	Drtn         DatePeriod3       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Drtn,omitempty"`
	FrstColltnDt ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 FrstColltnDt,omitempty"`
	FnlColltnDt  ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 FnlColltnDt,omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type MandateSuspension3 struct {
	SspnsnReqId Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SspnsnReqId"`
	OrgnlMsgInf OriginalMessageInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 OrgnlMsgInf,omitempty"`
	SspnsnRsn   MandateSuspensionReason2    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SspnsnRsn"`
	OrgnlMndt   OriginalMandate9Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 OrgnlMndt"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SplmtryData,omitempty"`
}

type MandateSuspensionReason1Choice struct {
	Cd    ExternalMandateSuspensionReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type MandateSuspensionReason2 struct {
	Orgtr    PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Orgtr,omitempty"`
	Rsn      MandateSuspensionReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Rsn"`
	AddtlInf []Max105Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 AddtlInf,omitempty"`
}

type MandateSuspensionRequestV03 struct {
	GrpHdr            GroupHeader80        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 GrpHdr"`
	UndrlygSspnsnDtls []MandateSuspension3 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 UndrlygSspnsnDtls"`
	SplmtryData       []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SplmtryData,omitempty"`
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CtgyPurp,omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Clssfctn,omitempty"`
}

// May be no more than 105 items long
type Max105Text string

// May be no more than 128 items long
type Max128Text string

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 2048 items long
type Max2048Text string

// May be no more than 34 items long
type Max34Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 4 items long
type Max4Text string

// May be no more than 70 items long
type Max70Text string

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type OriginalMandate9Choice struct {
	OrgnlMndtId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 OrgnlMndtId,omitempty"`
	OrgnlMndt   Mandate17 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 OrgnlMndt,omitempty"`
}

type OriginalMessageInformation1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 MsgNmId"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CreDtTm,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 OrgId,omitempty"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PrvtId,omitempty"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CtctDtls,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Issr,omitempty"`
}

type ReferredMandateDocument1 struct {
	Tp      ReferredDocumentType4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Tp,omitempty"`
	Nb      Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Nb,omitempty"`
	CdtrRef Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 CdtrRef,omitempty"`
	RltdDt  ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 RltdDt,omitempty"`
}

// May be one of RCUR, OOFF
type SequenceType2Code string

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Cd,omitempty"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Prtry,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
