<template>
	<div>
		<message v-for="msg in messages" :message="msg.message" :sending="msg.sent"></message>
	</div>
	<form @submit.prevent="submitMessage" class="new-message">
		<input placeholder="send something" v-model="newMessage"></input>
	</form>
</template>

<script setup lang="ts">
import message from '@/components/message.vue';
import axios from 'axios';
import { ref } from 'vue';
const props = defineProps({
	user: String
})

const messages = ref<{ message: string, sent: boolean }[]>([])

axios.post("/api/user/message", {
	sender: props.user
}).then(function (response) {
	messages.value = response.data
})

console.log(messages.value)

const newMessage = ref<string>()
async function submitMessage() {
	axios.post((import.meta.env.VITE_BACKEND_URL || "") + "/api/send", {
		message: newMessage.value,
		receiver: props.user
	})
}
</script>

<style scoped>
.new-message {
	position: absolute;
	bottom: 0;
	/* margin: 10px auto; */
	width: auto;

	input {
		font-family: "M PLUS Rounded 1c", sans-serif;
		font-size: 15px;
		border: none;
		border-radius: 10px;
		padding: 10px;
		margin: 10px;
		width: auto;
		z-index: 2;

		&:focus {
			outline: none;
			background-color: #37846824;
		}
	}
}
</style>
