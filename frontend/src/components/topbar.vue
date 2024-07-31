<template>
	<div class="topbar">
		<h1> {{ chat_name }} </h1>
		<i @click="toggleSidebar" class="fa-solid fa-bars fa-fw"></i>
	</div>
</template>

<script setup lang="ts">
import { inject } from 'vue';

const chat_name = inject('chat_name', '')

function toggleSidebar() {
	const sidebar = document.getElementById("sidebar")
	if (sidebar?.className == "sidebar") {
		sidebar.className = "sidebar open"
	}
	else {
		sidebar!.className = "sidebar"
	}
}
</script>

<style scoped>
.topbar {
	position: fixed;
	left: clamp(26%, 21rem, 21rem);
	right: 0.5rem;
	top: 0.5rem;
	background: var(--accent-color);
	height: 5rem;
	border-radius: 10px;
	display: flex;
	flex-direction: row;
	color: var(--text-secondary-color);

	i {
		visibility: hidden;
	}

	h1 {
		font-family: "M PLUS Rounded 1c", sans-serif;
		font-weight: bold;
		font-size: 30px;
		margin: auto 5px;
	}
}


@media screen and (max-width: 800px) {
	.topbar {
		left: 0.5rem;

		i {
			margin: auto;
			margin-right: 5px;
			font-size: 20px;
			visibility: visible;
		}
	}
}
</style>
