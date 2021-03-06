package enums

/**
 * The known statuses for a patron.
 *
 */
//public enum PatronStatus {
const (
	ST_CHARGE_PRIVILEGES_DENIED = iota
	ST_RENEWAL_PRIVILEGES_DENIED
	ST_RECALL_PRIVILEGES_DENIED
	ST_HOLD_PRIVILEGES_DENIED
	ST_CARD_REPORTED_LOST
	ST_TOO_MANY_ITEMS_CHARGED
	ST_TOO_MANY_ITEMS_OVERDUE
	ST_TOO_MANY_RENEWALS
	ST_TOO_MANY_CLAIMS_OF_ITEMS_RETURNED
	ST_TOO_MANY_ITEMS_LOST
	ST_EXCESSIVE_OUTSTANDING_FINES
	ST_EXCESSIVE_OUTSTANDING_FEES
	ST_RECALL_OVERDUE
	ST_TOO_MANY_ITEMS_BILLED
)
