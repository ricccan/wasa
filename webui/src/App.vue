<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {}
</script>

<template>

	<div>
		<RouterLink to="/"></RouterLink>
		<RouterLink to="/login"></RouterLink>
		<RouterView />
	</div>
</template>

<style>
</style>
