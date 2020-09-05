import Home from "./components/Home";
import Index from "./components/Index";
import Game from "./components/Game";

export default [
    {path:'/', component: Index, meta: {title:"That's Me!"}},
    {path:'/game', component: Game, meta: {title:"That's Me!"}},
    {path:'/ws', component: Home, meta:{title:'Home'}},
]