package carouse

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/carouse"
    carouseReq "github.com/flipped-aurora/gin-vue-admin/server/model/carouse/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CarouselImageApi struct {}



// CreateCarouselImage 创建carouselImage表
// @Tags CarouselImage
// @Summary 创建carouselImage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body carouse.CarouselImage true "创建carouselImage表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /carouselImage/createCarouselImage [post]
func (carouselImageApi *CarouselImageApi) CreateCarouselImage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var carouselImage carouse.CarouselImage
	err := c.ShouldBindJSON(&carouselImage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    carouselImage.CreatedBy = utils.GetUserID(c)
	err = carouselImageService.CreateCarouselImage(ctx,&carouselImage)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteCarouselImage 删除carouselImage表
// @Tags CarouselImage
// @Summary 删除carouselImage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body carouse.CarouselImage true "删除carouselImage表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /carouselImage/deleteCarouselImage [delete]
func (carouselImageApi *CarouselImageApi) DeleteCarouselImage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := carouselImageService.DeleteCarouselImage(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCarouselImageByIds 批量删除carouselImage表
// @Tags CarouselImage
// @Summary 批量删除carouselImage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /carouselImage/deleteCarouselImageByIds [delete]
func (carouselImageApi *CarouselImageApi) DeleteCarouselImageByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := carouselImageService.DeleteCarouselImageByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCarouselImage 更新carouselImage表
// @Tags CarouselImage
// @Summary 更新carouselImage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body carouse.CarouselImage true "更新carouselImage表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /carouselImage/updateCarouselImage [put]
func (carouselImageApi *CarouselImageApi) UpdateCarouselImage(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var carouselImage carouse.CarouselImage
	err := c.ShouldBindJSON(&carouselImage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    carouselImage.UpdatedBy = utils.GetUserID(c)
	err = carouselImageService.UpdateCarouselImage(ctx,carouselImage)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCarouselImage 用id查询carouselImage表
// @Tags CarouselImage
// @Summary 用id查询carouselImage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询carouselImage表"
// @Success 200 {object} response.Response{data=carouse.CarouselImage,msg=string} "查询成功"
// @Router /carouselImage/findCarouselImage [get]
func (carouselImageApi *CarouselImageApi) FindCarouselImage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	recarouselImage, err := carouselImageService.GetCarouselImage(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(recarouselImage, c)
}
// GetCarouselImageList 分页获取carouselImage表列表
// @Tags CarouselImage
// @Summary 分页获取carouselImage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query carouseReq.CarouselImageSearch true "分页获取carouselImage表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /carouselImage/getCarouselImageList [get]
func (carouselImageApi *CarouselImageApi) GetCarouselImageList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo carouseReq.CarouselImageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := carouselImageService.GetCarouselImageInfoList(ctx,pageInfo)
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

// GetCarouselImagePublic 不需要鉴权的carouselImage表接口
// @Tags CarouselImage
// @Summary 不需要鉴权的carouselImage表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /carouselImage/getCarouselImagePublic [get]
func (carouselImageApi *CarouselImageApi) GetCarouselImagePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    carouselImageService.GetCarouselImagePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的carouselImage表接口信息",
    }, "获取成功", c)
}
