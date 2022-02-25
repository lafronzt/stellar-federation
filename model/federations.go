package model

// Users is the details of the users that stellar requires to be able to send payments to the recipient.
type Users struct {
	ID       string `yaml:"id" json:"account_id"`                           // ID Public Key of the user, which Stellar refers to as the Account ID.
	Name     string `yaml:"name" json:"stellar_address"`                    // Name is the username of the stellar federation user. I.e. tyler in tyler*lafronz.com
	MemoType string `yaml:"memo_type,omitempty" json:"memo_type,omitempty"` // MemoType is the type of memo that is used for the transaction.
	Memo     string `yaml:"memo,omitempty" json:"memo,omitempty"`           // Memo is an optional memo that is added the transaction.
}

// Federations is a collection of users and their associated data points.
type Federations []Users
