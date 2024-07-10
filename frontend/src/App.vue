<template>
	<RouterView />
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'

const theme = localStorage.getItem("user-theme")
if (theme && theme != "system") {
	document.documentElement.className = theme;
} else {
	document.documentElement.className = (window.matchMedia("(prefers-color-scheme: dark)").matches) ? 'dark' : 'light'
}
</script>


<style scoped></style>
