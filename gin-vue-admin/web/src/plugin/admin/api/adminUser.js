import service from '@/utils/request'
// @Tags AdminUser
// @Summary 创建管理员表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AdminUser true "创建管理员表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /adminUser/createAdminUser [post]
export const createAdminUser = (data) => {
  return service({
    url: '/adminUser/createAdminUser',
    method: 'post',
    data
  })
}

// @Tags AdminUser
// @Summary 删除管理员表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AdminUser true "删除管理员表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adminUser/deleteAdminUser [delete]
export const deleteAdminUser = (params) => {
  return service({
    url: '/adminUser/deleteAdminUser',
    method: 'delete',
    params
  })
}

// @Tags AdminUser
// @Summary 批量删除管理员表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除管理员表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adminUser/deleteAdminUser [delete]
export const deleteAdminUserByIds = (params) => {
  return service({
    url: '/adminUser/deleteAdminUserByIds',
    method: 'delete',
    params
  })
}

// @Tags AdminUser
// @Summary 更新管理员表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AdminUser true "更新管理员表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adminUser/updateAdminUser [put]
export const updateAdminUser = (data) => {
  return service({
    url: '/adminUser/updateAdminUser',
    method: 'put',
    data
  })
}

// @Tags AdminUser
// @Summary 用id查询管理员表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AdminUser true "用id查询管理员表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adminUser/findAdminUser [get]
export const findAdminUser = (params) => {
  return service({
    url: '/adminUser/findAdminUser',
    method: 'get',
    params
  })
}

// @Tags AdminUser
// @Summary 分页获取管理员表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取管理员表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adminUser/getAdminUserList [get]
export const getAdminUserList = (params) => {
  return service({
    url: '/adminUser/getAdminUserList',
    method: 'get',
    params
  })
}
// @Tags AdminUser
// @Summary 不需要鉴权的管理员表接口
// @Accept application/json
// @Produce application/json
// @Param data query request.AdminUserSearch true "分页获取管理员表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /adminUser/getAdminUserPublic [get]
export const getAdminUserPublic = () => {
  return service({
    url: '/adminUser/getAdminUserPublic',
    method: 'get',
  })
}
