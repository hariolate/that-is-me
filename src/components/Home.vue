<template>
    <div>
        <div>
            <b-navbar
                    type="light"
                    variant="light"
                    id="nav">
                <b-navbar-nav class="ml-auto" >
                    <b-nav-item href="#">WebSocket Demo</b-nav-item>
                </b-navbar-nav>
            </b-navbar>
        </div>
        <b-img
                center
                :src="logoImg"
                alt="UNO Online"
                id="logo"
        ></b-img>
        <pre id="content">{{receive}}</pre>
    </div>
</template>

<script>
    import common from "../common";
    import root from "../protocol/protocol"

    export default {
        name: "Home",
        data() {
            return {
                logoImg: common.logoStatic,
                receive: "",
            }
        },
        mounted() {
            const host = location.host;
            const ws = new WebSocket('ws://'+host+'/ws');
            ws.binaryType = "arraybuffer";
            const rmsg = root.lookupType("RawMessage");
            const msg = root.lookupType("Message");
            const current = this;
            ws.onopen = function() {
                const m = rmsg.create({message:"Hello"});
                if(!rmsg.verify(m)) {
                    const buf = rmsg.encode(m).finish();
                    ws.send(buf);
                    current.receive += JSON.stringify(m) + "\n";
                }
            }
            ws.onmessage = function(evt) {
                const m = msg.decode(new Uint8Array(evt.data));
                current.receive += JSON.stringify(m) + "\n";
            }
        }
    }
</script>

<style scoped>
    #nav {
        border-bottom: solid lightgray 1px;
    }
    #logo {
        margin-top: 2vh;
        max-height: 300px;
    }
    #content {
        width: 70%;
        margin-left: 15%;
        margin-top: 2vh;
        padding: 30px;
        border: 1px solid black;
    }
</style>