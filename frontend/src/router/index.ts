import { createRouter, createWebHistory } from 'vue-router'
import Login from "@/views/login.vue"
import Signup from "@/views/signup.vue"
import PageNotFound from "@/views/404.vue"
import Chat from '@/views/chats.vue'
import Messages from '@/views/messages.vue'
import Defaults from '@/views/default.vue'
import Settings from '@/views/settings.vue'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			component: Chat,
			name: "home",
			children: [
				{
					path: "",
					component: Defaults,
					name: "home"
				},
				{
					path: "/u/:chat",
					component: Messages,
					props: true,
					name: "users"
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
			component: Login
		},
		{
			path: '/signup',
			name: 'signup',
			component: Signup
		},
		{
			path: '/settings',
			name: 'settings',
			component: Settings
		},
		{
			path: '/:pathMatch(.*)*',
			component: PageNotFound
		}

	]
})

export default router
