import request from '@/utils/request'

export function GetAllCategories () {
  return request.get('/api/category/list')
}

export function AddCategory ({ name }) {
  return request.post('/api/manage/category/add', { name })
}

export function UpdateCategory ({ id, name }) {
  return request.post('/api/manage/category/update', { id, name })
}

export function DelCategory ({ id }) {
  return request.get('/api/manage/category/del', { params: { id } })
}
