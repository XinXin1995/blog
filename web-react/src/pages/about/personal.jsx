import { Rate } from 'antd'
import { MailOutlined, WechatOutlined } from '@ant-design/icons'

const personal = () => {
  const skills = [
    {
      label: '具备扎实的 Javascript 基础，熟练使用 ES6+ 语法。',
      rate: 3
    },
    {
      label: '熟悉 VUE 并理解其原理，熟悉 React 框架及其用法。',
      rate: 3
    },
    {
      label: '熟练使用 Webpack 打包工具，熟悉常用工程化和模块化方案。',
      rate: 3
    },
    {
      label: '熟悉 GIN + GORM 接口的开发与设计！',
      rate: 2
    },
    {
      label: '熟悉 HTTP 协议，缓存、性能优化、安全等，了解浏览器原理。',
      rate: 2
    },
    {
      label: '了解 Docker 的使用方式',
      rate: 2
    },
    {
      label: '熟悉常用的算法与数据结构',
      rate: 2
    }
  ]
  return (
    <ul className='ul'>
      <li>姓名：WuChangxin</li>
      <li>学历专业：本科 计算机科学与技术</li>
      <li>
        联系方式：
        {/* <Icon type='qq' /> 434358603
         <Divider type='vertical' /> */}
        <MailOutlined/>
        <a href='mailto:1132425275@qq.com'>1132425275@qq.com</a>
        &nbsp;&nbsp;
        <WechatOutlined />
        <a href="phone:15896285705">15896285705</a>
      </li>
      <li>坐标：苏州市</li>
      <li>
        技能
        <ul>
          {skills.map((item, i) => (
            <li key={i}>
              {item.label}
              <Rate defaultValue={item.rate} disabled/>
            </li>
          ))}
        </ul>
      </li>
      <li>
        其他
        <ul>
          <li>常用开发工具： vscode、webstorm、goland、git...</li>
          <li>熟悉的 UI 框架： antd、element-ui、vux</li>
          <li>具备良好的编码风格和习惯，团队规范意识，乐于分享</li>
        </ul>
      </li>
      <li>
        个人
        <ul>
          <li>欢迎交流！</li>
        </ul>
      </li>
    </ul>
  )
}

export default personal
