
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UserSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Username  *string `json:"username" form:"username"` 
      Password  *string `json:"password" form:"password"` 
      Nickname  *string `json:"nickname" form:"nickname"` 
      Avatar  *string `json:"avatar" form:"avatar"` 
      Gender  *string `json:"gender" form:"gender"` 
      VoiceTag  *string `json:"voiceTag" form:"voiceTag"` 
      InterestTags  *string `json:"interestTags" form:"interestTags"` 
      Level  *int `json:"level" form:"level"` 
      VipStatus  *string `json:"vipStatus" form:"vipStatus"` 
      VipExpireTime  *time.Time `json:"vipExpireTime" form:"vipExpireTime"` 
      Balance  *float64 `json:"balance" form:"balance"` 
      Diamond  *int `json:"diamond" form:"diamond"` 
      LastLoginIp  *string `json:"lastLoginIp" form:"lastLoginIp"` 
      Status  *string `json:"status" form:"status"` 
      FreezeReason  *string `json:"freezeReason" form:"freezeReason"` 
      FreezeEndTime  *time.Time `json:"freezeEndTime" form:"freezeEndTime"` 
      IsAuth  *string `json:"isAuth" form:"isAuth"` 
    request.PageInfo
}
