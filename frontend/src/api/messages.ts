import axios from "axios"

// const socket = new WebSocket("ws://localhost:3000/api/messages")
const socket = new WebSocket("ws://" + location.host + "/api/messages")

export async function Send(target: string, message: string, attachment: File | null) {
	if (socket.readyState === 1) {
		var attachment_url: string | null = null
		if (attachment) {
			attachment_url = await UploadFile(attachment, target)
		}
		socket.send(JSON.stringify({
			chat_id: target,
			message: message,
			attachment: attachment_url
		}))
	}
}

export function Receive(): WebSocket {
	return socket
}

const path = (import.meta.env.VITE_BACKEND_URL || "") + '/api/messages/upload'
async function UploadFile(file: File, chat_id: string): Promise<string | null> {

	var formData = new FormData()
	formData.append("image", file)
	formData.append("image_name", chat_id)

	return axios.post(path, formData, {
		headers: {
			"Content-Type": "multipart/form-data"
		}
	}).then(function(response) {
		return response.data.url
	}).catch(function() {
		return null
	})
}
