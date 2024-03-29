import request from '@/utils/request'

export function GetArticles ({ pageNo = 1, pageSize = 10, keyword, category, tags, withContent = false, orderType = 1 }) {
  return request.post('/api/article/list', { pageNo, pageSize, keyword, category, tags, withContent, orderType })
}

export function GetArticleDetail (id) {
  return request.get('/api/article/detail', { params: { id } })
}
