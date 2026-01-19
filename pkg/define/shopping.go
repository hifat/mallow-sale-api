package define

import "errors"

// Code

const CodeInvalidShoppingStatus = "INVALID_SHOPPING_STATUS"

// Message

const MsgInvalidShoppingStatus = "invalid shopping status"

// Error

var ErrInvalidShoppingStatus = errors.New(MsgInvalidShoppingStatus)
