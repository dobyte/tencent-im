/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/28 1:14 上午
 * @Desc: TODO
 */

package enum

const (
	// 消息类型
	MsgText     = "TIMTextElem"      // 消息元素
	MsgLocation = "TIMLocationElem"  // 地理位置消息元素
	MsgFace     = "TIMFaceElem"      // 表情消息元素
	MsgCustom   = "TIMCustomElem"    // 自定义消息元素
	MsgSound    = "TIMSoundElem"     // 语音消息元素
	MsgImage    = "TIMImageElem"     // 图像消息元素
	MsgFile     = "TIMFileElem"      // 文件消息元素
	MsgVideo    = "TIMVideoFileElem" // 视频消息元素
	
	// 图片格式
	ImageFormatJPG   = 1   // JPG格式
	ImageFormatGIF   = 2   // GIF格式
	ImageFormatPNG   = 3   // PNG格式
	ImageFormatBMP   = 4   // BMP格式
	ImageFormatOTHER = 255 // 其他格式
	
	// 图片类型
	ImageTypeOriginal = 1 // 原图
	ImageTypePic      = 2 // 大图
	ImageTypeThumb    = 3 // 缩略图
)
