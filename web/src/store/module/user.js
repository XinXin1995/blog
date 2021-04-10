import types from '../mutationTypes'
import { Login, GetUser } from '@/api/user'
import { setToken } from '@/utils/auth'
const user = {
  state: {
    isAuthed: false,
    data: {}
  },
  mutations: {
    [types.SET_AUTH] (state, b) {
      state.isAuthed = b
    },
    [types.SET_USER_INFO] (state, data) {
      state.data = data
    },
    [types.CLEAR_USER_INFO] (state) {
      state.data = null
    }
  },
  actions: {
    login ({ commit }, param) {
      return new Promise((resolve) => {
        Login(param).then(res => {
          if (res.code === 0) {
            res.data.token && setToken(res.data.token)
            commit(types.SET_USER_INFO, res.data.user)
            commit(types.SET_AUTH, true)
          }
          resolve(res)
        })
      }).catch(error => {
        console.log(error)
      })
    },
    getUser ({ commit }) {
      return new Promise((resolve, reject) => {
        GetUser().then(res => {
          if (res.code === 0) {
            commit(types.SET_USER_INFO, res.data)
            commit(types.SET_AUTH, true)
          }
        })
      })
    }
  }
}
export default user
