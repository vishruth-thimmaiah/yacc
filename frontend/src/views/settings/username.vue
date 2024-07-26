<template>
	<h3>Set Up Your New Username</h3>
	<form @submit.prevent="updateUsername">
		<input placeholder="New Username" v-model="username" type="text"> </input>
	</form>
</template>

<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue';

const username = ref<string>()

async function updateUsername() {
	await axios.post("/api/auth/username", {
		username: username.value
	}).then(function () {
	})
}
</script>

<style scoped>
h3 {
	font-family: "Playwrite DE Grund", cursive;
	font-size: 25px;
	color: var(--text-primary-color);
}

input {
	border: none;
	border-radius: 10px;
	padding: 20px 10px;
	font-size: 16px;
}
</style>
