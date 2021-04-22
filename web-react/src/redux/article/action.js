import * as TYPES from '@/redux/types'
import { GetTags } from '@/api/tags'
import { GetCategories } from '@/api/categories'

export const getTagList = () => dispatch => {
  GetTags().then(res => {
    if (res.code === 0) {
      dispatch({
        type: TYPES.ARTICLE_GET_TAG_LIST,
        payload: res.data || []
      })
    }
  })
}

export const getCategoryList = () => dispatch => {
  GetCategories().then(res => {
    if (res.code === 0) {
      dispatch({
        type: TYPES.ARTICLE_GET_CATEGORY_LIST,
        payload: res.data || []
      })
    }
  })
}

