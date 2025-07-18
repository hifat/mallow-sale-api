package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeInvalidUsageUnit = "INVALID_USAGE_UNIT"

/* --------------------------------- Message -------------------------------- */

const MsgInvalidUsageUnit = "invalid usage unit"

/* ---------------------------------- Error --------------------------------- */

var ErrInvalidUsageUnit = errors.New(MsgInvalidUsageUnit)
