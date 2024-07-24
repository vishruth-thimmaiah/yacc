<template>
	<h3>UI</h3>
	<hr />
	<div class="dark-mode">
		<label for="dark-mode">Color Scheme</label>
		<select v-model="darkMode" @change="changeMode">
			<option selected="true" value="system">System</option>
			<option value="light">Light</option>
			<option value="dark">Dark</option>
		</select>
	</div>
</template>


<script setup lang="ts">
import { ref } from 'vue';

const darkMode = ref<string>(localStorage.getItem("user-theme") || "system")
function changeMode() {
	const theme = darkMode.value!
	localStorage.setItem("user-theme", theme)

	if (theme == 'system') {
		document.documentElement.className = window.matchMedia("(prefers-color-scheme: dark)").matches ? 'dark' : 'light'
	} else {
		document.documentElement.className = theme;
	}
}

</script>

<style scoped>
h3 {
	font-family: "Playwrite DE Grund", cursive;
	font-size: 25px;
	color: var(--text-primary-color);
}

.dark-mode {
	color: var(--text-primary-color);
	font-family: "M PLUS Rounded 1c", sans-serif;
	font-size: 15px;

	label {
		display: inline-block;
		width: 60%;
	}

	select {
		color: var(--text-primary-color);
		border: none;
		padding: 7px 10px;
		border-radius: 5px;
		background: var(--secondary-color);
		font-size: 15px;
		width: 40%;
	}
}
</style>
