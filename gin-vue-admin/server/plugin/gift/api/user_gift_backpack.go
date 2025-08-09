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

var UserGiftBackpack = new(userGiftBackpack)

type userGiftBackpack struct {}

// CreateUserGiftBackpack 创建用户背包礼物表
// @Tags UserGiftBackpack
// @Summary 创建用户背包礼物表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserGiftBackpack true "创建用户背包礼物表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userGiftBackpack/createUserGiftBackpack [post]
func (a *userGiftBackpack) CreateUserGiftBackpack(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.UserGiftBackpack
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.CreatedBy = utils.GetUserID(c)
	err = serviceUserGiftBackpack.CreateUserGiftBackpack(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUserGiftBackpack 删除用户背包礼物表
// @Tags UserGiftBackpack
// @Summary 删除用户背包礼物表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserGiftBackpack true "删除用户背包礼物表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userGiftBackpack/deleteUserGiftBackpack [delete]
func (a *userGiftBackpack) DeleteUserGiftBackpack(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := serviceUserGiftBackpack.DeleteUserGiftBackpack(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteUserGiftBackpackByIds 批量删除用户背包礼物表
// @Tags UserGiftBackpack
// @Summary 批量删除用户背包礼物表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userGiftBackpack/deleteUserGiftBackpackByIds [delete]
func (a *userGiftBackpack) DeleteUserGiftBackpackByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := serviceUserGiftBackpack.DeleteUserGiftBackpackByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateUserGiftBackpack 更新用户背包礼物表
// @Tags UserGiftBackpack
// @Summary 更新用户背包礼物表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserGiftBackpack true "更新用户背包礼物表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userGiftBackpack/updateUserGiftBackpack [put]
func (a *userGiftBackpack) UpdateUserGiftBackpack(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.UserGiftBackpack
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    info.UpdatedBy = utils.GetUserID(c)
	err = serviceUserGiftBackpack.UpdateUserGiftBackpack(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindUserGiftBackpack 用id查询用户背包礼物表
// @Tags UserGiftBackpack
// @Summary 用id查询用户背包礼物表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询用户背包礼物表"
// @Success 200 {object} response.Response{data=model.UserGiftBackpack,msg=string} "查询成功"
// @Router /userGiftBackpack/findUserGiftBackpack [get]
func (a *userGiftBackpack) FindUserGiftBackpack(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reuserGiftBackpack, err := serviceUserGiftBackpack.GetUserGiftBackpack(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reuserGiftBackpack, c)
}
// GetUserGiftBackpackList 分页获取用户背包礼物表列表
// @Tags UserGiftBackpack
// @Summary 分页获取用户背包礼物表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.UserGiftBackpackSearch true "分页获取用户背包礼物表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userGiftBackpack/getUserGiftBackpackList [get]
func (a *userGiftBackpack) GetUserGiftBackpackList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.UserGiftBackpackSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceUserGiftBackpack.GetUserGiftBackpackInfoList(ctx,pageInfo)
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
// GetUserGiftBackpackPublic 不需要鉴权的用户背包礼物表接口
// @Tags UserGiftBackpack
// @Summary 不需要鉴权的用户背包礼物表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userGiftBackpack/getUserGiftBackpackPublic [get]
func (a *userGiftBackpack) GetUserGiftBackpackPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceUserGiftBackpack.GetUserGiftBackpackPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的用户背包礼物表接口信息"}, "获取成功", c)
}
