package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeCreated = "CREATED"
const CodeInvalidForm = "INVALID_FORM"
const CodeRecordNotFound = "RECORD_NOT_FOUND"
const CodeInternalServerError = "INTERNAL_SERVER_ERROR"

/* --------------------------------- Message -------------------------------- */

const MsgCreated = "Created"
const MsgInvalidForm = "invalid form"
const MsgRecordNotFound = "record not found"
const MsgInternalServerError = "internal server error"

/* ---------------------------------- Error --------------------------------- */

var ErrRecordNotFound = errors.New(MsgRecordNotFound)
