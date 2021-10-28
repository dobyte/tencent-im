/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/28 19:24
 * @Desc: TODO
 */

package types

type (
	BaseResp struct {
		ErrorCode    int    `json:"ErrorCode"`
		ErrorInfo    string `json:"ErrorInfo"`
		ErrorDisplay string `json:"ErrorDisplay,omitempty"`
	}

	ActionBaseResp struct {
		BaseResp
		ActionStatus string `json:"ActionStatus"`
	}

	// AndroidInfo Android离线推送消息
	AndroidInfo struct {
		Sound                  string `json:"Sound,omitempty"`                  // （选填）Android 离线推送声音文件路径。
		HuaWeiChannelID        string `json:"HuaWeiChannelID,omitempty"`        // （选填）华为手机 EMUI 10.0 及以上的通知渠道字段。该字段不为空时，会覆盖控制台配置的 ChannelID 值；该字段为空时，不会覆盖控制台配置的 ChannelID 值。
		XiaoMiChannelID        string `json:"XiaoMiChannelID,omitempty"`        // （选填）小米手机 MIUI 10 及以上的通知类别（Channel）适配字段。该字段不为空时，会覆盖控制台配置的 ChannelID 值；该字段为空时，不会覆盖控制台配置的 ChannelID 值。
		OPPOChannelID          string `json:"OPPOChannelID,omitempty"`          // （选填）OPPO 手机 Android 8.0 及以上的 NotificationChannel 通知适配字段。该字段不为空时，会覆盖控制台配置的 ChannelID 值；该字段为空时，不会覆盖控制台配置的 ChannelID 值。
		GoogleChannelID        string `json:"GoogleChannelID,omitempty"`        // （选填）Google 手机 Android 8.0 及以上的通知渠道字段。Google 推送新接口（上传证书文件）支持 channel id，旧接口（填写服务器密钥）不支持。
		VIVOClassification     int    `json:"VIVOClassification,omitempty"`     // （选填）VIVO 手机推送消息分类，“0”代表运营消息，“1”代表系统消息，不填默认为1。
		HuaWeiImportance       string `json:"HuaWeiImportance,omitempty"`       // （选填）华为推送通知消息分类，取值为 LOW、NORMAL，不填默认为 NORMAL。
		ExtAsHuaweiIntentParam int    `json:"ExtAsHuaweiIntentParam,omitempty"` // （选填）在控制台配置华为推送为“打开应用内指定页面”的前提下，传“1”表示将透传内容 Ext 作为 Intent 的参数，“0”表示将透传内容 Ext 作为 Action 参数。不填默认为0。两种传参区别可参见 华为推送文档。
	}

	// ApnsInfo IOS离线推送消息
	ApnsInfo struct {
		BadgeMode      int    `json:"BadgeMode,omitempty"`      // （选填）这个字段缺省或者为0表示需要计数，为1表示本条消息不需要计数，即右上角图标数字不增加。
		Title          string `json:"Title,omitempty"`          // （选填）该字段用于标识 APNs 推送的标题，若填写则会覆盖最上层 Title。
		SubTitle       string `json:"SubTitle,omitempty"`       // （选填）该字段用于标识 APNs 推送的子标题。
		Image          string `json:"Image,omitempty"`          // （选填）该字段用于标识 APNs 携带的图片地址，当客户端拿到该字段时，可以通过下载图片资源的方式将图片展示在弹窗上。
		MutableContent int    `json:"MutableContent,omitempty"` // （选填）为1表示开启 iOS 10 的推送扩展，默认为0。
	}

	// OfflinePushInfo 离线推送消息
	OfflinePushInfo struct {
		PushFlag    int          `json:"PushFlag,omitempty"`    // （选填）推送标识。0表示推送，1表示不离线推送。
		Title       string       `json:"Title,omitempty"`       // （选填）离线推送标题。该字段为 iOS 和 Android 共用。
		Desc        string       `json:"Desc,omitempty"`        // （选填）离线推送内容。该字段会覆盖上面各种消息元素 TIMMsgElement 的离线推送展示文本。若发送的消息只有一个 TIMCustomElem 自定义消息元素，该 Desc 字段会覆盖 TIMCustomElem 中的 Desc 字段。如果两个 Desc 字段都不填，将收不到该自定义消息的离线推送。
		Ext         string       `json:"Ext,omitempty"`         // （选填）离线推送透传内容。由于国内各 Android 手机厂商的推送平台要求各不一样，请保证此字段为 JSON 格式，否则可能会导致收不到某些厂商的离线推送。
		AndroidInfo *AndroidInfo `json:"AndroidInfo,omitempty"` // （选填）Android 离线推送消息
		ApnsInfo    *ApnsInfo    `json:"ApnsInfo,omitempty"`    // （选填）IOS离线推送消息
	}

	// MsgBody 消息内容
	MsgBody struct {
		MsgType    string      `json:"MsgType"`
		MsgContent interface{} `json:"MsgContent"`
	}

	// TagPair 标签对
	TagPair struct {
		Tag   string      `json:"Tag"`   // 标签
		Value interface{} `json:"Value"` // 标签值
	}

	// MsgTextContent 文本消息内容
	MsgTextContent struct {
		Text string `json:"Text"` // （必填）消息内容。当接收方为 iOS 或 Android 后台在线时，作为离线推送的文本展示。
	}

	// MsgLocationContent 地理位置消息元素
	MsgLocationContent struct {
		Desc      string  `json:"Desc"`      // （必填）地理位置描述信息
		Latitude  float64 `json:"Latitude"`  // （必填）纬度
		Longitude float64 `json:"Longitude"` // （必填）经度
	}

	// MsgFaceContent 表情消息元素
	MsgFaceContent struct {
		Index int    `json:"Index"` // （必填）表情索引，用户自定义
		Data  string `json:"Data"`  // （选填）额外数据
	}

	// MsgCustomContent 自定义消息元素
	MsgCustomContent struct {
		Desc  string `json:"Desc"`  // （选填）自定义消息描述信息。当接收方为 iOS 或 Android 后台在线时，做离线推送文本展示。 若发送自定义消息的同时设置了 OfflinePushInfo.Desc 字段，此字段会被覆盖，请优先填 OfflinePushInfo.Desc 字段。
		Data  string `json:"Data"`  // （必填）自定义消息数据。 不作为 APNs 的 payload 字段下发，故从 payload 中无法获取 Data 字段
		Ext   string `json:"Ext"`   // （选填）扩展字段。当接收方为 iOS 系统且应用处在后台时，此字段作为 APNs 请求包 Payloads 中的 Ext 键值下发，Ext 的协议格式由业务方确定，APNs 只做透传。
		Sound string `json:"Sound"` // （选填）自定义 APNs 推送铃音。
	}

	// MsgSoundContent 语音消息元素
	MsgSoundContent struct {
		Url          string `json:"Url"`           // （必填）语音下载地址，可通过该 URL 地址直接下载相应语音
		Size         int    `json:"Size"`          // （必填）语音数据大小，单位：字节。
		Second       int    `json:"Second"`        // （必填）语音时长，单位：秒。
		DownloadFlag int    `json:"Download_Flag"` // （必填）语音下载方式标记。目前 Download_Flag 取值只能为2，表示可通过Url字段值的 URL 地址直接下载语音。
	}

	// MsgImageContent 图像消息元素
	MsgImageContent struct {
		UUID        string      `json:"UUID"`           // （必填）图片序列号。后台用于索引图片的键值。
		ImageFormat int         `json:"ImageFormat"`    // （必填）图片格式。JPG = 1，GIF = 2，PNG = 3，BMP = 4，其他 = 255。
		ImageInfos  []ImageInfo `json:"ImageInfoArray"` // （必填）原图、缩略图或者大图下载信息。
	}

	// MsgFileContent 文件消息元素
	MsgFileContent struct {
		Url          string `json:"Url"`           // （必填）文件下载地址，可通过该 URL 地址直接下载相应文件
		Size         int    `json:"FileSize"`      // （必填）文件数据大小，单位：字节
		Name         string `json:"FileName"`      // （必填）文件名称
		DownloadFlag int    `json:"Download_Flag"` // （必填）文件下载方式标记。目前 Download_Flag 取值只能为2，表示可通过Url字段值的 URL 地址直接下载文件。
	}

	// MsgVideoContent 视频消息元素
	MsgVideoContent struct {
		VideoUrl          string `json:"VideoUrl"`          // （必填）视频下载地址。可通过该 URL 地址直接下载相应视频
		VideoSize         int    `json:"VideoSize"`         // （必填）视频数据大小，单位：字节
		VideoSecond       int    `json:"VideoSecond"`       // （必填）视频时长，单位：秒
		VideoFormat       string `json:"VideoFormat"`       // （必填）视频格式，例如 mp4
		VideoDownloadFlag int    `json:"VideoDownloadFlag"` // （必填）视频下载方式标记。目前 VideoDownloadFlag 取值只能为2，表示可通过VideoUrl字段值的 URL 地址直接下载视频。
		ThumbUrl          string `json:"ThumbUrl"`          // （必填）视频缩略图下载地址。可通过该 URL 地址直接下载相应视频缩略图。
		ThumbSize         int    `json:"ThumbSize"`         // （必填）缩略图大小，单位：字节
		ThumbWidth        int    `json:"ThumbWidth"`        // （必填）缩略图宽度
		ThumbHeight       int    `json:"ThumbHeight"`       // （必填）缩略图高度
		ThumbFormat       string `json:"ThumbFormat"`       // （必填）缩略图格式，例如 JPG、BMP 等
		ThumbDownloadFlag int    `json:"ThumbDownloadFlag"` // （必填）视频缩略图下载方式标记。目前 ThumbDownloadFlag 取值只能为2，表示可通过ThumbUrl字段值的 URL 地址直接下载视频缩略图。
	}

	// ImageInfo 图片下载信息
	ImageInfo struct {
		Type   int    `json:"Type"`   // （必填）图片类型： 1-原图，2-大图，3-缩略图。
		Size   int    `json:"Size"`   // （必填）图片数据大小，单位：字节。
		Width  int    `json:"Width"`  // （必填）图片宽度。
		Height int    `json:"Height"` // （必填）图片高度。
		Url    string `json:"URL"`    // （必填）图片下载地址。
	}

	// GenderType 性别类型
	GenderType string

	// AllowType 加好友验证方式
	AllowType string

	// AdminForbidType 管理员禁止加好友标识类型
	AdminForbidType string

	// SyncOtherMachine 同步至其他设备
	SyncOtherMachine int

	// PushFlag 推送标识
	PushFlag int

	// HuaWeiImportance 华为推送通知消息分类
	HuaWeiImportance string

	// HuaweiIntentParam 华为推送为“打开应用内指定页面”的前提下透传参数行为
	HuaweiIntentParam int

	// VivoClassification VIVO手机推送消息分类
	VivoClassification int

	// BadgeMode IOS徽章计数模式
	BadgeMode int

	// MutableContent IOS10的推送扩展开关
	MutableContent int
)
