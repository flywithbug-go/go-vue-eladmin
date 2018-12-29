package model

type AppVersion struct {
	AppId         string `json:"app_id"` //所属AppId
	Version       string `json:"version"`
	ParentVersion string `json:"parent_version"`
	Platform      string `json:"platform"`      //(iOS,Android,H5,Server)["iOS","Android","H5","Server"]
	Status        int    `json:"status"`        //状态    0(准备中) 1(开发中) 2(灰度) 3(已发布)
	ApprovalTime  int    `json:"approval_time"` //立项时间
	LockTime      int    `json:"lock_time"`     //锁版时间
	GrayTime      int    `json:"gray_time"`     //灰度时间
	CreateTime    int    `json:"create_time"`   //添加时间
}
