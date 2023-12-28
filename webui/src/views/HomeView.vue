<script>
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			photos: {
				records: [],
				lastDate: null,
				lastId: null,
				hasNext: true
			}
		}
	},
	methods: {
		async loadPhotos() {
			this.loading = true
			try {
				let response = await this.$axios.get("/stream", {
					params: {
						startDate: this.photos.lastDate,
						startId: this.photos.lastId,
					}
				})
				if (response.data.records) {
					this.photos.records = this.photos.records.concat(response.data.records)
					this.photos.lastDate = response.data.lastDate
					this.photos.lastId = response.data.lastId
					if (response.data.records.length < this.$paginationLimit) {
						this.photos.hasNext = false
					}
				} else {
					this.photos.hasNext = false
				}

			} catch (e) {
				if (e.response && e.response.status === 400) {
					this.errormsg = "Bad request."
				} else if (e.response && e.response.status === 401) {
					this.errormsg = "Unauthorized."
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "An internal error occurred. Please try again later."
				} else {
					this.errormsg = e.toString()
				}
			}
			this.loading = false
		},
		handlePostDeleted() {
			this.photos = {
				records: [],
				lastDate: null,
				lastId: null,
				hasNext: true
			}
			this.loadPhotos()
		}
	},
	mounted() {
		this.$emit("logged-in")
		this.loadPhotos()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Stream</h1>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<PostList :posts-data="photos" @load-more="loadPhotos" @post-deleted="handlePostDeleted"></PostList>
		<LoadingSpinner :loading="loading"></LoadingSpinner>
	</div>
</template>

<style scoped></style>
