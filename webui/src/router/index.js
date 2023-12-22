import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import SearchView from '../views/SearchView.vue'
import PostView from '../views/PostView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/login'},
		{path: '/home', component: HomeView, name: 'Home'},
		{path: '/login', component: LoginView, name: 'Login'},
		{path: '/search', component: SearchView},
		{path: '/post', component: PostView},
	]
})

export default router
