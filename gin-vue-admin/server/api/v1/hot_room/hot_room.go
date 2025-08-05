package hot_room

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/hot_room"
    hot_roomReq "github.com/flipped-aurora/gin-vue-admin/server/model/hot_room/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type HotRoomApi struct {}



// CreateHotRoom 创建hotRoom表
// @Tags HotRoom
// @Summary 创建hotRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body hot_room.HotRoom true "创建hotRoom表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /hotRoom/createHotRoom [post]
func (hotRoomApi *HotRoomApi) CreateHotRoom(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var hotRoom hot_room.HotRoom
	err := c.ShouldBindJSON(&hotRoom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    hotRoom.CreatedBy = utils.GetUserID(c)
	err = hotRoomService.CreateHotRoom(ctx,&hotRoom)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteHotRoom 删除hotRoom表
// @Tags HotRoom
// @Summary 删除hotRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body hot_room.HotRoom true "删除hotRoom表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /hotRoom/deleteHotRoom [delete]
func (hotRoomApi *HotRoomApi) DeleteHotRoom(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := hotRoomService.DeleteHotRoom(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteHotRoomByIds 批量删除hotRoom表
// @Tags HotRoom
// @Summary 批量删除hotRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /hotRoom/deleteHotRoomByIds [delete]
func (hotRoomApi *HotRoomApi) DeleteHotRoomByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := hotRoomService.DeleteHotRoomByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateHotRoom 更新hotRoom表
// @Tags HotRoom
// @Summary 更新hotRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body hot_room.HotRoom true "更新hotRoom表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /hotRoom/updateHotRoom [put]
func (hotRoomApi *HotRoomApi) UpdateHotRoom(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var hotRoom hot_room.HotRoom
	err := c.ShouldBindJSON(&hotRoom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    hotRoom.UpdatedBy = utils.GetUserID(c)
	err = hotRoomService.UpdateHotRoom(ctx,hotRoom)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindHotRoom 用id查询hotRoom表
// @Tags HotRoom
// @Summary 用id查询hotRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询hotRoom表"
// @Success 200 {object} response.Response{data=hot_room.HotRoom,msg=string} "查询成功"
// @Router /hotRoom/findHotRoom [get]
func (hotRoomApi *HotRoomApi) FindHotRoom(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rehotRoom, err := hotRoomService.GetHotRoom(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rehotRoom, c)
}
// GetHotRoomList 分页获取hotRoom表列表
// @Tags HotRoom
// @Summary 分页获取hotRoom表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query hot_roomReq.HotRoomSearch true "分页获取hotRoom表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /hotRoom/getHotRoomList [get]
func (hotRoomApi *HotRoomApi) GetHotRoomList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo hot_roomReq.HotRoomSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := hotRoomService.GetHotRoomInfoList(ctx,pageInfo)
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

// GetHotRoomPublic 不需要鉴权的hotRoom表接口
// @Tags HotRoom
// @Summary 不需要鉴权的hotRoom表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /hotRoom/getHotRoomPublic [get]
func (hotRoomApi *HotRoomApi) GetHotRoomPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    hotRoomService.GetHotRoomPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的hotRoom表接口信息",
    }, "获取成功", c)
}
