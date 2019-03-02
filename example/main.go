package main

import (
	"fmt"
	"os"

	"Packages/anticaptcha"
)

var (
	ClientKey = "..."
)

func main() {

	client := anticaptcha.NewClient(ClientKey)
	client.SetDebug(true)

	//ReCaptcha(client)

	ImageToText(client)

	//Stat(client)

	//Reseller(client)

}

func ReCaptcha(client *anticaptcha.Client) {
	url := "https://...."
	key := "..."
	task := anticaptcha.NewTask_NoCaptchaProxyless(url, key)

	r, ac_err := client.CreateTask(task)
	if ac_err.Error != nil {
		fmt.Println(ac_err.ErrMsg())
		os.Exit(1)
	}

	fmt.Println("TaskID:", r.TaskID)

	result, ac_err := client.WaitResult(r.TaskID, 600)
	if ac_err.Error != nil {
		fmt.Println(ac_err.ErrMsg())
		os.Exit(1)
	}
	fmt.Println("Result:", result.ResultRecaptcha())
}

func ImageToText(client *anticaptcha.Client) {

	task, err := anticaptcha.NewTask_ImageFromFile("./test.png")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	r, ac_err := client.CreateTask(task)
	if ac_err.Error != nil {
		fmt.Println(ac_err.ErrMsg())
		os.Exit(1)
	}

	fmt.Println("TaskID:", r.TaskID)
	result, ac_err := client.WaitResult(r.TaskID, 600)
	if ac_err.Error != nil {
		fmt.Println(ac_err.ErrMsg())
		os.Exit(1)
	}
	fmt.Println("Result:")
	fmt.Println(result.ResultImage())
}

func Stat(client *anticaptcha.Client) {
	b, ac_err := client.GetBalance()
	if ac_err.Error != nil {
		fmt.Println(ac_err.ErrMsg())
		fmt.Println(ac_err.ErrMsgRu())
		os.Exit(1)
	}
	fmt.Println(b.Balance)

	q, ac_err := client.GetQueueStats(anticaptcha.Queue_RecaptchaProxyless)
	if ac_err.Error != nil {
		fmt.Println(ac_err.ErrMsg())
		os.Exit(1)
	}
	fmt.Printf("%#v\n", q)

	s, ac_err := client.GetSpendingStats(0, "", 0, "")
	if ac_err.Error != nil {
		fmt.Println(ac_err.ErrMsg())
		os.Exit(1)
	}
	fmt.Printf("%#v\n", s)

}

func Reseller(client *anticaptcha.Client) {
	// reseller
	c, ac_err := client.GenerateCoupons(1, 5, "http://")
	if ac_err.Error != nil {
		fmt.Println(ac_err.ErrMsg())
		os.Exit(1)
	}
	fmt.Printf("%#v\n", c)

	r, ac_err := client.GetResellerData(0)
	if ac_err.Error != nil {
		fmt.Println(ac_err.ErrMsg())
		os.Exit(1)
	}
	fmt.Printf("%#v\n", r)
}
