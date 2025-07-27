package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeInvalidSupplierID = "INVALID_SUPPLIER_ID"

/* --------------------------------- Message -------------------------------- */

const MsgInvalidSupplierID = "invalid supplier id"

/* ---------------------------------- Error --------------------------------- */

var ErrInvalidSupplierID = errors.New(MsgInvalidSupplierID)
