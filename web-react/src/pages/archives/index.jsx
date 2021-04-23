import { useDispatch, useSelector } from 'react-redux'
import { getCategoryList } from '@/redux/article/action'
import { Tag, Badge } from 'antd'
import useMount from '@/hooks/useMount'
import { useHistory } from 'react-router-dom'
import './index.less'

function Archives () {
  const history = useHistory()
  const dispatch = useDispatch()
  const categories = useSelector(state => state.article.categoryList)
  const colors = [
    'magenta',
    'blue',
    'red',
    'volcano',
    'orange',
    'gold',
    'lime',
    'green',
    'cyan',
    'geekblue',
    'purple'
  ]
  useMount(() => {
    dispatch(getCategoryList())
  })
  const handleRedirect = (id) => {
    history.push('/list?pageNo=1&category=' + id)
  }
  return <div className={'archives'}>
    {
      categories.map((v, i) => (
        <Badge key={i} count={v.count}>
          <Tag color={colors[i]} onClick={() => handleRedirect(v.id)}>{v.name}</Tag>
        </Badge>
      ))
    }
  </div>
}

export default Archives
