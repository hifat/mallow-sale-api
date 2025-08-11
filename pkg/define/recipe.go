package define

import "errors"

/* ---------------------------------- Code ---------------------------------- */

const CodeOrderNoMustBeUnique = "ORDER_NO_MUST_BE_UNIQUE"
const CodeInvalidRecipeType = "INVALID_RECIPE_TYPE"

/* --------------------------------- Message -------------------------------- */

const MsgOrderNoMustBeUnique = "order no must be unique"
const MsgInvalidRecipeType = "invalid recipe type"

/* ---------------------------------- Error --------------------------------- */

var ErrOrderNoMustBeUnique = errors.New(MsgOrderNoMustBeUnique)
var ErrInvalidRecipeType = errors.New(MsgInvalidRecipeType)
