package stream

type UuserEntity struct {
	Id         int    `json:"id"`
	TenantId   int    `json:"tenant_id"`
	CustomerId int    `json:"customer_id"`
	BizId      string `json:"biz_id"`
	UserName   string `json:"username"`
	Mobile     string `json:"mobile"`
	Password   string `json:"password"`
	NickName   string `json:"nickname"`
	IconUrl    string `json:"icon_url"`
	Email      string `json:"email"`
	RoleType   int    `json:"role_type"`
	Type       int    `json:"type"`
	Status     int    `json:"status"`
	OnOff      string `json:"on_off"`
	InsertTime int64  `json:"inserttime"`
	UpdateTime int64  `json:"updatetime"`
	IsActive   bool   `json:"isactive"`
}

type UuserResponse struct {
	Result        int           `json:"result"`
	ResultMessage string        `json:"resultMessage"`
	Content       []UuserEntity `json:"content"`
}



type LimitInitLogEntity struct {
		Id                      int     `json:"id"`
		UserId                  int     `json:"user_id"`
		BizType                 int     `json:"biz_type"`
		FlowNo                  string  `json:"flow_no"`
		TotalCreditAmount       float64 `json:"total_credit_amount"`
		PdlCreditAmount         float64 `json:"pdl_credit_amount"`
		InstallmentCreditAmount float64 `json:"installment_credit_amount"`
		LaunchTime              int64   `json:"launch_time"`
		CompleteTime            int64   `json:"complete_time"`
		CompleteFlag            bool    `json:"complete_flag"`
		MqMsg                   string  `json:"mq_msg"`
		InsertTime              int64   `json:"inserttime"`
		UpdateTime              int64   `json:"updatetime"`
		IsActive                bool    `json:"isactive"`
		PaylaterCreditAmount    float64 `json:"paylater_credit_amount"`
		CashloanCreditAmount    float64 `json:"cashloan_credit_amount"`
}

type LimitInitLogResponse struct {
		Result        int                  `json:"result"`
		ResultMessage string               `json:"resultMessage"`
		Content       []LimitInitLogEntity `json:"content"`
}
