import { GetArticles } from '@/api/articles'
import useFetchList from '@/hooks/useFetchList'
import Pagination from '@/components/pagination'
import ArticleItem from './articleItem'
import { useLocation } from 'react-router-dom'
import './index.less'
import { translateMarkdown } from '@/utils'
import { Skeleton } from 'antd'

const List = () => {
  const location = useLocation()
  const { list, loading, pagination } = useFetchList({
    requestMethod: GetArticles,
    queryParams: { pageSize: 5, category: 0, orderType: 3, withContent: true },
    fetchDependence: [location.pathname, location.search]
  })
  return (
    <div className={'article-list'}>
      {
        loading ?
          <Skeleton active/>
          :
          list.map(v => {
            v.content = translateMarkdown(v.content)
            return (
              <ArticleItem key={v.id} item={v}/>
            )
          })
      }
      <Pagination {...pagination} />
    </div>
  )
}

export default List
