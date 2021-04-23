import { Menu } from 'antd'
import { HomeOutlined, BookOutlined, FolderOutlined, MessageOutlined, UserOutlined } from '@ant-design/icons'
import { useLocation, useHistory } from 'react-router-dom'

function Nav () {
  const location = useLocation()
  const history = useHistory()
  const handleSetCurrent = (e) => {
    history.push(e.key)
  }
  return (
    <Menu  mode="horizontal" id="nav" selectedKeys={[location.pathname]} onClick={handleSetCurrent}>
      <Menu.Item key="/"  icon={<HomeOutlined />}>
        首页
      </Menu.Item>
      <Menu.Item key="/list?pageNo=1&category=5" icon={<BookOutlined />}>
        笔记
      </Menu.Item>
      <Menu.Item key="/archives" icon={<FolderOutlined />}>
        归档
      </Menu.Item>
      <Menu.Item key="/message" icon={<MessageOutlined />}>
        留言
      </Menu.Item>
      <Menu.Item key="/about" icon={<UserOutlined />}>
        关于
      </Menu.Item>
    </Menu>
  )
}

export default Nav
