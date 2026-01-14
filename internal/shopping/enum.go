package shoppingModule

type EnumCodeShoppingStatusType string

const (
	EnumCodeShoppingPending    EnumCodeShoppingStatusType = "PENDING"
	EnumCodeShoppingInProgress EnumCodeShoppingStatusType = "IN_PROGRESS"
	EnumCodeShoppingSuccess    EnumCodeShoppingStatusType = "SUCCESS"
	EnumCodeShoppingCancel     EnumCodeShoppingStatusType = "CANCEL"
)

type EnumCodeInventoryStatusType string

const (
	EnumCodeInventoryPending   EnumCodeInventoryStatusType = "PENDING"
	EnumCodeInventorySuccess   EnumCodeInventoryStatusType = "SUCCESS"
	EnumCodeInventoryMoveOrder EnumCodeInventoryStatusType = "MOVE_ORDER"
	EnumCodeInventoryCancel    EnumCodeInventoryStatusType = "CANCEL"
)
