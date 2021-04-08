<template>
    <div class="logo" @click="handleToHome">
        <p class="name">JACOB WOO</p>
        <p class="tip">never say never</p>
    </div>
</template>

<script>
export default {
  methods: {
    handleToHome () {
      this.$router.push({ name: 'home' })
    }
  }
}
</script>

<style lang="scss" scoped>
    .logo{
        display: inline-block;
        color: #fff;
        text-align: left;
        font-family: 'Playfair Display', serif;
        p{
            margin: 0;
        }
        .name{
            font-size: 20px;
            margin-bottom: $-margin/2;
            font-weight: 700;
        }
        .tip{
            color: #828282;
        }
    }
</style>
