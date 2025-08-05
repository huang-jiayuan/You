import service from '@/utils/request'
// @Tags HotRoom
// @Summary 创建hotRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HotRoom true "创建hotRoom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hotRoom/createHotRoom [post]
export const createHotRoom = (data) => {
  return service({
    url: '/hotRoom/createHotRoom',
    method: 'post',
    data
  })
}

// @Tags HotRoom
// @Summary 删除hotRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HotRoom true "删除hotRoom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hotRoom/deleteHotRoom [delete]
export const deleteHotRoom = (params) => {
  return service({
    url: '/hotRoom/deleteHotRoom',
    method: 'delete',
    params
  })
}

// @Tags HotRoom
// @Summary 批量删除hotRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hotRoom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hotRoom/deleteHotRoom [delete]
export const deleteHotRoomByIds = (params) => {
  return service({
    url: '/hotRoom/deleteHotRoomByIds',
    method: 'delete',
    params
  })
}

// @Tags HotRoom
// @Summary 更新hotRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HotRoom true "更新hotRoom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hotRoom/updateHotRoom [put]
export const updateHotRoom = (data) => {
  return service({
    url: '/hotRoom/updateHotRoom',
    method: 'put',
    data
  })
}

// @Tags HotRoom
// @Summary 用id查询hotRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.HotRoom true "用id查询hotRoom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hotRoom/findHotRoom [get]
export const findHotRoom = (params) => {
  return service({
    url: '/hotRoom/findHotRoom',
    method: 'get',
    params
  })
}

// @Tags HotRoom
// @Summary 分页获取hotRoom表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hotRoom表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hotRoom/getHotRoomList [get]
export const getHotRoomList = (params) => {
  return service({
    url: '/hotRoom/getHotRoomList',
    method: 'get',
    params
  })
}

// @Tags HotRoom
// @Summary 不需要鉴权的hotRoom表接口
// @Accept application/json
// @Produce application/json
// @Param data query hot_roomReq.HotRoomSearch true "分页获取hotRoom表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /hotRoom/getHotRoomPublic [get]
export const getHotRoomPublic = () => {
  return service({
    url: '/hotRoom/getHotRoomPublic',
    method: 'get',
  })
}
