import {Fragment} from 'react'
import {GetArticles} from '@/api/articles'
import useFetchList from '@/hooks/useFetchList'
import Pagination from '@/components/pagination'
import {Link} from 'react-router-dom'
import {Timeline} from 'antd'
import {ClockCircleOutlined} from '@ant-design/icons'
import {useLocation} from 'react-router-dom'
import './index.less'
import {decodeQuery} from "@/utils";

export const groupBy = (arr, f) => {
    const groups = {}
    arr.forEach(item => {
        let key = item.year = item.createdAt.slice(0, 4)
        item.date = formatDate(item.createdAt)
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
const Note = props => {
    const location = useLocation()
    const {tag} = decodeQuery(location.search)
    const {list, pagination} = useFetchList({
        requestMethod: GetArticles,
        queryParams: {pageSize: 15, category: 0, orderType: 1, tags: tag ? [Number(tag)] : []},
        fetchDependence: [props.location.pathname, props.location.search]
    })
    return (
        <div className={'article-list'}>
            {
                <Timeline mode={'left'}>
                    {
                        list && groupBy(list).map((item, index) => (
                            <Fragment key={index}>
                                <Timeline.Item color="red" dot={<ClockCircleOutlined style={{fontSize: '22px'}}/>}>
                                    <span className={'label-year'}>{[item[0].year]}</span>
                                </Timeline.Item>
                                {
                                    item.map(v => (
                                        <Timeline.Item key={`${index}_${v.id}`}>
                                            <span className={'label-date'}>{v.date}</span>
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
