<template>
	<div class="background">
		<form class="signup-box" method="post" action="/api/user/signup">
			<label class="heading">Sign Up</label>
			<input class="email" name="email" placeholder="email"> </input>
			<input class="password" name="passwd" :type="showPasswd ? 'text' : 'password'" placeholder="password">
			</input>
			<div class="show-passwd">
				<input type="checkbox" id="passwd" v-model="showPasswd"></input>
				<label for="test">Show Password</label>
			</div>
			<button>Sign Up</button>
			<label class="signup">Already have an account? <router-link to="/login">Login</router-link></label>
			<label class="forgotpasswd">Forgot <router-link to="/forgot">Password</router-link>?</label>
		</form>
	</div>
</template>

<script setup lang="ts">
const path = (import.meta.env.VITE_BACKEND_URL || "") + '/api/user/signup'

import { ref } from 'vue';

const showPasswd = ref(false)
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

.signup-box {
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
		grid-area: signup;
		border: none;
		border-radius: 10px;
		padding: 10px 5px;
		width: 10rem;
		margin: auto 0 2rem 0;
		background-color: #6DB9EF;
	}

	label {
		font-family: "M PLUS Rounded 1c", sans-serif;
		font-weight: bold;
		font-size: 14px;

		a {
			color: #6DB9EF;
			text-decoration: none;
		}
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
		color: #6DB9EF;
	}
}
</style>
