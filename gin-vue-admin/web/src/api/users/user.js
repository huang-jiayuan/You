import service from '@/utils/request'
// @Tags User
// @Summary 创建user表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.User true "创建user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /user1/createUser [post]
export const createUser = (data) => {
  return service({
    url: '/user1/createUser',
    method: 'post',
    data
  })
}

// @Tags User
// @Summary 删除user表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.User true "删除user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user1/deleteUser [delete]
export const deleteUser = (params) => {
  return service({
    url: '/user1/deleteUser',
    method: 'delete',
    params
  })
}

// @Tags User
// @Summary 批量删除user表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user1/deleteUser [delete]
export const deleteUserByIds = (params) => {
  return service({
    url: '/user1/deleteUserByIds',
    method: 'delete',
    params
  })
}

// @Tags User
// @Summary 更新user表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.User true "更新user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user1/updateUser [put]
export const updateUser = (data) => {
  return service({
    url: '/user1/updateUser',
    method: 'put',
    data
  })
}

// @Tags User
// @Summary 用id查询user表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.User true "用id查询user表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user1/findUser [get]
export const findUser = (params) => {
  return service({
    url: '/user1/findUser',
    method: 'get',
    params
  })
}

// @Tags User
// @Summary 分页获取user表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取user表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user1/getUserList [get]
export const getUserList = (params) => {
  return service({
    url: '/user1/getUserList',
    method: 'get',
    params
  })
}

// @Tags User
// @Summary 不需要鉴权的user表接口
// @Accept application/json
// @Produce application/json
// @Param data query usersReq.UserSearch true "分页获取user表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /user1/getUserPublic [get]
export const getUserPublic = () => {
  return service({
    url: '/user1/getUserPublic',
    method: 'get',
  })
}
