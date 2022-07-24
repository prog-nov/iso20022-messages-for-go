# ISO 20022 message elements Go language

For SEPA Credit Transfer, Domestic Transfer, Interbank systems, Multilateral clearing, Real-Time Gross Settlement. Miscellaneous message types pain.001, pacs.002, pacs.003, pacs.004, pacs.007, pacs.008, pacs.009, pacs.010, pacs.028, camt.026, camt.029, camt.030, camt.056, camt.087, head.001 and others.

## Large ISO 20022 message container: ISO 20022 Business Areas

| Business area       | Description                 | 
| ---------|:--------------------------------------:|
| caam     | ATM Management                         |
| camt     | Cash Management                        |
| catm     | Terminal Management                    |
| catp     | ATM Card Transactions                  |
| fxtr     | Foreign Exchange Trade                 |
| head     | Business Application Header            |
| pacs     | Payments Clearing and Settlement       |
| pain     | Payment Initiation                     |
| reda     | Reference Data                         |
| remt     | Payments Remittance advice             |
| secl     | Security Clearing                      |
| seev     | Security Events                        |
| semt     | Security Management                    |
| sese     | Securities Settlement                  |
| setr     | Securities Trade                       |
| tsin     | Trade Service Initiation               |
| tsmt     | Trade Services Management              |
| tsrv     | Trade Services                         | 
| admi     | Administration                         |
 	
## Popular business areas

The messages head, pacs, camt, pain, semt and others, are a large container in Go.

## Navigation

* ISO 20022 Message XSD Schema is located in the `XSD` directory.
> This catalog provides formal XML examples for ISO 20022 messages: head.001.001.02, pain.001.001.09, pain.002.001.10, pain.013.001.01, pacs.002.001.10, pacs.004.001.09, pacs.008.001.08, pacs.009.001.08, camt.016.001.04, camt.017.001.05, camt.026.001.08, camt.029.001.10, camt.030.001.05, camt.056.001.09, camt.087.001.07.

* Large container of **XSD schemes** in the `XSDtoGo/xsd` directory.

* ISO 20022 **XML Message Templates** is located in the `XML` directory.
> XSD templates are provided for messages: head.001.001.02, pain.001.001.09, pain.002.001.10, pain.013.001.01, pacs.002.001.10, pacs.004.001.09, pacs.008.001.08, pacs.009.001.08, camt.016.001.04, camt.017.001.05, camt.026.001.08, camt.029.001.10, camt.030.001.05, camt.056.001.09, camt.087.001.07.

* **Parser fo XSD to GO** ISO 20022 Messages is located in the `XSDtoGo` directory.
> It contains golang-library ISO20022 XML parsers from XSD Shema. This tool will help you implement this repository.

## Very Large Container of GO Files

This is a very large container of GO files for financial messages, so it is divided into 10 sub-repositories. 

* [iso20022-go-prefix-A](https://github.com/prog-nov/iso20022-go-prefix-A)
> ISO 20022 financial message structures, **prefix A** (Only A)
* [iso20022-go-prefix-B](https://github.com/prog-nov/iso20022-go-prefix-B)
> ISO 20022 financial message structures, **prefix B** (A, B, C)
* [iso20022-go-prefix-C2](https://github.com/prog-nov/iso20022-go-prefix-C2)
> ISO 20022 financial message structures, **prefix C2** (Only C)
* [iso20022-go-prefix-D](https://github.com/prog-nov/iso20022-go-prefix-D)
> ISO 20022 financial message structures, **prefix D** (C, D, E, F)
* [iso20022-go-prefix-F](https://github.com/prog-nov/iso20022-go-prefix-F)
> ISO 20022 financial message structures, **prefix F** (F, G, H, I, L, M)
* [iso20022-go-prefix-M](https://github.com/prog-nov/iso20022-go-prefix-M)
> ISO 20022 financial message structures, **prefix M** (M, O, P)
* [iso20022-go-prefix-P](https://github.com/prog-nov/iso20022-go-prefix-P)
> ISO 20022 financial message structures, **prefix P** (P, Q, R)
* [iso20022-go-prefix-R](https://github.com/prog-nov/iso20022-go-prefix-R)
> ISO 20022 financial message structures, **prefix R** (R, S)
* [iso20022-go-prefix-S](https://github.com/prog-nov/iso20022-go-prefix-S)
> ISO 20022 financial message structures, **prefix S** (S, T)
* [iso20022-go-prefix-T](https://github.com/prog-nov/iso20022-go-prefix-T)
> ISO 20022 financial message structures, **prefix T** (T, U, V, W, Y)

## Code Development Note

The official ISO 20022 catalog is constantly evolving, with new versions of ISO financial messages. This repository contains older versions, this only applies to the Go implementation, but examples and XSD shema are provided for more current versions.

You can use this repository to understand the business logic of the Go implementation and to update the financial messages to your version. Your welcome!

### PACS Catalog: Payments Clearing and Settlement — CreditTransfer

PACS Messages ISO 20022

* pacs.002.001 — FIToFIPaymentStatusReport —> FIToFIPaymentStatusReportV03
* pacs.003.001 — FIToFICustomerDirectDebit —> FIToFICustomerDirectDebitV01
* pacs.004.001 — PaymentReturn —> PaymentReturnV02
* pacs.007.001 — FIToFIPaymentReversal —> FIToFIPaymentReversalV02
* pacs.008.001 — FIToFICustomerCreditTransfer —> FIToFICustomerCreditTransferV02
* pacs.009.001 — FinancialInstitutionCreditTransfer —> FinancialInstitutionCreditTransferV02
* pacs.010.001 — FinancialInstitutionDirectDebit —> FinancialInstitutionDirectDebitV02
* pacs.028.001 — FIToFIPaymentStatusRequest —> FIToFIPaymentStatusRequestV01
* and more.

### PAIN catalog: Payments Initiation

PAIN Messages ISO 20022

* CreditorPaymentActivationRequestStatusReportV01
* CreditorPaymentActivationRequestStatusReportV02
* CreditorPaymentActivationRequestStatusReportV03
* CreditorPaymentActivationRequestStatusReportV04
* CreditorPaymentActivationRequestStatusReportV05
* CreditorPaymentActivationRequestStatusReportV06
* CreditorPaymentActivationRequestV01
* CreditorPaymentActivationRequestV02
* CreditorPaymentActivationRequestV03
* CreditorPaymentActivationRequestV04
* CreditorPaymentActivationRequestV05
* CreditorPaymentActivationRequestV06
* CustomerCreditTransferInitiationV02
* CustomerCreditTransferInitiationV03
* CustomerCreditTransferInitiationV04
* CustomerCreditTransferInitiationV05
* CustomerCreditTransferInitiationV06
* CustomerCreditTransferInitiationV07
* CustomerCreditTransferInitiationV08
* CustomerDirectDebitInitiationV01
* CustomerDirectDebitInitiationV02
* CustomerDirectDebitInitiationV03
* CustomerDirectDebitInitiationV04
* CustomerDirectDebitInitiationV05
* CustomerDirectDebitInitiationV06
* CustomerDirectDebitInitiationV07
* CustomerPaymentReversalV01
* CustomerPaymentReversalV02
* CustomerPaymentReversalV03
* CustomerPaymentReversalV04
* CustomerPaymentReversalV05
* CustomerPaymentReversalV06
* CustomerPaymentReversalV07
* CustomerPaymentStatusReportV03
* CustomerPaymentStatusReportV04
* CustomerPaymentStatusReportV05
* CustomerPaymentStatusReportV06
* CustomerPaymentStatusReportV07
* CustomerPaymentStatusReportV08
* MandateAcceptanceReportV01
* MandateAcceptanceReportV02
* MandateAcceptanceReportV03
* MandateAcceptanceReportV04
* MandateAcceptanceReportV05
* MandateAmendmentRequestV01
* MandateAmendmentRequestV02
* MandateAmendmentRequestV03
* MandateAmendmentRequestV04
* MandateAmendmentRequestV05
* MandateCancellationRequestV01
* MandateCancellationRequestV02
* MandateCancellationRequestV03
* MandateCancellationRequestV04
* MandateCancellationRequestV05
* MandateCopyRequestV01
* MandateInitiationRequestV01
* MandateInitiationRequestV02
* MandateInitiationRequestV03
* MandateInitiationRequestV04
* MandateInitiationRequestV05
* MandateSuspensionRequestV01
* PaymentCancellationRequestV01
* PaymentStatusReportV02

### Cash Management — CAMT

Camt Messages ISO 20022

* camt.003.001 — GetAccount
* camt.004.001 — ReturnAccount
* camt.005.001 — GetTransaction
* camt.006.001 — ReturnTransaction
* camt.007.001 — ModifyTransaction
* camt.008.001 — CancelTransaction
* camt.009.001— GetLimit
* camt.010.001 — ReturnLimit
* camt.011.001 — ModifyLimit
* camt.012.001 — DeleteLimit
* camt.013.001 — GetMemberV
* camt.014.001 — ReturnMember
* camt.015.001— ModifyMember
* camt.016.001 — GetCurrencyExchangeRate
* camt.017.001 — ReturnCurrencyExchangeRate
* camt.018.001 — GetBusinessDayInformation
* camt.019.001 — ReturnBusinessDayInformation
* camt.020.001 — GetGeneralBusinessInformation
* camt.021.001— ReturnGeneralBusinessInformation
* camt.023.001 — BackupPayment
* camt.024.001 — ModifyStandingOrder
* camt.025.001 — ReceiptV
* camt.026.001 — UnableToApply
* camt.027.001 — ClaimNonReceipt
* camt.028.001 — AdditionalPaymentInformation
* camt.029.001 — ResolutionOfInvestigation
* camt.030.001 — NotificationOfCaseAssignment
* camt.031.001 — RejectInvestigation
* camt.032.001 — CancelCaseAssignment
* camt.033.001 — RequestForDuplicate
* camt.034.001 — Duplicate
* camt.035.001 — ProprietaryFormatInvestigation
* camt.036.001 — DebitAuthorisationResponse
* camt.037.001 — DebitAuthorisationRequest
* camt.038.001 — CaseStatusReportRequest
* camt.039.001 — CaseStatusReport
* camt.040.001 — FundEstimatedCashForecastReport
* camt.041.001 — FundConfirmedCashForecastReport
* camt.042.001 — FundDetailedEstimatedCashForecastReport
* camt.043.001 — FundDetailedConfirmedCashForecastReport
* camt.044.001 — FundConfirmedCashForecastReportCancellation
* camt.045.001 — FundDetailedConfirmedCashForecastReportCancellation
* camt.046.001 — GetReservation
* camt.047.001 — ReturnReservation
* camt.048.001 — ModifyReservation
* camt.049.001 — DeleteReservation
* camt.050.001 — LiquidityCreditTransfer
* camt.051.001 — LiquidityDebitTransfer
* camt.052.001 — BankToCustomerAccountReport
* camt.053.001 — BankToCustomerStatement
* camt.054.001 — BankToCustomerDebitCreditNotification
* camt.055.001 — CustomerPaymentCancellationRequest
* camt.056.001 — FIToFIPaymentCancellationRequest
* camt.057.001 — NotificationToReceive
* camt.058.001 — NotificationToReceiveCancellationAdvice
* camt.059.001 — NotificationToReceiveStatusReport
* camt.060.001 — AccountReportingRequest
* camt.061.001 — PayInCall
* camt.062.001 — PayInSchedule
* camt.063.001 — PayInEventAcknowledgement
* camt.066.001 — IntraBalanceMovementInstruction
* camt.067.001 — IntraBalanceMovementStatusAdvice
* camt.068.001 — IntraBalanceMovementConfirmation
* camt.069.001 — GetStandingOrder
* camt.070.001 — ReturnStandingOrder
* camt.071.001 — DeleteStandingOrder
* camt.072.001 — IntraBalanceMovementModificationRequest
* camt.073.001 — IntraBalanceMovementModificationRequestStatusAdvice
* camt.074.001 — IntraBalanceMovementCancellationRequest
* camt.075.001 — IntraBalanceMovementCancellationRequestStatusAdvice
* camt.078.001 — IntraBalanceMovementQuery
* camt.079.001 — IntraBalanceMovementQueryResponse
* camt.080.001 — IntraBalanceMovementModificationQuery
* camt.081.001 — IntraBalanceMovementModificationReport
* camt.082.001 — IntraBalanceMovementCancellationQuery
* camt.083.001 — IntraBalanceMovementCancellationReport
* camt.084.001 — IntraBalanceMovementPostingReport
* camt.085.001 — IntraBalanceMovementPendingReport
* camt.086.001 — BankServicesBillingStatement
* camt.087.001 — RequestToModifyPayment
* camt.088.001 — NetReport
* camt.101.001 — CreateLimit
* camt.102.001 — CreateStandingOrder
* camt.103.001 — CreateReservation
* camt.104.001 — CreateMember
* camt.105.001 — ChargesPaymentNotification
* camt.106.001 — ChargesPaymentRequest
* camt.107.001 — ChequePresentmentNotification
* camt.108.001 — ChequeCancellationOrStopRequest
* camt.109.001 — ChequeCancellationOrStopReport

### Other ISO Messages catalog

All directories are intuitively named, so you can easily find the ISO 20022 message group you need.