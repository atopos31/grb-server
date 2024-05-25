package errcode

var (
	Success           = NewError(200, "Ok")
	ErrInternalServer = NewError(500, "Internal server error")
	ErrInvalidParam   = NewError(400, "Invalid params")
	ErrAccessDenied   = NewError(401, "Access denied")
)

var (
	ErrDataIsExits = NewError(1400, "Data is exists") // 数据已存在
	ErrNotFound    = NewError(1401, "Data not found")
)
