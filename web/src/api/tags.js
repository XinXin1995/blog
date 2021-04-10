import request from '@/utils/request'

export function GetAllTags () {
  return request.get('/api/tag/list')
}

export function DelTag ({ id }) {
  return request.delete('/api/manager/tag/delete', { params: { id } })
}

export function AddTag ({ name, color }) {
  return request.post('/api/manager/tag/create', { name, color })
}

export function UpdateTag ({ id, name, color }) {
  return request.put('/api/manager/tag/update', { id, name, color })
}
