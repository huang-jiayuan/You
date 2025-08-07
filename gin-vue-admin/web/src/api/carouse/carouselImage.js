import service from '@/utils/request'
// @Tags CarouselImage
// @Summary 创建carouselImage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CarouselImage true "创建carouselImage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /carouselImage/createCarouselImage [post]
export const createCarouselImage = (data) => {
  return service({
    url: '/carouselImage/createCarouselImage',
    method: 'post',
    data
  })
}

// @Tags CarouselImage
// @Summary 删除carouselImage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CarouselImage true "删除carouselImage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carouselImage/deleteCarouselImage [delete]
export const deleteCarouselImage = (params) => {
  return service({
    url: '/carouselImage/deleteCarouselImage',
    method: 'delete',
    params
  })
}

// @Tags CarouselImage
// @Summary 批量删除carouselImage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除carouselImage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carouselImage/deleteCarouselImage [delete]
export const deleteCarouselImageByIds = (params) => {
  return service({
    url: '/carouselImage/deleteCarouselImageByIds',
    method: 'delete',
    params
  })
}

// @Tags CarouselImage
// @Summary 更新carouselImage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CarouselImage true "更新carouselImage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /carouselImage/updateCarouselImage [put]
export const updateCarouselImage = (data) => {
  return service({
    url: '/carouselImage/updateCarouselImage',
    method: 'put',
    data
  })
}

// @Tags CarouselImage
// @Summary 用id查询carouselImage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.CarouselImage true "用id查询carouselImage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /carouselImage/findCarouselImage [get]
export const findCarouselImage = (params) => {
  return service({
    url: '/carouselImage/findCarouselImage',
    method: 'get',
    params
  })
}

// @Tags CarouselImage
// @Summary 分页获取carouselImage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取carouselImage表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carouselImage/getCarouselImageList [get]
export const getCarouselImageList = (params) => {
  return service({
    url: '/carouselImage/getCarouselImageList',
    method: 'get',
    params
  })
}

// @Tags CarouselImage
// @Summary 不需要鉴权的carouselImage表接口
// @Accept application/json
// @Produce application/json
// @Param data query carouseReq.CarouselImageSearch true "分页获取carouselImage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /carouselImage/getCarouselImagePublic [get]
export const getCarouselImagePublic = () => {
  return service({
    url: '/carouselImage/getCarouselImagePublic',
    method: 'get',
  })
}
