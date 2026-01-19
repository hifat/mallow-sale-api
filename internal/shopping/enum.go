package shoppingModule

type EnumCodeShoppingStatusType string

const (
	EnumCodeShoppingPending    EnumCodeShoppingStatusType = "PENDING"
	EnumCodeShoppingInProgress EnumCodeShoppingStatusType = "IN_PROGRESS"
	EnumCodeShoppingSuccess    EnumCodeShoppingStatusType = "SUCCESS"
	EnumCodeShoppingCancel     EnumCodeShoppingStatusType = "CANCEL"
)

func (c EnumCodeShoppingStatusType) GetShoppingStatusName() string {
	switch c {
	case EnumCodeShoppingPending:
		return "pending"
	case EnumCodeShoppingInProgress:
		return "in_progress"
	case EnumCodeShoppingSuccess:
		return "success"
	case EnumCodeShoppingCancel:
		return "cancel"
	default:
		return ""
	}
}

type EnumCodeInventoryStatusType string

const (
	EnumCodeInventoryPending   EnumCodeInventoryStatusType = "PENDING"
	EnumCodeInventorySuccess   EnumCodeInventoryStatusType = "SUCCESS"
	EnumCodeInventoryMoveOrder EnumCodeInventoryStatusType = "MOVE_ORDER"
	EnumCodeInventoryCancel    EnumCodeInventoryStatusType = "CANCEL"
)
