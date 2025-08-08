import service from '@/utils/request'
// @Tags GiftSendRecord
// @Summary 创建giftSendRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftSendRecord true "创建giftSendRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /giftSendRecord/createGiftSendRecord [post]
export const createGiftSendRecord = (data) => {
  return service({
    url: '/giftSendRecord/createGiftSendRecord',
    method: 'post',
    data
  })
}

// @Tags GiftSendRecord
// @Summary 删除giftSendRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftSendRecord true "删除giftSendRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /giftSendRecord/deleteGiftSendRecord [delete]
export const deleteGiftSendRecord = (params) => {
  return service({
    url: '/giftSendRecord/deleteGiftSendRecord',
    method: 'delete',
    params
  })
}

// @Tags GiftSendRecord
// @Summary 批量删除giftSendRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除giftSendRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /giftSendRecord/deleteGiftSendRecord [delete]
export const deleteGiftSendRecordByIds = (params) => {
  return service({
    url: '/giftSendRecord/deleteGiftSendRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags GiftSendRecord
// @Summary 更新giftSendRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GiftSendRecord true "更新giftSendRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /giftSendRecord/updateGiftSendRecord [put]
export const updateGiftSendRecord = (data) => {
  return service({
    url: '/giftSendRecord/updateGiftSendRecord',
    method: 'put',
    data
  })
}

// @Tags GiftSendRecord
// @Summary 用id查询giftSendRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.GiftSendRecord true "用id查询giftSendRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /giftSendRecord/findGiftSendRecord [get]
export const findGiftSendRecord = (params) => {
  return service({
    url: '/giftSendRecord/findGiftSendRecord',
    method: 'get',
    params
  })
}

// @Tags GiftSendRecord
// @Summary 分页获取giftSendRecord表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取giftSendRecord表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /giftSendRecord/getGiftSendRecordList [get]
export const getGiftSendRecordList = (params) => {
  return service({
    url: '/giftSendRecord/getGiftSendRecordList',
    method: 'get',
    params
  })
}
// @Tags GiftSendRecord
// @Summary 不需要鉴权的giftSendRecord表接口
// @Accept application/json
// @Produce application/json
// @Param data query request.GiftSendRecordSearch true "分页获取giftSendRecord表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /giftSendRecord/getGiftSendRecordPublic [get]
export const getGiftSendRecordPublic = () => {
  return service({
    url: '/giftSendRecord/getGiftSendRecordPublic',
    method: 'get',
  })
}
