import Vue from 'vue'
import VueRouter from 'vue-router'
import routes from './routes'
import axios from 'axios'
import {BootstrapVue,BootstrapVueIcons} from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import App from "./App";

Vue.config.productionTip = false;
Vue.use(BootstrapVue);
Vue.use(BootstrapVueIcons);
Vue.use(VueRouter);
Vue.prototype.axios = axios;
//Vue.prototype.ws = new WebSocket('ws://158.247.202.150:7000/ws');
Vue.prototype.ws = new WebSocket('wss://echo.websocket.org/');
Vue.prototype.ws.binaryType = "arraybuffer";
Vue.prototype.global = {
    status: 0,
    role: -1,
}

const router = new VueRouter({
    mode: 'hash',
    routes
});

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next();
});

new Vue({
    router,
    components: {App},
    template: '<App/>'
}).$mount('#app');