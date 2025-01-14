<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			token: localStorage.getItem("token"),
			username: localStorage.getItem("username")
		}
	},
	methods: {
		async newPage() {
            this.loading = true;
			this.errormsg = null;
			try {
                this.$router.push({ path: "/profile" });
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
        },
	},
	mounted() {
		
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2"> Hello, {{ username }}!</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newPage">
						Profilo
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
