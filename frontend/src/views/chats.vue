<template>
	<sidebar />
	<div class="messages">
		<Topbar />
		<RouterView />
	</div>
</template>

<script setup lang="ts">
import sidebar from '@/components/sidebar.vue';
import Topbar from '@/components/topbar.vue';
import { provide, ref } from 'vue';


const chat_name = ref('')
provide('chat_name', chat_name)
</script>

<style scoped>
.messages {
	position: fixed;
	left: clamp(26%, 21rem, 21rem);
	top: 0;
	bottom: 0;
	right: 0.5rem;
	overflow-y: scroll;
	padding: 4rem 0;
}

@media screen and (max-width: 800px) {
	.sidebar {
		visibility: hidden;
	}


	.sidebar.open {
		visibility: visible;
		z-index: 5;
	}

	.messages {
		left: 0;
	}
}
</style>
