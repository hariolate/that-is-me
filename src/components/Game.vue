<template>
    <div>
        <div id="background">
            <h2
                    align="center"
                    id="title"
            >{{title}}</h2>
            <div id="view">
                <canvas
                        id="content"
                ></canvas>
            </div>
        </div>
    </div>
</template>

<script>
    import common from "../common";

    export default {
        name: "Game",
        data() {
            return {
                logoImg: common.logoStatic,
                title: "Ant",
                length: 0,
                ctx: null,
            }
        },
        mounted() {
            window.onload = this.fixLayout;
            window.onresize = this.fixLayout;
            window.onkeydown = this.keyInput;
            this.ctx = document.querySelector("#content").getContext("2d");
            setInterval(this.animate, 50);
        },
        methods: {
            fixLayout() {
                const v = document.querySelector("#view");
                const c = document.querySelector("#content");
                const s = Math.min(window.innerHeight, window.innerWidth);
                this.length = Math.round(s * 0.6);
                v.style.width = this.length + "px";
                v.style.height = this.length + "px";
                c.width = this.length;
                c.height = this.length;
            },
            animate() {
                const t = document.querySelector("#title");
                const v = document.querySelector("#view");
                t.style.opacity = 70 + Math.random() * 30 + "%";
                v.style.opacity = 80 + Math.random() * 20 + "%";
                this.ctx.fillStyle = "#FFFFFF";
                this.ctx.fillRect(0, 0, this.length, this.length);
            },
            keyInput(evt) {
                alert(evt.code);
            },
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
        margin-top: 15vh;
    }

    #content {
        border: lightgrey 1px dashed;
    }

    #view {
        margin-left: 50%;
        transform: translateX(-50%);
    }
</style>