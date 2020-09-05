import Home from "./components/Home";
import Index from "./components/Index";

export default [
    {path:'/', component: Index, meta: {title:"That's me!"}},
    {path:'/ws', component: Home, meta:{title:'Home'}},
]