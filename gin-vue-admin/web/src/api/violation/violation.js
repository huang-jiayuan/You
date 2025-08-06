import service from '@/utils/request'
// @Tags Violation
// @Summary 创建violation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Violation true "创建violation表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /violations/createViolation [post]
export const createViolation = (data) => {
  return service({
    url: '/violations/createViolation',
    method: 'post',
    data
  })
}

// @Tags Violation
// @Summary 删除violation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Violation true "删除violation表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /violations/deleteViolation [delete]
export const deleteViolation = (params) => {
  return service({
    url: '/violations/deleteViolation',
    method: 'delete',
    params
  })
}

// @Tags Violation
// @Summary 批量删除violation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除violation表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /violations/deleteViolation [delete]
export const deleteViolationByIds = (params) => {
  return service({
    url: '/violations/deleteViolationByIds',
    method: 'delete',
    params
  })
}

// @Tags Violation
// @Summary 更新violation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Violation true "更新violation表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /violations/updateViolation [put]
export const updateViolation = (data) => {
  return service({
    url: '/violations/updateViolation',
    method: 'put',
    data
  })
}

// @Tags Violation
// @Summary 用id查询violation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Violation true "用id查询violation表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /violations/findViolation [get]
export const findViolation = (params) => {
  return service({
    url: '/violations/findViolation',
    method: 'get',
    params
  })
}

// @Tags Violation
// @Summary 分页获取violation表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取violation表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /violations/getViolationList [get]
export const getViolationList = (params) => {
  return service({
    url: '/violations/getViolationList',
    method: 'get',
    params
  })
}

// @Tags Violation
// @Summary 不需要鉴权的violation表接口
// @Accept application/json
// @Produce application/json
// @Param data query violationReq.ViolationSearch true "分页获取violation表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /violations/getViolationPublic [get]
export const getViolationPublic = () => {
  return service({
    url: '/violations/getViolationPublic',
    method: 'get',
  })
}
