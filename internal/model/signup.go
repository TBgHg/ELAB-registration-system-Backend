package model

type SignupSubmitReq struct {
	Name         string `json:"name"`
	StudentID    string `json:"student_id"`
	Gender       int32  `json:"gender"`
	Class        string `json:"class"`
	Position     string `json:"position"`
	Mobile       string `json:"mobile"`
	Group        string `json:"group"`
	Introduction string `json:"introduction"`
	Awards       string `json:"awards"`
	Reason       string `json:"reason"`
}

type SignupSubmitResp struct {
	*CommonResp
}

type SignupGetReq struct {
}

type SignupGetResp struct {
	*CommonResp
	Name         string `json:"name"`
	StudentID    string `json:"student_id"`
	Gender       int32  `json:"gender"`
	Class        string `json:"class"`
	Position     string `json:"position"`
	Mobile       string `json:"mobile"`
	Group        string `json:"group"`
	Introduction string `json:"introduction"`
	Awards       string `json:"awards"`
	Reason       string `json:"reason"`
}

type SignupUpdateReq struct {
	Name         string `json:"name"`
	StudentID    string `json:"student_id"`
	Gender       int32  `json:"gender"`
	Class        string `json:"class"`
	Position     string `json:"position"`
	Mobile       string `json:"mobile"`
	Group        string `json:"group"`
	Introduction string `json:"introduction"`
	Awards       string `json:"awards"`
	Reason       string `json:"reason"`
}

type SignupUpdateResp struct {
	*CommonResp
}

type SignupDeleteReq struct {
}

type SignupDeleteResp struct {
	*CommonResp
}
