
// 自动生成模板CarouselImage
package carouse
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// carouselImage表 结构体  CarouselImage
type CarouselImage struct {
    global.GVA_MODEL
  Image  *string `json:"image" form:"image" gorm:"comment:图片预览;column:image;size:255;" binding:"required"`  //图片预览
  Title  *string `json:"title" form:"title" gorm:"comment:标题;column:title;size:60;" binding:"required"`  //标题
  Url  *string `json:"url" form:"url" gorm:"comment:跳转链接;column:url;size:255;" binding:"required"`  //跳转链接
  OrderId  *int `json:"orderId" form:"orderId" gorm:"comment:排序序号;column:order_id;size:19;" binding:"required"`  //排序序号
  Status  *string `json:"status" form:"status" gorm:"column:status;size:50;" binding:"required"`  //状态
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName carouselImage表 CarouselImage自定义表名 carousel_image
func (CarouselImage) TableName() string {
    return "carousel_image"
}





