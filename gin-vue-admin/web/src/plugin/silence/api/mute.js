import service from '@/utils/request'
// @Tags Mute
// @Summary 创建mute表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Mute true "创建mute表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mute/createMute [post]
export const createMute = (data) => {
  return service({
    url: '/mute/createMute',
    method: 'post',
    data
  })
}

// @Tags Mute
// @Summary 删除mute表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Mute true "删除mute表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mute/deleteMute [delete]
export const deleteMute = (params) => {
  return service({
    url: '/mute/deleteMute',
    method: 'delete',
    params
  })
}

// @Tags Mute
// @Summary 批量删除mute表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mute表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mute/deleteMute [delete]
export const deleteMuteByIds = (params) => {
  return service({
    url: '/mute/deleteMuteByIds',
    method: 'delete',
    params
  })
}

// @Tags Mute
// @Summary 更新mute表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Mute true "更新mute表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mute/updateMute [put]
export const updateMute = (data) => {
  return service({
    url: '/mute/updateMute',
    method: 'put',
    data
  })
}

// @Tags Mute
// @Summary 用id查询mute表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Mute true "用id查询mute表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mute/findMute [get]
export const findMute = (params) => {
  return service({
    url: '/mute/findMute',
    method: 'get',
    params
  })
}

// @Tags Mute
// @Summary 分页获取mute表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mute表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mute/getMuteList [get]
export const getMuteList = (params) => {
  return service({
    url: '/mute/getMuteList',
    method: 'get',
    params
  })
}
// @Tags Mute
// @Summary 不需要鉴权的mute表接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MuteSearch true "分页获取mute表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mute/getMutePublic [get]
export const getMutePublic = () => {
  return service({
    url: '/mute/getMutePublic',
    method: 'get',
  })
}
