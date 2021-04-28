import {Divider} from 'antd'
import './index.less'
import Personal from './personal'

function About () {
  return (
    <div  className={'about'}>
      <Divider orientation={'left'}>网站简述</Divider>
      <ul className={'ul'}>
        <li>
          本站用到的技术栈为：前端：react hooks + antd   后端：gin + gorm + mysql
        </li>
        <li>本站笔记为博主平时学习过程中，摘录网上的一些技术文章，仅供学习参考，如有侵权，请联系删除</li>
        <li>本站仅供学习参考使用，不做任何商业用途</li>
      </ul>
      <Divider orientation={'left'}>关于我</Divider>
      <Personal />
    </div>
  )
}

export default About
