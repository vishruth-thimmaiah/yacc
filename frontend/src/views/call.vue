<template>
	<div>
		<video autoplay playsinline id="local"></video>
		<video autoplay playsinline id="remote"></video>
	</div>
</template>

<script setup lang="ts">
import { Receive, Send } from '@/api/messages';
import { type Ref, inject, onMounted } from 'vue';

const props = defineProps<{
	chat: string
}>()

const rtcrequest = inject<Ref<string>>("rtcrequest")
const r = Receive()

onMounted(async function () {

	const localStream = document.getElementById("local") as HTMLVideoElement
	const remoteStream = document.getElementById("remote") as HTMLVideoElement

	const localMedia = await navigator.mediaDevices.getUserMedia({ video: true, audio: false })
	localStream.srcObject = localMedia

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
		if (conn.connectionState === 'connected') {
			alert('connected')
		}
	});

	conn.addEventListener('track', async function (event) {
		const [remotev] = event.streams;
		remoteStream.srcObject = remotev
	})

	localMedia.getTracks().forEach(function (stream) {
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
	background: var(--secondary-color);
}

video {
	width: 400px;
	height: 400px;
}
</style>
