// Code generated by download. DO NOT EDIT.

package iso20022_seev_027_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AgentCAStandingInstructionStatusAdviceV01 struct {
	Id                    DocumentIdentification8                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Id"`
	AgtCAStgInstrReqId    DocumentIdentification8                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 AgtCAStgInstrReqId,omitempty"`
	AgtCAStgInstrCxlReqId DocumentIdentification8                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 AgtCAStgInstrCxlReqId,omitempty"`
	StgInstrGnlInf        CorporateActionStandingInstructionGeneralInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 StgInstrGnlInf"`
	StgInstrReqSts        StandingInstructionStatus1Choice                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 StgInstrReqSts,omitempty"`
	StgInstrCxlReqSts     StandingInstructionCancellationStatus1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 StgInstrCxlReqSts,omitempty"`
}

type AlternateSecurityIdentification3 struct {
	Id         Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Id"`
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 DmstIdSrc,omitempty"`
	PrtryIdSrc Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 PrtryIdSrc,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, INCR, INTR, LIQU, MCAL, MRGR, ODLT, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, OTHR
type CorporateActionEventType2Code string

type CorporateActionEventType2FormatChoice struct {
	Cd    CorporateActionEventType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Cd,omitempty"`
	Prtry GenericIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Prtry,omitempty"`
}

type CorporateActionStandingInstructionCancellationProcessingStatus1 struct {
	Sts      ProcessedStatus4FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Sts"`
	AddtlInf Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 AddtlInf,omitempty"`
}

type CorporateActionStandingInstructionCancellationRejectionStatus1 struct {
	Rsn      []RejectionReason10FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Rsn"`
	AddtlInf Max350Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 AddtlInf,omitempty"`
}

type CorporateActionStandingInstructionGeneralInformation1 struct {
	StgInstrTp     StandingInstructionType1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 StgInstrTp"`
	EvtTp          []CorporateActionEventType2FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 EvtTp,omitempty"`
	InstgPtyId     PartyIdentification2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 InstgPtyId"`
	ClntStgInstrId Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 ClntStgInstrId"`
	AcctDtls       []IncludedAccount1                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 AcctDtls,omitempty"`
	UndrlygScty    FinancialInstrumentDescription3         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 UndrlygScty,omitempty"`
}

type CorporateActionStandingInstructionProcessingStatus1 struct {
	Sts      ProcessedStatus3FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Sts"`
	AddtlInf Max350Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 AddtlInf,omitempty"`
}

type CorporateActionStandingInstructionRejectionStatus1 struct {
	Rsn      []RejectionReason20FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Rsn"`
	AddtlInf Max350Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 AddtlInf,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	AgtCAStgInstrStsAdvc AgentCAStandingInstructionStatusAdviceV01 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 AgtCAStgInstrStsAdvc"`
}

type DocumentIdentification8 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 CreDtTm,omitempty"`
}

type FinancialInstrumentDescription3 struct {
	SctyId     SecurityIdentification7    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 SctyId"`
	PlcOfListg MICIdentifier              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 PlcOfListg,omitempty"`
	SfkpgPlc   PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 SfkpgPlc,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Issr"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IncludedAccount1 struct {
	SctiesAcctId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 SctiesAcctId"`
	InclInd      bool      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 InclInd"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// May be no more than 70 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Adr,omitempty"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 BICOrBEI,omitempty"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 NmAndAdr,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Ctry"`
}

// May be one of RECE, PEND, PACK
type ProcessedStatus3Code string

type ProcessedStatus3FormatChoice struct {
	Cd    ProcessedStatus3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Cd,omitempty"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Prtry,omitempty"`
}

// May be one of RECE, COMP, PEND
type ProcessedStatus4Code string

type ProcessedStatus4FormatChoice struct {
	Cd    ProcessedStatus4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Cd,omitempty"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Prtry,omitempty"`
}

// May be one of FAIL
type RejectionReason10Code string

type RejectionReason10FormatChoice struct {
	Cd    RejectionReason10Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Cd,omitempty"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Prtry,omitempty"`
}

// May be one of FAIL, CASA, CORR, STAN, NOHO
type RejectionReason20Code string

type RejectionReason20FormatChoice struct {
	Cd    RejectionReason20Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Cd,omitempty"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Prtry,omitempty"`
}

type SecurityIdentification7 struct {
	ISIN   ISINIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 ISIN,omitempty"`
	OthrId AlternateSecurityIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 OthrId,omitempty"`
	Desc   Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Desc,omitempty"`
}

type StandingInstructionCancellationStatus1Choice struct {
	PrcdSts  CorporateActionStandingInstructionCancellationProcessingStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 PrcdSts,omitempty"`
	RjctdSts CorporateActionStandingInstructionCancellationRejectionStatus1  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 RjctdSts,omitempty"`
}

type StandingInstructionStatus1Choice struct {
	PrcdSts  CorporateActionStandingInstructionProcessingStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 PrcdSts,omitempty"`
	RjctdSts CorporateActionStandingInstructionRejectionStatus1  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 RjctdSts,omitempty"`
}

// May be one of CASH, PAYM, SECU
type StandingInstructionType1Code string

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