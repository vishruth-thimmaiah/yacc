<template>
	<div :class="sending ? 'sending' : 'receiving'">
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
	background-color: red;
	border-radius: 10px 0px 10px 10px;
}

.receiving {
	background-color: blue;
	border-radius: 0px 10px 10px 10px;
}
</style>
