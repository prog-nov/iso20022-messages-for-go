// Code generated by download. DO NOT EDIT.

package iso20022_seev_025_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AgentCAStandingInstructionRequestV01 struct {
	Id             DocumentIdentification8                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Id"`
	StgInstrGnlInf CorporateActionStandingInstructionGeneralInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 StgInstrGnlInf"`
	StgInstrDtls   CorporateActionStandingInstruction1                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 StgInstrDtls"`
	CtctDtls       ContactPerson1                                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 CtctDtls,omitempty"`
}

type AlternateSecurityIdentification3 struct {
	Id         Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Id"`
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 DmstIdSrc,omitempty"`
	PrtryIdSrc Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 PrtryIdSrc,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type CashAccount17 struct {
	AcctId     CashAccountIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 AcctId"`
	PmtCcy     ActiveCurrencyCode               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 PmtCcy"`
	AcctOwnrId PartyIdentification2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 AcctOwnrId,omitempty"`
	CrspdtBkId BICIdentifier                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 CrspdtBkId"`
}

type CashAccountIdentification1Choice struct {
	IBAN     IBANIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 IBAN,omitempty"`
	BBAN     BBANIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 BBAN,omitempty"`
	UPIC     UPICIdentifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 UPIC,omitempty"`
	DmstAcct SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 DmstAcct,omitempty"`
}

type ContactIdentification4 struct {
	Nm       Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Nm"`
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 NmPrfx,omitempty"`
	GvnNm    Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 GvnNm,omitempty"`
	Role     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Role,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 PhneNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 EmailAdr,omitempty"`
}

type ContactPerson1 struct {
	CtctPrsn ContactIdentification4     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 CtctPrsn"`
	InstnId  PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 InstnId,omitempty"`
}

// May be one of ACTV, ATTI, BIDS, BONU, BPUT, BRUP, CAPG, CAPI, CERT, CHAN, CLSA, CONS, CONV, COOP, DECR, DETI, DFLT, DLST, DRAW, DRIP, DSCL, DTCH, DVCA, DVOP, DVSC, DVSE, EXOF, EXRI, EXTM, EXWA, INCR, INTR, LIQU, MCAL, MRGR, ODLT, PARI, PCAL, PDEF, PINK, PLAC, PPMT, PRED, PRII, PRIO, REDM, REDO, REMK, RHDI, RHTS, SHPR, SMAL, SOFF, SPLF, SPLR, SUSP, TEND, TREC, WRTH, WTRC, OTHR
type CorporateActionEventType2Code string

type CorporateActionEventType2FormatChoice struct {
	Cd    CorporateActionEventType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Cd,omitempty"`
	Prtry GenericIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Prtry,omitempty"`
}

type CorporateActionStandingInstruction1 struct {
	NetOrGrss         StandingInstructionGrossNet1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 NetOrGrss,omitempty"`
	CshDstrbtnDtls    CashAccount17                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 CshDstrbtnDtls,omitempty"`
	SctiesDstrbtnDtls SecuritiesAccount6               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 SctiesDstrbtnDtls,omitempty"`
	AddtlInf          Max350Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 AddtlInf,omitempty"`
}

type CorporateActionStandingInstructionGeneralInformation1 struct {
	StgInstrTp     StandingInstructionType1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 StgInstrTp"`
	EvtTp          []CorporateActionEventType2FormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 EvtTp,omitempty"`
	InstgPtyId     PartyIdentification2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 InstgPtyId"`
	ClntStgInstrId Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 ClntStgInstrId"`
	AcctDtls       []IncludedAccount1                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 AcctDtls,omitempty"`
	UndrlygScty    FinancialInstrumentDescription3         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 UndrlygScty,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	AgtCAStgInstrReq AgentCAStandingInstructionRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 AgtCAStgInstrReq"`
}

type DocumentIdentification8 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 CreDtTm,omitempty"`
}

type FinancialInstrumentDescription3 struct {
	SctyId     SecurityIdentification7    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 SctyId"`
	PlcOfListg MICIdentifier              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 PlcOfListg,omitempty"`
	SfkpgPlc   PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 SfkpgPlc,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Issr"`
}

// Must match the pattern [a-zA-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBANIdentifier string

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
	SctiesAcctId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 SctiesAcctId"`
	InclInd      bool      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 InclInd"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 256 items long
type Max256Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// May be no more than 70 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 BICOrBEI,omitempty"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 NmAndAdr,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Ctry"`
}

type SecuritiesAccount6 struct {
	SctyId       SecurityIdentification7    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 SctyId"`
	SctiesAcctId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 SctiesAcctId"`
	AcctOwnrId   PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 AcctOwnrId,omitempty"`
	SfkpgPlc     PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 SfkpgPlc"`
	RegnDtls     Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 RegnDtls,omitempty"`
}

type SecurityIdentification7 struct {
	ISIN   ISINIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 ISIN,omitempty"`
	OthrId AlternateSecurityIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 OthrId,omitempty"`
	Desc   Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Desc,omitempty"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Id"`
}

// May be one of GROS, NETT
type StandingInstructionGrossNet1Code string

// May be one of CASH, PAYM, SECU
type StandingInstructionType1Code string

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

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
