<template>
	<RouterLink :to="'/u/' + chat"> {{ name }} </RouterLink>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';

defineProps({
	name: String,
	chat: String
});

</script>

<style scoped>
a {
	width: 90%;
	background-color: color-mix(in srgb, var(--accent-color), #00000060);
	margin: 10px 0;
	padding: 10px;
	border: none;
	border-radius: 10px;
	font-family: "M PLUS Rounded 1c", sans-serif;
	transition: 300ms background-color;
	text-decoration: none;
	color: var(--text-secondary-color);

	&:hover {
		background-color: color-mix(in srgb, var(--accent-color), #00000090);
	}
}

.router-link-active {
	background-color: color-mix(in srgb, var(--accent-color), #000000cc);
}
</style>
