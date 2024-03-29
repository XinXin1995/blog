import request from '@/utils/request'

export function AddArticle ({ title, category, tags, content, thumb }) {
  return request.post('/api/manager/article/add', {
    title,
    category,
    tags,
    content,
    thumb
  })
}

export function UpdateArticle ({ id, title, category, tags, content, thumb }) {
  return request.put('/api/manager/article/update', {
    id,
    title,
    category,
    tags,
    content,
    thumb
  })
}

export function GetArticle ({ id }) {
  return request.get('/api/article/detail', { params: { id } })
}

export function GetArticleUnion ({ pageNo = 0, pageSize = 10, keyword = '', category = 0, tags = [], orderType = 1 }) {
  return request.post('/api/article/list', {
    pageNo,
    pageSize,
    keyword,
    category,
    orderType,
    tags
  })
}

export function DelArticle ({ id }) {
  return request.delete('/api/manager/article/delete', { params: { id } })
}

export function TopArticle ({ id }) {
  return request.get('/api/article/top', { params: { id } })
}

export function like (id) {
  return request.get('/api/article/like', { params: { id } })
}
