import * as TYPES from '@/redux/types'

const defaultState = {
  tagList: [],
  categoryList: []
}

export default function articleReducer (state = defaultState, action) {
  const { type, payload } = action
  switch (type) {
    case TYPES.ARTICLE_GET_TAG_LIST:
      return { ...state, tagList: payload }
    case TYPES.ARTICLE_GET_CATEGORY_LIST:
      return { ...state, categoryList: payload }
    default:
      return state
  }
}
