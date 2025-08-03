package users

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/users"
    usersReq "github.com/flipped-aurora/gin-vue-admin/server/model/users/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type UserApi struct {}



// CreateUser 创建user表
// @Tags User
// @Summary 创建user表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body users.User true "创建user表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /user1/createUser [post]
func (user1Api *UserApi) CreateUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var user1 users.User
	err := c.ShouldBindJSON(&user1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    user1.CreatedBy = utils.GetUserID(c)
	err = user1Service.CreateUser(ctx,&user1)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUser 删除user表
// @Tags User
// @Summary 删除user表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body users.User true "删除user表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /user1/deleteUser [delete]
func (user1Api *UserApi) DeleteUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := user1Service.DeleteUser(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserByIds 批量删除user表
// @Tags User
// @Summary 批量删除user表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /user1/deleteUserByIds [delete]
func (user1Api *UserApi) DeleteUserByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := user1Service.DeleteUserByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUser 更新user表
// @Tags User
// @Summary 更新user表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body users.User true "更新user表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /user1/updateUser [put]
func (user1Api *UserApi) UpdateUser(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var user1 users.User
	err := c.ShouldBindJSON(&user1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    user1.UpdatedBy = utils.GetUserID(c)
	err = user1Service.UpdateUser(ctx,user1)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUser 用id查询user表
// @Tags User
// @Summary 用id查询user表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询user表"
// @Success 200 {object} response.Response{data=users.User,msg=string} "查询成功"
// @Router /user1/findUser [get]
func (user1Api *UserApi) FindUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reuser1, err := user1Service.GetUser(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reuser1, c)
}
// GetUserList 分页获取user表列表
// @Tags User
// @Summary 分页获取user表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query usersReq.UserSearch true "分页获取user表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /user1/getUserList [get]
func (user1Api *UserApi) GetUserList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo usersReq.UserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := user1Service.GetUserInfoList(ctx,pageInfo)
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

// GetUserPublic 不需要鉴权的user表接口
// @Tags User
// @Summary 不需要鉴权的user表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /user1/getUserPublic [get]
func (user1Api *UserApi) GetUserPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    user1Service.GetUserPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的user表接口信息",
    }, "获取成功", c)
}
