import request from '@/utils/request'

export function GetCommentList({articleId, pageNo = 1, pageSize = 10}) {
    return request.post('/api/comment/list', {articleId, pageNo, pageSize})
}

export function AddComment({articleId, parentId, content}) {
    return request.post('/api/manager/comment/add',{articleId, parentId, content})
}