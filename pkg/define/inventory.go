package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeInvalidInventoryID = "INVALID_INVENTORY_ID"
const CodeInventoryNameAlreadyExists = "INVENTORY_NAME_ALREADY_EXISTS"

/* --------------------------------- Message -------------------------------- */

const MsgInvalidInventoryID = "invalid inventory id"
const MsgInventoryNameAlreadyExists = "inventory name already exists"

/* ---------------------------------- Error --------------------------------- */

var ErrInvalidInventory = errors.New(MsgInvalidInventoryID)
var ErrInventoryNameAlreadyExists = errors.New(MsgInventoryNameAlreadyExists)
