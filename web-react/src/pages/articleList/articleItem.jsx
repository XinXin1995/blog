import { Divider, Tag } from 'antd'
import { TagsOutlined, PartitionOutlined, LikeOutlined, EyeOutlined } from '@ant-design/icons'
import { useHistory } from 'react-router-dom'
import { useMediaQuery } from 'react-responsive'
import './articleItem.less'

function formatDate (date) {
  date = new Date(date)
  let y = date.getFullYear()
  let m = date.getMonth() + 1
  let d = date.getDate()
  return `${y}/${m < 10 ? '0' + m : m}/${d < 10 ? '0' + d : d}`
}


const ArticleItem = props => {
  const isLessThan736 = useMediaQuery({
    query: '(max-width: 736px)'
  })

  const { item } = props

  const history = useHistory()

  const handleRedirect = () => {
    history.push('/detail/' + item.id)
  }
  return (
    <div className={'article-item'} onClick={handleRedirect}>
      {
        isLessThan736 ?
          (
            <h1 className={'title'}>
              {item.title}
              <p className={'sub-title'}>{formatDate(item.createdAt)}</p>
            </h1>
          )
          :
          (<Divider orientation={'left'}>
            <h1 className={'title'}>
              {item.title}
              <span className={'sub-title'}>{formatDate(item.createdAt)}</span>
            </h1>
          </Divider>)
      }
      <div className={'article-content article-detail'}
           dangerouslySetInnerHTML={{ __html: item.content }}
      >

      </div>
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
  )
}

export default ArticleItem
