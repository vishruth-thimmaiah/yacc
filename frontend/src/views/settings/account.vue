<template>
	<div class="settings">
		<RouterLink to="/settings/username">
			<label>Change Username</label>
			<i class="fa-solid fa-arrow-right"></i>
		</RouterLink>
		<hr />
		<div>
			<label>Logout</label>
			<button class="imp" @click="Logout">Logout</button>
		</div>
		<div>
			<label>Delete Account</label>
			<button class="imp" @click="Delete">Delete</button>
		</div>
	</div>
</template>

<script setup lang="ts">
import { setLoggedIn } from '@/middleware';
import axios from 'axios';
import { useRouter } from 'vue-router';


const router = useRouter()
async function Logout() {
	await axios.get("/api/auth/logout").then(function () {
		setLoggedIn(false)
		router.push("/login")
	})
}


async function Delete() {
	await axios.get("/api/auth/delete").then(function () {
		setLoggedIn(false)
		router.push("/login")
	})
}

</script>

<style scoped>
.settings {
	color: var(--text-primary-color);
	font-family: "M PLUS Rounded 1c", sans-serif;
	font-size: 15px;

	div,
	a {
		margin: 1rem 0;
	}

	a {
		color: var(--text-primary-color);
	}

	label {
		display: inline-block;
		width: 60%;
	}

	button {
		border: none;
		padding: 7px 10px;
		border-radius: 5px;
		font-size: 15px;
		background: transparent;
		color: var(--text-primary-color);
		text-align: right;
		width: 40%;
		transition: background-color 400ms;

		&:hover {
			color: color-mix(in srgb, var(--text-primary-color), #00000060);
		}
	}

	.imp {
		color: var(--text-primary-color);
		background-color: #e6665c;
		text-align: center;

		&:hover {
			background-color: #e6665c80;
		}
	}

}
</style>
