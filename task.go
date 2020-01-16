package anticaptcha

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

type Task_ImageToText struct {
	Type      string `json:"type"`
	Body      string `json:"body"`
	Phrase    bool   `json:"phrase,omitempty"`
	Case      bool   `json:"case,omitempty"`
	Numeric   bool   `json:"numeric,omitempty"`
	Math      int    `json:"math,omitempty"`
	MinLength int    `json:"minLength,omitempty"`
	MaxLength int    `json:"maxLength,omitempty"`
}

type Task_NoCaptcha struct {
	Type          string `json:"type"`                    //Yes    NoCaptchaTask
	WebsiteURL    string `json:"websiteURL"`              //Yes    Address of target web page
	WebsiteKey    string `json:"websiteKey"`              //Yes    Recaptcha website key
	WebsiteSToken string `json:"websiteSToken,omitempty"` //No
	ProxyType     string `json:"proxyType"`               //Yes
	ProxyAddress  string `json:"proxyAddress"`            //Yes
	ProxyPort     int    `json:"proxyPort"`               //Yes
	ProxyLogin    string `json:"proxyLogin,omitempty"`    //No
	ProxyPassword string `json:"proxyPassword,omitempty"` //No
	UserAgent     string `json:"userAgent"`               //Yes
	Cookies       string `json:"cookies,omitempty"`       //No
	IsInvisible   bool   `json:"isInvisible,omitempty"`   //Specify if Recaptcha is invisible.
}

type Task_NoCaptchaProxyless struct {
	Type          string `json:"type"`                    //Yes    NoCaptchaTask
	WebsiteURL    string `json:"websiteURL"`              //Yes    Address of target web page
	WebsiteKey    string `json:"websiteKey"`              //Yes    Recaptcha website key
	WebsiteSToken string `json:"websiteSToken,omitempty"` //No
	IsInvisible   bool   `json:"isInvisible,omitempty"`   //Specify if Recaptcha is invisible.
}

type Task_FunCaptcha struct {
	Type                     string `json:"type"`
	WebsiteURL               string `json:"websiteURL"`
	FuncaptchaAPIJSSubdomain string `json:"funcaptchaApiJSSubdomain"`
	WebsitePublicKey         string `json:"websitePublicKey"`
	ProxyType                string `json:"proxyType"`               //Yes
	ProxyAddress             string `json:"proxyAddress"`            //Yes
	ProxyPort                int    `json:"proxyPort"`               //Yes
	ProxyLogin               string `json:"proxyLogin,omitempty"`    //No
	ProxyPassword            string `json:"proxyPassword,omitempty"` //No
	UserAgent                string `json:"userAgent"`               //Yes
	Cookies                  string `json:"cookies,omitempty"`       //No
}

type Task_FunCaptchaTaskProxyless struct {
	Type                     string `json:"type"`
	WebsiteURL               string `json:"websiteURL"`
	FuncaptchaAPIJSSubdomain string `json:"funcaptchaApiJSSubdomain,omitempty"`
	WebsitePublicKey         string `json:"websitePublicKey"`
}

type Task_GeeTest struct {
	Type                      string `json:"type"`
	WebsiteURL                string `json:"websiteURL"`
	Gt                        string `json:"gt"`
	Challenge                 string `json:"challenge"`
	GeetestApiServerSubdomain string `json:"geetestApiServerSubdomain,omitempty"`
	ProxyType                 string `json:"proxyType"`
	ProxyAddress              string `json:"proxyAddress"`
	ProxyPort                 int    `json:"proxyPort"`
	ProxyLogin                string `json:"proxyLogin,omitempty"`
	ProxyPassword             string `json:"proxyPassword,omitempty"`
	UserAgent                 string `json:"userAgent"`
	Cookies                   string `json:"cookies,omitempty"`
}
type Task_GeeTestProxyless struct {
	Type                      string `json:"type"`
	WebsiteURL                string `json:"websiteURL"`
	Gt                        string `json:"gt"`
	Challenge                 string `json:"challenge"`
	GeetestApiServerSubdomain string `json:"geetestApiServerSubdomain,omitempty"`
}

type Task_CustomCaptcha struct {
	Type       string `json:"type"`
	ImageURL   string `json:"imageUrl"`
	Assignment string `json:"assignment"`
	Forms      []struct {
		Label        string      `json:"label"`
		LabelHint    string      `json:"labelHint"`
		ContentType  string      `json:"contentType,omitempty"`
		Content      interface{} `json:"contentType,omitempty"`
		Name         string      `json:"name"`
		InputType    string      `json:"inputType"`
		InputOptions struct {
			Width       string `json:"width"`
			Rows        string `json:"rows,omitempty"`
			PlaceHolder string `json:"placeHolder"`
		} `json:"inputOptions"`
	} `json:"forms"`
}

type Task_SquareNetText struct {
	Type         string `json:"type"`
	Body         string `json:"body"`
	ObjectName   string `json:"objectName"`
	RowsCount    int    `json:"rowsCount"`
	ColumnsCount int    `json:"columnsCount"`
}

// NewTask_ImageFromFile
func NewTask_ImageFromFile(file string) (task Task_ImageToText, err error) {
	task = Task_ImageToText{}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return task, err
	}
	encoded := base64.StdEncoding.EncodeToString(b)
	task.Type = Type_ImageToText
	task.Body = encoded
	return task, err
}

// NewTask_ImageFromURL
func NewTask_ImageFromURL(url string) (task Task_ImageToText, err error) {
	task = Task_ImageToText{}
	resp, err := http.Get(url)
	if err != nil {
		return task, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return task, err
	}
	encoded := base64.StdEncoding.EncodeToString(b)
	task.Type = Type_ImageToText
	task.Body = encoded
	return task, err
}

// NewTask_ImageFromBytes
func NewTaskImageFromBytes(b []byte) Task_ImageToText {
	encoded := base64.StdEncoding.EncodeToString(b)
	return Task_ImageToText{
		Type: Type_ImageToText,
		Body: encoded,
	}
}

// NewTask_NoCaptcha
func NewTask_NoCaptcha(websiteURL, websiteKey, proxyType, proxyAddress string, proxyPort int, userAgent string) Task_NoCaptcha {
	return Task_NoCaptcha{
		Type:         Type_NoCaptcha,
		WebsiteURL:   websiteURL,
		WebsiteKey:   websiteKey,
		ProxyType:    proxyType,
		ProxyAddress: proxyAddress,
		ProxyPort:    proxyPort,
		UserAgent:    userAgent,
	}
}

// NewNoCaptchaProxyless
func NewTask_NoCaptchaProxyless(websiteURL, websiteKey string) Task_NoCaptchaProxyless {
	return Task_NoCaptchaProxyless{
		Type:       Type_NoCaptchaProxyless,
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
	}
}

// NewTask_FunCaptcha
func NewTask_FunCaptcha(websiteURL, WebsitePublicKey, proxyType, proxyAddress string, proxyPort int, userAgent string) Task_FunCaptcha {
	return Task_FunCaptcha{
		Type:             Type_FunCaptcha,
		WebsiteURL:       websiteURL,
		WebsitePublicKey: WebsitePublicKey,
		ProxyType:        proxyType,
		ProxyAddress:     proxyAddress,
		ProxyPort:        proxyPort,
		UserAgent:        userAgent,
	}
}

// NewTask_FunCaptchaProxyless
func NewTask_FunCaptchaProxyless(websiteURL, WebsitePublicKey string) Task_FunCaptchaTaskProxyless {
	return Task_FunCaptchaTaskProxyless{
		Type:             Type_FunCaptchaProxyless,
		WebsiteURL:       websiteURL,
		WebsitePublicKey: WebsitePublicKey,
	}
}

// NewTask_SquareNetTextFromFile
func NewTask_SquareNetTextFromFile(file, objectName string, rowsCount, columnsCount int) (task Task_SquareNetText, err error) {
	task = Task_SquareNetText{}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return task, err
	}
	encoded := base64.StdEncoding.EncodeToString(b)
	task.Type = Type_SquareNetText
	task.Body = encoded
	task.ObjectName = objectName
	task.RowsCount = rowsCount
	task.ColumnsCount = columnsCount
	return task, err
}

// NewTask_SquareNetTextFromURL
func NewTask_SquareNetTextFromURL(url, objectName string, rowsCount, columnsCount int) (task Task_SquareNetText, err error) {
	task = Task_SquareNetText{}
	resp, err := http.Get(url)
	if err != nil {
		return task, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return task, err
	}
	encoded := base64.StdEncoding.EncodeToString(b)
	task.Type = Type_SquareNetText
	task.Body = encoded
	task.ObjectName = objectName
	task.RowsCount = rowsCount
	task.ColumnsCount = columnsCount
	return task, err
}

// NewTask_SquareNetTextFromBytes
func NewTask_SquareNetTextFromBytes(b []byte, objectName string, rowsCount, columnsCount int) Task_SquareNetText {
	encoded := base64.StdEncoding.EncodeToString(b)
	return Task_SquareNetText{
		Type:         Type_SquareNetText,
		Body:         encoded,
		ObjectName:   objectName,
		RowsCount:    rowsCount,
		ColumnsCount: columnsCount,
	}
}

// NewTask_GeeTest
func NewTaskGeeTest(websiteURL, gt, challenge, proxyType, proxyAddress string, proxyPort int, userAgent string) Task_GeeTest {
	return Task_GeeTest{
		Type:         Type_GeeTest,
		WebsiteURL:   websiteURL,
		Gt:           gt,
		Challenge:    challenge,
		ProxyType:    proxyType,
		ProxyAddress: proxyAddress,
		ProxyPort:    proxyPort,
		UserAgent:    userAgent,
	}
}

// NewTask_GeeTestProxyless
func NewTaskGeeTestProxyless(websiteURL, gt, challenge string) Task_GeeTestProxyless {
	return Task_GeeTestProxyless{
		Type:       Type_GeeTestProxyless,
		WebsiteURL: websiteURL,
		Gt:         gt,
		Challenge:  challenge,
	}
}

// NewTask_CustomCaptcha
func NewTask_CustomCaptcha(imageUrl string) Task_CustomCaptcha {
	return Task_CustomCaptcha{
		Type:     Type_CustomCaptcha,
		ImageURL: imageUrl,
	}
}

// CreateTask
// This method creates a task for solving selected captcha type.
func (c *Client) CreateTask(task interface{}) (*Response_CreateTask, *Error) {
	req := &Reqest{}
	req.Task = task
	res := &Response_CreateTask{}
	err := c.request(c.URL+EP_createTask, req, res)
	return res, err
}

// GetTaskResult
// request task result
func (c *Client) GetTaskResult(taskId int) (*Response_GetTaskResult, *Error) {
	req := &Reqest{}
	req.TaskID = taskId
	res := &Response_GetTaskResult{}
	err := c.request(c.URL+EP_getTaskResult, req, res)
	return res, err
}

// ReportIncorrectImageCaptcha
//Complaints are accepted only for image captchas. Your complaint will be checked by 5 workers, 3 of them must confirm it. Only then you get full refund. If you have less than 20% confirmation ratio, your reports will be ignored.
func (c *Client) ReportIncorrectImageCaptcha(taskId int) (*ResponseReportIncorrectImageCaptcha, *Error) {
	req := &Reqest{}
	req.TaskID = taskId
	res := &ResponseReportIncorrectImageCaptcha{}
	err := c.request(c.URL+EP_reportIncorrectImageCaptcha, req, res)
	return res, err
}
