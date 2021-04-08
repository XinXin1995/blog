
var Loading = (function () {
  function Loading () {
    this.id = 'iyyMainAppLoading'
    this.loadingGif = './loading/puff.svg'
    this.createLoadingDOM = function () {
      if (this.id && document.getElementById(this.id)) {
        this.show()
        return
      }
      let content = document.body
      content.style.position = 'relative'
      let div = document.createElement('div')
      div.id = this.id
      div.innerHTML = `<img src="${this.loadingGif}"/>`
      const styleSt = {
        position: 'fixed',
        top: '0',
        left: '0',
        bottom: '0',
        right: '0',
        zIndex: '10000',
        backgroundColor: '#34495E',
        display: 'flex',
        justifyContent: 'space-around',
        alignItems: 'center'
      }
      for (let key in styleSt) {
        div.style[key] = styleSt[key]
      }
      content.appendChild(div)
    }
    this.hide = function () {
      const loading = document.getElementById(this.id)
      if (loading) {
        loading.style.display = 'none'
      }
    }
    this.show = function () {
      const loading = document.getElementById(this.id)
      loading ? loading.style.display = 'flex' : this.createLoadingDOM()
    }
    this.createLoadingDOM()
  }
  return Loading
}())
window.mainAppLoading = new Loading()
