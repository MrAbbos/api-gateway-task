package handlers

type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type HasAccessModel struct {
	ClientPlatformID string `json:"client_platform_id"`
	ClientTypeID     string `json:"client_type_id"`
	UserID           string `json:"user_id"`
	CompanyID        string `json:"company_id"`
	BranchID         string `json:"branch_id"`
	ID               string `json:"id"`
	RoleID           string `json:"role_id"`
	Data             string `json:"data"`
}
