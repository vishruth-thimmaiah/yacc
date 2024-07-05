<template>
	<div :class="sending ? 'receiving' : 'sending'">
		<label>{{ message }}</label>
	</div>
</template>

<script setup lang="ts">
defineProps({
	message: String,
	sending: Boolean
})
</script>

<style scoped>
.sending,
.receiving {
	font-family: "M PLUS Rounded 1c", sans-serif;
	font-size: 15px;
	font-weight: bold;
	margin: 2rem 1rem;
	padding: 2rem;
}

.sending {
	background-color: var(--chat-received-color);
	border-radius: 10px 0px 10px 10px;
	margin-left: 4rem;
}

.receiving {
	background-color: var(--chat-sent-color);
	border-radius: 0px 10px 10px 10px;
	margin-right: 4rem;
}
</style>
