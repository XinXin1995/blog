import {Divider, Tag} from 'antd'
import {useSelector, useDispatch} from 'react-redux'
import {getTagList} from '@/redux/article/action'
import useMount from '@/hooks/useMount'
import {useHistory} from 'react-router-dom'

function Tags() {
    const history = useHistory()
    const dispatch = useDispatch()
    const tags = useSelector(state => state.article.tagList)
    useMount(() => {
        dispatch(getTagList())
    })
    const handleSearch = tag => {
        history.push(`/list?pageNo=1&pageSize=15&tag=` + tag.id)
    }
    return (
        <div className={'tags'}>
            <Divider orientation={'left'}>
                标签
            </Divider>
            <div className={'tags-wrap'}>
                {
                    tags.map((tag, i) => (
                        <Tag key={i} color={tag.color} onClick={() => handleSearch(tag)}>{tag.name}</Tag>
                    ))
                }
            </div>
        </div>
    )
}

export default Tags
