<template>
	<div id="sidebar" class="sidebar">
		<h1>Contacts</h1>
		<div class="contacts">
			<contact v-for="ctc in contacts" :name="ctc.name" :chat="ctc.chat_id" />
		</div>
		<div class="options">
			<form class="new-contact" @submit.prevent="addContact" v-show="newContact">
				<input type="email" v-model="email" placeholder="email id"> </input>
				<button>
					<i class="fa-solid fa-plus fa-fw fa-xl" />
				</button>
			</form>
			<RouterLink to="/settings">
				<i class="fa-solid fa-gear fa-fw fa-xl"></i>
			</RouterLink>
			<i @click="newContact = !newContact" class="fa-solid fa-plus fa-fw fa-xl">
			</i>
		</div>
	</div>
</template>

<script setup lang="ts">
import contact from '@/components/contact.vue';

import axios from 'axios';
import { ref } from 'vue';

const newContact = ref<boolean>(false)
const contacts = ref<{ name: string, chat_id: string }[]>([])
const email = ref()

async function loadContacts() {
	axios.get((import.meta.env.VITE_BACKEND_URL || "") + "/api/user/contacts").then(function (response) {
		contacts.value = response.data
	})
}

function addContact() {
	axios.post((import.meta.env.VITE_BACKEND_URL || "") + "/api/addcontact", {
		email: email.value
	}).then(function (_) {
		loadContacts()
	})
}

loadContacts()

</script>

<style scoped>
.sidebar {
	position: fixed;
	width: clamp(25%, 20rem, 20rem);
	background-color: var(--accent-color);
	left: 0;
	bottom: 0;
	top: 0;
	margin: 0.5rem;
	border-radius: 10px;
	user-select: none;
	color: var(--text-secondary-color);
	box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;

	h1 {
		font-family: "Playwrite DE Grund", cursive;
		font-weight: bold;
		font-size: 40px;
	}

	.contacts {
		position: absolute;
		right: 0;
		left: 0;
		top: 6rem;
		bottom: 3rem;
		display: flex;
		flex-direction: column;
		overflow-y: scroll;
		align-items: center;
	}

	.options {
		position: absolute;
		bottom: 0;
		right: 0;
		left: 0;
		height: 3rem;
		display: flex;

		.new-contact {
			position: absolute;
			background: #88888888;
			padding: 10px;
			border-radius: 10px;
			top: -2rem;
			left: 0.5rem;
			right: 0.5rem;

			input {
				background: none;
				border: none;
				width: 100%;

				&:focus {
					outline: none;
				}
			}

			button {
				background: none;
				border: none;
				position: absolute;
				right: 0.5rem;
			}
		}

		a {
			margin: auto 0;
			margin-right: auto;
			color: var(--text-secondary-color);
		}

		.fa-plus {
			margin: auto 0;
			color: var(--text-secondary-color);
		}
	}
}
</style>
