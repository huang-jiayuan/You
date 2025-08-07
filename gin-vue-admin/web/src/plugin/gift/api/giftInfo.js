import service from '@/utils/request'
// @Tags GiftInfo
// @Summary 创建giftInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftInfo true "创建giftInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /giftInfo/createGiftInfo [post]
export const createGiftInfo = (data) => {
  return service({
    url: '/giftInfo/createGiftInfo',
    method: 'post',
    data
  })
}

// @Tags GiftInfo
// @Summary 删除giftInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftInfo true "删除giftInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /giftInfo/deleteGiftInfo [delete]
export const deleteGiftInfo = (params) => {
  return service({
    url: '/giftInfo/deleteGiftInfo',
    method: 'delete',
    params
  })
}

// @Tags GiftInfo
// @Summary 批量删除giftInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除giftInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /giftInfo/deleteGiftInfo [delete]
export const deleteGiftInfoByIds = (params) => {
  return service({
    url: '/giftInfo/deleteGiftInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags GiftInfo
// @Summary 更新giftInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftInfo true "更新giftInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /giftInfo/updateGiftInfo [put]
export const updateGiftInfo = (data) => {
  return service({
    url: '/giftInfo/updateGiftInfo',
    method: 'put',
    data
  })
}

// @Tags GiftInfo
// @Summary 用id查询giftInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.GiftInfo true "用id查询giftInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /giftInfo/findGiftInfo [get]
export const findGiftInfo = (params) => {
  return service({
    url: '/giftInfo/findGiftInfo',
    method: 'get',
    params
  })
}

// @Tags GiftInfo
// @Summary 分页获取giftInfo表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取giftInfo表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /giftInfo/getGiftInfoList [get]
export const getGiftInfoList = (params) => {
  return service({
    url: '/giftInfo/getGiftInfoList',
    method: 'get',
    params
  })
}
// @Tags GiftInfo
// @Summary 不需要鉴权的giftInfo表接口
// @Accept application/json
// @Produce application/json
// @Param data query request.GiftInfoSearch true "分页获取giftInfo表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /giftInfo/getGiftInfoPublic [get]
export const getGiftInfoPublic = () => {
  return service({
    url: '/giftInfo/getGiftInfoPublic',
    method: 'get',
  })
}
