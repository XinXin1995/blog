import Personal from './personalInfo'
import Hot from './hot'
import Tags from './tag'

function Sidebar (props) {
  return (
    <div id={'sidebar'} style={{height: `${props.width}px`}} className={`${props.fixed ? 'fixed' : ''}`}>
      <Personal/>
      <Hot/>
      <Tags/>
    </div>
  )
}

export default Sidebar
