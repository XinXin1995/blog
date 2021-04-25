import { useState, useEffect, useRef } from 'react'
import { Row, Col, BackTop } from 'antd'
import Nav from './Nav'
import SearchBox from '@/layout/searchBox'
import Sidebar from './sidebar'
import logoSrc from '@/assets/img/logo.svg'
import './layout.less'
import { getScrollTop} from '@/utils'
import PublicComponent from "@/components/public";


const Layout = props => {
  const [sidebarFixed, setSideBarFix] = useState(false)
  const col = useRef()
  useEffect(() => {
    let flag = false
    window.addEventListener('scroll', () => {
      if (flag) {
        return
      }
      flag = true
      let scrollTop = getScrollTop()
      setSideBarFix(scrollTop >= 104)
      flag = false
    })
    document.getElementById('sidebar').style.width = col.current.clientWidth + 'px'
  }, [setSideBarFix])

  return (
    <>
      <header id="header">
        <Row>
          <Col xs={20} sm={20} md={6} lg={6} xl={5} xxl={4}>
            <h1>
              <a id="logo" href="/" >
                <img src={logoSrc} alt="logo"/>
                JACBO WOO
              </a>
            </h1>
          </Col>
          <Col className={'menu-row'} xs={4} sm={4} md={18} lg={18} xl={19} xxl={20}>
            <SearchBox/>
            <Nav/>
          </Col>
        </Row>
      </header>
      <main id={'main'}>
        <Row>
          <Col ref={col} xs={0} sm={0} md={6} lg={6} xl={5} xxl={4}>
            <Sidebar fixed={sidebarFixed} width={col && col.clientWidth}/>
          </Col>
          <Col xs={24} sm={24} md={18} lg={18} xl={19} xxl={20}>
            <div className={'main-wrap'}>
              {props.children}
            </div>
          </Col>
        </Row>
      </main>
      <PublicComponent />
      <BackTop />
    </>
  )
}

export default Layout
