package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeInvalidUsageUnit = "INVALID_USAGE_UNIT"
const CodeInvalidPurchaseUnit = "INVALID_PURCHASE_UNIT"

/* --------------------------------- Message -------------------------------- */

const MsgInvalidUsageUnit = "invalid usage unit"
const MsgInvalidPurchaseUnit = "invalid purchase unit"

/* ---------------------------------- Error --------------------------------- */

var ErrInvalidUsageUnit = errors.New(MsgInvalidUsageUnit)
var ErrInvalidPurchaseUnit = errors.New(MsgInvalidUsageUnit)
