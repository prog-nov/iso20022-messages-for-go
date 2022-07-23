// Code generated by download. DO NOT EDIT.

package iso20022_seev_047_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of USUF, OWNR, BOWN
type AccountOwnershipType5Code string

type AccountSubLevel22 struct {
	NonDscldShrhldgQty   FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NonDscldShrhldgQty,omitempty"`
	BlwThrshldShrhldgQty FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 BlwThrshldShrhldgQty,omitempty"`
	Dsclsr               []AccountSubLevel23                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Dsclsr,omitempty"`
}

type AccountSubLevel23 struct {
	SfkpgAcct  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 SfkpgAcct,omitempty"`
	AcctHldr   PartyIdentification243 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 AcctHldr"`
	ShrhldgBal []ShareholdingBalance1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 ShrhldgBal"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type ContactIdentification2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NmPrfx,omitempty"`
	GvnNm    Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 GvnNm,omitempty"`
	Nm       Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Nm"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 FaxNb,omitempty"`
	EmailAdr Max256Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 EmailAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Dt,omitempty"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 DtTm,omitempty"`
}

type DateAndPlaceOfBirth2 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 CityOfBirth,omitempty"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 CtryOfBirth,omitempty"`
}

type DateCode20Choice struct {
	Cd    DateType1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Prtry,omitempty"`
}

type DateFormat46Choice struct {
	Dt   DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Dt,omitempty"`
	DtCd DateCode20Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 DtCd,omitempty"`
}

type DateFormat57Choice struct {
	Dt   ISODate          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Dt,omitempty"`
	DtCd DateCode20Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 DtCd,omitempty"`
}

// May be one of UKWN
type DateType1Code string

type Disclosure2Choice struct {
	NoDsclsr          NoReasonCode           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NoDsclsr,omitempty"`
	SfkpgAcctAndHldgs []SafekeepingAccount11 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 SfkpgAcctAndHldgs,omitempty"`
}

type DisclosureRequestIdentification1 struct {
	IssrDsclsrReqId      Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 IssrDsclsrReqId"`
	FinInstrmId          SecurityIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 FinInstrmId"`
	ShrhldrsDsclsrRcrdDt DateFormat46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 ShrhldrsDsclsrRcrdDt"`
}

type Document struct {
	ShrhldrsIdDsclsrRspn ShareholdersIdentificationDisclosureResponseV02 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 ShrhldrsIdDsclsrRspn"`
}

// May be one of ELIG, RETL, PROF
type Eligibility1Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity18Choice struct {
	Unit    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Unit,omitempty"`
	FaceAmt float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 FaceAmt,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 SchmeNm,omitempty"`
}

// Must match the pattern [A-U]{1,1}[0-9]{0,4}
type ISICIdentifier string

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type ISOYear time.Time

func (t *ISOYear) UnmarshalText(text []byte) error {
	return (*xsdGYear)(t).UnmarshalText(text)
}
func (t ISOYear) MarshalText() ([]byte, error) {
	return xsdGYear(t).MarshalText()
}

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Cd,omitempty"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Prtry,omitempty"`
}

type IdentificationType45Choice struct {
	Cd    TypeOfIdentification4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Cd,omitempty"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Prtry,omitempty"`
}

type InvestorType1Choice struct {
	Cd    Eligibility1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

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

// May be no more than 50 items long
type Max50Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// May be no more than 70 items long
type Max70Text string

type NameAndAddress17 struct {
	Nm  Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Nm"`
	Adr PostalAddress26 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type NaturalPersonIdentification1 struct {
	Id   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Id"`
	IdTp IdentificationType45Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 IdTp,omitempty"`
}

// May be one of NORE
type NoReasonCode string

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Tp"`
}

type Ownership1 struct {
	OwnrshTp   OwnershipType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 OwnrshTp,omitempty"`
	OwnrshPctg float64              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 OwnrshPctg,omitempty"`
	UsfrctPctg float64              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 UsfrctPctg,omitempty"`
}

type OwnershipType3Choice struct {
	Cd    AccountOwnershipType5Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Cd,omitempty"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Prtry,omitempty"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 LastPgInd"`
}

type PartyIdentification195Choice struct {
	AnyBIC  AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 AnyBIC,omitempty"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 PrtryId,omitempty"`
	LEI     LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 LEI,omitempty"`
}

type PartyIdentification198Choice struct {
	NtlRegnNb Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NtlRegnNb,omitempty"`
	LEI       LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 LEI,omitempty"`
	AnyBIC    AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 AnyBIC,omitempty"`
	ClntId    Max50Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 ClntId,omitempty"`
	PrtryId   GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 PrtryId,omitempty"`
}

type PartyIdentification201 struct {
	NmAndAdr PersonName2                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NmAndAdr,omitempty"`
	Id       PartyIdentification198Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Id"`
}

type PartyIdentification202 struct {
	NmAndAdr PersonName1                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NmAndAdr"`
	Id       NaturalPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Id"`
}

type PartyIdentification205Choice struct {
	LglPrsn  PartyIdentification201 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 LglPrsn,omitempty"`
	NtrlPrsn PartyIdentification202 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NtrlPrsn,omitempty"`
}

type PartyIdentification217 struct {
	NmAndAdr        PersonName3                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NmAndAdr"`
	EmailAdr        Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 EmailAdr,omitempty"`
	Id              NaturalPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Id"`
	Ntlty           CountryCode                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Ntlty,omitempty"`
	DtAndPlcOfBirth DateAndPlaceOfBirth2         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 DtAndPlcOfBirth,omitempty"`
	InvstrTp        InvestorType1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 InvstrTp,omitempty"`
	Ownrsh          Ownership1                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Ownrsh,omitempty"`
}

type PartyIdentification218 struct {
	Role PartyRole6Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Role"`
	Id   PartyIdentification205Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Id"`
}

type PartyIdentification219 struct {
	NmAndAdr PersonName2                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NmAndAdr"`
	Id       PartyIdentification195Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Id"`
	CtctPrsn ContactIdentification2       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 CtctPrsn,omitempty"`
}

type PartyIdentification237 struct {
	NmAndAdr        NameAndAddress17             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NmAndAdr"`
	EmailAdr        Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 EmailAdr,omitempty"`
	Id              PartyIdentification198Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Id"`
	CtryOfIncorprtn CountryCode                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 CtryOfIncorprtn,omitempty"`
	YrOfIncorprtn   ISOYear                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 YrOfIncorprtn,omitempty"`
	ActvtyInd       ISICIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 ActvtyInd,omitempty"`
	InvstrTp        InvestorType1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 InvstrTp,omitempty"`
	Ownrsh          Ownership1                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Ownrsh,omitempty"`
}

type PartyIdentification243 struct {
	LglPrsn  []PartyIdentification237 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 LglPrsn,omitempty"`
	NtrlPrsn []PartyIdentification217 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NtrlPrsn,omitempty"`
}

// May be one of DIST, LGRD, DECM
type PartyRole2Code string

type PartyRole6Choice struct {
	Cd    PartyRole2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Prtry,omitempty"`
}

type PersonName1 struct {
	FrstNm Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 FrstNm"`
	Srnm   Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Srnm"`
	Adr    PostalAddress26 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Adr,omitempty"`
}

type PersonName2 struct {
	Nm  Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Nm"`
	Adr PostalAddress26 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Adr,omitempty"`
}

type PersonName3 struct {
	NmPrfx NamePrefix2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 NmPrfx,omitempty"`
	FrstNm Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 FrstNm"`
	Srnm   Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Srnm"`
	Adr    PostalAddress26 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Adr,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress26 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 BldgNb,omitempty"`
	PstBx       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 PstBx,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Ctry"`
}

type SafekeepingAccount11 struct {
	SfkpgAcct            Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 SfkpgAcct"`
	AcctSvcr             PartyIdentification195Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 AcctSvcr"`
	ShrhldgBalOnOwnAcct  FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 ShrhldgBalOnOwnAcct"`
	ShrhldgBalOnClntAcct FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 ShrhldgBalOnClntAcct"`
	TtlShrhldgBal        FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 TtlShrhldgBal"`
	AcctSubLvl           AccountSubLevel22                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 AcctSubLvl,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Desc,omitempty"`
}

type ShareholdersIdentificationDisclosureResponseV02 struct {
	Pgntn            Pagination1                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Pgntn,omitempty"`
	IssrDsclsrReqRef DisclosureRequestIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 IssrDsclsrReqRef"`
	DsclsrRspnId     Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 DsclsrRspnId"`
	RspndgIntrmy     PartyIdentification219           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 RspndgIntrmy"`
	DsclsrInf        Disclosure2Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 DsclsrInf"`
	SplmtryData      []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 SplmtryData,omitempty"`
}

type ShareholdingBalance1 struct {
	ShrhldgTp        ShareholdingType1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 ShrhldgTp"`
	Qty              FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Qty"`
	InitlDtOfShrhldg DateFormat57Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 InitlDtOfShrhldg,omitempty"`
	ThrdPty          []PartyIdentification218            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 ThrdPty,omitempty"`
	SplmtryData      []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 SplmtryData,omitempty"`
}

// May be one of BENE, NOMI, OOAC, UKWN
type ShareholdingType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.047.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of ARNU, CUST, CORP, DRLC, IDCD, NRIN, CCPT, SOCS, TXID
type TypeOfIdentification4Code string

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

type xsdGYear time.Time

func (t *xsdGYear) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006")
}
func (t xsdGYear) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006")
}
func (t xsdGYear) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYear) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
