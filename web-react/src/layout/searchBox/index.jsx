import { Input } from 'antd'
import { SearchOutlined } from '@ant-design/icons'
import { useHistory } from 'react-router-dom'


function SearchBox () {
  const history = useHistory()
  const handleChange = (v) => {
    history.push('/list?pageNo=1&keyword=' + v)
  }
  return (
    <div id={'searchBox'}>
      <SearchOutlined className={'search-icon'}/>
      <Input.Search placeholder={'搜索文章'} onSearch={handleChange}/>
    </div>
  )
}

export default SearchBox
