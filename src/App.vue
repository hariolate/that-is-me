<template>
    <div id="app">
        <router-view/>
    </div>
</template>

<script>
    export default {
        name: 'App',
        methods:{
            reload () {
                this.isRouterAlive = false
                this.$nextTick(() => (this.isRouterAlive = true))
            }
        }
    }
</script>

<style scoped>

</style>