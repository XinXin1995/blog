import marked from 'marked'
import hljs from 'highlight.js'

export function translate (plainText) {
  return marked(plainText, {
    renderer: new marked.Renderer(),
    gfm: true,
    pedantic: false,
    sanitize: false,
    tables: true,
    breaks: true,
    smartLists: true,
    smartypants: true,
    highlight: function (code) {
      return hljs.highlightAuto(code).value
    }
  })
}
