import Personal from './personalInfo'
import Hot from './hot'
import Tags from './tag'
import { useMediaQuery } from 'react-responsive'

function Sidebar (props) {
  const isLessThan736 = useMediaQuery({
    query: '(max-width: 736px)'
  })
  if (isLessThan736) {
    return (
      <></>
    )
  }
  return (
    (
      <div id={'sidebar'} style={{ height: `${props.width}px` }} className={`${props.fixed ? 'fixed' : ''}`}>
        <Personal/>
        <Hot/>
        <Tags/>
      </div>
    )
  )
}

export default Sidebar
