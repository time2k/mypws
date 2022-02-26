package mytypedef

type URPCRiskCheckRequest string
type URPCRiskCheckResponse int
type URPCValidateRequest string
type URPCValidateResponse int
type URPCGetUserinfoRequest string
type URPCGetUserinfoResponse struct {
	Users []UserInfo `json:"Users"`
}
type URPCGetUserinfoWithGeoRequest struct {
	Users    string `json:"users"`
	MyUserid int    `json:"myuserid"`
}
type URPCGetUserinfoWithGeoResponse struct {
	Users []UserInfoWithGeo `json:"users"`
}
type URPCGetUserinfoDetailRequest struct {
	Users    string `json:"users"`
	MyUserid int    `json:"myuserid"`
}
type URPCGetUserinfoDetailResponse struct {
	Users []UserInfoDetail `json:"users"`
}

type URPCGetUserGeoRequest int
type URPCGetUserGeoResponse UserGeoInfo
type URPCGetUserStatusRequest int
type URPCGetUserStatusResponse UserStatus

type URPCIncrUserLevelRequest struct {
	UserID int     `json:"userid"` //用户ID
	Coins  int     `json:"coins"`  //金币
	Point  float64 `json:"point"`  //积分
}

type URPCGetUserConcernRequest int
type URPCGetUserConcernResponse MultiUserInfo

//UserInfo db结构体
type UserInfo struct {
	ID               int     `json:"userid"`
	Username         string  `json:"username"`
	RongYunToken     string  `json:"rongyun_token"`
	Mobile           string  `json:"mobile"`
	Gender           string  `json:"gender"`
	Birthday         string  `json:"birthday"`
	AvatarImgurl     string  `json:"avatar_imgurl"`
	Height           string  `json:"height"`
	Weight           string  `json:"weight"`
	Profession       string  `json:"profession"`
	Education        string  `json:"education"`
	AnnualIncome     string  `json:"annual_income"`
	ResidencePlace   string  `json:"residence_place"`
	Signature        string  `json:"signature"`
	WealthLevel      int     `json:"wealth_level"`
	WealthLevelPoint int     `json:"wealth_levelpoint"`
	WealthLevelPic   string  `json:"wealth_levelpic"`
	CharmLevel       int     `json:"charm_level"`
	CharmLevelPoint  float64 `json:"charm_levelpoint"`
	CharmLevelPic    string  `json:"charm_levelpic"`
	VIPRight         int     `json:"vipright"`
	StatusFlag       int     `json:"status_flag"` //状态 0-正常 1-ban 2-注销'
	RegTime          string  `json:"reg_time"`
	RealPersonAuth   int     `json:"realperson_auth"` //实人认证
	RealNameAuth     int     `json:"realname_auth"`   //实名认证
	OnlineStatus     int     `json:"online_status"`   //在线状态 0-未知 1-在线 2-不在线
}

//UserStatus 用户状态
type UserStatus struct {
	UserID    int   `json:"userid"`
	Status    int   `json:"status"` //0-未知 1-在线 2-不在线
	TimeStamp int64 `json:"timestamp"`
}

//MultiUserInfo 返回
type MultiUserInfo struct {
	Users []UserInfo `json:"users"`
}

//用户声音
type UserVoice struct {
	Voice         string `json:"voice_url"`
	VoiceDuration int    `json:"voice_duration"`
}

//用户相册信息
type UserPhotoAlbum struct {
	PhotoAlbum []string `json:"photo_album"`
}

//用户相册单条
type UserPhoto struct {
	Photo    string `json:"photo"`
	Position int    `json:"position"`
}

//用户标签
type UserTags map[string][]string

//用户地理信息
type UserGeo struct {
	Location       string  `json:"location"`
	Distance       float64 `json:"distance"`
	DistanceString string  `json:"distance_string"`
}

//用户完整信息
type UserInfoDetail struct {
	UserInfo  `json:"userinfo"`
	UserVoice `json:"voice"`
	UserPhotoAlbum
	UserTags `json:"tags"`
	UserGeo  `json:"geoinfo"`
}

//UserGeoInfo 用户地理属性
type UserGeoInfo struct {
	UserID   int     `json:"userid"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Province string  `json:"province"`
	City     string  `json:"city"`
}

//用户信息+地理位置
type UserInfoWithGeo struct {
	UserInfo `json:"userinfo"`
	UserGeo  `json:"geoinfo"`
}
