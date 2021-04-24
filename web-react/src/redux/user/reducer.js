import * as TYPES from '@/redux/types'
import {GetUser, RmToken, RMUser, SetToken, SetUser} from "@/utils/auth";

let defaultState = {
    id: '',
    username: '',
    avatar: '',
    role: ''
}
const userInfo = GetUser()
if (userInfo) {
    defaultState = {...defaultState, ...userInfo}
}
export default function UserReducer(state = defaultState, action) {
    const {type, payload} = action
    switch (type) {
        case TYPES.USER_LOGIN:
            let {user, token} = payload
            SetToken(token)
            SetUser(user)
            return {...state, ...user}
        case TYPES.USER_LOGIN_OUT:
            RMUser()
            RmToken()
            return {...state}

        default:
            return state
    }
}