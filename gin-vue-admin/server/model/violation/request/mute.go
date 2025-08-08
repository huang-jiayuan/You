
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MuteSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      UserId  *string `json:"userId" form:"userId"` 
      MuteType  *string `json:"muteType" form:"muteType"` 
      Reason  *string `json:"reason" form:"reason"` 
      OperatorId  *int `json:"operatorId" form:"operatorId"` 
      Status  *string `json:"status" form:"status"` 
      MuteDay  *string `json:"muteDay" form:"muteDay"` 
      MuteResult  *string `json:"muteResult" form:"muteResult"` 
    request.PageInfo
}
