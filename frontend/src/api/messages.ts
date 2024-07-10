const socket = new WebSocket("ws://localhost:3000/api/messages")

export function Send(target: string, message: string) {
	if (socket.readyState === 1) {
		socket.send(JSON.stringify({
			chat_id: target,
			message: message
		}))
	}
}
