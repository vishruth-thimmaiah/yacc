<template>
	<div class="sidebar">
		<label>Contacts</label>
		<div>
			<contact v-for="_ in 10" name="rando name" />
		</div>
	</div>
	<div class="messages">
		<message v-for="i in 10" message="rando message" :sending="i % 2 == 1" />
	</div>
</template>

<script setup lang="ts">
import contact from '@/components/contact.vue';
import message from '@/components/message.vue';

</script>

<style scoped>
.sidebar {
	position: fixed;
	width: clamp(25%, 25rem, 25rem);
	background-color: #6DB9EF;
	left: 0;
	bottom: 0;
	top: 0;

	label {
		font-family: "Playwrite DE Grund", cursive;
		font-weight: bold;
		font-size: 40px;
	}

	>div {
		margin: 0 auto;
	}
}

.messages {
	margin-left: clamp(25%, 25rem, 25rem);
}
</style>
