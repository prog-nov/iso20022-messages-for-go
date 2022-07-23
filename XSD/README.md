# ISO 20022 Message XSD Schema

The catalog contains ISO 20022 XSD message templates. These are the most common ISO messages that are used. There is also an explanation of each message.

In the main directory of this repository there is a large GO implementation for Business Areas: camt, catm, fxtr, pacs, pain, semt, tsin and more.

## Overview

XSD templates are provided for messages: head.001.001.02, pain.001.001.09, pain.002.001.10, pain.013.001.01, pacs.002.001.10, pacs.004.001.09, pacs.008.001.08, pacs.009.001.08, camt.016.001.04, camt.017.001.05, camt.026.001.08, camt.029.001.10, camt.030.001.05, camt.056.001.09, camt.087.001.07.

* BusinessApplicationHeaderV02
* CustomerCreditTransferInitiationV09
* CreditorPaymentActivationRequestV01
* CustomerPaymentStatusReportV10
* FIToFIPaymentStatusReportV10
* PaymentReturnV09
* FIToFICustomerCreditTransferV08
* FinancialInstitutionCreditTransferV08
* GetCurrencyExchangeRateV04
* ReturnCurrencyExchangeRateV05
* UnableToApplyV08
* ResolutionOfInvestigationV10
* NotificationOfCaseAssignmentV05
* FIToFIPaymentCancellationRequestV09
* RequestToModifyPaymentV07

### Business application header (head.001.001.002)

The Business Application Header (BusinessApplicationHeaderV02) is a header that has been defined by the ISO 20022 community, that can form part of an ISO 20022 business message. Specifically, the BAH is an ISO20022 message definition (head.001.001.0x) which can be combined with any other ISO20022 message definition to form a business message.

### Customer Credit Transfer Initiation (pain.001.001.09)

The CustomerCreditTransferInitiation message can contain one or more customer credit transfer instructions.

The CustomerCreditTransferInitiation message is sent by the initiating party to the forwarding agent or debtor agent. It is used to request movement of funds from the debtor account to a creditor.

### Customer Payment Status Report (pain.002.001.10)

The CustomerPaymentStatusReport message is exchanged between an agent and a non-financial institution customer to provide status information on instructions previously sent. Its usage will always be governed by a bilateral agreement between the agent and the non-financial institution customer.

The CustomerPaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.

### Creditor Payment Activation Request (pain.013.001.01)

The CreditorPaymentActivationRequest message is sent by the Creditor sending party to the Debtor receiving party, directly or through agents. It is used by a Creditor to request movement of funds from the debtor account to a creditor.

### FI To FI Payment Status Report (pacs.002.001.10)

The FIToFIPaymentStatusReport message is exchanged between agents to provide status information about instructions previously sent. Its usage will always be governed by a bilateral agreement between the agents.

The FIToFIPaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.

### Payment Return (pacs.004.001.09)

The PaymentReturn message is sent by an agent to the previous agent in the payment chain to undo a payment previously settled.

The PaymentReturn message is exchanged between agents to return funds after settlement of credit transfer instructions (that is {mx0567|FIToFICustomerCreditTransfer} message and {mx0568|FinancialInstitutionCreditTransfer} message) or direct debit instructions ({mx0564|FIToFICustomerDirectDebit} message).

### FI To FI Customer Credit Transfer (pacs.008.001.08)

The FinancialInstitutionToFinancialInstitutionCustomerCreditTransfer message is sent by the debtor agent to the creditor agent, directly or through other agents and/or a payment clearing and settlement system. It is used to move funds from a debtor account to a creditor.

The FIToFICustomerCreditTransfer message is exchanged between agents and can contain one or more customer credit transfer instructions.

### Financial Institution Credit Transfer (pacs.009.001.08)

The FinancialInstitutionCreditTransfer message is sent by a debtor financial institution to a creditor financial institution, directly or through other agents and/or a payment clearing and settlement system. It is used to move funds from a debtor account to a creditor, where both debtor and creditor are financial institutions.

The FinancialInstitutionCreditTransfer message is exchanged between agents and can contain one or more credit transfer instructions where debtor and creditor are both financial institutions.

### Get Currency Exchange Rate (camt.016.001.04)

The GetCurrencyExchangeRate message is sent by a member to the transaction administrator. It is used to request information on static data maintained by the transaction administrator and related to currency exchange details as maintained for the system operations by the transaction administrator.

The transaction administrator is in charge of providing the members with business information. The term business information covers all information related to the management of the system, i.e., not related to the transactions created into the system. The type of business information available can vary depending on the system.

### Return Currency Exchange Rate (camt.017.001.05)

The ReturnCurrencyExchangeRate message is sent by the transaction administrator to a member of the system. It is used to provide information on static data and related to currency exchange details as maintained for system operations by the transaction administrator.

The transaction administrator is in charge of providing the members with business information. The term business information covers all information related to the management of the system, that is, not related to the transactions or requests created in the system. The type of business information available can vary depending on the system.

### Unable To Apply (camt.026.001.08)

The UnableToApply message is sent by a case creator or a case assigner to a case assignee. This message is used to initiate an investigation of a payment instruction that cannot be executed or reconciled.

The UnableToApply case occurs in two situations:

* an agent cannot execute the payment instruction due to missing or incorrect information;

* a creditor cannot reconcile the payment entry as it is received unexpectedly, or missing or incorrect information prevents reconciliation.

### Resolution Of Investigation (camt.029.001.10)

The ResolutionOfInvestigation message is sent by a case assignee to a case creator/case assigner. This message is used to inform of the resolution of a case, and optionally provides details about.

### Notification Of Case Assignment (camt.030.001.05)

NotificationOfCaseAssignment this message notifies the case assigner that a case has been assigned to a next party in the payment chain or that the assignee will do the cancellation, modification or correction himself.


### FI To FI Payment Cancellation Request (camt.056.001.09)

The FIToFIPaymentCancellationRequest message is sent by a case creator/case assigner to a case assignee.

The FIToFIPaymentCancellationRequest message supports both the request for cancellation (the instructed agent — or assignee — has not yet processed and forwarded the payment instruction) as well as the request for refund (payment has been fully processed already by the instructed agent — or assignee).

### Request To Modify Payment (camt.087.001.07)

The RequestToModifyPayment message is sent by a case creator/case assigner to a case assignee. This message is used to request the modification of characteristics of an original payment instruction.










