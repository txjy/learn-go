package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	//user 模块
	ErrorExistUser                 = 30001
	ErrorFailEncryption            = 30002
	ErrorExistUserNotFound         = 30003
	ErrorNotCompare                = 30004
	ErrorAuthorToken               = 30005
	ErrorAuthCheckTokenTimeout     = 30006
	ErrorUpLoadFail                = 30007
	ErrorSendEmail                 = 30008
	ErrorAuthorCheckTokenFail      = 30009
	ErrorAuthInsufficientAuthority = 30010
	//

	//product模块
	ErrorProductImgUpload = 40001
)
