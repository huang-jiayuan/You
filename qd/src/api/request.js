/**
 * HTTP请求工具类
 * 基于fetch API封装，用于与后端接口通信
 */

import { API_CONFIG, HTTP_STATUS, BUSINESS_CODE } from './config.js'
import ResponseHandler from './responseHandler.js'

class HttpRequest {
  constructor() {
    this.baseURL = API_CONFIG.baseURL
    this.timeout = API_CONFIG.timeout
    this.token = this.getToken()
  }

  // 获取存储的token
  getToken() {
    return localStorage.getItem('access_token') || ''
  }

  // 设置token
  setToken(token) {
    this.token = token
    localStorage.setItem('access_token', token)
  }

  // 清除token
  clearToken() {
    this.token = ''
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
  }

  // 获取默认请求头
  getDefaultHeaders() {
    const headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    }

    if (this.token) {
      headers['Authorization'] = `Bearer ${this.token}`
    }

    return headers
  }

  // 处理请求URL
  getFullURL(url) {
    if (url.startsWith('http')) {
      return url
    }
    return `${this.baseURL}${url}`
  }

  // 处理响应
  async handleResponse(response) {
    const contentType = response.headers.get('content-type')
    let data

    console.log('HTTP响应状态:', {
      status: response.status,
      statusText: response.statusText,
      ok: response.ok,
      contentType: contentType
    })

    try {
      if (contentType && contentType.includes('application/json')) {
        data = await response.json()
      } else {
        data = await response.text()
      }
      
      console.log('解析后的响应数据:', data)
    } catch (parseError) {
      console.error('Response parsing error:', parseError)
      throw new Error('响应数据解析失败')
    }

    // 处理 HTTP 错误状态码
    if (!response.ok) {
      console.error('HTTP错误响应:', {
        status: response.status,
        data: data
      })
      
      // 如果是 JSON 响应且包含错误信息，优先使用业务错误信息
      if (typeof data === 'object' && data.msg) {
        throw new Error(data.msg)
      } else if (typeof data === 'object' && data.message) {
        throw new Error(data.message)
      } else {
        const errorMap = {
          400: '请求参数错误',
          401: '未授权访问，请重新登录',
          403: '禁止访问',
          404: '请求的资源不存在',
          500: '服务器内部错误',
          502: '网关错误',
          503: '服务暂时不可用'
        }
        throw new Error(errorMap[response.status] || `HTTP Error: ${response.status}`)
      }
    }

    return data
  }

  // 处理错误
  handleError(error) {
    console.error('API Request Error:', error)
    
    // 根据错误类型进行处理
    if (error.name === 'AbortError') {
      throw new Error('请求超时，请重试')
    }
    
    if (error.message.includes('Failed to fetch')) {
      throw new Error('网络连接失败，请检查网络')
    }
    
    throw error
  }

  // 通用请求方法
  async request(url, options = {}) {
    const controller = new AbortController()
    const timeoutId = setTimeout(() => controller.abort(), this.timeout)

    try {
      const config = {
        method: 'GET',
        headers: this.getDefaultHeaders(),
        signal: controller.signal,
        ...options
      }

      // 处理请求体
      if (config.body && typeof config.body === 'object' && !(config.body instanceof FormData)) {
        config.body = JSON.stringify(config.body)
      }

      const response = await fetch(this.getFullURL(url), config)
      clearTimeout(timeoutId)
      
      const data = await this.handleResponse(response)
      
      // 使用 ResponseHandler 处理业务逻辑
      try {
        return ResponseHandler.handleApiResponse(data)
      } catch (businessError) {
        // 处理需要重新登录的错误
        if (ResponseHandler.shouldRelogin(businessError.code)) {
          this.clearToken()
          window.location.href = '/'
        }
        throw businessError
      }
    } catch (error) {
      clearTimeout(timeoutId)
      return this.handleError(error)
    }
  }

  // GET请求
  get(url, params = {}) {
    const searchParams = new URLSearchParams()
    Object.keys(params).forEach(key => {
      if (params[key] !== undefined && params[key] !== null) {
        searchParams.append(key, params[key])
      }
    })
    
    const queryString = searchParams.toString()
    const fullURL = queryString ? `${url}?${queryString}` : url
    
    return this.request(fullURL, { method: 'GET' })
  }

  // POST请求
  post(url, data = {}) {
    return this.request(url, {
      method: 'POST',
      body: data
    })
  }

  // PUT请求
  put(url, data = {}) {
    return this.request(url, {
      method: 'PUT',
      body: data
    })
  }

  // DELETE请求
  delete(url) {
    return this.request(url, { method: 'DELETE' })
  }

  // 文件上传
  upload(url, file, onProgress) {
    return new Promise((resolve, reject) => {
      const formData = new FormData()
      formData.append('file', file)

      const xhr = new XMLHttpRequest()
      
      // 上传进度
      if (onProgress) {
        xhr.upload.addEventListener('progress', (event) => {
          if (event.lengthComputable) {
            const progress = Math.round((event.loaded / event.total) * 100)
            onProgress(progress)
          }
        })
      }

      xhr.addEventListener('load', () => {
        if (xhr.status === HTTP_STATUS.SUCCESS) {
          try {
            const response = JSON.parse(xhr.responseText)
            resolve(response)
          } catch (error) {
            reject(new Error('响应解析失败'))
          }
        } else {
          reject(new Error(`上传失败: ${xhr.status}`))
        }
      })

      xhr.addEventListener('error', () => {
        reject(new Error('上传失败'))
      })

      xhr.open('POST', this.getFullURL(url))
      
      // 设置认证头
      if (this.token) {
        xhr.setRequestHeader('Authorization', `Bearer ${this.token}`)
      }
      
      xhr.send(formData)
    })
  }
}

// 创建请求实例
const http = new HttpRequest()

export default http