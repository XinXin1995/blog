import {Divider} from 'antd'
import './index.less'
import Personal from './personal'

function About () {
  return (
    <div  className={'about'}>
      <Divider orientation={'left'}>网站简述</Divider>
      <Divider orientation={'left'}>关于我</Divider>
      <Personal />
    </div>
  )
}

export default About
