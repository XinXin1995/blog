import {Comment, Avatar, message, Spin} from 'antd'
import useFetchList from "@/hooks/useFetchList";
import {AddComment, GetCommentList} from "@/api/comment";
import moment from "moment";
import {Divider, Menu, Dropdown} from "antd";
import {DownOutlined} from '@ant-design/icons'
import {useSelector, useDispatch} from "react-redux";
import useBus from '@/hooks/useBus'
import {loginout} from "@/redux/user/actions";
import CommentForm from './commentForm'
import {useState} from "react";
import Pagination from '@/components/pagination'
import {translateMarkdown} from "@/utils";

const CommentItem = ({item, id, onsubmit}) => {
    const [showCommentForm, setShowCommentForm] = useState(false)
    const handleSwitchCommentForm = () => {
        setShowCommentForm(!showCommentForm)
    }
    function handleComment(params) {
        handleSwitchCommentForm()
        onsubmit(params)
    }
    return (
        <div>
            <Comment
                size={'large'}
                actions={[<span onClick={handleSwitchCommentForm}>回复</span>]}
                author={item.username}
                avatar={
                    <Avatar size={'large'} alt={item.username}>{item.username.slice(0, 1)}</Avatar>
                }
                content={
                    <div className={'article-detail'} dangerouslySetInnerHTML={{__html:translateMarkdown(item.content)}}>
                    </div>
                }
                datetime={
                    <span>{moment().subtract(0, 'days').from(item.createdAt)}</span>
                }
            >

                {showCommentForm && <CommentForm id={id} parentId={item.id} onsubmit={handleComment} />}
                {
                    item.children && <List id={id} list={item.children} onsubmit={onsubmit}/>
                }
            </Comment>
        </div>
    )
}

const List = ({list, id, onsubmit}) => {
    return (
        list && list.map(item => (
            <CommentItem key={item.id} id={id} item={item} onsubmit={onsubmit}/>
        ))
    )

}


const CommentList = props => {
    const bus = useBus()
    const userInfo = useSelector(state => state.user)
    const {username} = userInfo
    const dispatch = useDispatch()
    const renderDropdownMenu = () => {
        return username ? (
            <Menu onClick={handleMenuClick}>
                <Menu.Item key='loginout'>注销</Menu.Item>
            </Menu>
        ) : (
            <Menu onClick={handleMenuClick}>
                <Menu.Item key='login'>登录</Menu.Item>
                <Menu.Item key='register'>注册</Menu.Item>
            </Menu>
        )
    }

    const handleMenuClick = e => {
        switch (e.key) {
            case 'login':
                bus.emit('openSignModal', 'login')

                break

            case 'register':
                bus.emit('openSignModal', 'register')

                break

            case 'loginout':
                dispatch(loginout())
                break

            default:
                break
        }
    }

    const {list, onFetch, loading, pagination} = useFetchList({
        requestMethod: GetCommentList,
        queryParams: {
            pageNo: 1,
            pageSize: 10,
            articleId: Number(props.id)
        },
        usePathParams: false,
        fetchDependence: [props.id]
    })


    const handleSubmitForm = params => {
        AddComment(params).then(res => {
            if (res.code === 0) {
                message.success("评论提交成功")
                onFetch()
            }
        })
    }
    return (
        <Spin spinning={loading}>
        <div className={'comments'}>
            <div className={'comments-header'}>
                <span className={'title'}>评论</span>
                <div className={'right'}>
                    <Dropdown overlay={renderDropdownMenu()} trigger={['click']}>
                    <span className={'dropdown'}>
                      {username || '未登录用户'} <DownOutlined/>
                    </span>
                    </Dropdown>
                </div>
                <Divider/>
            </div>
            <CommentForm id={props.id} parentId={0} onsubmit={handleSubmitForm}/>
            <List id={props.id} list={list}  onsubmit={handleSubmitForm}/>
        </div>
            <Pagination {...pagination} />
        </Spin>
    )
}
export default CommentList