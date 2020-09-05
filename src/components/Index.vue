<template>
    <div>
        <div id="background">
            <div id="view">
                <b-img
                        center
                        :src="logoImg"
                        id="logo"
                ></b-img>
                <h4 id="title">{{title}}</h4>
            </div>
        </div>
    </div>
</template>

<script>
    import common from "../common";

    export default {
        name: "Index",
        data() {
            return {
                logoImg: common.logoStatic,
                title: "Enter 2 Enter",
            }
        },
        mounted() {
            window.blur();
            window.onload = this.fixLayout;
            window.onresize = this.fixLayout;
            window.onkeydown = this.enter;
            window.onblur = this.blur;
            window.onfocus = this.focus;
            setInterval(this.twinkle, 50);
        },
        methods: {
            fixLayout() {
                const v = document.querySelector("#view");
                const s = Math.min(window.innerHeight, window.innerWidth);
                v.style.width = s * 0.6 + "px";
                v.style.height = s * 0.6 + "px";
            },
            twinkle() {
                const l = document.querySelector("#logo");
                const t = document.querySelector("#title");
                l.style.opacity = 80 + Math.random() * 20 + "%";
                t.style.opacity = 70 + Math.random() * 30 + "%";
            },
            enter(evt) {
                if (evt.code === "Enter") {
                    this.$router.push("/game");
                    this.$router.go(0);
                }
            },
            blur() {
                this.title = "Eyes on me!";
            },
            focus() {
                this.title = "Enter 2 Enter";
            }
        }
    }
</script>

<style scoped>
    * {
        user-select: none;
        -webkit-user-drag: none;
    }

    #background {
        width: 100%;
        height: 100%;
        position: absolute;
        background: url("../assets/noise.gif") repeat;
    }

    #title {
        width: 50%;
        margin-left: 50%;
        transform: translateX(-50%);
    }

    #view {
        margin-left: 50%;
        margin-top: 20vh;
        transform: translateX(-50%);
    }
</style>