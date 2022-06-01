package enums

/**
 * Indicates the type of the {@code Hold} message.
 *
 */
const (
	//public enum HoldType {
	HT_OTHER               = iota /** Another hold mechanism. */
	HT_ANY_COPY_TITLE             /** Any copy or a title. */
	HT_SPECIFIC_COPY_TITLE        /** A specific copy of a title. */
	HT_ANY_COPY_LOCATION          /** Any copy at a single branch or sublocation. */
)
