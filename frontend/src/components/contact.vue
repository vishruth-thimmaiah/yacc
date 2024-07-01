<template>
	<div>
		{{name}}
	</div>
</template>

<script setup lang="ts">
defineProps({
	name: String
});

</script>

<style scoped>
div {
	width: 90%;
	background-color: #37846824;
	margin: 10px;
	padding: 10px;
	border-radius: 10px;
}
</style>
