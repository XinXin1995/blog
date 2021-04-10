import request from '@/utils/request'

export function DeleteImg (id) {
  return request.delete('/api/file/delete', { params: { id } })
}
