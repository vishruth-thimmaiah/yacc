<template>
	<RouterLink :to="'/u/' + uuid"> {{ name }} </RouterLink>
</template>

<script setup lang="ts">
import { RouterLink} from 'vue-router';

defineProps({
	name: String,
	uuid: String
});

</script>

<style scoped>
a {
	width: 90%;
	background-color: #37846824;
	margin: 10px;
	padding: 10px;
	border: none;
	border-radius: 10px;
	font-family: "M PLUS Rounded 1c", sans-serif;
	transition: 300ms background-color;
	text-decoration: none;

	&:hover {
		background-color: #37846847;
	}
}
</style>
