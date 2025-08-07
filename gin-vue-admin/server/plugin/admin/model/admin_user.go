
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AdminUser 管理员表 结构体
type AdminUser struct {
    global.GVA_MODEL
  UserId  *int `json:"userId" form:"userId" gorm:"comment:用户的ID;column:user_id;" binding:"required"`  //用户的ID
  AdminPassword  *string `json:"adminPassword" form:"adminPassword" gorm:"comment:管理员的密码;column:admin_password;size:32;" binding:"required"`  //管理员的密码
  AdminName  *string `json:"adminName" form:"adminName" gorm:"comment:管理员账号;column:admin_name;size:25;" binding:"required"`  //管理员账号
  Status  *string `json:"status" form:"status" gorm:"comment:用户的账号状态(0-封禁,1-冻结,2-禁言);column:status;size:2;" binding:"required"`  //用户的账号状态(0-封禁,1-冻结,2-禁言)
  IdAdmin  *string `json:"idAdmin" form:"idAdmin" gorm:"comment:管理员等级（1普通,2超级）;column:id_admin;size:2;" binding:"required"`  //管理员等级（1普通,2超级）
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 管理员表 AdminUser自定义表名 admin_user
func (AdminUser) TableName() string {
    return "admin_user"
}







