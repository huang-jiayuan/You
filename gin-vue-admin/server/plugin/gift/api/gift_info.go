package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var GiftInfo = new(giftInfo)

type giftInfo struct {}

// CreateGiftInfo 创建giftInfo表
// @Tags GiftInfo
// @Summary 创建giftInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftInfo true "创建giftInfo表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /giftInfo/createGiftInfo [post]
func (a *giftInfo) CreateGiftInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.GiftInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.CreatedBy = utils.GetUserID(c)
	err = serviceGiftInfo.CreateGiftInfo(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteGiftInfo 删除giftInfo表
// @Tags GiftInfo
// @Summary 删除giftInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftInfo true "删除giftInfo表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /giftInfo/deleteGiftInfo [delete]
func (a *giftInfo) DeleteGiftInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	giftId := c.Query("giftId")
    userID := utils.GetUserID(c)
	err := serviceGiftInfo.DeleteGiftInfo(ctx,giftId,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteGiftInfoByIds 批量删除giftInfo表
// @Tags GiftInfo
// @Summary 批量删除giftInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /giftInfo/deleteGiftInfoByIds [delete]
func (a *giftInfo) DeleteGiftInfoByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	giftIds := c.QueryArray("giftIds[]")
    userID := utils.GetUserID(c)
	err := serviceGiftInfo.DeleteGiftInfoByIds(ctx,giftIds,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateGiftInfo 更新giftInfo表
// @Tags GiftInfo
// @Summary 更新giftInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftInfo true "更新giftInfo表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /giftInfo/updateGiftInfo [put]
func (a *giftInfo) UpdateGiftInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.GiftInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.UpdatedBy = utils.GetUserID(c)
	err = serviceGiftInfo.UpdateGiftInfo(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindGiftInfo 用id查询giftInfo表
// @Tags GiftInfo
// @Summary 用id查询giftInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param giftId query int true "用id查询giftInfo表"
// @Success 200 {object} response.Response{data=model.GiftInfo,msg=string} "查询成功"
// @Router /giftInfo/findGiftInfo [get]
func (a *giftInfo) FindGiftInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	giftId := c.Query("giftId")
	regiftInfo, err := serviceGiftInfo.GetGiftInfo(ctx,giftId)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(regiftInfo, c)
}
// GetGiftInfoList 分页获取giftInfo表列表
// @Tags GiftInfo
// @Summary 分页获取giftInfo表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.GiftInfoSearch true "分页获取giftInfo表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /giftInfo/getGiftInfoList [get]
func (a *giftInfo) GetGiftInfoList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.GiftInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceGiftInfo.GetGiftInfoInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}
// GetGiftInfoPublic 不需要鉴权的giftInfo表接口
// @Tags GiftInfo
// @Summary 不需要鉴权的giftInfo表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /giftInfo/getGiftInfoPublic [get]
func (a *giftInfo) GetGiftInfoPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceGiftInfo.GetGiftInfoPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的giftInfo表接口信息"}, "获取成功", c)
}
