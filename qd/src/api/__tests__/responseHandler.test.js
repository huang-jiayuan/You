/**
 * ResponseHandler 单元测试
 */

import { describe, test, expect } from 'vitest'
import ResponseHandler from '../responseHandler.js'
import { BUSINESS_CODE } from '../config.js'

describe('ResponseHandler', () => {
  test('should handle successful response', () => {
    const mockResponse = {
      code: 200,
      msg: 'success',
      data: { id: 1, name: 'test' }
    }
    
    const result = ResponseHandler.handleApiResponse(mockResponse)
    expect(result).toEqual({ id: 1, name: 'test' })
  })

  test('should handle successful response without data field', () => {
    const mockResponse = {
      code: 200,
      msg: 'success'
    }
    
    const result = ResponseHandler.handleApiResponse(mockResponse)
    expect(result).toEqual(mockResponse)
  })

  test('should throw error for business error response', () => {
    const mockResponse = {
      code: 1004,
      msg: '密码错误',
      data: null
    }
    
    expect(() => {
      ResponseHandler.handleApiResponse(mockResponse)
    }).toThrow('密码错误')
  })

  test('should handle non-standard response format', () => {
    const mockResponse = { message: 'direct response' }
    
    const result = ResponseHandler.handleApiResponse(mockResponse)
    expect(result).toEqual(mockResponse)
  })

  test('should identify token expired error', () => {
    expect(ResponseHandler.shouldRelogin(BUSINESS_CODE.TOKEN_EXPIRED)).toBe(true)
    expect(ResponseHandler.shouldRelogin(BUSINESS_CODE.UNAUTHORIZED)).toBe(true)
    expect(ResponseHandler.shouldRelogin(BUSINESS_CODE.PASSWORD_ERROR)).toBe(false)
  })

  test('should format error messages correctly', () => {
    const networkError = new Error('Failed to fetch')
    expect(ResponseHandler.formatErrorMessage(networkError)).toBe('网络连接失败，请检查网络连接')

    const timeoutError = new Error('timeout')
    expect(ResponseHandler.formatErrorMessage(timeoutError)).toBe('请求超时，请重试')

    const businessError = new Error('密码错误')
    businessError.code = 1004
    expect(ResponseHandler.formatErrorMessage(businessError)).toBe('密码错误')
  })
})