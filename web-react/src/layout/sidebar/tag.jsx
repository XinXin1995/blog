import { Divider, Tag } from 'antd'
import { useSelector, useDispatch } from 'react-redux'
import { getCategoryList, getTagList } from '@/redux/article/action'
import useMount from '@/hooks/useMount'

function Tags () {
  const dispatch = useDispatch()
  const tags = useSelector(state => state.article.tagList)
  const categories = useSelector(state => state.article.categoryList)
  useMount(() => {
    dispatch(getTagList())
    dispatch(getCategoryList())
  })
  return (
    <div className={'tags'}>
      <Divider orientation={'left'}>
        标签
      </Divider>
      <div className={'tags-wrap'}>
        {
          tags.map((tag, i) => (
            <Tag key={i} color={tag.color}>{tag.name}</Tag>
          ))
        }
      </div>
      <div className={'tags-wrap'}>
        {
          categories.map((tag, i) => (
            <Tag key={i} >{tag.name}</Tag>
          ))
        }
      </div>
    </div>
  )
}

export default Tags
