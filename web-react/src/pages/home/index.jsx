import { useState } from 'react'
import ArticleItem from './articleItem'
import { translateMarkdown } from '@/utils'
import { GetArticles } from '@/api/articles'
import useMount from '@/hooks/useMount'
import { Skeleton } from 'antd'
import './index.less'
import '@/assets/style/article.less'

function Home () {
  const [list, setList] = useState([])
  const [loading, setLoading] = useState(false)
  useMount(() => {
    setLoading(true)
    GetArticles({ widthContent: true, pageSize:5, order: 2 }).then(res => {
      if (res.code === 0) {
        setList(res.data.list || [])
      }
      setLoading(false)
    }).catch(() => {
      setLoading(false)
    })
  })
  return (
    <div className={'home'}>
      {
        loading
          ?
          <Skeleton active />
          :
          list.map(v => {
            v.content = translateMarkdown(v.content)
            return (
              <div key={v.id}>
                <ArticleItem unique={v.id} item={v}/>
              </div>
            )
          })
      }
    </div>
  )
}

export default Home
