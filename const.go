package anticaptcha

const (
	URL_createTask                  = "https://api.anti-captcha.com/createTask"
	URL_getTaskResult               = "https://api.anti-captcha.com/getTaskResult"
	URL_getBalance                  = "https://api.anti-captcha.com/getBalance"
	URL_getQueueStats               = "https://api.anti-captcha.com/getQueueStats"
	URL_reportIncorrectImageCaptcha = "https://api.anti-captcha.com/reportIncorrectImageCaptcha"
	URL_getSpendingStats            = "https://api.anti-captcha.com/getSpendingStats"
	URL_getAppStats                 = "https://api.anti-captcha.com/getAppStats"
	URL_generateCoupons             = "https://api.anti-captcha.com/generateCoupons"

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
