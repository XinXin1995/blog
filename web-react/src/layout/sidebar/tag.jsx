import { Divider, Tag } from 'antd'
import { useSelector, useDispatch } from 'react-redux'
import { getTagList } from '@/redux/article/action'
import useMount from '@/hooks/useMount'

function Tags () {
  const dispatch = useDispatch()
  const tags = useSelector(state => state.article.tagList)
  useMount(() => {
    dispatch(getTagList())
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
    </div>
  )
}

export default Tags
