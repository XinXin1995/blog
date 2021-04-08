import axios from 'axios'
import { Message } from 'element-ui'
import { BASE_URL } from '@/config'
import { getToken } from '@/utils/auth'

const service = axios.create({
  baseURL: BASE_URL,
  timeout: 1000000 // request timeout
})
service.interceptors.request.use(
  config => {
    config.headers.token = getToken()
    return config
  },
  error => {
    Promise.reject(error)
  }
)
// response interceptor
service.interceptors.response.use(
  response => {
    if (response.data.code !== 0) {
      Message.error(response.data.msg)
    }
    return response.data
  },
  error => {
    return Promise.reject(error)
  }
)
export default service
