
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CarouselImageSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Image  *string `json:"image" form:"image"` 
      Title  *string `json:"title" form:"title"` 
      Url  *string `json:"url" form:"url"` 
      OrderId  *int `json:"orderId" form:"orderId"` 
      Status  *string `json:"status" form:"status"` 
    request.PageInfo
}
