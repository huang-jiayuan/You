
// 自动生成模板Violation
package violation
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// violation表 结构体  Violation
type Violation struct {
    global.GVA_MODEL
  ViolationMean  *string `json:"violationMean" form:"violationMean" gorm:"comment:违规原因;column:violation_mean;size:100;" binding:"required"`  //违规原因
  Result  *string `json:"result" form:"result" gorm:"comment:处理结果;column:result;size:100;" binding:"required"`  //处理结果
  ResultTime  *time.Time `json:"resultTime" form:"resultTime" gorm:"comment:处理时间;column:result_time;" binding:"required"`  //处理时间
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName violation表 Violation自定义表名 violation
func (Violation) TableName() string {
    return "violation"
}





