<template>
    <div>
        <div id="background">
            <h3
                    align="center"
                    id="title"
            >{{title}}</h3>
            <div id="view">
                <div id="chat">
                    <b-form-textarea
                            readonly
                            no-resize
                            v-model="msg"
                            id="message"
                    ></b-form-textarea>
                </div>
                <div id="type">
                    <b-input
                            v-model="inMsg"
                            placeholder="Chat !"
                            id="textarea"
                    ></b-input>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import common from "../common";

    export default {
        name: "Chat",
        data() {
            return {
                logoImg: common.logoStatic,
                title: "00:35",
                msg: "",
                inMsg: "",
            }
        },
        mounted() {
            window.blur();
            window.onload = this.fixLayout;
            window.onresize = this.fixLayout;
            window.onkeydown = this.keyInput;
            setInterval(this.twinkle, 50);
        },
        methods: {
            fixLayout() {
                const v = document.querySelector("#view");
                const c = document.querySelector("#chat");
                const s = Math.min(window.innerHeight, window.innerWidth);
                v.style.width = s * 0.6 + "px";
                v.style.height = s * 0.6 + "px";
                c.style.height = s * 0.4 + "px";
                v.style.visibility = "visible";
            },
            twinkle() {
                const v = document.querySelector("#view");
                const t = document.querySelector("#title");
                v.style.opacity = 80 + Math.random() * 20 + "%";
                t.style.opacity = 70 + Math.random() * 30 + "%";
            },
            keyInput(evt) {
                if (evt.code === "Enter") {
                    if (this.inMsg.length > 0) {
                        this.msg += "A: " + this.inMsg + "\n";
                        this.inMsg = "";
                        const m = document.querySelector("#message");
                        m.scrollTop = m.scrollHeight;
                    }
                }
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
        margin-top: 20vh;
        color: grey;
        transform: translateX(-50%);
    }

    #view {
        visibility: hidden;
        margin-left: 50%;
        transform: translateX(-50%);
    }

    #chat {
        width: 100%;
    }

    #message {
        width: 100%;
        height: 100%;
        background: white;
        border: lightgrey 1px solid;
        border-bottom: none;
        border-radius: 0;
    }

    #message:focus {
        outline: 0px !important;
        -webkit-appearance: none;
        box-shadow: none !important;
    }

    #type {
        width: 100%;
    }

    #textarea {
        width: 100%;
        height: 100%;
        border: lightgrey 1px solid;
        border-radius: 0;
    }

    #textarea:focus {
        outline: 0px !important;
        -webkit-appearance: none;
        box-shadow: none !important;
    }
</style>