import { createRouter, createWebHistory } from 'vue-router'
import login from "@/views/login.vue"
import signup from "@/views/signup.vue"
import pageNotFound from "@/views/404.vue"
import chat from '@/views/chats.vue'
import messages from '@/views/messages.vue'
import defaults from '@/views/default.vue'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			component: chat,
			name: "home",
			children: [
				{
					path: "",
					component: defaults
				},
				{
					path: "/u/:chat",
					component: messages,
					props: true
				}
			]
		},
		{
			path: "/u",
			redirect: {
				name: "home"
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
			path: '/:pathMatch(.*)*',
			component: pageNotFound
		}

	]
})

export default router
