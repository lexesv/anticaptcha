package anticaptcha

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	ClientKey string
	URL       string
	Debug     bool
}

type Reqest struct {
	ClientKey string `json:"clientKey"`

	// tasks
	Task         interface{} `json:"task,omitempty"`
	SoftId       int         `json:"softId,omitempty"`
	LanguagePool string      `json:"languagePool,omitempty"`
	CallbackUrl  string      `json:"callbackUrl,omitempty"`
	TaskID       int         `json:"taskId,omitempty"`
	IsExtended   bool        `json:"isExtended,omitempty"`

	// queue stats
	QueueId int `json:"queueId,omitempty"`

	// spending stats
	Queue string `json:"queue,omitempty"`
	Date  int    `json:"date,omitempty"`
	IP    string `json:"ip,omitempty"`

	// app stat
	Mode string `json:"mode,omitempty"`

	// reseller
	Count         int     `json:"count,omitempty"`
	Amount        float32 `json:"amount,omitempty"`
	PurchaseLink  string  `json:"purchaseLink,omitempty"`
	MinCreateDate int     `json:"minCreateDate,omitempty"`
}

type Response_CreateTask struct {
	ErrorID int `json:"errorId"`
	TaskID  int `json:"taskId,omitempty"`
}

type Response_GetTaskResult struct {
	ErrorID  int    `json:"errorId"`
	Status   string `json:"status,omitempty"`
	Solution struct {
		// ImageToText
		Text string `json:"text,omitempty"`
		URL  string `json:"url,omitempty"`
		// NoCaptcha
		GRecaptchaResponse    string `json:"gRecaptchaResponse,omitempty"`
		GRecaptchaResponseMD5 string `json:"gRecaptchaResponseMD5,omitempty"`
		// FunCaptcha
		Token string `json:"token,omitempty"`
		// SquareNetText
		CellNumbers []int `json:"cellNumbers"`
		// CustomCaptcha
		TaskID  int               `json:"taskId,omitempty"`
		Status  string            `json:"status,omitempty"`
		Answers map[string]string `json:"answers,omitempty"`
	} `json:"solution,omitempty"`
	Cost       string `json:"cost,omitempty"`
	IP         string `json:"ip,omitempty"`
	CreateTime int    `json:"createTime,omitempty"`
	EndTime    int    `json:"endTime,omitempty"`
	SolveCount int    `json:"solveCount,omitempty"`
}

func (r Response_GetTaskResult) ResultImage() (text, url string) {
	return r.Solution.Text, r.Solution.URL
}
func (r Response_GetTaskResult) ResultRecaptcha() string {
	return r.Solution.GRecaptchaResponse
}
func (r Response_GetTaskResult) ResultFunCaptcha() string {
	return r.Solution.Token
}
func (r Response_GetTaskResult) ResultSquareNetText() []int {
	return r.Solution.CellNumbers
}

type ResponseReportIncorrectImageCaptcha struct {
	ErrorID int    `json:"errorId"`
	Status  string `json:"status,omitempty"`
}

type Response_GetBalance struct {
	ErrorID int     `json:"errorId"`
	Balance float32 `json:"balance,omitempty"`
}

type Response_GetQueueStats struct {
	ErrorID int     `json:"errorId"`
	Waiting int     `json:"waiting,omitempty"`
	Load    float64 `json:"load,omitempty"`
	Bid     float64 `json:"bid,omitempty"`
	Speed   float64 `json:"speed,omitempty"`
	Total   int     `json:"total,omitempty"`
}

type Response_GetSpendingStats struct {
	ErrorID int `json:"errorId"`
	Data    []struct {
		DateFrom int     `json:"dateFrom"`
		DateTill int     `json:"dateTill"`
		Volume   int     `json:"volume"`
		Money    float64 `json:"money"`
	} `json:"data"`
}

type Response_GetAppStats struct {
	ErrorID   int `json:"errorId"`
	ChartData []struct {
		Name string `json:"name"`
		Data []struct {
			Date       string `json:"date,omitempty"`
			Shortdate  string `json:"shortdate,omitempty"`
			Y          int    `json:"y,omitempty"`
			Beginstamp int    `json:"beginstamp,omitempty"`
			Endstamp   int    `json:"endstamp,omitempty"`
			Stamp      int    `json:"stamp,omitempty"`
		} `json:"data"`
		Itemname    string `json:"itemname"`
		ErrorID     int    `json:"errorId"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Count       int    `json:"count,omitempty"`
	} `json:"chartData"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

type Response_GenerateCoupons struct {
	ErrorID int      `json:"errorId"`
	Coupons []string `json:"coupons"`
}

type Response_GetResellerData struct {
	ErrorID         int     `json:"errorId"`
	EligibleBalance float64 `json:"eligibleBalance"`
	Coupons         []struct {
		ID     int    `json:"id"`
		Amount int    `json:"amount"`
		Code   string `json:"code"`
		Link   string `json:"link"`
		Status string `json:"status"`
	} `json:"coupons"`
}

type Response struct {
	ErrorID int `json:"errorId"`
}

// NewClient
// Create new anti-captcha client
func NewClient(ClientKey string) *Client {
	return &Client{ClientKey: ClientKey, URL: URL}
}

func (c *Client) ChengeURL(url string) {
	s := []rune(url)
	if string(s[len(s)-1]) == "/" {
		url = string(s[0 : len(s)-1])
	}
	c.URL = url
}

func (c *Client) SetDebug(v bool) {
	c.Debug = v
}

func (c *Client) request(url string, req_data *Reqest, resp_data interface{}) (error *Error) {
	error = &Error{}
	req_data.ClientKey = c.ClientKey

	payloadBytes, err := json.Marshal(req_data)
	if err != nil {
		error.set(err)
		return error
	}

	if c.Debug {
		fmt.Println("Request URL:", url)
		fmt.Println("Request JSON:", string(payloadBytes))
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		error.set(err)
		return error
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	//resp, err := http.DefaultClient.Do(req)
	if err != nil {
		error.set(err)
		return error
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		error.set(err)
		return error
	}

	if c.Debug {
		fmt.Println("Response JSON:", string(b))
	}

	err = json.Unmarshal(b, resp_data)
	if err != nil {
		error.set(err)
		return error
	}

	switch resp_data.(type) {
	case *Response_CreateTask:
		error.ErrorID = resp_data.(*Response_CreateTask).ErrorID
	case *Response_GetTaskResult:
		error.ErrorID = resp_data.(*Response_GetTaskResult).ErrorID
	case *ResponseReportIncorrectImageCaptcha:
		error.ErrorID = resp_data.(*ResponseReportIncorrectImageCaptcha).ErrorID
	case *Response_GetBalance:
		error.ErrorID = resp_data.(*Response_GetBalance).ErrorID
	case *Response_GetQueueStats:
		error.ErrorID = resp_data.(*Response_GetQueueStats).ErrorID
	case *Response_GetSpendingStats:
		error.ErrorID = resp_data.(*Response_GetSpendingStats).ErrorID
	case *Response_GetAppStats:
		error.ErrorID = resp_data.(*Response_GetAppStats).ErrorID
	case *Response_GenerateCoupons:
		error.ErrorID = resp_data.(*Response_GenerateCoupons).ErrorID
	case *Response_GetResellerData:
		error.ErrorID = resp_data.(*Response_GetResellerData).ErrorID
	default:
		error.setString("Response structure not found")
	}
	if error.ErrorID != 0 {
		error.Error = error.ErrMsg()
	}

	return error
}

// WaitResult
func (c *Client) WaitResult(taskID int, timeout int64) (result *Response_GetTaskResult, error *Error) {
	error = &Error{}
	if timeout == 0 {
		timeout = timeout + 120
	}
	endTime := time.Now().Unix() + timeout
	for {
		if time.Now().Unix() >= endTime {
			error.setString("Timeout")
			return result, error
		}
		time.Sleep(time.Second * 10)
		result, error = c.GetTaskResult(taskID)
		if error.Error != nil {
			return result, error
		}
		if result.Status == TaskStatus_Ready {
			return result, error
		} else {
			if c.Debug {
				fmt.Println("Not ready. Wait 10s")
			}
		}
	}
}
