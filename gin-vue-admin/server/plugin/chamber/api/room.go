package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var Room = new(room)

type room struct {}

// CreateRoom 创建房间表
// @Tags Room
// @Summary 创建房间表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Room true "创建房间表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /room/createRoom [post]
func (a *room) CreateRoom(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Room
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.CreatedBy = utils.GetUserID(c)
	err = serviceRoom.CreateRoom(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteRoom 删除房间表
// @Tags Room
// @Summary 删除房间表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Room true "删除房间表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /room/deleteRoom [delete]
func (a *room) DeleteRoom(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := serviceRoom.DeleteRoom(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteRoomByIds 批量删除房间表
// @Tags Room
// @Summary 批量删除房间表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /room/deleteRoomByIds [delete]
func (a *room) DeleteRoomByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := serviceRoom.DeleteRoomByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateRoom 更新房间表
// @Tags Room
// @Summary 更新房间表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Room true "更新房间表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /room/updateRoom [put]
func (a *room) UpdateRoom(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Room
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.UpdatedBy = utils.GetUserID(c)
	err = serviceRoom.UpdateRoom(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindRoom 用id查询房间表
// @Tags Room
// @Summary 用id查询房间表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询房间表"
// @Success 200 {object} response.Response{data=model.Room,msg=string} "查询成功"
// @Router /room/findRoom [get]
func (a *room) FindRoom(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reroom, err := serviceRoom.GetRoom(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reroom, c)
}
// GetRoomList 分页获取房间表列表
// @Tags Room
// @Summary 分页获取房间表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.RoomSearch true "分页获取房间表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /room/getRoomList [get]
func (a *room) GetRoomList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.RoomSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceRoom.GetRoomInfoList(ctx,pageInfo)
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
// GetRoomPublic 不需要鉴权的房间表接口
// @Tags Room
// @Summary 不需要鉴权的房间表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /room/getRoomPublic [get]
func (a *room) GetRoomPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceRoom.GetRoomPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的房间表接口信息"}, "获取成功", c)
}
