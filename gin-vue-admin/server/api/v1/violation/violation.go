package violation

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/violation"
    violationReq "github.com/flipped-aurora/gin-vue-admin/server/model/violation/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ViolationApi struct {}



// CreateViolation 创建violation表
// @Tags Violation
// @Summary 创建violation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body violation.Violation true "创建violation表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /violations/createViolation [post]
func (violationsApi *ViolationApi) CreateViolation(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var violations violation.Violation
	err := c.ShouldBindJSON(&violations)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    violations.CreatedBy = utils.GetUserID(c)
	err = violationsService.CreateViolation(ctx,&violations)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteViolation 删除violation表
// @Tags Violation
// @Summary 删除violation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body violation.Violation true "删除violation表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /violations/deleteViolation [delete]
func (violationsApi *ViolationApi) DeleteViolation(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := violationsService.DeleteViolation(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteViolationByIds 批量删除violation表
// @Tags Violation
// @Summary 批量删除violation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /violations/deleteViolationByIds [delete]
func (violationsApi *ViolationApi) DeleteViolationByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := violationsService.DeleteViolationByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateViolation 更新violation表
// @Tags Violation
// @Summary 更新violation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body violation.Violation true "更新violation表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /violations/updateViolation [put]
func (violationsApi *ViolationApi) UpdateViolation(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var violations violation.Violation
	err := c.ShouldBindJSON(&violations)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    violations.UpdatedBy = utils.GetUserID(c)
	err = violationsService.UpdateViolation(ctx,violations)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindViolation 用id查询violation表
// @Tags Violation
// @Summary 用id查询violation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询violation表"
// @Success 200 {object} response.Response{data=violation.Violation,msg=string} "查询成功"
// @Router /violations/findViolation [get]
func (violationsApi *ViolationApi) FindViolation(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reviolations, err := violationsService.GetViolation(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reviolations, c)
}
// GetViolationList 分页获取violation表列表
// @Tags Violation
// @Summary 分页获取violation表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query violationReq.ViolationSearch true "分页获取violation表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /violations/getViolationList [get]
func (violationsApi *ViolationApi) GetViolationList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo violationReq.ViolationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := violationsService.GetViolationInfoList(ctx,pageInfo)
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

// GetViolationPublic 不需要鉴权的violation表接口
// @Tags Violation
// @Summary 不需要鉴权的violation表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /violations/getViolationPublic [get]
func (violationsApi *ViolationApi) GetViolationPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    violationsService.GetViolationPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的violation表接口信息",
    }, "获取成功", c)
}
