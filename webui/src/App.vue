<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	methods: {
		logout() {
			this.loading = true
			this.$axios.defaults.headers.common['Authorization'] = null
			localStorage.clear()
			this.$router.push("login")
			this.loading = false
		},
	}
}
</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPhoto</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								Home
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/search" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
								Search
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/post" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#plus-square"/></svg>
								Post
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/" class="nav-link" @click="logout">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
								Logout
							</RouterLink>
						</li>
					</ul>
				</div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style scoped>
</style>
