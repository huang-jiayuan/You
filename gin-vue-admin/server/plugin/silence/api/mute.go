package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/silence/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/silence/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var Mute = new(mute)

type mute struct {}

// CreateMute 创建mute表
// @Tags Mute
// @Summary 创建mute表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Mute true "创建mute表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mute/createMute [post]
func (a *mute) CreateMute(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Mute
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.CreatedBy = utils.GetUserID(c)
	err = serviceMute.CreateMute(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMute 删除mute表
// @Tags Mute
// @Summary 删除mute表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Mute true "删除mute表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mute/deleteMute [delete]
func (a *mute) DeleteMute(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := serviceMute.DeleteMute(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteMuteByIds 批量删除mute表
// @Tags Mute
// @Summary 批量删除mute表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mute/deleteMuteByIds [delete]
func (a *mute) DeleteMuteByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := serviceMute.DeleteMuteByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateMute 更新mute表
// @Tags Mute
// @Summary 更新mute表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Mute true "更新mute表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mute/updateMute [put]
func (a *mute) UpdateMute(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Mute
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.UpdatedBy = utils.GetUserID(c)
	err = serviceMute.UpdateMute(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindMute 用id查询mute表
// @Tags Mute
// @Summary 用id查询mute表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mute表"
// @Success 200 {object} response.Response{data=model.Mute,msg=string} "查询成功"
// @Router /mute/findMute [get]
func (a *mute) FindMute(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remute, err := serviceMute.GetMute(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(remute, c)
}
// GetMuteList 分页获取mute表列表
// @Tags Mute
// @Summary 分页获取mute表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.MuteSearch true "分页获取mute表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mute/getMuteList [get]
func (a *mute) GetMuteList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.MuteSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMute.GetMuteInfoList(ctx,pageInfo)
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
// GetMutePublic 不需要鉴权的mute表接口
// @Tags Mute
// @Summary 不需要鉴权的mute表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mute/getMutePublic [get]
func (a *mute) GetMutePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceMute.GetMutePublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的mute表接口信息"}, "获取成功", c)
}
