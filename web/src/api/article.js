import request from '@/utils/request'

export function AddArticle ({ title, category, tags, content, thumb }) {
  return request.post('/api/manage/article/add', {
    title,
    category,
    tags,
    content,
    thumb
  })
}
export function UpdateArticle ({ id, title, category, tags, content, thumb }) {
  return request.post('/api/manage/article/update', {
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

export function GetArticleList ({ pageNo = 0, pageSize = 10, keyword = '', category = 0, tags = [] }) {
  return request.get('/api/article/list', {
    params: {
      pageNo,
      pageSize,
      keyword,
      category,
      tags
    }
  })
}
export function GetArticleUnion ({ pageNo = 0, pageSize = 10, keyword = '', category = 0, tags = [] }) {
  return request.get('/api/article/union', {
    params: {
      pageNo,
      pageSize,
      keyword,
      category,
      tags
    }
  })
}

export function DelArticle ({ id }) {
  return request.get('/api/manage/article/del', { params: { id } })
}

export function like (id) {
  return request.get('/api/article/like', { params: { id } })
}
