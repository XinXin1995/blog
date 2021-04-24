import * as store from './localStorage'

const TokenKey = 'JACOB_BLOG_TOKEN'
const UserKey = 'JACOB_BLOG_USER'

export function SetToken(token) {
    store.save(TokenKey, token)
}

export function GetToken() {
    return store.get(TokenKey)
}

export function RmToken() {
    store.remove(TokenKey)
}


export function SetUser(userinfo) {
    store.save(UserKey, userinfo)
}

export function GetUser() {
    return store.get(UserKey)
}

export function RMUser() {
    store.remove(UserKey)
}