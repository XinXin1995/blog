import * as TYPES from '@/redux/types'
import {Login, Register} from "@/api/user";
import {message} from "antd";

export const login = params => {
    return dispatch =>
        Login(params).then(res => {
            if (res.code === 0) {
                dispatch({
                    type: TYPES.USER_LOGIN,
                    payload: res.data
                })
                message.success(`登录成功, 欢迎您 ${res.data.user.username}`)
                return res
            }
        })
}

export const register = params => {
    return dispatch =>
        new Promise((resolve, reject) => {
            Register(params).then(res => {
                if (res.code === 0) {
                    message.success('注册成功，请重新登录您的账号！')
                    resolve()
                }
                reject(res)
            }).catch(err => {
                console.error(err)
                reject(err)
            })
        })

}

export const loginout = () => ({
    type: TYPES.USER_LOGIN_OUT
})
