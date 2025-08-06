
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ViolationSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      ViolationMean  *string `json:"violationMean" form:"violationMean"` 
      Result  *string `json:"result" form:"result"` 
    request.PageInfo
}
