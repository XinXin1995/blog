import request from '@/utils/request'

export function Login ({ username, password }) {
  return request.post('/api/user/login', { username, password })
}

export function GetUser () {
  return request.get('/api/manage/user/detail')
}

export function AddUser ({ username, password, avatar }) {
  return request.post('/api/user/register', { username, password, avatar })
}
