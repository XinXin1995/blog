import request from '@/utils/request'

export function Login ({ username, password, captcha, captchaId }) {
  return request.post('/api/user/login', { username, password, captcha, captchaId })
}

export function GetUser () {
  return request.get('/api/manager/user/detail')
}

export function AddUser ({ username, password, avatar }) {
  return request.post('/api/user/register', { username, password, avatar })
}
