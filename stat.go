package anticaptcha

// GetBalance
// retrieve account balance
func (c *Client) GetBalance() (*Response_GetBalance, *Error) {
	req := &Reqest{}
	res := &Response_GetBalance{}
	err := c.request(c.URL+EP_getBalance, req, res)
	return res, err
}

// GetQueueStats
// obtain queue load statistics
func (c *Client) GetQueueStats(queueId int) (*Response_GetQueueStats, *Error) {
	req := &Reqest{}
	req.QueueId = queueId
	res := &Response_GetQueueStats{}
	err := c.request(c.URL+EP_getQueueStats, req, res)
	return res, err
}

// GetSpendingStats
// retrieve account spending stats
func (c *Client) GetSpendingStats(date int, queue string, softId int, ip string) (*Response_GetSpendingStats, *Error) {
	req := &Reqest{}
	if date != 0 {
		req.Date = date
	}
	if queue != "" {
		req.Queue = queue
	}
	if softId != 0 {
		req.SoftId = softId
	}
	if ip != "" {
		req.IP = ip
	}
	res := &Response_GetSpendingStats{}
	err := c.request(c.URL+EP_getSpendingStats, req, res)
	return res, err
}

// getAppStats
// retrieve application statistics
func (c *Client) GetAppStats(softId int, mode string) (*Response_GetAppStats, *Error) {
	e := NewError()
	req := &Reqest{}
	if softId == 0 {
		e.setString("softId required")
		return nil, e
	}
	req.SoftId = softId
	if mode != "" {
		req.Mode = mode
	}
	res := &Response_GetAppStats{}
	err := c.request(c.URL+EP_getAppStats, req, res)
	return res, err
}
