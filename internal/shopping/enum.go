package shoppingModule

type EnumCodeInventoryStatusType string

const EnumCodePending EnumCodeInventoryStatusType = "PENDING"        // รอดำเนินการ
const EnumCodeInProgress EnumCodeInventoryStatusType = "IN_PROGRESS" // กำลังดำเนินการ
const EnumCodeSuccess EnumCodeInventoryStatusType = "SUCCESS"        // เสร็จสิน
const EnumCodeCancel EnumCodeInventoryStatusType = "CANCEL"          // ยกเลิก
