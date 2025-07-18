package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeInvalidInventoryID = "INVALID_INVENTORY_ID"

/* --------------------------------- Message -------------------------------- */

const MsgInvalidInventoryID = "invalid inventory id"

/* ---------------------------------- Error --------------------------------- */

var ErrInvalidInventory = errors.New(MsgInvalidInventoryID)
