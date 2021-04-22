const TokenKey = 'JACBO_BLOG_TOKEN'

export function SetToken (token) {
  localStorage.setItem(TokenKey, token)
}

export function GetToken () {
  localStorage.getItem(TokenKey)
}

export function RmToken () {
  localStorage.removeItem(TokenKey)
}
