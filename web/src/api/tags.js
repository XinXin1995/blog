import request from '@/utils/request'

export function GetAllTags () {
  return request.get('/api/tag/list')
}

export function DelTag ({ id }) {
  return request.get('/api/manage/tag/del', { params: { id } })
}

export function AddTag ({ name, color }) {
  return request.post('/api/manage/tag/add', { name, color })
}

export function UpdateTag ({ id, name, color }) {
  return request.post('/api/manage/tag/update', { id, name, color })
}
