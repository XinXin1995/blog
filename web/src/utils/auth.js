const TokenKey = 'BLOG_TOKEN'

export function setToken (token) {
  localStorage.setItem(TokenKey, token)
}

export function getToken () {
  return localStorage.getItem(TokenKey)
}
