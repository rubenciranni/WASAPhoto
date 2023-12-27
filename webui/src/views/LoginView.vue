<script>
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			username: null,
		}
	},
	methods: {
		async login() {
			this.loading = true
			this.errormsg = null
			try {
				let response = await this.$axios.post("/session", { username: this.username })
				let userId = response.data.userId
				localStorage.setItem("userId", userId)
				localStorage.setItem("username", this.username)
				this.$axios.defaults.headers.common['Authorization'] = `Bearer ${userId}`
				this.$emit("logged-in")
				this.$router.push("home")
			} catch (e) {
				if (e.response && e.response.status === 400) {
					this.errormsg = "Invalid username."
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "An internal error occurred. Please try again later."
				} else {
					this.errormsg = e.toString()
				}
			} finally {
				this.loading = false
			}
		},
		isUsernameValid() {
			const usernameRegex = /^[a-zA-Z0-9_-]{3,16}$/
			return usernameRegex.test(this.username)
		}
	}
}
</script>

<template>
	<div class="container">
		<div class="d-flex justify-content-center align-items-center vh-100">
			<div class="text-center">
				<h2 class="mb-4">Login</h2>
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<form @submit.prevent="login" style="max-width: 300px; margin: auto;">
					<div class="mb-3">
						<label for="username" class="form-label">Username</label>
						<input type="text" class="form-control" id="username" aria-describedby="usernameHelp"
							v-model="username" :class="{ 'is-invalid': !isUsernameValid() }">
						<div id="usernameHelp" class="form-text">
							Your username must be 3-16 characters long, containing only letters (a-z, A-Z), numbers (0-9),
							hyphens (-), and underscores (_).
						</div>
					</div>
					<button :disabled="!username || !isUsernameValid() || loading" type="submit"
						class="btn btn-primary">Login</button>
					<LoadingSpinner :loading="loading"></LoadingSpinner>
				</form>
			</div>
		</div>
	</div>
</template>

<style scoped></style>
  