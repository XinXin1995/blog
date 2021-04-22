import request from '@/utils/request'
export function GetCategories () {
  return request.get('api/category/statistic')
}
