package anticaptcha

// GenerateCoupons
// generate coupons for specified value and count
func (c *Client) GenerateCoupons(count int, amount float32, purchaseLink string) (*Response_GenerateCoupons, *Error) {
	e := NewError()
	req := &Reqest{}
	if count == 0 {
		e.setString("Count required")
		return nil, e
	}
	req.Count = count
	if amount == 0 {
		e.setString("Amount required")
		return nil, e
	}
	req.Amount = amount
	if purchaseLink == "" {
		e.setString("PurchaseLink required")
		return nil, e
	}
	req.PurchaseLink = purchaseLink
	res := &Response_GenerateCoupons{}
	err := c.request(URL_generateCoupons, req, res)
	return res, err
}

// GetResellerData
// retrieve coupons list and eligible balance
func (c *Client) GetResellerData(minCreateDate int) (*Response_GetResellerData, *Error) {
	req := &Reqest{}
	if minCreateDate != 0 {
		req.MinCreateDate = minCreateDate
	}
	res := &Response_GetResellerData{}
	err := c.request(URL_generateCoupons, req, res)
	return res, err
}
