/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/9/7 11:09
 * @Desc: TODO
 */

package account

// ImportedStatusType 导入状态
type ImportedStatusType string

const (
	ImportedStatusNo  ImportedStatusType = "NotImported" // 未导入
	ImportedStatusYes ImportedStatusType = "Imported"    // 已导入
)
