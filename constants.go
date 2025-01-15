package main

const DatabasePath = "il-campaign-disclosures.db"

var (
	Candidates = Table{
		Name: "candidates",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/Candidates.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "last_name", RawName: "FirstName", Type: ColumnTypeString},
			{Name: "first_name", RawName: "LastName", Type: ColumnTypeString},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "office", RawName: "Office", Type: ColumnTypeString},
			{Name: "district_type", RawName: "DistrictType", Type: ColumnTypeString},
			{Name: "district", RawName: "District", Type: ColumnTypeString},
			{Name: "residence_county", RawName: "ResidenceCounty", Type: ColumnTypeString},
			{Name: "party_affiliation", RawName: "PartyAffiliation", Type: ColumnTypeString},
			{Name: "redaction_requested", RawName: "RedactionRequested", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "last_name", "first_name"},
	}
	CandidateCommittees = Table{
		Name: "candidate_committees",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/CmteCandidateLinks.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "candidate_id", RawName: "CandidateID", Type: ColumnTypeInt, NotNullable: true},
		},
		IndexedColumns: []string{"id", "committee_id", "candidate_id"},
	}
	CandidateElections = Table{
		Name: "candidate_elections",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/CanElections.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "candidate_id", RawName: "CandidateID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "election_type", RawName: "ElectionType", Type: ColumnTypeString},
			{Name: "election_year", RawName: "ElectionYear", Type: ColumnTypeInt},
			{Name: "incumbent_challenger_open", RawName: "IncChallOpen", Type: ColumnTypeString},
			{Name: "won_lost", RawName: "WonLost", Type: ColumnTypeString},
			{Name: "fair_campaign", RawName: "FairCampaign", Type: ColumnTypeBool},
			{Name: "limits_off", RawName: "LimitsOff", Type: ColumnTypeBool},
			{Name: "limits_off_reason", RawName: "LimitsOffReason", Type: ColumnTypeString},
		},
		IndexedColumns: []string{"id", "candidate_id"},
	}
	Committees = Table{
		Name: "committees",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/Committees.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "type_of_committee", RawName: "TypeOfCommittee", Type: ColumnTypeString},
			{Name: "state_committee", RawName: "StateCommittee", Type: ColumnTypeString},
			{Name: "state_id", RawName: "StateID", Type: ColumnTypeString},
			{Name: "local_committee", RawName: "LocalCommittee", Type: ColumnTypeString},
			{Name: "local_id", RawName: "LocalID", Type: ColumnTypeString},
			{Name: "refer_name", RawName: "ReferName", Type: ColumnTypeString},
			{Name: "name", RawName: "Name", Type: ColumnTypeString},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "address3", RawName: "Address3", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "status", RawName: "Status", Type: ColumnTypeString},
			{Name: "status_date", RawName: "StatusDate", Type: ColumnTypeString},
			{Name: "creation_date", RawName: "CreationDate", Type: ColumnTypeString},
			{Name: "creation_amount", RawName: "CreationAmount", Type: ColumnTypeString}, // TODO: Could be a decimal
			{Name: "disp_funds_return", RawName: "DispFundsReturn", Type: ColumnTypeString},
			{Name: "disp_funds_pol_comm", RawName: "DispFundsPolComm", Type: ColumnTypeString},
			{Name: "disp_funds_charity", RawName: "DispFundsCharity", Type: ColumnTypeString},
			{Name: "disp_funds_95", RawName: "DispFunds95", Type: ColumnTypeBool},
			{Name: "disp_funds_descrip", RawName: "DispFundsDescrip", Type: ColumnTypeString},
			{Name: "can_supp_opp", RawName: "CanSuppOpp", Type: ColumnTypeString},
			{Name: "policy_supp_opp", RawName: "PolicySuppOpp", Type: ColumnTypeString},
			{Name: "party_affiliation", RawName: "PartyAffiliation", Type: ColumnTypeString},
			{Name: "purpose", RawName: "Purpose", Type: ColumnTypeString},
		},
		IndexedColumns: []string{"id"},
	}
	CommitteeOfficers = Table{
		Name: "committee_officers",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/CmteOfficerLinks.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "officer_id", RawName: "OfficerID", Type: ColumnTypeInt, NotNullable: true},
		},
		IndexedColumns: []string{"id", "committee_id", "officer_id"},
	}
	D2Totals = Table{
		Name: "d2_totals",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/D2Totals.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "filed_doc_id", RawName: "FiledDocID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "beginning_funds_available", RawName: "BegFundsAvail", Type: ColumnTypeDecimal},
			{Name: "individual_contributions_itemized", RawName: "IndivContribI", Type: ColumnTypeDecimal},
			{Name: "individual_contributions_not_itemized", RawName: "IndivContribNI", Type: ColumnTypeDecimal},
			{Name: "transfers_in_itemized", RawName: "XferInI", Type: ColumnTypeDecimal},
			{Name: "transfers_in_not_itemized", RawName: "XferInNI", Type: ColumnTypeDecimal},
			{Name: "loan_received_itemized", RawName: "LoanRcvI", Type: ColumnTypeDecimal},
			{Name: "loan_received_not_itemized", RawName: "LoanRcvNI", Type: ColumnTypeDecimal},
			{Name: "other_receipts_itemized", RawName: "OtherRctI", Type: ColumnTypeDecimal},
			{Name: "other_receipts_not_itemized", RawName: "OtherRctNI", Type: ColumnTypeDecimal},
			{Name: "total_receipts", RawName: "TotalReceipts", Type: ColumnTypeDecimal},
			{Name: "in_kind_itemized", RawName: "InKindI", Type: ColumnTypeString},
			{Name: "in_kind_not_itemized", RawName: "InKindNI", Type: ColumnTypeDecimal},
			{Name: "total_in_kind", RawName: "TotalInKind", Type: ColumnTypeDecimal},
			{Name: "transfers_out_itemized", RawName: "XferOutI", Type: ColumnTypeDecimal},
			{Name: "transfers_out_not_itemized", RawName: "XferOutNI", Type: ColumnTypeDecimal},
			{Name: "loan_made_itemized", RawName: "LoanMadeI", Type: ColumnTypeString},
			{Name: "loan_made_not_itemized", RawName: "LoanMadeNI", Type: ColumnTypeDecimal},
			{Name: "expenditures_itemized", RawName: "ExpendI", Type: ColumnTypeDecimal},
			{Name: "expenditures_not_itemized", RawName: "ExpendNI", Type: ColumnTypeDecimal},
			{Name: "independent_exp_itemized", RawName: "IndependentExpI", Type: ColumnTypeDecimal},
			{Name: "independent_exp_not_itemized", RawName: "IndependentExpNI", Type: ColumnTypeDecimal},
			{Name: "total_expenditures", RawName: "TotalExpend", Type: ColumnTypeDecimal},
			{Name: "debts_itemized", RawName: "DebtsI", Type: ColumnTypeDecimal},
			{Name: "debts_not_itemized", RawName: "DebtsNI", Type: ColumnTypeDecimal},
			{Name: "total_debts", RawName: "TotalDebts", Type: ColumnTypeDecimal},
			{Name: "total_investments", RawName: "TotalInvest", Type: ColumnTypeDecimal},
			{Name: "end_funds_available", RawName: "EndFundsAvail", Type: ColumnTypeDecimal},
			{Name: "archived", RawName: "Archived", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "committee_id", "filed_doc_id"},
	}
	Expenditures = Table{
		Name: "expenditures",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/Expenditures.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "filed_doc_id", RawName: "FiledDocID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "etrans_id", RawName: "ETransID", Type: ColumnTypeInt},
			{Name: "last_only_name", RawName: "LastOnlyName", Type: ColumnTypeString},
			{Name: "first_name", RawName: "FirstName", Type: ColumnTypeString},
			{Name: "expended_date", RawName: "ExpendedDate", Type: ColumnTypeString},
			{Name: "amount", RawName: "Amount", Type: ColumnTypeDecimal},
			{Name: "aggregate_amount", RawName: "AggregateAmount", Type: ColumnTypeDecimal},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "d2_part", RawName: "D2Part", Type: ColumnTypeString},
			{Name: "purpose", RawName: "Purpose", Type: ColumnTypeString},
			{Name: "candidate_name", RawName: "CandidateName", Type: ColumnTypeString},
			{Name: "office", RawName: "Office", Type: ColumnTypeString},
			{Name: "supporting", RawName: "Supporting", Type: ColumnTypeBool},
			{Name: "opposing", RawName: "Opposing", Type: ColumnTypeBool},
			{Name: "archived", RawName: "Archived", Type: ColumnTypeBool},
			{Name: "country", RawName: "Country", Type: ColumnTypeString},
			{Name: "redaction_requested", RawName: "RedactionRequested", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "committee_id", "filed_doc_id"},
	}
	FiledDocs = Table{
		Name: "filed_docs",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/FiledDocs.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "filed_doc_type", RawName: "FiledDocType", Type: ColumnTypeInt},
			{Name: "doc_name", RawName: "DocName", Type: ColumnTypeString},
			{Name: "amend", RawName: "Amend", Type: ColumnTypeBool},
			{Name: "comment", RawName: "Comment", Type: ColumnTypeString},
			{Name: "pages", RawName: "Pages", Type: ColumnTypeInt},
			{Name: "election_type", RawName: "ElectionType", Type: ColumnTypeString},
			{Name: "election_year", RawName: "ElectionYear", Type: ColumnTypeInt},
			{Name: "report_period_begin_date", RawName: "RptPdBegDate", Type: ColumnTypeString},
			{Name: "report_period_end_date", RawName: "RptPdEndDate", Type: ColumnTypeString},
			{Name: "received_at", RawName: "RcvdAt", Type: ColumnTypeString},
			{Name: "received_date_time", RawName: "RecvdDateTime", Type: ColumnTypeString},
			{Name: "source", RawName: "Source", Type: ColumnTypeString},
			{Name: "provider", RawName: "Provider", Type: ColumnTypeString},
			{Name: "signer_last_name", RawName: "SingerLastOnlyName", Type: ColumnTypeString},
			{Name: "signer_first_name", RawName: "SignerFirstName", Type: ColumnTypeString},
			{Name: "submitter_last_name", RawName: "SbmttrLastOnlyName", Type: ColumnTypeString},
			{Name: "submitter_first_name", RawName: "SbmttrFirstName", Type: ColumnTypeString},
			{Name: "submitter_address1", RawName: "SbmttrAddress1", Type: ColumnTypeString},
			{Name: "submitter_address2", RawName: "SbmttrAddress2", Type: ColumnTypeString},
			{Name: "submitter_city", RawName: "SbmttrCity", Type: ColumnTypeString},
			{Name: "submitter_state", RawName: "SbmttrState", Type: ColumnTypeString},
			{Name: "submitter_zip", RawName: "SbmttrZip", Type: ColumnTypeString},
			{Name: "b9_signer_last_name", RawName: "B9SignerLastOnlyName", Type: ColumnTypeString},
			{Name: "b9_signer_first_name", RawName: "B9SignerFirstName", Type: ColumnTypeString},
			{Name: "archived", RawName: "Archived", Type: ColumnTypeBool},
			{Name: "clarification", RawName: "Clarification", Type: ColumnTypeString},
			{Name: "redaction_requested", RawName: "RedactionRequested", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "committee_id"},
	}
	Investments = Table{
		Name: "investments",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/Investments.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "filed_doc_id", RawName: "FiledDocID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "description", RawName: "Description", Type: ColumnTypeString},
			{Name: "purchase_date", RawName: "PurchaseDate", Type: ColumnTypeString},
			{Name: "purchase_shares", RawName: "PurchaseShares", Type: ColumnTypeDecimal},
			{Name: "purchase_price", RawName: "PurchasePrice", Type: ColumnTypeDecimal},
			{Name: "current_value", RawName: "CurrentValue", Type: ColumnTypeDecimal},
			{Name: "liquid_value", RawName: "LiquidValue", Type: ColumnTypeDecimal},
			{Name: "liquid_date", RawName: "LiquidDate", Type: ColumnTypeString},
			{Name: "last_only_name", RawName: "LastOnlyName", Type: ColumnTypeString},
			{Name: "first_name", RawName: "FirstName", Type: ColumnTypeString},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "archived", RawName: "Archived", Type: ColumnTypeBool},
			{Name: "country", RawName: "Country", Type: ColumnTypeString},
		},
		IndexedColumns: []string{"id", "committee_id", "filed_doc_id"},
	}
	Officers = Table{
		Name: "officers",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/Officers.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "last_name", RawName: "LastName", Type: ColumnTypeString},
			{Name: "first_name", RawName: "FirstName", Type: ColumnTypeString},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "title", RawName: "Title", Type: ColumnTypeString},
			{Name: "phone", RawName: "Phone", Type: ColumnTypeString},
			{Name: "redaction_requested", RawName: "RedactionRequested", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "last_name", "first_name"},
	}
	PreviousOfficers = Table{
		Name: "previous_officers",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/PrevOfficers.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "last_name", RawName: "LastName", Type: ColumnTypeString},
			{Name: "first_name", RawName: "FirstName", Type: ColumnTypeString},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "title", RawName: "Title", Type: ColumnTypeString},
			{Name: "resign_date", RawName: "ResignDate", Type: ColumnTypeString},
			{Name: "redaction_requested", RawName: "RedactionRequested", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "committee_id", "last_name", "first_name"},
	}
	Receipts = Table{
		Name: "receipts",
		URL:  "https://elections.il.gov/campaigndisclosuredatafiles/Receipts.txt",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "filed_doc_id", RawName: "FiledDocID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "etrans_id", RawName: "ETransID", Type: ColumnTypeInt},
			{Name: "last_only_name", RawName: "LastOnlyName", Type: ColumnTypeString},
			{Name: "first_name", RawName: "FirstName", Type: ColumnTypeString},
			{Name: "receive_date", RawName: "RcvDate", Type: ColumnTypeString},
			{Name: "amount", RawName: "Amount", Type: ColumnTypeDecimal},
			{Name: "aggregate_amount", RawName: "AggregateAmount", Type: ColumnTypeDecimal},
			{Name: "loan_amount", RawName: "LoanAmount", Type: ColumnTypeDecimal},
			{Name: "occupation", RawName: "Occupation", Type: ColumnTypeString},
			{Name: "employer", RawName: "Employer", Type: ColumnTypeString},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "d2_part", RawName: "D2Part", Type: ColumnTypeString},
			{Name: "description", RawName: "Description", Type: ColumnTypeString},
			{Name: "vendor_last_only_name", RawName: "VendorLastOnlyName", Type: ColumnTypeString},
			{Name: "vendor_first_name", RawName: "VendorFirstName", Type: ColumnTypeString},
			{Name: "vendor_address1", RawName: "VendorAddress1", Type: ColumnTypeString},
			{Name: "vendor_address2", RawName: "VendorAddress2", Type: ColumnTypeString},
			{Name: "vendor_city", RawName: "VendorCity", Type: ColumnTypeString},
			{Name: "vendor_state", RawName: "VendorState", Type: ColumnTypeString},
			{Name: "vendor_zip", RawName: "VendorZip", Type: ColumnTypeString},
			{Name: "archived", RawName: "Archived", Type: ColumnTypeBool},
			{Name: "country", RawName: "Country", Type: ColumnTypeString},
			{Name: "redaction_requested", RawName: "RedactionRequested", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "committee_id", "filed_doc_id"},
	}
)
