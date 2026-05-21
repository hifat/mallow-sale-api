package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeInvalidServiceCode = "INVALID_SERVICE_CODE"
const CodeFileTooLarge = "FILE_TOO_LARGE"

/* --------------------------------- Message -------------------------------- */

const MsgInvalidServiceCode = "invalid service code"
const MsgFileTooLarge = "file size exceeds 2MB limit"

/* ---------------------------------- Error --------------------------------- */

var ErrInvalidServiceCode = errors.New(MsgInvalidServiceCode)
var ErrFileTooLarge = errors.New(MsgFileTooLarge)
