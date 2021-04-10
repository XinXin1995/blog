import Vue from 'vue'

export default (function (Vue) {
  const filters = {
    formatDateTime: (value) => {
      if (value) {
        let d = new Date(value)
        let year = d.getFullYear()
        let month = d.getMonth() + 1
        month = month < 10 ? '0' + month : month
        let date = d.getDate()
        date = date < 10 ? '0' + date : date
        let h = d.getHours()
        h = h < 10 ? '0' + h : h
        let m = d.getMinutes()
        m = m < 10 ? '0' + m : m
        return `${year}-${month}-${date} ${h}:${m}`
      } else {
        return '- -'
      }
    }
  }
  Object.keys(filters).forEach(filterName => {
    Vue.filter(filterName, filters[filterName])
  })
})(Vue)
