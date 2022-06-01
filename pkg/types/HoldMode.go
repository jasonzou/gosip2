package enums

/**
 * Hold operations supported by the SC.
 *
 */
//public enum HoldMode {
const (
	HOLD_ADD    = iota /** Add patron the the hold queue for the item. */
	HOLD_DELETE        /** Delete patron from the hold queue for the item. */
	HOLD_CHANGE        /** * Change the hold to match the {@code Hold} message parameters.  */
)
