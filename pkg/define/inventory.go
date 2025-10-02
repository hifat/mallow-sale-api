package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeInvalidInventoryID = "INVALID_INVENTORY_ID"
const CodeDuplicatedInventoryName = "INVENTORY_NAME_ALREADY_EXISTS"

/* --------------------------------- Message -------------------------------- */

const MsgInvalidInventoryID = "invalid inventory id"
const MsgDuplicatedInventoryName = "inventory name already exists"

/* ---------------------------------- Error --------------------------------- */

var ErrInvalidInventory = errors.New(MsgInvalidInventoryID)
var ErrDuplicatedInventoryName = errors.New(MsgDuplicatedInventoryName)
