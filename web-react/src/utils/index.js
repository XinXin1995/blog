import marked from 'marked'

export function getScrollTop () {
  let scroll_top = 0
  if (document.documentElement && document.documentElement.scrollTop) {
    scroll_top = document.documentElement.scrollTop
  } else if (document.body) {
    scroll_top = document.body.scrollTop
  }
  return scroll_top
}

// 转化 md 语法为 html
export const translateMarkdown = (plainText, isGuardXss = false) => {
  return marked(isGuardXss ? xss(plainText) : plainText, {
    renderer: new marked.Renderer(),
    gfm: true,
    pedantic: false,
    sanitize: false,
    tables: true,
    breaks: true,
    smartLists: true,
    smartypants: true,
    highlight: function(code) {
      /*eslint no-undef: "off"*/
      return hljs.highlightAuto(code).value
    }
  })
}

export const decodeQuery = url => {
  const params = {}
  const paramsStr = url.replace(/\.*\?/, '') // a=1&b=2&c=&d=xxx&e
  paramsStr.split('&').forEach(v => {
    const d = v.split('=')
    if (d[1] && d[0]) params[d[0]] = d[1]
  })
  return params
}

export function isExternal(path) {
  return /^(https?:|mailto:|tel:|http:)/.test(path)
}
