package enums

/**
 * Defined circulation statuses.
 *
 */
const (
	//enum CirculationStatus {
	CIRC_OTHER = iota
	CIRC_ON_ORDER
	CIRC_AVAILABLE
	CIRC_CHARGED
	CIRC_CHARGED_NOT_TO_BE_RECALLED_UNTIL_EARLIEST_RECALL_DATE
	CIRC_IN_PROCESS
	CIRC_RECALLED
	CIRC_WAITING_ON_HOLD_SHELF
	CIRC_WAITING_TO_BE_RESHELVED
	CIRC_IN_TRANSIT_BETWEEN_LIBRARY_LOCATIONS
	CIRC_CLAIMED_RETURNED
	CIRC_LOST
	CIRC_MISSING
)