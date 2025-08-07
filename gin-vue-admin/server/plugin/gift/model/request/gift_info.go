
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)
type GiftInfoSearch struct{
       GiftName  *string `json:"giftName" form:"giftName"` 
    request.PageInfo
}
