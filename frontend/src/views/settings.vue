<template>
	<div class="sidebar">
		<h1> Settings </h1>
		<h2>General</h2>
	</div>
	<div class="settings">
		<h3>UI</h3>
		<hr />
		<div class="dark-mode">
			<label for="dark-mode">Dark Mode</label>
			<select v-model="darkMode" @change="changeMode">
				<option selected="true" value="system">System</option>
				<option value="light">Light</option>
				<option value="dark">Dark</option>
			</select>
		</div>

		<h3>Account</h3>
		<hr />
		<div class="logout">
			<label for="logout">Logout</label>
			<button @click="Logout"> Logout</button>
		</div>
	</div>
</template>

<script setup lang="ts">
import { setLoggedIn } from '@/middleware';
import axios from 'axios';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const darkMode = ref<string>(localStorage.getItem("user-theme") || "system")
function changeMode() {
	const theme = darkMode.value!
	localStorage.setItem("user-theme", theme)

	if (theme == 'system') {
		document.documentElement.className = window.matchMedia("(prefers-color-scheme: dark)").matches ? 'dark' : 'light'
	} else {
		document.documentElement.className = theme;
	}
}

const router = useRouter()
async function Logout() {
	await axios.get("/api/auth/logout").then(function () {
		setLoggedIn(false)
		router.push("/login")
	})
}

</script>

<style scoped>
.sidebar {
	position: fixed;
	width: clamp(25%, 20rem, 20rem);
	background-color: var(--accent-color);
	left: 0;
	bottom: 0;
	top: 0;
	margin: 0.5rem;
	border-radius: 10px;
	user-select: none;
	color: var(--text-secondary-color);

	h1 {
		font-family: "Playwrite DE Grund", cursive;
		font-weight: bold;
		font-size: 40px;
	}

	h2 {
		font-family: "Playwrite DE Grund", cursive;
		font-weight: bold;
		font-size: 25px;
	}
}

.settings {
	position: fixed;
	left: clamp(26%, 21rem, 21rem);
	top: 0;
	bottom: 0;
	right: 0.5rem;
	user-select: none;

	h3 {
		font-family: "Playwrite DE Grund", cursive;
		font-size: 25px;
		color: var(--text-primary-color);
	}

	.dark-mode {
		color: var(--text-primary-color);
		font-family: "M PLUS Rounded 1c", sans-serif;
		font-size: 15px;

		label {
			display: inline-block;
			width: 60%;
		}

		select {
			color: var(--text-primary-color);
			border: none;
			padding: 7px 10px;
			border-radius: 5px;
			background: var(--secondary-color);
			font-size: 15px;
			width: 40%;
		}
	}

	.logout {
		color: var(--text-primary-color);
		font-family: "M PLUS Rounded 1c", sans-serif;
		font-size: 15px;

		label {
			display: inline-block;
			width: 60%;
		}

		button {
			color: var(--text-primary-color);
			border: none;
			padding: 7px 10px;
			border-radius: 5px;
			background-color: #e6665c;
			font-size: 15px;
			width: 40%;
		}
	}
}
</style>
