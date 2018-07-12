package transaction

import "bmvs.io/ynab/api"

// Transaction represents a full transaction for a budget
type Transaction struct {
	ID   string   `json:"id"`
	Date api.Date `json:"date"`
	// Amount Transaction amount in milliunits format
	Amount    int64          `json:"amount"`
	Cleared   ClearingStatus `json:"cleared"`
	Approved  bool           `json:"approved"`
	AccountID string         `json:"account_id"`
	// Deleted Deleted transactions will only be included in delta requests
	Deleted         bool              `json:"deleted"`
	AccountName     string            `json:"account_name"`
	SubTransactions []*SubTransaction `json:"subtransactions"`

	Memo              *string    `json:"memo"`
	FlagColor         *FlagColor `json:"flag_color"`
	PayeeID           *string    `json:"payee_id"`
	CategoryID        *string    `json:"category_id"`
	TransferAccountID *string    `json:"transfer_account_id"`
	// ImportID If the Transaction was imported, this field is a unique (by account) import
	// identifier. If this transaction was imported through File Based Import or
	// Direct Import and not through the API, the import_id will have the format:
	// 'YNAB:[milliunit_amount]:[iso_date]:[occurrence]'. For example, a transaction
	// dated 2015-12-30 in the amount of -$294.23 USD would have an import_id of
	// 'YNAB:-294230:2015-12-30:1’. If a second transaction on the same account
	// was imported and had the same date and same amount, its import_id would
	// be 'YNAB:-294230:2015-12-30:2’.
	ImportID     *string `json:"import_id"`
	PayeeName    *string `json:"payee_name"`
	CategoryName *string `json:"category_name"`
}

// Summary represents the summary of a transaction for a budget
type Summary struct {
	ID   string   `json:"id"`
	Date api.Date `json:"date"`
	// Amount Transaction amount in milliunits format
	Amount    int64          `json:"amount"`
	Cleared   ClearingStatus `json:"cleared"`
	Approved  bool           `json:"approved"`
	AccountID string         `json:"account_id"`
	// Deleted Deleted transactions will only be included in delta requests
	Deleted bool `json:"deleted"`

	Memo              *string    `json:"memo"`
	FlagColor         *FlagColor `json:"flag_color"`
	PayeeID           *string    `json:"payee_id"`
	CategoryID        *string    `json:"category_id"`
	TransferAccountID *string    `json:"transfer_account_id"`

	// ImportID If the Transaction was imported, this field is a unique (by account) import
	// identifier. If this transaction was imported through File Based Import or
	// Direct Import and not through the API, the import_id will have the format:
	// 'YNAB:[milliunit_amount]:[iso_date]:[occurrence]'. For example, a transaction
	// dated 2015-12-30 in the amount of -$294.23 USD would have an import_id of
	// 'YNAB:-294230:2015-12-30:1’. If a second transaction on the same account
	// was imported and had the same date and same amount, its import_id would
	// be 'YNAB:-294230:2015-12-30:2’.
	ImportID *string `json:"import_id"`
}

// SubTransaction represents a sub-transaction for a transaction
type SubTransaction struct {
	ID            string `json:"id"`
	TransactionID string `json:"transaction_id"`
	// Amount sub-transaction amount in milliunits format
	Amount int64 `json:"amount"`
	// Deleted Deleted sub-transactions will only be included in delta requests.
	Deleted bool `json:"deleted"`

	Memo       *string `json:"memo"`
	PayeeID    *string `json:"payee_id"`
	CategoryID *string `json:"category_id"`
	// TransferAccountID If a transfer, the account_id which the
	// sub-transaction transfers to
	TransferAccountID *string `json:"transfer_account_id"`
}

// Scheduled represents a scheduled transaction for a budget
type Scheduled struct {
	ID        string             `json:"id"`
	DateFirst api.Date           `json:"date_first"`
	DateNext  api.Date           `json:"date_next"`
	Frequency ScheduledFrequency `json:"frequency"`
	// Amount The scheduled transaction amount in milliunits format
	Amount    int64  `json:"amount"`
	AccountID string `json:"account_id"`
	// Deleted Deleted scheduled transactions will only be included in delta requests.
	Deleted         bool                       `json:"deleted"`
	AccountName     string                     `json:"account_name"`
	SubTransactions []*ScheduledSubTransaction `json:"subtransactions"`

	Memo       *string    `json:"memo"`
	FlagColor  *FlagColor `json:"flag_color"`
	PayeeID    *string    `json:"payee_id"`
	CategoryID *string    `json:"category_id"`
	// TransferAccountID If a transfer, the account_id which the scheduled
	// transaction transfers to
	TransferAccountID *string `json:"transfer_account_id"`
	PayeeName         *string `json:"payee_name"`
	CategoryName      *string `json:"category_name"`
}

// ScheduledSummary represents the summary of a scheduled transaction for a budget
type ScheduledSummary struct {
	ID        string             `json:"id"`
	DateFirst api.Date           `json:"date_first"`
	DateNext  api.Date           `json:"date_next"`
	Frequency ScheduledFrequency `json:"frequency"`
	// Amount The scheduled transaction amount in milliunits format
	Amount    int64  `json:"amount"`
	AccountID string `json:"account_id"`
	// Deleted Deleted scheduled transactions will only be included in delta requests.
	Deleted bool `json:"deleted"`

	Memo       *string    `json:"memo"`
	FlagColor  *FlagColor `json:"flag_color"`
	PayeeID    *string    `json:"payee_id"`
	CategoryID *string    `json:"category_id"`
	// TransferAccountID If a transfer, the account_id which the scheduled
	// transaction transfers to
	TransferAccountID *string `json:"transfer_account_id"`
}

// ScheduledSubTransaction represents a scheduled sub-transaction for
// a scheduled transaction
type ScheduledSubTransaction struct {
	ID                     string `json:"id"`
	ScheduledTransactionID string `json:"scheduled_transaction_id"`
	// Amount The scheduled sub-transaction amount in milliunits format
	Amount int64 `json:"amount"`
	// Deleted Deleted scheduled sub-transactions will only be included in delta requests
	Deleted bool `json:"deleted"`

	Memo       *string `json:"memo"`
	PayeeID    *string `json:"payee_id"`
	CategoryID *string `json:"category_id"`
	// TransferAccountID If a transfer, the account_id which the scheduled subtransaction transfers to
	TransferAccountID *string `json:"transfer_account_id"`
}