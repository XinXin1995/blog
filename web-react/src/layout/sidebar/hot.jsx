import { Divider } from 'antd'
import { GetArticles } from '@/api/articles'
import { useState } from 'react'
import { useHistory } from 'react-router-dom'
import useMount from '@/hooks/useMount'


function Hot () {
  const history = useHistory()
  const [list, setList] = useState([])
  useMount(() => {
    GetArticles({ pageSize: 5, pageNo: 1, orderType: 1 }).then(res => {
      if (res.code === 0) {
        let list = res.data.list || []
        setList(list)
      }
    })
  })
  const handleDirect = (v) => {
    history.push('/detail/' + v.id)
  }
  return (
    <div className={'hot'}>
      <Divider orientation={'left'}>TOP</Divider>

      <ul>
        {list.map((v, i) => <li key={i} onClick={() => handleDirect(v)}>{v.title}</li>)}
      </ul>
    </div>
  )
}

export default Hot
