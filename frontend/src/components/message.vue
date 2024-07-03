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
.sending, .receiving {
	margin: 2rem 1rem;
	padding: 2rem;
}

.sending {
	background-color: #AAD7D9;
	border-radius: 10px 0px 10px 10px;
	margin-left: 4rem;
}

.receiving {
	background-color: #ffa096;
	border-radius: 0px 10px 10px 10px;
	margin-right: 4rem;
}
</style>
