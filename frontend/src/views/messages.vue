<template>
	<div>
		<message v-for="msg in messages" :message="msg.message" :sending="msg.sent"></message>
	</div>

	<form @submit.prevent="submitMessage" class="new-message">
		<input placeholder="send something" v-model="newMessage"></input>
		<button><i class="fa-solid fa-reply fa-rotate-180 fa-fw"></i></button>
	</form>
</template>

<script setup lang="ts">
import message from '@/components/message.vue';
import axios from 'axios';
import { onMounted, ref, watch } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
const props = defineProps({
	user: String
})
const messages = ref<{ message: string, sent: boolean }[]>([])

async function loadMessages() {
	messages.value = []
	axios.post((import.meta.env.VITE_BACKEND_URL || "") + "/api/user/message", {
		sender: props.user
	}).then(function (response) {
		messages.value = response.data
	})

}

onMounted(async () => {
	loadMessages()
})

watch(() => route.params, async function () {
	loadMessages()
})

const newMessage = ref<string>()
async function submitMessage() {
	axios.post((import.meta.env.VITE_BACKEND_URL || "") + "/api/send", {
		message: newMessage.value,
		receiver: props.user
	}).then(function (response) {
		messages.value.push({ message: newMessage.value!, sent: false })
		newMessage.value = ""
	})

}
</script>

<style scoped>
.new-message {
	position: fixed;
	bottom: 0;
	right: 0;
	left: clamp(25%, 20rem, 20rem);
	margin: 1rem auto;
	width: max-content;
	border-radius: 10px;
	padding: 10px;
	background: white;

	input {
		font-family: "M PLUS Rounded 1c", sans-serif;
		font-size: 15px;
		border: none;
		width: 50vw;
		z-index: 2;
		color: grey;

		&:focus {
			outline: none;
			color: black;
		}
	}

	button {
		background: none;
		border: none;
		transition: 300ms transform;

		&:hover {
			transform: scale(1.5);
		}
	}
}
</style>
