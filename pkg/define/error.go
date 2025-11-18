package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeCreated = "CREATED"
const CodeUpdated = "UPDATED"
const CodeDeleted = "DELETED"
const CodeInvalidForm = "INVALID_FORM"
const CodeRecordNotFound = "RECORD_NOT_FOUND"
const CodeInternalServerError = "INTERNAL_SERVER_ERROR"
const CodeUnauthorized = "UNAUTHORIZED"
const CodeInvalidCredentials = "INVALID_CREDENTIALS"
const CodeInvalidToken = "INVALID_TOKEN"
const CodeTokenExpired = "TOKEN_EXPIRED"

/* --------------------------------- Message -------------------------------- */

const MsgCreated = "Created"
const MsgUpdated = "Updated"
const MsgDeleted = "Deleted"
const MsgInvalidForm = "invalid form"
const MsgRecordNotFound = "record not found"
const MsgInternalServerError = "internal server error"
const MsgUnauthorized = "unauthorized"
const MsgInvalidCredentials = "invalid username or password"
const MsgInvalidToken = "invalid token"
const MsgTokenExpired = "token expired"

/* ---------------------------------- Error --------------------------------- */

var ErrRecordNotFound = errors.New(MsgRecordNotFound)
