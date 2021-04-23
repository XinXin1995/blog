import '@/assets/style/article.less'
import './index.less'
import { EyeOutlined, LikeOutlined, PartitionOutlined, TagsOutlined } from '@ant-design/icons'
import { Divider, Tag } from 'antd'
import { getScrollTop, translateMarkdown } from '@/utils'
import Navigation from './navigation'
import { GetArticleDetail } from '@/api/articles'
import { useParams } from 'react-router-dom'
import { useState, useEffect } from 'react'

const Detail = () => {
  let { id } = useParams()
  const [item, setItem] = useState({})
  const [fixed, setFixed] = useState(false)

  const scrollCallBack = () => {
    let scrollTop = getScrollTop()
    setFixed(scrollTop > 104)
  }
  useEffect(() => {
    window.addEventListener('scroll', scrollCallBack)

    return () => {
      window.removeEventListener('scroll', scrollCallBack)
    }
  })
  useEffect(() => {
    const init = () => {
      GetArticleDetail(id).then(res => {
        if (res.code === 0) {
          let data = res.data
          data.content = translateMarkdown(data.content)
          setItem(data)
        }
      })
    }
    init()
  }, [id])
  return (
    <div className={'article'}>
      <div className={'article-wrap'}>
        <div className={'article-title'}>
          <h1 className={'title'}>{item.title}</h1>
          <div className={'article-other'}>
            <EyeOutlined className={'icon'}/>
            {item.view}
            <Divider type={'vertical'}/>
            <LikeOutlined className={'icon'}/>
            {item.likes}
            <Divider type={'vertical'}/>
            <PartitionOutlined className={'icon'}/>
            <Tag color={'blue'}>{item.category && item.category.name}</Tag>
            <Divider type={'vertical'}/>
            <TagsOutlined className={'icon'}/>
            {item.tags && item.tags.map(it => (
              <Tag key={it.id} color={it.color}>{it.name}</Tag>
            ))}
          </div>
        </div>
        <div
          className={'article-content article-detail'}
          dangerouslySetInnerHTML={{ __html: item.content || '' }}
        />
      </div>
      <div className={'article-index'}>
        <div className={`index ${fixed ? 'fixed' : ''}`}>
          <Navigation content={item.content || ''}/>
        </div>
      </div>
    </div>
  )
}

export default Detail
