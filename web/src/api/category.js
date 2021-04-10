import request from '@/utils/request'

export function GetAllCategories () {
  return request.get('/api/category/list')
}

export function AddCategory ({ name }) {
  return request.post('/api/manager/category/create', { name })
}

export function UpdateCategory ({ id, name }) {
  return request.put('/api/manager/category/update', { id, name })
}

export function DelCategory ({ id }) {
  return request.delete('/api/manager/category/delete', { params: { id } })
}
