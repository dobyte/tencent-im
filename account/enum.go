/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/9/7 11:09
 * @Desc: TODO
 */

package account

// ImportedStatus 导入状态
type ImportedStatus string

const (
	ImportedStatusNo  ImportedStatus = "NotImported" // 未导入
	ImportedStatusYes ImportedStatus = "Imported"    // 已导入
)
