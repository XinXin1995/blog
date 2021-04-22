import request from '@/utils/request'

export function GetTags () {
  return request.get('/api/tag/list')
}
