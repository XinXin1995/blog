<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  mounted () {
    // setTimeout(() => {
    window.mainAppLoading.hide()
    // }, 2000)
  }
}
</script>

<style lang="scss">
  #app{
    height: 100%;
    overflow: hidden;
  }
</style>
