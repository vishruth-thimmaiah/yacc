<template>
	<div id="messages">
		<message v-for="msg in messages" :message="msg.message" :sending="msg.sent" :date="msg.date"
			:attachment_url="msg.attachment"></message>
	</div>

	<form @submit.prevent="submitMessage" class="new-message">
		<input placeholder="send something" v-model="newMessage"></input>
		<input :disabled="attachmentUrl != ''" accept="image/*,video/*,audio/*" @change="attachFile($event)" id="file"
			class="file-upload-input" type="file" />
		<label :class="'upload ' + (attachmentUrl != '' ? 'disabled' : '')" for="file">
			<i class="fa-solid fa-paperclip"></i>
		</label>
		<RouterLink :to="'/u/' + chat + '/call'">
			<i class="fa-solid fa-video"></i>
		</RouterLink>
		<button><i class="fa-solid fa-reply fa-rotate-180 fa-fw"></i></button>
	</form>
</template>

<script setup lang="ts">
import message from '@/components/message.vue';
import axios from 'axios';
import { inject, onMounted, ref, watch, type Ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { Reciever, Send } from "@/api/messages"

const props = defineProps({
	chat: String
})

const route = useRoute()
const router = useRouter()

const messages = ref<{ message: string, sent: boolean, date: string | number, attachment: string }[]>([])
const newMessage = ref<string>("")
var attachment: File | null = null
const attachmentUrl = ref<string>("")

async function loadMessages() {
	messages.value = []
	axios.post((import.meta.env.VITE_BACKEND_URL || "") + "/api/user/message", {
		chat_id: props.chat
	}).then(function (response) {
		messages.value = response.data || []
	})
}

function attachFile(event: Event) {

	const target = event.target as HTMLInputElement
	if (target.files) {
		attachment = target.files[0]
		attachmentUrl.value = URL.createObjectURL(target.files[0])
	}
}

function submitMessage() {
	if ((newMessage.value && newMessage.value != "") || attachmentUrl.value != "") {
		Send("message", props.chat!, newMessage.value, attachment)
		messages.value.push({ message: newMessage.value, sent: false, date: Date.now(), attachment: attachmentUrl.value })
		newMessage.value = ""
		attachment = null
		attachmentUrl.value = ""
	}
}

function notify(message_body: string) {
	if ("Notification" in window) {
		if (Notification.permission === "granted") {
			new Notification(message_body);
		} else if (Notification.permission !== "denied") {
			Notification.requestPermission().then((permission) => {
				if (permission === "granted") {
					new Notification(message_body);
				}
			});
		}
	}
}

const recieve = Reciever

const rtcrequest = inject<Ref<string>>("rtcrequest")

onMounted(async () => {
	loadMessages()

	recieve.addEventListener('message', function (event: any) {
		const response: { type: string, message: string, attachment: string, date: string, chat_id: string } = event.detail
		if (response.chat_id === props.chat!) {
			messages.value.push({ message: response.message, sent: true, date: Date.parse(response.date), attachment: response.attachment })
		}
		else {
			notify(response.message)
		}
	})
	recieve.addEventListener('rtc', function (event: any) {
		if (event.detail.type === "rtcreq") {
			rtcrequest!.value = event.detail.desc
			router.push('/u/' + props.chat! + '/call')
		}
	})
})

watch(() => route.params, async function () {
	loadMessages()
})

</script>

<style scoped>
.new-message {
	position: fixed;
	bottom: 0;
	right: 0.5rem;
	left: clamp(26%, 21rem, 21rem);
	margin: 1rem auto;
	border-radius: 10px;
	padding: 10px;
	background: var(--secondary-color);
	display: flex;

	input {
		font-family: "M PLUS Rounded 1c", sans-serif;
		font-size: 15px;
		border: none;
		width: 50vw;
		z-index: 2;
		color: grey;
		background: none;

		&:focus {
			outline: none;
			color: var(--text-primary-color);
		}
	}

	.file-upload-input {
		display: none;
	}

	.upload {
		color: var(--text-primary-color);
		margin: 0 8px;
		margin-left: auto;

		&.disabled {
			color: grey;
		}
	}

	.fa-video {
		color: var(--text-primary-color);
		margin: auto 8px;
	}

	button {
		background: none;
		border: none;
		transition: 300ms transform;
		color: var(--text-primary-color);

		&:hover {
			transform: scale(1.5);
		}
	}
}

@media screen and (max-width: 800px) {
	.new-message {
		left: 0.5rem;
	}
}
</style>
