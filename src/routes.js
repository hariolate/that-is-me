import Index from "./components/Index";
import Game from "./components/Game";
import Chat from "./components/Chat";

export default [
    {path:'/', component: Index, meta: {title:"That's Me!"}},
    {path:'/game', component: Game, meta: {title:"That's Me!"}},
    {path:'/chat', component: Chat, meta: {title:"That's Me!"}},
]