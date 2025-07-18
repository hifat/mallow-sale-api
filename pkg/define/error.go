package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

var CodeInvalidForm = "INVALID_FORM"
var CodeRecordNotFound = "RECORD_NOT_FOUND"
var CodeInternalServerError = "INTERNAL_SERVER_ERROR"

/* --------------------------------- Message -------------------------------- */

var MsgInvalidForm = "invalid form"
var MsgRecordNotFound = "record not found"
var MsgInternalServerError = "internal server error"

/* ---------------------------------- Error --------------------------------- */

var ErrRecordNotFound = errors.New("record not found")
