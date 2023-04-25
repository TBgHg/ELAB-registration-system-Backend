package model

type UserSubmitReq struct {
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
	OpenID       string
	Email        string
}

type UserSubmitResp struct {
	*CommonResp
}

type UserGetResp struct {
	*CommonResp
	Name         string `json:"name"`
	StudentID    string `json:"student_id"`
	Gender       int32  `json:"gender"`
	Avatar       string `json:"avatar"`
	IsELABer     int32  `json:"is_elaber"`
	Class        string `json:"class"`
	Position     string `json:"position"`
	Mobile       string `json:"mobile"`
	Group        string `json:"group"`
	Introduction string `json:"introduction"`
	Awards       string `json:"awards"`
	Reason       string `json:"reason"`
}

type UploadAvatarResp struct {
	*CommonResp
	Avatar string `json:"avatar"`
}

type UserUpdateReq struct {
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
	OpenID       string
}

type UserUpdateResp struct {
	*CommonResp
}
