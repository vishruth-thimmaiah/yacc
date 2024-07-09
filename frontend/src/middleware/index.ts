import axios from "axios";

var LoggedIn: boolean | null = null

async function VerifyLogin(): Promise<boolean> {

	const path = (import.meta.env.VITE_BACKEND_URL || "") + '/api/auth/verify'
	return await axios.get(path).then(function(_) {
		return true
	}).catch(function() {
		return false
	})
}

export async function getLoggedIn(): Promise<boolean> {
	if (LoggedIn === null) {
		LoggedIn = await VerifyLogin()
	}
	return LoggedIn
}

export function setLoggedIn(isLoggedIn: boolean) {
	LoggedIn = isLoggedIn
}
