import { createRouter, createWebHistory } from "vue-router"

const routes = [
    {
        path: '/ui',
        alias: '/ui/home',
        name: 'home',
        component: () => import('./components/Home.vue')
    },
    {
        path: '/ui/about',
        alias: '/ui/about-us',
        name: 'about',
        component: () => import('./components/About.vue')
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes: routes
})

export default router;
