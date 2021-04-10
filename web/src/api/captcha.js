import request from '@/utils/request'

export function Captcha () {
  return request.get('/api/base/captcha')
}
