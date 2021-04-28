import { Menu } from 'antd'
import { HomeOutlined, BookOutlined, FolderOutlined,  UserOutlined } from '@ant-design/icons'
import { useLocation, useHistory } from 'react-router-dom'

function Nav () {
  const location = useLocation()
  const history = useHistory()
  const handleSetCurrent = (e) => {
    let path = e.key
    if(e.key === '/list'){
      path = path + '?pageNo=1&category=5'
    }
    history.push(path)
  }
  return (
    <Menu  mode="horizontal" id="nav" selectedKeys={[location.pathname]} onClick={handleSetCurrent}>
      <Menu.Item key="/"  icon={<HomeOutlined />}>
        首页
      </Menu.Item>
      <Menu.Item key="/list" icon={<BookOutlined />}>
        笔记
      </Menu.Item>
      <Menu.Item key="/archives" icon={<FolderOutlined />}>
        归档
      </Menu.Item>
      <Menu.Item key="/about" icon={<UserOutlined />}>
        我的
      </Menu.Item>
    </Menu>
  )
}

export default Nav
