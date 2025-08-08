import service from '@/utils/request'
// @Tags HostRoom
// @Summary 创建hostRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HostRoom true "创建hostRoom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hostRoom/createHostRoom [post]
export const createHostRoom = (data) => {
  return service({
    url: '/hostRoom/createHostRoom',
    method: 'post',
    data
  })
}

// @Tags HostRoom
// @Summary 删除hostRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HostRoom true "删除hostRoom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hostRoom/deleteHostRoom [delete]
export const deleteHostRoom = (params) => {
  return service({
    url: '/hostRoom/deleteHostRoom',
    method: 'delete',
    params
  })
}

// @Tags HostRoom
// @Summary 批量删除hostRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除hostRoom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hostRoom/deleteHostRoom [delete]
export const deleteHostRoomByIds = (params) => {
  return service({
    url: '/hostRoom/deleteHostRoomByIds',
    method: 'delete',
    params
  })
}

// @Tags HostRoom
// @Summary 更新hostRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.HostRoom true "更新hostRoom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hostRoom/updateHostRoom [put]
export const updateHostRoom = (data) => {
  return service({
    url: '/hostRoom/updateHostRoom',
    method: 'put',
    data
  })
}

// @Tags HostRoom
// @Summary 用id查询hostRoom表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.HostRoom true "用id查询hostRoom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hostRoom/findHostRoom [get]
export const findHostRoom = (params) => {
  return service({
    url: '/hostRoom/findHostRoom',
    method: 'get',
    params
  })
}

// @Tags HostRoom
// @Summary 分页获取hostRoom表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取hostRoom表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hostRoom/getHostRoomList [get]
export const getHostRoomList = (params) => {
  return service({
    url: '/hostRoom/getHostRoomList',
    method: 'get',
    params
  })
}

// @Tags HostRoom
// @Summary 不需要鉴权的hostRoom表接口
// @Accept application/json
// @Produce application/json
// @Param data query hot_roomReq.HostRoomSearch true "分页获取hostRoom表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /hostRoom/getHostRoomPublic [get]
export const getHostRoomPublic = () => {
  return service({
    url: '/hostRoom/getHostRoomPublic',
    method: 'get',
  })
}
