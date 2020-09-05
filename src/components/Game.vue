<template>
    <div>
        <b-modal
                centered
                hide-header
                hide-footer
                id="round">
            <b-img
                    fluid
                    :src="logoCat"></b-img>
        </b-modal>
        <b-modal
                hide-header
                hide-footer
                id="success">
            You survive!
        </b-modal>
        <b-modal
                hide-header
                hide-footer
                id="fail">
            You lose!
        </b-modal>
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

    class Sprite {
        constructor(sprite) {
            this.frames = {};
            this.dir = "idle";
            this.lastdir = "down";
            this.count = {
                up: 0,
                down: 0,
                left: 0,
                right: 0,
            };
            for (const dir in sprite) {
                this.frames[dir] = []
                for (const idx in sprite[dir]) {
                    this.frames[dir].push(new Image());
                    this.frames[dir][idx].src = sprite[dir][idx];
                }
            }
        }

        Dir(dir) {
            this.lastdir = this.dir;
            this.dir = dir;
        }

        Move() {
            if (this.dir === "idle") {
                if (this.lastdir === "idle") {
                    return this.frames['down'][0];
                }
                return this.frames[this.lastdir][0];
            } else {
                if (this.count[this.dir] === 4) this.count[this.dir] = 0;
                const res = this.frames[this.dir][this.count[this.dir]];
                this.count[this.dir]++;
                return res;
            }
        }
    }

    class Player {
        constructor(spriteObj) {
            this.sprite = spriteObj;
            this.speed = 0.02;
            this.padding = 0.05;
            this.x = 0;
            this.y = 0;
        }

        Key(k) {
            if (k === null) {
                this.sprite.Dir("idle");
            }
            switch (k) {
                case "KeyW":
                    this.sprite.Dir("up");
                    break;
                case "KeyS":
                    this.sprite.Dir("down");
                    break;
                case "KeyA":
                    this.sprite.Dir("left");
                    break;
                case "KeyD":
                    this.sprite.Dir("right");
                    break;
                default:
                    break;
            }
        }

        Img() {
            switch (this.sprite.dir) {
                case "idle":
                    break;
                case "up":
                    if (this.y > this.padding) this.y -= this.speed;
                    break;
                case "down":
                    if (this.y < 1 - this.padding * 3) this.y += this.speed;
                    break;
                case "left":
                    if (this.x > this.padding) this.x -= this.speed;
                    break;
                case "right":
                    if (this.x < 1 - this.padding * 2) this.x += this.speed;
                    break;
                default:
                    break;
            }
            return this.sprite.Move();
        }

        Del() {
            this.sprite = null;
        }
    }

    export default {
        name: "Game",
        data() {
            return {
                logoImg: common.logoStatic,
                logoCat: common.logoCat,
                player: new Player(new Sprite(common.sprite.cat02)),
                title: "00:60",
                timeout: 60,
                length: 0,
                ctx: null,
                NPC: [],
                pressed: false,
                NPCdir: [],
            }
        },
        mounted() {
            // if (!(this.global.status === 1) || this.global.role === -1) {
            //     this.$router.push("/");
            //     this.$router.go(0);
            // }
            this.$bvModal.show("round");
            this.fixLayout();
            const cs = new Sprite(common.sprite.cat02);
            for (let i = 0; i < 5; i++) {
                this.NPC.push(new Player(cs));
                this.NPC[i].x = Math.random() * 0.5 + 0.3;
                this.NPC[i].y = Math.random() * 0.5 + 0.3;
            }
            window.onload = this.fixLayout;
            window.onresize = this.fixLayout;
            window.onkeydown = this.keyInput;
            window.onkeyup = this.keyOut;
            document.querySelector("#content").onclick = this.mouseClick;
            this.ctx = document.querySelector("#content").getContext("2d");
            setInterval(this.twinkle, 30);
            setInterval(this.animate, 80);
            const that = this;
            setInterval(() => {
                that.timeout--;
                that.title = "00:" + ('0' + that.timeout).slice(-2);
                if (that.timeout === 0) {
                    that.$bvModal.show("fail");
                    setTimeout(() => {
                        that.$router.push("/");
                        that.$router.go(0);
                    }, 2000);
                }
            }, 1000);
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
                v.style.visibility = "visible";
            },
            twinkle() {
                const t = document.querySelector("#title");
                const v = document.querySelector("#view");
                t.style.opacity = 70 + Math.random() * 30 + "%";
                v.style.opacity = 80 + Math.random() * 20 + "%";
            },
            animate() {
                this.ctx.fillStyle = "#FFFFFF";
                this.ctx.fillRect(0, 0, this.length, this.length);
                this.ctx.drawImage(
                    this.player.Img(),
                    this.player.x * this.length,
                    this.player.y * this.length);
                for (let i = 0; i < 5; i++) {
                    this.ctx.drawImage(
                        this.NPC[i].Img(),
                        this.NPC[i].x * this.length,
                        this.NPC[i].y * this.length);
                }
            },
            keyInput(evt) {
                if (evt.code === "Enter") {
                    this.$bvModal.show("success");
                    const that = this;
                    setTimeout(() => {
                        that.$router.push("/chat");
                        that.$router.go(0);
                    }, 2000);
                } else {
                    this.player.Key(evt.code);
                    if (!this.pressed) {
                        this.pressed = true;
                        this.NPCdir = [];
                        for (let i = 0; i < 5; i++) {
                            const r = Math.floor(Math.random() * 4);
                            if (r === 0) this.NPCdir.push("KeyW");
                            if (r === 1) this.NPCdir.push("KeyA");
                            if (r === 2) this.NPCdir.push("KeyS");
                            if (r === 3) this.NPCdir.push("KeyD");
                        }
                    } else {
                        this.keyNPC();
                    }
                }
            },
            keyOut() {
                this.player.Key(null);
                for (let i = 0; i < 5; i++) {
                    this.NPC[i].Key(null);
                }
                this.pressed = false;
            },
            keyNPC() {
                for (let i = 0; i < 5; i++) {
                    this.NPC[i].Key(this.NPCdir[i]);
                }
            },
            mouseClick(e) {
                if (e.offsetX || e.layerX) {
                    let x = e.offsetX === undefined ? e.layerX : e.offsetX;
                    let y = e.offsetY === undefined ? e.layerY : e.offsetY;
                    x /= this.length;
                    y /= this.length;
                    for (let i=0; i<5; i++) {
                        if (Math.sqrt(Math.pow(this.NPC[i].x-x, 2) + Math.pow(this.NPC[i].y-y,2)) < 0.04) {
                            this.NPC[i].Del();
                        }
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
        margin-top: 15vh;
        color: gray;
    }

    #content {
        border: lightgrey 1px dashed;
    }

    #view {
        visibility: hidden;
        margin-left: 50%;
        background: white;
        transform: translateX(-50%);
    }
</style>