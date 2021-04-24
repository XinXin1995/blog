import {Button, Input} from "antd";
import {InfoCircleOutlined} from "@ant-design/icons";
import {useState} from 'react'


const CommentForm = props => {
    const [replyContent, setReplyContent] = useState('')

    function handleCommentChange(e) {
        setReplyContent(e.currentTarget.value)
    }

    function handleSubmit() {
        let content = replyContent
        let articleId = Number(props.id)
        let parentId = props.parentId
        props.onsubmit({parentId, articleId, content})
    }

    return (
        <div className={'comment-form'}>
            <Input.TextArea value={replyContent} onChange={handleCommentChange}/>
            <div className={'comment-form-btn'}>
                <InfoCircleOutlined/>
                <span className='controls-tip'>支持 Markdown 语法</span>
                <Button type={"primary"} onClick={handleSubmit}>发表评论</Button>
            </div>
        </div>
    )
}

export default CommentForm