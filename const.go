package anticaptcha

const (
	URL                            = "https://api.anti-captcha.com"
	EP_createTask                  = "/createTask"
	EP_getTaskResult               = "/getTaskResult"
	EP_getBalance                  = "/getBalance"
	EP_getQueueStats               = "/getQueueStats"
	EP_reportIncorrectImageCaptcha = "/reportIncorrectImageCaptcha"
	EP_getSpendingStats            = "/getSpendingStats"
	EP_getAppStats                 = "/getAppStats"
	EP_generateCoupons             = "/generateCoupons"

	//Captcha Task Types

	// solve usual image captcha
	Type_ImageToText = "ImageToTextTask"
	//Google Recaptcha puzzle solving
	Type_NoCaptcha = "NoCaptchaTask"
	//Google Recaptcha puzzle solving without proxies
	Type_NoCaptchaProxyless = "NoCaptchaTaskProxyless"
	//rotating captcha funcaptcha.com
	Type_FunCaptcha = "FunCaptchaTask"
	// funcaptcha without proxy
	Type_FunCaptchaProxyless = "FunCaptchaTaskProxyless"
	//select objects on image with an overlay grid
	Type_SquareNetText = "SquareNetTextTask"
	//  captcha from geetest.com
	Type_GeeTest = "GeeTestTask"
	//captcha from geetest.com without proxy
	Type_GeeTestProxyless = "GeeTestTaskProxyless"
	//image captcha with custom form
	Type_CustomCaptcha = "CustomCaptchaTask"

	Queue_ImageToTextEn       = 1
	Queue_ImageToTextRu       = 2
	Queue_Recaptcha           = 5
	Queue_RecaptchaProxyless  = 6
	Queue_Funcaptcha          = 7
	Queue_FuncaptchaProxyless = 10

	TaskStatus_Ready      = "ready"
	TaskStatus_Processing = "processing"
)
