import request from '@/utils/request'

export function AddComment ({ articleId, parentId, content }) {
  return request.post('/api/manager/comment/add', { articleId, parentId, content })
}

export function GetCommentsTree ({ articleId, pageNo, pageSize }) {
  return request.post('/api/comment/list', { articleId, pageNo, pageSize })
}
