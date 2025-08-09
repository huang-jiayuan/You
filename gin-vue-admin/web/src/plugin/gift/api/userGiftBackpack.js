import service from '@/utils/request'
// @Tags UserGiftBackpack
// @Summary 创建用户背包礼物表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserGiftBackpack true "创建用户背包礼物表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userGiftBackpack/createUserGiftBackpack [post]
export const createUserGiftBackpack = (data) => {
  return service({
    url: '/userGiftBackpack/createUserGiftBackpack',
    method: 'post',
    data
  })
}

// @Tags UserGiftBackpack
// @Summary 删除用户背包礼物表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserGiftBackpack true "删除用户背包礼物表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userGiftBackpack/deleteUserGiftBackpack [delete]
export const deleteUserGiftBackpack = (params) => {
  return service({
    url: '/userGiftBackpack/deleteUserGiftBackpack',
    method: 'delete',
    params
  })
}

// @Tags UserGiftBackpack
// @Summary 批量删除用户背包礼物表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户背包礼物表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userGiftBackpack/deleteUserGiftBackpack [delete]
export const deleteUserGiftBackpackByIds = (params) => {
  return service({
    url: '/userGiftBackpack/deleteUserGiftBackpackByIds',
    method: 'delete',
    params
  })
}

// @Tags UserGiftBackpack
// @Summary 更新用户背包礼物表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserGiftBackpack true "更新用户背包礼物表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userGiftBackpack/updateUserGiftBackpack [put]
export const updateUserGiftBackpack = (data) => {
  return service({
    url: '/userGiftBackpack/updateUserGiftBackpack',
    method: 'put',
    data
  })
}

// @Tags UserGiftBackpack
// @Summary 用id查询用户背包礼物表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserGiftBackpack true "用id查询用户背包礼物表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userGiftBackpack/findUserGiftBackpack [get]
export const findUserGiftBackpack = (params) => {
  return service({
    url: '/userGiftBackpack/findUserGiftBackpack',
    method: 'get',
    params
  })
}

// @Tags UserGiftBackpack
// @Summary 分页获取用户背包礼物表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户背包礼物表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userGiftBackpack/getUserGiftBackpackList [get]
export const getUserGiftBackpackList = (params) => {
  return service({
    url: '/userGiftBackpack/getUserGiftBackpackList',
    method: 'get',
    params
  })
}
// @Tags UserGiftBackpack
// @Summary 不需要鉴权的用户背包礼物表接口
// @Accept application/json
// @Produce application/json
// @Param data query request.UserGiftBackpackSearch true "分页获取用户背包礼物表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userGiftBackpack/getUserGiftBackpackPublic [get]
export const getUserGiftBackpackPublic = () => {
  return service({
    url: '/userGiftBackpack/getUserGiftBackpackPublic',
    method: 'get',
  })
}
