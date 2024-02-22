package structs

type DougaBaseInfo struct {
	Title           string `json:"title"`
	BananaCountShow string `json:"bananaCountShow"`
	IsLike          bool   `json:"isLike"`
	CommentCount    int    `json:"commentCount"`
	BananaCount     int    `json:"bananaCount"`
	Mkey            string `json:"mkey"`
	IsFavorite      bool   `json:"isFavorite"`
	IsThrowBanana   bool   `json:"isThrowBanana"`
	StowCount       int    `json:"stowCount"`
}

type DougaInfo struct {
	DougaBaseInfo
	CurrentVideoID              int              `json:"currentVideoId"`
	Pctr                        float64          `json:"pctr"`
	CommentCountRealValue       int              `json:"commentCountRealValue"`
	GroupID                     string           `json:"groupId"`
	StowCountShow               string           `json:"stowCountShow"`
	GiftPeachCountShow          string           `json:"giftPeachCountShow"`
	Channel                     Channel          `json:"channel"`
	LikeCount                   int              `json:"likeCount"`
	HasHotComment               bool             `json:"hasHotComment"`
	ShareCountShow              string           `json:"shareCountShow"`
	IsDislike                   bool             `json:"isDislike"`
	ShareCount                  int              `json:"shareCount"`
	PicShareURL                 string           `json:"picShareUrl"`
	DanmakuCount                int              `json:"danmakuCount"`
	CurrentVideoInfo            CurrentVideoInfo `json:"currentVideoInfo"`
	ViewCountShow               string           `json:"viewCountShow"`
	DougaID                     string           `json:"dougaId"`
	CommentCountTenThousandShow string           `json:"commentCountTenThousandShow"`
	CoverImgInfo                CoverImgInfo     `json:"coverImgInfo"`
	CoverURL                    string           `json:"coverUrl"`
	DisableEdit                 bool             `json:"disableEdit"`
	CreateTimeMillis            uint64           `json:"createTimeMillis"`
	SuperUbb                    bool             `json:"superUbb"`
	ShareURL                    string           `json:"shareUrl"`
	Status                      int              `json:"status"`
	RecoReason                  RecoReason       `json:"recoReason"`
	Description                 string           `json:"description"`
	BelongToSpecifyArubamu      bool             `json:"belongToSpecifyArubamu"`
	Result                      int              `json:"result"`
	OriginalDeclare             int              `json:"originalDeclare"`
	VideoList                   []VideoList      `json:"videoList"`
	ViewCount                   int              `json:"viewCount"`
	LikeCountShow               string           `json:"likeCountShow"`
	CoverCdnUrls                []CoverCdnUrls   `json:"coverCdnUrls"`
	GiftPeachCount              int              `json:"giftPeachCount"`
	IsRewardSupportted          bool             `json:"isRewardSupportted"`
	CommentCountShow            string           `json:"commentCountShow"`
	DurationMillis              int              `json:"durationMillis"`
	HostName                    string           `json:"host-name"`
	TagList                     []TagList        `json:"tagList"`
	DanmakuCountShow            string           `json:"danmakuCountShow"`
	CreateTime                  string           `json:"createTime"`
	User                        User             `json:"user"`
	Priority                    int              `json:"priority"`
}

type BangumiInfo struct {
	DougaBaseInfo
	CoverImgVInfo          CoverImgVInfo       `json:"coverImgVInfo"`
	CoverImgHInfo          CoverImgHInfo       `json:"coverImgHInfo"`
	FirstPlayDate          int64               `json:"firstPlayDate"`
	UpdateWeekDay          int                 `json:"updateWeekDay"`
	UpdateDayTime          int                 `json:"updateDayTime"`
	LastUpdateItemName     string              `json:"lastUpdateItemName"`
	WebUpdateInfoShow      string              `json:"webUpdateInfoShow"`
	UpdateInfoShow         string              `json:"updateInfoShow"`
	CommentCountShow       string              `json:"commentCountShow"`
	PlayCountShow          string              `json:"playCountShow"`
	ShareURL               string              `json:"shareUrl"`
	StowCountShow          string              `json:"stowCountShow"`
	CurrentVideoInfo       CurrentVideoInfo    `json:"currentVideoInfo"`
	PlayCount              int                 `json:"playCount"`
	AcfunOnly              bool                `json:"acfunOnly"`
	BangumiPaymentType     BangumiPaymentType  `json:"bangumiPaymentType"`
	UpdateTime             string              `json:"updateTime"`
	Image                  string              `json:"image"`
	ItemID                 int                 `json:"itemId"`
	BelongResource         BelongResource      `json:"belongResource"`
	PaidForUser            bool                `json:"paidForUser"`
	BangumiTitle           string              `json:"bangumiTitle"`
	BelongType             int                 `json:"belongType"`
	NeedPay                bool                `json:"needPay"`
	VideoSizeType          int                 `json:"videoSizeType"`
	ImgInfo                ImgInfo             `json:"imgInfo"`
	Introduction           string              `json:"introduction"`
	EpisodeName            string              `json:"episodeName"`
	PaymentType            PaymentType         `json:"paymentType"`
	VideoID                int                 `json:"videoId"`
	Priority               int                 `json:"priority"`
	Sidelights             []Sidelights        `json:"sidelights"`
	RelatedBangumis        []RelatedBangumis   `json:"relatedBangumis"`
	RecommendBangumis      []RecommendBangumis `json:"recommendBangumis"`
	HotTags                []HotTags           `json:"hotTags"`
	BangumiLikeCount       int                 `json:"bangumiLikeCount"`
	BangumiBananaCount     int                 `json:"bangumiBananaCount"`
	BangumiIsLike          bool                `json:"bangumiIsLike"`
	BangumiIntro           string              `json:"bangumiIntro"`
	BangumiCoverImageV     string              `json:"bangumiCoverImageV"`
	BangumiCoverImageH     string              `json:"bangumiCoverImageH"`
	CommentParted          bool                `json:"commentParted"`
	BangumiStyleList       []BangumiStyleList  `json:"bangumiStyleList"`
	UpdateDayTimeStr       string              `json:"updateDayTimeStr"`
	AllowDownload          bool                `json:"allowDownload"`
	AllowComment           bool                `json:"allowComment"`
	Online                 int                 `json:"online"`
	BangumiID              int                 `json:"bangumiId"`
	UpdateStatus           int                 `json:"updateStatus"`
	OnlineTime             int64               `json:"onlineTime"`
	ItemCount              int                 `json:"itemCount"`
	ExtendsStatus          string              `json:"extendsStatus"`
	LatestItem             string              `json:"latestItem"`
	UpdateFrequency        string              `json:"updateFrequency"`
	ShowTitle              string              `json:"showTitle"`
	BangumiLikeCountShow   string              `json:"bangumiLikeCountShow"`
	BangumiBananaCountShow string              `json:"bangumiBananaCountShow"`
}

type Channel struct {
	ParentName string `json:"parentName"`
	ParentID   int    `json:"parentId"`
	Name       string `json:"name"`
	ID         int    `json:"id"`
}
type Representation struct {
	ID           int      `json:"id"`
	URL          string   `json:"url"`
	BackupURL    []string `json:"backupUrl"`
	M3U8Slice    string   `json:"m3u8Slice"`
	MaxBitrate   int      `json:"maxBitrate"`
	AvgBitrate   int      `json:"avgBitrate"`
	Codecs       string   `json:"codecs"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	FrameRate    float64  `json:"frameRate"`
	Quality      float64  `json:"quality"`
	QualityType  string   `json:"qualityType"`
	QualityLabel string   `json:"qualityLabel"`
	Comment      string   `json:"comment"`
	HdrType      int      `json:"hdrType"`
}
type AdaptationSet struct {
	ID             int              `json:"id"`
	Duration       int              `json:"duration"`
	Representation []Representation `json:"representation"`
}

type TranscodeInfos struct {
	QualityType string `json:"qualityType"`
	SizeInBytes int    `json:"sizeInBytes"`
	Hdr         bool   `json:"hdr"`
}
type CurrentVideoInfo struct {
	Priority             int              `json:"priority"`
	Title                string           `json:"title"`
	DanmakuCountShow     string           `json:"danmakuCountShow"`
	UserPlayedSeconds    int              `json:"userPlayedSeconds"`
	SizeType             int              `json:"sizeType"`
	DanmakuGuidePosition int              `json:"danmakuGuidePosition"`
	SourceStatus         int              `json:"sourceStatus"`
	DurationMillis       int              `json:"durationMillis"`
	UploadTime           int64            `json:"uploadTime"`
	VisibleType          int              `json:"visibleType"`
	DanmakuCount         int              `json:"danmakuCount"`
	FileName             string           `json:"fileName"`
	IsKsManifest         bool             `json:"isKsManifest"`
	TranscodeInfos       []TranscodeInfos `json:"transcodeInfos"`
	ID                   string           `json:"id"`
}
type CdnUrls struct {
	URL                            string `json:"url"`
	FreeTrafficProductAbbreviation string `json:"freeTrafficProductAbbreviation"`
	FreeTrafficCdn                 bool   `json:"freeTrafficCdn"`
}
type ThumbnailImage struct {
	CdnUrls  []CdnUrls `json:"cdnUrls"`
	Animated int       `json:"animated"`
}
type ExpandedImage struct {
	CdnUrls []CdnUrls `json:"cdnUrls"`
}
type CoverImgInfo struct {
	Width                int            `json:"width"`
	Height               int            `json:"height"`
	Size                 int            `json:"size"`
	Type                 int            `json:"type"`
	Animated             bool           `json:"animated"`
	ThumbnailImage       ThumbnailImage `json:"thumbnailImage"`
	ExpandedImage        ExpandedImage  `json:"expandedImage"`
	ThumbnailImageCdnURL string         `json:"thumbnailImageCdnUrl"`
}
type RecoReason struct {
	Desc       string `json:"desc"`
	Href       string `json:"href"`
	Tag        string `json:"tag"`
	LayoutType int    `json:"layoutType"`
	DescType   int    `json:"descType"`
	Type       int    `json:"type"`
}
type VideoList struct {
	Priority             int    `json:"priority"`
	Title                string `json:"title"`
	DanmakuCountShow     string `json:"danmakuCountShow"`
	UserPlayedSeconds    int    `json:"userPlayedSeconds"`
	SizeType             int    `json:"sizeType"`
	DanmakuGuidePosition int    `json:"danmakuGuidePosition"`
	SourceStatus         int    `json:"sourceStatus"`
	DurationMillis       int    `json:"durationMillis"`
	UploadTime           int64  `json:"uploadTime"`
	VisibleType          int    `json:"visibleType"`
	DanmakuCount         int    `json:"danmakuCount"`
	FileName             string `json:"fileName"`
	ID                   string `json:"id"`
}
type CoverCdnUrls struct {
	URL            string `json:"url"`
	FreeTrafficCdn bool   `json:"freeTrafficCdn"`
}
type TagList struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
type HeadCdnUrls struct {
	URL            string `json:"url"`
	FreeTrafficCdn bool   `json:"freeTrafficCdn"`
}
type SocialMedal struct {
}
type UserHeadImgInfo struct {
	Width                int            `json:"width"`
	Height               int            `json:"height"`
	Size                 int            `json:"size"`
	Type                 int            `json:"type"`
	Animated             bool           `json:"animated"`
	ThumbnailImage       ThumbnailImage `json:"thumbnailImage"`
	ExpandedImage        ExpandedImage  `json:"expandedImage"`
	ThumbnailImageCdnURL string         `json:"thumbnailImageCdnUrl"`
}
type User struct {
	Action               int             `json:"action"`
	Href                 string          `json:"href"`
	HeadCdnUrls          []HeadCdnUrls   `json:"headCdnUrls"`
	IsJoinUpCollege      bool            `json:"isJoinUpCollege"`
	FollowingCountValue  int             `json:"followingCountValue"`
	ContributeCountValue int             `json:"contributeCountValue"`
	FanCount             string          `json:"fanCount"`
	SocialMedal          SocialMedal     `json:"socialMedal"`
	AvatarImage          string          `json:"avatarImage"`
	UserHeadImgInfo      UserHeadImgInfo `json:"userHeadImgInfo"`
	IsFollowed           bool            `json:"isFollowed"`
	IsFollowing          bool            `json:"isFollowing"`
	AvatarFrame          int             `json:"avatarFrame"`
	ContributeCount      string          `json:"contributeCount"`
	AvatarFramePcImg     string          `json:"avatarFramePcImg"`
	AvatarFrameMobileImg string          `json:"avatarFrameMobileImg"`
	SexTrend             int             `json:"sexTrend"`
	VerifiedTypes        []int           `json:"verifiedTypes"`
	HeadURL              string          `json:"headUrl"`
	FanCountValue        int             `json:"fanCountValue"`
	FollowingStatus      int             `json:"followingStatus"`
	NameColor            int             `json:"nameColor"`
	FollowingCount       string          `json:"followingCount"`
	Gender               int             `json:"gender"`
	Name                 string          `json:"name"`
	Signature            string          `json:"signature"`
	ID                   string          `json:"id"`
	NameStyle            string          `json:"nameStyle"`
}

type CoverImgVInfo struct {
	Width                int            `json:"width"`
	Height               int            `json:"height"`
	Size                 int            `json:"size"`
	Type                 int            `json:"type"`
	Animated             bool           `json:"animated"`
	ThumbnailImage       ThumbnailImage `json:"thumbnailImage"`
	ExpandedImage        ExpandedImage  `json:"expandedImage"`
	ThumbnailImageCdnURL string         `json:"thumbnailImageCdnUrl"`
}
type CoverImgHInfo struct {
	Width                int            `json:"width"`
	Height               int            `json:"height"`
	Size                 int            `json:"size"`
	Type                 int            `json:"type"`
	Animated             bool           `json:"animated"`
	ThumbnailImage       ThumbnailImage `json:"thumbnailImage"`
	ExpandedImage        ExpandedImage  `json:"expandedImage"`
	ThumbnailImageCdnURL string         `json:"thumbnailImageCdnUrl"`
}
type BangumiPaymentType struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}
type BelongResource struct {
	CoverImageV string `json:"coverImageV"`
}
type ImgInfo struct {
	Width                int            `json:"width"`
	Height               int            `json:"height"`
	Size                 int            `json:"size"`
	Type                 int            `json:"type"`
	Animated             bool           `json:"animated"`
	ThumbnailImage       ThumbnailImage `json:"thumbnailImage"`
	ThumbnailImageCdnURL string         `json:"thumbnailImageCdnUrl"`
}
type PaymentType struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

type Sidelights struct {
	GroupID                     string       `json:"groupId"`
	LikeCount                   int          `json:"likeCount"`
	CoverImgInfo                CoverImgInfo `json:"coverImgInfo,omitempty"`
	BananaCountShow             string       `json:"bananaCountShow"`
	IsLike                      bool         `json:"isLike"`
	CommentCountTenThousandShow string       `json:"commentCountTenThousandShow,omitempty"`
	ViewCount                   string       `json:"viewCount"`
	StowCount                   string       `json:"stowCount"`
	BananaCount                 int          `json:"bananaCount"`
	Caption                     string       `json:"caption"`
	SourceID                    string       `json:"sourceId,omitempty"`
	CoverUrls                   string       `json:"coverUrls"`
	ContributeTime              string       `json:"contributeTime"`
	IsThrowBanana               bool         `json:"isThrowBanana"`
	DisplayPlayCount            string       `json:"displayPlayCount,omitempty"`
	VideoSizeType               int          `json:"videoSizeType,omitempty"`
	CommentCount                string       `json:"commentCount"`
	Type                        int          `json:"type"`
	ContentID                   string       `json:"contentId"`
	VideoID                     string       `json:"videoId,omitempty"`
	Vindex                      int          `json:"vindex,omitempty"`
	UserInfo                    User         `json:"user,omitempty"`
}
type RelatedBangumis struct {
	GroupID string `json:"groupId"`
	Name    string `json:"name"`
	ID      int    `json:"id"`
	Sort    int    `json:"sort"`
}
type RecommendBangumis struct {
	GroupID            string        `json:"groupId"`
	CoverImgVInfo      CoverImgVInfo `json:"coverImgVInfo"`
	CoverImgHInfo      CoverImgHInfo `json:"coverImgHInfo"`
	StowCountShow      string        `json:"stowCountShow"`
	PaymentType        PaymentType   `json:"paymentType"`
	Title              string        `json:"title"`
	LastUpdateItemName string        `json:"lastUpdateItemName"`
	StowCount          int           `json:"stowCount"`
	CoverImageV        string        `json:"coverImageV"`
	CoverImageH        string        `json:"coverImageH"`
	UpdateStatus       int           `json:"updateStatus"`
	ID                 int           `json:"id"`
	ItemCount          int           `json:"itemCount"`
}
type HotTags struct {
	GroupID          string `json:"groupId"`
	TagID            int    `json:"tagId"`
	TagCover         string `json:"tagCover"`
	TagName          string `json:"tagName"`
	TagBackGround    string `json:"tagBackGround"`
	TagCountStr      string `json:"tagCountStr"`
	StowCount        int    `json:"stowCount"`
	Summary          string `json:"summary"`
	TagResourceCount int    `json:"tagResourceCount"`
}
type BangumiStyleList struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Priority        int    `json:"priority"`
	Status          string `json:"status"`
	CreateTime      int64  `json:"createTime"`
	UpdateTime      int64  `json:"updateTime"`
	CategoryEntryID int    `json:"categoryEntryId"`
}
