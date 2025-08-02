import service from '@/utils/request'
// @Tags Room
// @Summary 创建房间表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Room true "创建房间表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /room/createRoom [post]
export const createRoom = (data) => {
  return service({
    url: '/room/createRoom',
    method: 'post',
    data
  })
}

// @Tags Room
// @Summary 删除房间表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Room true "删除房间表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /room/deleteRoom [delete]
export const deleteRoom = (params) => {
  return service({
    url: '/room/deleteRoom',
    method: 'delete',
    params
  })
}

// @Tags Room
// @Summary 批量删除房间表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除房间表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /room/deleteRoom [delete]
export const deleteRoomByIds = (params) => {
  return service({
    url: '/room/deleteRoomByIds',
    method: 'delete',
    params
  })
}

// @Tags Room
// @Summary 更新房间表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Room true "更新房间表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /room/updateRoom [put]
export const updateRoom = (data) => {
  return service({
    url: '/room/updateRoom',
    method: 'put',
    data
  })
}

// @Tags Room
// @Summary 用id查询房间表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Room true "用id查询房间表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /room/findRoom [get]
export const findRoom = (params) => {
  return service({
    url: '/room/findRoom',
    method: 'get',
    params
  })
}

// @Tags Room
// @Summary 分页获取房间表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取房间表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /room/getRoomList [get]
export const getRoomList = (params) => {
  return service({
    url: '/room/getRoomList',
    method: 'get',
    params
  })
}
// @Tags Room
// @Summary 不需要鉴权的房间表接口
// @Accept application/json
// @Produce application/json
// @Param data query request.RoomSearch true "分页获取房间表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /room/getRoomPublic [get]
export const getRoomPublic = () => {
  return service({
    url: '/room/getRoomPublic',
    method: 'get',
  })
}
