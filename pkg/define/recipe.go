package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeOrderNoMustBeUnique = "ORDER_NO_MUST_BE_UNIQUE"

/* --------------------------------- Message -------------------------------- */

const MsgOrderNoMustBeUnique = "order no must be unique"

/* ---------------------------------- Error --------------------------------- */

var ErrOrderNoMustBeUnique = errors.New(MsgOrderNoMustBeUnique)
