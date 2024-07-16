import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import PageNotFoundView from '../views/PageNotFound.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/',   redirect: '/login'},
        {path: '/login',  component: LoginView},
        {path: '/homepage',   component: HomeView},
		{path: '/:catchAll(.*)', component: PageNotFoundView}		
	]
})

export default router
