<template>
	<div class="background">
		<form class="login-box" @submit.prevent="login">
			<label class="heading">Login</label>
			<input class="email" v-model="email" placeholder="email"> </input>
			<input class="password" v-model="passwd" :type="showPasswd ? 'text' : 'password'" placeholder="password">
			</input>
			<div class="show-passwd">
				<input type="checkbox" id="passwd" v-model="showPasswd"></input>
				<label for="passwd">Show Password</label>
			</div>
			<a href="/api/auth/google"><i class="fa-brands fa-google"></i></a>
			<label class="error">{{ error }}</label>
			<button>Login</button>
			<label class="signup">Don't have an account? <router-link to="/signup">sign up</router-link> </label>
			<label class="forgotpasswd">Forgot <router-link to="/forgot">Password</router-link>? </label>

		</form>

	</div>
</template>

<script setup lang="ts">
import { setLoggedIn } from '@/middleware';
import axios, { AxiosError } from 'axios';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const path = (import.meta.env.VITE_BACKEND_URL || "") + '/api/auth/login'
const router = useRouter()

const showPasswd = ref<boolean>(false)
const error = ref<string>()


const email = ref<string>()
const passwd = ref<string>()
async function login() {
	await axios.post(path, {
		email: email.value,
		passwd: passwd.value
	}).then(function () {
		setLoggedIn(true)
		router.push("/");
		error.value = ""

	}).catch(function (res: AxiosError) {
		switch (res.response?.status) {
			case 401:
				error.value = "invalid credentials."
				break;

			default:
				error.value = "an error occured."
				break;
		}
	})

}

</script>

<style scoped>
.background {
	background-image: url("../assets/bg_1.svg");
	position: fixed;
	top: 0;
	bottom: 0;
	right: 0;
	left: 0;
	background-position: center;
	background-repeat: no-repeat;
	background-size: cover;
	display: flex;
	align-items: center;
	justify-content: center;
}

.login-box {
	padding: 5%;
	width: clamp(50%, 15rem, 15rem);
	height: clamp(70%, 25rem, 25rem);
	background-color: rgba(0, 0, 0, 0.3);
	border-radius: 10px;
	box-shadow: 0 0 20px rgba(0, 0, 0, 0.2);
	user-select: none;
	backdrop-filter: blur(10px);
	display: flex;
	flex-direction: column;
	align-items: center;

	input:not([type=checkbox]) {
		border: none;
		border-radius: 10px;
		padding: 20px 10px;
		width: 100%;
		font-size: 16px;
	}

	button {
		grid-area: Login;
		border: none;
		border-radius: 10px;
		padding: 10px 5px;
		width: 10rem;
		margin: 0.5rem 0 2rem 0;
		background-color: var(--accent-color);
	}

	label {
		font-family: "M PLUS Rounded 1c", sans-serif;
		font-weight: bold;
		font-size: 14px;
		color: var(--text-primary-color);

		a {
			color: var(--accent-color);
			text-decoration: none;
		}
	}

	.error {
		margin: auto 0 2rem 0;
		color: red;
	}

	.email {
		margin: 1rem;
	}

	.show-passwd {
		margin-right: auto;
	}

	.heading {
		grid-area: Heading;
		font-family: "Playwrite DE Grund", cursive;
		font-weight: bold;
		font-size: 40px;
		margin-bottom: 5vh;
		color: var(--accent-color);
	}
}
</style>
