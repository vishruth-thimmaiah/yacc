<template>
	<div class="sidebar">
		<h1>Contacts</h1>
		<div class="contacts">
			<contact v-for="ctc in contacts" :name="ctc.name" :chat="ctc.chat_id" />
		</div>
		<RouterLink to="/settings" class="options">
			<i class="fa-solid fa-gear fa-fw fa-xl"></i>
		</RouterLink>
	</div>
	<div class="messages">
		<RouterView />
	</div>
</template>

<script setup lang="ts">
import contact from '@/components/contact.vue';

import axios from 'axios';
import { ref } from 'vue';

const contacts = ref<{ name: string, chat_id: string }[]>([])
axios.get((import.meta.env.VITE_BACKEND_URL || "") + "/api/user/contacts").then(function (response) {
	contacts.value = response.data
})

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

	h1 {
		font-family: "Playwrite DE Grund", cursive;
		font-weight: bold;
		font-size: 40px;
	}

	.contacts {
		position: absolute;
		right: 0;
		left: 0;
		top: 6rem;
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.options {
		position: absolute;
		bottom: 0;
		margin: 0 0 10px 5px;
		color: black;
	}
}

.messages {
	position: fixed;
	left: clamp(26%, 21rem, 21rem);
	top: 0;
	bottom: 0;
	right: 0.5rem;
}
</style>
