<template>
	<div :class="sending ? 'receiving' : 'sending'">
		<img v-if="attachment_type === 'image'" :src="attachment" />
		<video v-if="attachment_type === 'video'" controls :src="attachment" />
		<audio v-if="attachment_type === 'audio'" controls :src="attachment" />
		<label class="message">{{ message }}</label>
		<label class="date">{{ cleanDate }}</label>
	</div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue';

const SEC = 1000
const MIN = 60 * SEC
const HRS = 60 * MIN
const DAYS = HRS * 24

const props = defineProps({
	message: String,
	date: [String, Number],
	sending: Boolean,
	attachment_url: String
})

const attachment = ref<string>()
const attachment_type = ref<string>()

async function DownloadMedia() {
	if (props.attachment_url) {
		axios.get(props.attachment_url, {
			responseType: "blob",
		}).then(function (response) {
			attachment_type.value = response.headers["content-type"].split("/")[0];
			attachment.value = URL.createObjectURL(response.data)

		}).catch(function () {
			console.warn("error loading attachment:", props.attachment_url)
		})
	}
}
DownloadMedia()

const cleanDate = ref<string>("")
const diff = Date.now() - Date.parse(props.date!.toString())

let days = Math.floor(diff / DAYS)
let hrs = Math.floor((diff % DAYS) / HRS)
let min = Math.floor((diff % HRS) / MIN)
let sec = Math.floor((diff % MIN) / SEC)

switch (true) {
	case days > 0:
		cleanDate.value = days + (days == 1 ? "day" : "days") + " ago."
		break;
	case hrs > 0:
		cleanDate.value = hrs + (hrs == 1 ? "hr" : "hrs") + " ago."
		break;
	case min > 0:
		cleanDate.value = min + "min ago."
		break;
	case sec > 0:
		cleanDate.value = sec + "sec ago."
		break;
	default:
		cleanDate.value = "a few moments ago."
		break;
}


</script>

<style scoped>
.sending,
.receiving {
	font-family: "M PLUS Rounded 1c", sans-serif;
	font-size: 15px;
	font-weight: bold;
	margin: 2rem 1rem;
	display: flex;
	flex-direction: column;
}

.sending {
	background-color: var(--chat-received-color);
	border-radius: 10px 0px 10px 10px;
	margin-left: 4rem;
}

.receiving {
	background-color: var(--chat-sent-color);
	border-radius: 0px 10px 10px 10px;
	margin-right: 4rem;
}

.message {
	padding: 1rem;
}

.date {
	font-size: 10px;
	padding: 0.5rem;
	padding-top: 0;
	color: #222222;
}

img,
video,
audio {
	margin: 1rem;
	border-radius: 10px;
}
</style>
