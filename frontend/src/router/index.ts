import { createRouter, createWebHistory } from 'vue-router'
import login from "@/views/login.vue"
import signup from "@/views/signup.vue"
import pageNotFound from "@/views/404.vue"
import messages from '@/views/messages.vue'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			redirect: {
				name: 'login'
			}
		},
		{
			path: '/login',
			name: 'login',
			component: login
		},
		{
			path: '/signup',
			name: 'signup',
			component: signup
		},
		{
			path: '/messages',
			name: 'messages',
			component: messages
		},
		{
			path: '/:pathMatch(.*)*',
			component: pageNotFound
		}

	]
})

export default router
