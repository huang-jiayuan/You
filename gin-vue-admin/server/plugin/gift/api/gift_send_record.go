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

var GiftSendRecord = new(giftSendRecord)

type giftSendRecord struct {}

// CreateGiftSendRecord 创建giftSendRecord表
// @Tags GiftSendRecord
// @Summary 创建giftSendRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftSendRecord true "创建giftSendRecord表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /giftSendRecord/createGiftSendRecord [post]
func (a *giftSendRecord) CreateGiftSendRecord(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.GiftSendRecord
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.CreatedBy = utils.GetUserID(c)
	err = serviceGiftSendRecord.CreateGiftSendRecord(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteGiftSendRecord 删除giftSendRecord表
// @Tags GiftSendRecord
// @Summary 删除giftSendRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftSendRecord true "删除giftSendRecord表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /giftSendRecord/deleteGiftSendRecord [delete]
func (a *giftSendRecord) DeleteGiftSendRecord(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := serviceGiftSendRecord.DeleteGiftSendRecord(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteGiftSendRecordByIds 批量删除giftSendRecord表
// @Tags GiftSendRecord
// @Summary 批量删除giftSendRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /giftSendRecord/deleteGiftSendRecordByIds [delete]
func (a *giftSendRecord) DeleteGiftSendRecordByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := serviceGiftSendRecord.DeleteGiftSendRecordByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateGiftSendRecord 更新giftSendRecord表
// @Tags GiftSendRecord
// @Summary 更新giftSendRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftSendRecord true "更新giftSendRecord表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /giftSendRecord/updateGiftSendRecord [put]
func (a *giftSendRecord) UpdateGiftSendRecord(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.GiftSendRecord
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.UpdatedBy = utils.GetUserID(c)
	err = serviceGiftSendRecord.UpdateGiftSendRecord(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindGiftSendRecord 用id查询giftSendRecord表
// @Tags GiftSendRecord
// @Summary 用id查询giftSendRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询giftSendRecord表"
// @Success 200 {object} response.Response{data=model.GiftSendRecord,msg=string} "查询成功"
// @Router /giftSendRecord/findGiftSendRecord [get]
func (a *giftSendRecord) FindGiftSendRecord(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	regiftSendRecord, err := serviceGiftSendRecord.GetGiftSendRecord(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(regiftSendRecord, c)
}
// GetGiftSendRecordList 分页获取giftSendRecord表列表
// @Tags GiftSendRecord
// @Summary 分页获取giftSendRecord表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.GiftSendRecordSearch true "分页获取giftSendRecord表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /giftSendRecord/getGiftSendRecordList [get]
func (a *giftSendRecord) GetGiftSendRecordList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.GiftSendRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceGiftSendRecord.GetGiftSendRecordInfoList(ctx,pageInfo)
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
// GetGiftSendRecordPublic 不需要鉴权的giftSendRecord表接口
// @Tags GiftSendRecord
// @Summary 不需要鉴权的giftSendRecord表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /giftSendRecord/getGiftSendRecordPublic [get]
func (a *giftSendRecord) GetGiftSendRecordPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceGiftSendRecord.GetGiftSendRecordPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的giftSendRecord表接口信息"}, "获取成功", c)
}
