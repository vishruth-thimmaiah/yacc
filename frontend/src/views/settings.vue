<template>
	<div>
		<h1> Settings </h1>
		<h2>General</h2>
		<div class="dark-mode">
			<label for="dark-mode">Dark Mode</label>
			<select v-model="darkMode" @change="changeMode">
				<option selected="true" value="system">System</option>
				<option value="light">Light</option>
				<option value="dark">Dark</option>
			</select>
		</div>

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
h1 {
	font-family: "Playwrite DE Grund", cursive;
	font-weight: bold;
	font-size: 40px;
}

h2 {
	font-family: "Playwrite DE Grund", cursive;
	font-weight: bold;
	font-size: 25px;
}

.dark-mode {
	label {
		user-select: none;
	}

	select {
		border: none;
		padding: 5px;
		border-radius: 5px;
	}
}
</style>
