import { createRouter, createWebHistory } from 'vue-router'
import login from "@/views/login.vue"
import signup from "@/views/signup.vue"
import pageNotFound from "@/views/404.vue"
import chat from '@/views/chats.vue'
import messages from '@/views/messages.vue'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			component: chat,
			children: [
				{
					path: "/u/:user",
					component: messages,
					props: true
				}
			]
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
			path: '/:pathMatch(.*)*',
			component: pageNotFound
		}

	]
})

export default router
