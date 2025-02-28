import './assets/main.css'

import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import HomePage from './components/HomePage.vue'
import AboutPage from './components/AboutPage.vue'
import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

// Routes registration
const routes: RouteRecordRaw[] = [
    { path: '/', component: HomePage },
    { path: '/about', component: AboutPage },
]

// Create Router
const router = createRouter({
    history: createWebHistory(),
    routes,
})

// Create app
createApp(App) // Root Component
    .use(router) // Use Router
    .mount('#app') // Html root element

const app = createApp(App)

app.use(router)

app.mount('#app')
