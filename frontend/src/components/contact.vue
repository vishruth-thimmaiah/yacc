<template>
	<RouterLink @click="updateTopbar" :to="'/u/' + chat"> {{ name }} </RouterLink>
</template>

<script setup lang="ts">
import { inject, type Ref } from 'vue';
import { RouterLink, useRoute } from 'vue-router';

const props = defineProps({
	name: String,
	chat: String
});

const contact: Ref = inject('chat_name')!

const r = useRoute()
console.log(r.path, `/u/${props.chat}`)
if (r.path === `/u/${props.chat}`) {
	contact.value = props.name
}

function updateTopbar() {
	contact.value = props.name
}

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
