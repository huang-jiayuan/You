
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type GiftInfoSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       GiftName  *string `json:"giftName" form:"giftName"` 
    request.PageInfo
}
