import {Input} from 'antd'
import {SearchOutlined} from '@ant-design/icons'

function SearchBox () {
  return (
    <div id={'searchBox'}>
      <SearchOutlined className={'search-icon'}/>
      <Input placeholder={'搜索文章'}/>
    </div>
  )
}
export default SearchBox
