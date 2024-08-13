import { createRouter, createWebHistory } from 'vue-router'
import Login from "@/views/login.vue"
import Signup from "@/views/signup.vue"
import PageNotFound from "@/views/404.vue"
import Chat from '@/views/chats.vue'
import Messages from '@/views/messages.vue'
import Defaults from '@/views/default.vue'
import Settings from '@/views/settings.vue'
import General from '@/views/settings/general.vue'
import Account from '@/views/settings/account.vue'
import Username from '@/views/settings/username.vue'
import Call from '@/views/call.vue'
import { getLoggedIn } from '@/middleware'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			component: Chat,
			name: "home",
			meta: {
				requiresAuth: true
			},
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
					name: "users",

				},
				{
					path: "/u/:chat/call",
					component: Call,
					props: true,
					name: "Call"
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
			path: '/settings',
			name: 'settings',
			component: Settings,
			meta: {
				requiresAuth: true
			},
			children: [
				{
					path: "general",
					component: General
				},
				{
					path: "username",
					component: Username
				},
				{
					path: "account",
					component: Account
				}
			]
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
			path: '/:pathMatch(.*)*',
			component: PageNotFound
		}

	]
})

router.beforeEach(async (to) => {
	if (to.meta.requiresAuth) {
		const loggedIn = await getLoggedIn()
		if (!loggedIn) {
			return { name: "login" };
		}
	}
});

export default router
