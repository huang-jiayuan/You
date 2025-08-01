import service from '@/utils/request'
// @Tags RoomTagDict
// @Summary 创建房间话题表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RoomTagDict true "创建房间话题表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /roomTagDict/createRoomTagDict [post]
export const createRoomTagDict = (data) => {
  return service({
    url: '/roomTagDict/createRoomTagDict',
    method: 'post',
    data
  })
}

// @Tags RoomTagDict
// @Summary 删除房间话题表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RoomTagDict true "删除房间话题表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /roomTagDict/deleteRoomTagDict [delete]
export const deleteRoomTagDict = (params) => {
  return service({
    url: '/roomTagDict/deleteRoomTagDict',
    method: 'delete',
    params
  })
}

// @Tags RoomTagDict
// @Summary 批量删除房间话题表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除房间话题表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /roomTagDict/deleteRoomTagDict [delete]
export const deleteRoomTagDictByIds = (params) => {
  return service({
    url: '/roomTagDict/deleteRoomTagDictByIds',
    method: 'delete',
    params
  })
}

// @Tags RoomTagDict
// @Summary 更新房间话题表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RoomTagDict true "更新房间话题表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /roomTagDict/updateRoomTagDict [put]
export const updateRoomTagDict = (data) => {
  return service({
    url: '/roomTagDict/updateRoomTagDict',
    method: 'put',
    data
  })
}

// @Tags RoomTagDict
// @Summary 用id查询房间话题表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.RoomTagDict true "用id查询房间话题表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /roomTagDict/findRoomTagDict [get]
export const findRoomTagDict = (params) => {
  return service({
    url: '/roomTagDict/findRoomTagDict',
    method: 'get',
    params
  })
}

// @Tags RoomTagDict
// @Summary 分页获取房间话题表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取房间话题表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /roomTagDict/getRoomTagDictList [get]
export const getRoomTagDictList = (params) => {
  return service({
    url: '/roomTagDict/getRoomTagDictList',
    method: 'get',
    params
  })
}
// @Tags RoomTagDict
// @Summary 不需要鉴权的房间话题表接口
// @Accept application/json
// @Produce application/json
// @Param data query request.RoomTagDictSearch true "分页获取房间话题表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /roomTagDict/getRoomTagDictPublic [get]
export const getRoomTagDictPublic = () => {
  return service({
    url: '/roomTagDict/getRoomTagDictPublic',
    method: 'get',
  })
}
