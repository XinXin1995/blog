import { useState, Fragment } from 'react'
import { GetArticles } from '@/api/articles'
import useMount from '@/hooks/useMount'
import Pagination from '@/components/pagination'
import { Link } from 'react-router-dom'
import { Timeline } from 'antd'
import { ClockCircleOutlined } from '@ant-design/icons'
import './index.less'

export const groupBy = (arr, f) => {
  const groups = {}
  arr.forEach(item => {
    let key = item.year = item.createdAt.slice(0, 4)
    item.createdAt = formatDate(item.createdAt)
    groups[key] = groups[key] || []
    groups[key].push(item)
  })
  return Object.values(groups)
}

const formatDate = (dateStr) => {
  const d = new Date(dateStr)
  let m = d.getMonth()
  let dd = d.getDate()
  m = m >= 10 ? m : `0${m}`
  dd = dd >= 10 ? dd : `0${dd}`
  return `${m}-${dd}`
}
const Note = () => {
  const init = () => {
    GetArticles(pagination).then(res => {
      if (res.code === 0) {
        let data = res.data.list || []
        setList(groupBy(data))
        setPagination({ ...pagination, total: res.data.total })
      }
    })
  }
  const [list, setList] = useState([])
  const [pagination, setPagination] = useState({
    pageNo: 1,
    pageSize: 20,
    keyword: '',
    category: 5,
    order: 1,
    widthContent: false,
    total: 0,
    onChange () {
    }
  })

  useMount(() => {
    init()
  })
  return (
    <div className={'notes'}>
      {
        <Timeline mode={'left'}>
          {
            list.map((item, index) => (
              <Fragment key={index}>
                <Timeline.Item color="red" dot={<ClockCircleOutlined style={{ fontSize: '22px' }}/>}>
                  <span className={'label-year'}>{[item[0].year]}</span>
                </Timeline.Item>
                {
                  item.map(v => (
                    <Timeline.Item key={`${index}_${v.id}`}>
                      <span className={'label-date'}>{v.createdAt}</span>
                      <Link to={`/detail/${v.id}`}>{v.title}</Link>
                    </Timeline.Item>
                  ))
                }
              </Fragment>
            ))
          }
        </Timeline>
      }
      <Pagination {...pagination} />
    </div>
  )
}

export default Note
