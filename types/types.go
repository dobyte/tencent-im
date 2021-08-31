/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/5/28 19:24
 * @Desc: TODO
 */

package types

type (
	BaseResp struct {
		ErrorCode int    `json:"ErrorCode"`
		ErrorInfo string `json:"ErrorInfo"`
	}
	
	ActionBaseResp struct {
		BaseResp
		ActionStatus string `json:"ActionStatus"`
	}
	
	AndroidInfo struct {
		Sound string `json:"Sound"`
	}
	
	ApnsInfo struct {
		Sound     string `json:"Sound"`
		BadgeMode int    `json:"BadgeMode"`
		Title     string `json:"Title"`
		SubTitle  string `json:"SubTitle"`
		Image     string `json:"Image"`
	}
	
	OfflinePushInfo struct {
		PushFlag    int         `json:"PushFlag"`
		Desc        string      `json:"Desc"`
		Ext         string      `json:"Ext"`
		AndroidInfo AndroidInfo `json:"AndroidInfo"`
		ApnsInfo    ApnsInfo    `json:"ApnsInfo"`
	}
	
	MsgBodyContentImageInfo struct {
		Type   int    `json:"Type"`
		Size   int    `json:"Size"`
		Width  int    `json:"Width"`
		Height int    `json:"Height"`
		URL    string `json:"URL"`
	}
	
	MsgBodyContent struct {
		Text           string                    `json:"Text"`
		Desc           string                    `json:"Desc"`
		Latitude       float64                   `json:"Latitude"`
		Longitude      float64                   `json:"Longitude"`
		Index          int                       `json:"Index"`
		Data           string                    `json:"Data"`
		Ext            string                    `json:"Ext"`
		Sound          string                    `json:"Sound"`
		Url            string                    `json:"Url"`
		Size           int                       `json:"Size"`
		Second         int                       `json:"Second"`
		DownloadFlag   int                       `json:"Download_Flag"`
		UUID           string                    `json:"UUID"`
		ImageFormat    int                       `json:"ImageFormat"`
		ImageInfoArray []MsgBodyContentImageInfo `json:"ImageInfoArray"`
	}
	
	MsgBody struct {
		MsgType    string      `json:"MsgType"`
		MsgContent interface{} `json:"MsgContent"`
	}
	
	// 标签对
	TagPair struct {
		Tag   string      `json:"Tag"`   // 标签
		Value interface{} `json:"Value"` // 标签值
	}
	
	// 文本消息内容
	MsgTextContent struct {
		Text string `json:"Text"` // （必填）消息内容。当接收方为 iOS 或 Android 后台在线时，作为离线推送的文本展示。
	}
	
	// 地理位置消息元素
	MsgLocationContent struct {
		Desc      string  `json:"Desc"`      // （必填）地理位置描述信息
		Latitude  float64 `json:"Latitude"`  // （必填）纬度
		Longitude float64 `json:"Longitude"` // （必填）经度
	}
	
	// 表情消息元素
	MsgFaceContent struct {
		Index int    `json:"Index"` // （必填）表情索引，用户自定义
		Data  string `json:"Data"`  // （选填）额外数据
	}
	
	// 自定义消息元素
	MsgCustomContent struct {
		Desc  string `json:"Desc"`  // （选填）自定义消息描述信息。当接收方为 iOS 或 Android 后台在线时，做离线推送文本展示。 若发送自定义消息的同时设置了 OfflinePushInfo.Desc 字段，此字段会被覆盖，请优先填 OfflinePushInfo.Desc 字段。
		Data  string `json:"Data"`  // （必填）自定义消息数据。 不作为 APNs 的 payload 字段下发，故从 payload 中无法获取 Data 字段
		Ext   string `json:"Ext"`   // （选填）扩展字段。当接收方为 iOS 系统且应用处在后台时，此字段作为 APNs 请求包 Payloads 中的 Ext 键值下发，Ext 的协议格式由业务方确定，APNs 只做透传。
		Sound string `json:"Sound"` // （选填）自定义 APNs 推送铃音。
	}
	
	// 语音消息元素
	MsgSoundContent struct {
		Url          string `json:"Url"`           // （必填）语音下载地址，可通过该 URL 地址直接下载相应语音
		Size         int    `json:"Size"`          // （必填）语音数据大小，单位：字节。
		Second       int    `json:"Second"`        // （必填）语音时长，单位：秒。
		DownloadFlag int    `json:"Download_Flag"` // （必填）语音下载方式标记。目前 Download_Flag 取值只能为2，表示可通过Url字段值的 URL 地址直接下载语音。
	}
	
	// 图像消息元素
	MsgImageContent struct {
		UUID        string      `json:"UUID"`           // （必填）图片序列号。后台用于索引图片的键值。
		ImageFormat int         `json:"ImageFormat"`    // （必填）图片格式。JPG = 1，GIF = 2，PNG = 3，BMP = 4，其他 = 255。
		ImageInfos  []ImageInfo `json:"ImageInfoArray"` // （必填）原图、缩略图或者大图下载信息。
	}
	
	// 文件消息元素
	MsgFileContent struct {
		Url          string `json:"Url"`           // （必填）文件下载地址，可通过该 URL 地址直接下载相应文件
		Size         int    `json:"FileSize"`      // （必填）文件数据大小，单位：字节
		Name         string `json:"FileName"`      // （必填）文件名称
		DownloadFlag int    `json:"Download_Flag"` // （必填）文件下载方式标记。目前 Download_Flag 取值只能为2，表示可通过Url字段值的 URL 地址直接下载文件。
	}
	
	// 视频消息元素
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
	
	// 图片下载信息
	ImageInfo struct {
		Type   int    `json:"Type"`   // （必填）图片类型： 1-原图，2-大图，3-缩略图。
		Size   int    `json:"Size"`   // （必填）图片数据大小，单位：字节。
		Width  int    `json:"Width"`  // （必填）图片宽度。
		Height int    `json:"Height"` // （必填）图片高度。
		Url    string `json:"URL"`    // （必填）图片下载地址。
	}
)
