<template>
    <div>
        <div id="background">
            <div id="view">
                <b-img
                        center
                        :src="logoImg"
                        id="logo"
                ></b-img>
                <h4 id="title">{{title}}
                    <b-spinner
                            v-show="matchmaking"
                            type="grow"></b-spinner>
                </h4>
            </div>
        </div>
    </div>
</template>

<script>
    import common from "../common";
    import root from "../protocol/protocol"

    export default {
        name: "Index",
        data() {
            return {
                logoImg: common.logoStatic,
                title: "Enter 2 Enter",
                matchmaking: false,
                Request: {
                    Wrapper: root.lookupType("Wrapper"),
                    MatchMakingRequest: root.lookupType("MatchMakingRequest"),
                    MatchMakingResponse: root.lookupType("MatchMakingResponse"),
                    MatchStateChangeEvent: root.lookupType("MatchStateChangeEvent"),
                    Any: root.lookupType("google.protobuf.Any"),
                },
                Wrapper: root.lookupType("Wrapper"),
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
            const that = this;
            this.global.status = 0;
            this.global.role = -1;
            this.ws.onmessage = function (evt) {
                const wrp = that.Request.Wrapper.decode(new Uint8Array(evt.data));
                console.log(JSON.stringify(wrp));
                if (wrp.type === 0) {
                    console.log("Looped!");
                }
                if (wrp.type === 1) {
                    const mrr = that.Request.MatchMakingResponse.decode(wrp.message.value);
                    console.log(JSON.stringify(mrr));
                    if (mrr.type === 0 && this.matchmaking) {
                        if (that.MatchMakingRequest("Accept")) {
                            console.log("Accepted as " + mrr.role);
                            this.global.status = 1;
                            this.global.role = mrr.role;
                            this.$router.push("/game");
                            this.$router.go(0);
                        }
                    }
                }
                if (wrp.type === 2) {
                    const msce = that.Request.MatchStateChangeEvent.decode(wrp.message.value);
                    console.log(JSON.stringify(msce));
                }
            }
        },
        methods: {
            MatchMakingRequest(type) {
                const mmr = this.Request.MatchMakingRequest.fromObject({type: type});
                const mmr_e = this.Request.MatchMakingRequest.encode(mmr).finish();
                const any = this.Request.Any.create({type_url: "MatchMakingRequest", value: mmr_e});
                const wrp = this.Request.Wrapper.fromObject({type: "MatchMakingRequest", message: any});
                const err = this.Request.Wrapper.verify(wrp);
                if (!err) {
                    const wrp_e = this.Request.Wrapper.encode(wrp).finish();
                    this.ws.send(wrp_e);
                    return true;
                } else {
                    console.log(err);
                    return false;
                }
            },
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
                if (evt.code === "Enter" && !this.matchmaking && this.ws.readyState === 1) {
                    if (this.MatchMakingRequest("Begin")) {
                        this.title = "MatchMaking ";
                        this.matchmaking = true;
                    }
                } else if (evt.code === "Escape" && this.matchmaking) {
                    if (this.MatchMakingRequest("Cancel")) {
                        this.title = "Enter 2 Enter";
                        this.matchmaking = false;
                    }
                }
            },
            blur() {
                if (!this.matchmaking) this.title = "Eyes on me!";
            },
            focus() {
                if (!this.matchmaking) this.title = "Enter 2 Enter";
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