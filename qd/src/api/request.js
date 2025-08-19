/**
 * HTTP请求工具类
 * 基于fetch API封装，用于与后端接口通信
 */

import { API_CONFIG, HTTP_STATUS, BUSINESS_CODE } from './config.js'
import { isUserLoggedIn, clearUserAuth } from '../utils/userUtils.js'

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

    // 添加认证头，确保用户ID能够被后端正确识别
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

    try {
      if (contentType && contentType.includes('application/json')) {
        data = await response.json()
      } else {
        data = await response.text()
      }
    } catch (parseError) {
      console.error('Response parsing error:', parseError)
      throw new Error('响应数据解析失败')
    }

    // 处理 HTTP 错误状态码
    if (!response.ok) {
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
    
    // 如果是认证错误，清除无效的token
    if (error.message.includes('未授权') || error.message.includes('401')) {
      console.warn('Authentication failed, clearing tokens')
      clearUserAuth()
      // 可以在这里触发跳转到登录页面
      // window.location.href = '/login'
    }
    
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
      
      // 直接返回完整的响应数据，让上层组件处理业务逻辑
      // 不要在这里根据 code 判断成功失败，因为后端可能返回 code !== 200 但仍然是有效响应
      console.log('HTTP请求完成，返回数据:', data)
      return data
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