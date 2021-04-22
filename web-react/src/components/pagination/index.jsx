import React from 'react'
import PropTypes from 'prop-types'
import { Pagination } from 'antd'

function WebPagination({ total, pageNo, onChange, pageSize, style = {} }) {

  return (
    <div className='app-pagination' style={style}>
      <Pagination
        hideOnSinglePage
        current={pageNo}
        onChange={onChange}
        total={total}
        pageSize={pageSize}
      />
    </div>
  )
}

WebPagination.propTypes = {
  total: PropTypes.number.isRequired,
  onChange: PropTypes.func.isRequired,
  pageNo: PropTypes.number.isRequired,
  pageSize: PropTypes.number
}

WebPagination.defaultProps = {
  pageSize: 10
}

export default WebPagination
