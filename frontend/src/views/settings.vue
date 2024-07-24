<template>
	<div class="sidebar">
		<h1>Settings</h1>
		<div>
			<router-link to="/settings/general">General</router-link>
			<router-link to="/settings/account">Account</router-link>
		</div>
	</div>

	<div class="settings">
		<RouterView></RouterView>
	</div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';
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

	div {
		display: flex;
		flex-direction: column;
	}

	h1 {
		font-family: "Playwrite DE Grund", cursive;
		font-weight: bold;
		font-size: 40px;
	}

	a {
		font-family: "Playwrite DE Grund", cursive;
		font-weight: bold;
		font-size: 25px;
		text-decoration: none;
	}
}

.settings {
	position: fixed;
	left: clamp(26%, 21rem, 21rem);
	top: 0.5rem;
	bottom: 0;
	right: 0.5rem;
	user-select: none;
}
</style>
