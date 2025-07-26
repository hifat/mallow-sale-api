package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeInvalidForm = "INVALID_FORM"
const CodeRecordNotFound = "RECORD_NOT_FOUND"
const CodeInternalServerError = "INTERNAL_SERVER_ERROR"
const CodeInvalidSupplierID = "INVALID_SUPPLIER_ID"

/* --------------------------------- Message -------------------------------- */

const MsgInvalidForm = "invalid form"
const MsgRecordNotFound = "record not found"
const MsgInternalServerError = "internal server error"
const MsgInvalidSupplierID = "invalid supplier id"

/* ---------------------------------- Error --------------------------------- */

var ErrRecordNotFound = errors.New("record not found")
