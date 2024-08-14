<template>
	<div>
		<video autoplay playsinline id="local"></video>
		<video autoplay playsinline id="remote"></video>
		<div class="controls">
			<i @click="endCall" class="fa-solid fa-phone fa-fw"></i>
			<!-- <i @click="toggleV" class="fa-solid fa-video fa-fw"></i> -->
			<!-- <i class="fa-solid fa-microphone fa-fw"></i> -->
		</div>
	</div>
</template>

<script setup lang="ts">
import { Receive, Send } from '@/api/messages';
import { type Ref, inject, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const props = defineProps<{
	chat: string
}>()
const router = useRouter()

const rtcrequest = inject<Ref<string>>("rtcrequest")
const r = Receive()
var localMedia: MediaStream

const conn = new RTCPeerConnection({
	iceServers: [
		{ urls: 'stun:stun.l.google.com:19302' },
		{ urls: "stun:stun.l.google.com:19302" },
		{ urls: "stun:stun.l.google.com:5349" },
		{ urls: "stun:stun1.l.google.com:3478" },
	]
})

conn.addEventListener('icecandidate', function (event) {
	Send("rtccan", props.chat!, JSON.stringify(event.candidate), null)
})

r.addEventListener('message', function (event) {
	const res = JSON.parse(event.data)
	if (res.type === "rtccan") {
		conn.addIceCandidate(JSON.parse(res.message))
	}
})

conn.addEventListener('connectionstatechange', _ => {
	switch (conn.connectionState) {
		case 'connected':
			alert('connected')
			break
		case 'disconnected':
			router.push({ name: 'users' })
			break
	}
});

var trackid: MediaStreamTrack | null

function endCall() {
	conn.close()
	router.push({ name: 'users' })
}


onMounted(async function () {

	const localStream = document.getElementById("local") as HTMLVideoElement
	const remoteStream = document.getElementById("remote") as HTMLVideoElement

	localMedia = await navigator.mediaDevices.getUserMedia({ video: true, audio: true })
	localStream.srcObject = localMedia

	conn.addEventListener('track', async function (event) {
		const [remotev] = event.streams;
		remoteStream.srcObject = remotev
	})

	localMedia.getTracks().forEach(function (stream) {
		trackid = stream
		conn.addTrack(stream, localMedia)
	})

	if (rtcrequest?.value !== '') {
		await conn.setRemoteDescription(new RTCSessionDescription(JSON.parse(rtcrequest!.value)))
		const ans = await conn.createAnswer()
		await conn.setLocalDescription(ans)
		Send("rtcans", props.chat!, JSON.stringify(ans), null)
	}
	else {

		r.addEventListener('message', async function (event) {

			const res = JSON.parse(event.data)
			if (res.type === "rtcans") {
				const response = JSON.parse(res.message)

				const remote = new RTCSessionDescription(response)
				await conn.setRemoteDescription(remote)

			}
		})

		const offer = await conn.createOffer()
		await conn.setLocalDescription(offer)

		Send("rtcreq", props.chat!, JSON.stringify(offer), null)
	}
})

</script>

<style scoped>
div {
	position: fixed;
	left: 1rem;
	right: 1rem;
	top: 1rem;
	bottom: 1rem;
	background: black;
	border-radius: 10px;
}

video {
	position: absolute;
}

#local {
	height: 25%;
	bottom: 1em;
	left: 1em;
	z-index: 10;
	border-radius: 10px;
	border: var(--accent-color) 2px solid;
}

#remote {
	top: 0.5em;
	bottom: 0.5em;
	left: 0.5em;
	right: 0.5em;
	width: max-content;
	height: max-content;
	max-width: 100%;
	max-height: 100%;
	object-fit: cover;
	margin: auto;
}

.controls {
	position: absolute;
	right: 0;
	z-index: 15;
	background: transparent;

	i {
		padding: 8px;
		border-radius: 20px;
		background: aquamarine;
		margin: 2px;
	}

	i:first-of-type {
		background: #e6665c;
		padding: 8px 20px;
	}
}
</style>
