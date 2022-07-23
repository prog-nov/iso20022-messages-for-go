package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02400208 struct {
	XMLName xml.Name                                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.024.002.08 Document"`
	Message *SecuritiesSettlementTransactionStatusAdvice002V08 `xml:"SctiesSttlmTxStsAdvc"`
}

func (d *Document02400208) AddMessage() *SecuritiesSettlementTransactionStatusAdvice002V08 {
	d.Message = new(SecuritiesSettlementTransactionStatusAdvice002V08)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionStatusAdvice to an account owner to advise the status of a securities settlement transaction instruction previously sent by the account owner or the status of a settlement transaction existing in the books of the servicer for the account of the owner. The status may be a processing, pending processing, internal matching, matching and/or settlement status.
// The status advice may be sent as a response to the request of the account owner or not.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionStatusAdvice002V08 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *iso20022.TransactionIdentifications38 `xml:"TxId"`

	// Link to another transaction - provided for information only.
	Linkages *iso20022.Linkages50 `xml:"Lnkgs,omitempty"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *iso20022.ProcessingStatus63Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *iso20022.MatchingStatus32Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *iso20022.MatchingStatus32Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *iso20022.SettlementStatus22Choice `xml:"SttlmSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.TransactionDetails100 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionStatusAdvice002V08) AddTransactionIdentification() *iso20022.TransactionIdentifications38 {
	s.TransactionIdentification = new(iso20022.TransactionIdentifications38)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionStatusAdvice002V08) AddLinkages() *iso20022.Linkages50 {
	s.Linkages = new(iso20022.Linkages50)
	return s.Linkages
}

func (s *SecuritiesSettlementTransactionStatusAdvice002V08) AddProcessingStatus() *iso20022.ProcessingStatus63Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus63Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdvice002V08) AddInferredMatchingStatus() *iso20022.MatchingStatus32Choice {
	s.InferredMatchingStatus = new(iso20022.MatchingStatus32Choice)
	return s.InferredMatchingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdvice002V08) AddMatchingStatus() *iso20022.MatchingStatus32Choice {
	s.MatchingStatus = new(iso20022.MatchingStatus32Choice)
	return s.MatchingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdvice002V08) AddSettlementStatus() *iso20022.SettlementStatus22Choice {
	s.SettlementStatus = new(iso20022.SettlementStatus22Choice)
	return s.SettlementStatus
}

func (s *SecuritiesSettlementTransactionStatusAdvice002V08) AddTransactionDetails() *iso20022.TransactionDetails100 {
	s.TransactionDetails = new(iso20022.TransactionDetails100)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionStatusAdvice002V08) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
