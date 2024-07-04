<template>
	<div class="sidebar">
		<h1>Contacts</h1>
		<div class="contacts">
			<contact v-for="ctc in contacts" :name="ctc.name" :uuid="ctc.id" />
		</div>
		<div class="options">
			<i class="fa-solid fa-user fa-fw fa-2xl"></i>
			<label>user</label>
		</div>
	</div>
	<div class="messages">
		<RouterView />
		<form class="new-message">
			<input ></input>
		</form>
	</div>
</template>

<script setup lang="ts">
import contact from '@/components/contact.vue';

import axios from 'axios';
import { ref } from 'vue';

const contacts = ref<{ name: string, id: string }[]>([])
axios.get((import.meta.env.VITE_BACKEND_URL || "") + "/api/user/contacts").then(function (response) {
	contacts.value = response.data
})

</script>

<style scoped>
.sidebar {
	position: fixed;
	width: clamp(25%, 20rem, 20rem);
	background-color: #fc8277;
	left: 0;
	bottom: 0;
	top: 0;

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
		height: 80%;
		display: flex;
		flex-direction: column;
	}

	.options {
		position: absolute;
		bottom: 0;
		margin: 0 0 10px 5px;

		label {
			font-family: "M PLUS Rounded 1c", sans-serif;
			font-weight: bold;
			font-size: 20px;
		}
	}
}

.messages {
	position: fixed;
	bottom: 0;
	top: 0;
	left: clamp(25%, 20rem, 20rem);
	right: 0;
	width: auto;
}

.new-message {
	position: absolute;
	bottom: 0;
	/* margin: 10px auto; */
	margin-left: auto;
}
</style>
