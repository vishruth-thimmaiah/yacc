<template>
	<div>
		<message v-for="msg in messages" :message="msg.message" :sending="msg.sent"></message>
	</div>
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

</script>

<style scoped></style>
