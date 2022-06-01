package enums

/**
 * Known summary options for patron information.
 *
 */
const (
	SUM_HOLD_ITEMS = iota
	SUM_OVERDUE_ITEMS
	SUM_CHARGED_ITEMS
	SUM_FINE_ITEMS
	SUM_RECALL_ITEMS
	SUM_UNAVAILABLE_HOLDS
)
