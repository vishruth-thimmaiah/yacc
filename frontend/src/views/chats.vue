<template>
	<sidebar></sidebar>
	<div class="messages">
		<RouterView />
	</div>
</template>

<script setup lang="ts">
import sidebar from '@/components/sidebar.vue';
</script>

<style scoped>

.messages {
	position: fixed;
	left: clamp(26%, 21rem, 21rem);
	top: 0;
	bottom: 0;
	right: 0.5rem;
	overflow-y: scroll;
	padding-bottom: 4rem;
}
</style>
