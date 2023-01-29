package constant

type ErrorCode struct {
	Code string
	Msg  string
}

func NewErrorCode(code string, msg string) ErrorCode {
	return ErrorCode{
		Code: code,
		Msg:  msg,
	}
}

var (
	OK                       = NewErrorCode("000000", "OK")
	Err                      = NewErrorCode("000001", "Fail")
	ParamErr                 = NewErrorCode("000002", "Parameters are empty or miss match")
	RequestReadErr           = NewErrorCode("000003", "network error")
	JsonStringToInterfaceErr = NewErrorCode("000004", "please read our API document")
	FileReadError            = NewErrorCode("000005", "read file error")
	TimeFormatError          = NewErrorCode("000006", "time format error")
	DBInsertError            = NewErrorCode("000006", "Save file error")

	PasswordMismatch    = NewErrorCode("010000", "Please type the correct username or password")
	TokenGenerateFailed = NewErrorCode("010001", "Please type the correct username or password")
	StoreFileFailed     = NewErrorCode("010002", "Store file to local file system failed")
	UsernameNotFound    = NewErrorCode("010003", "Please type the correct username or password")
)
