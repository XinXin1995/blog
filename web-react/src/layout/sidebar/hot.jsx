import { Divider } from 'antd'

const list = [
  {
    title: '用 node 写命令行工具'
  },
  {
    title: '用 node 写命令行工具'
  },
  {
    title: '用 node 写命令行工具'
  },
  {
    title: '用 node 写命令行工具'
  },
  {
    title: '用 node 写命令行工具'
  }
]

function Hot () {
  return (
    <div className={'hot'}>
      <Divider orientation={'left'}>Hot</Divider>

      <ul>
        {list.map((v, i) => <li key={i}>{v.title}</li>)}
      </ul>
    </div>
  )
}

export default Hot
